package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/example/mini-tiktok/proto"
	"github.com/example/mini-tiktok/pkg/auth"
	"os"
)

func (s *server) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	u := &User{ID: uuid.NewString(), Email: in.Email, Nickname: in.Nickname, PasswordHash: sha256Hex(in.Password)}
	s.db.Create(u)
	return &proto.RegisterResponse{UserId: u.ID}, nil
}

func (s *server) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	var u User
	if err := s.db.Where("email = ?", in.Email).First(&u).Error; err != nil { return nil, err }
	if u.PasswordHash != sha256Hex(in.Password) { return nil, context.Canceled }
	secret := os.Getenv("JWT_SECRET"); if secret == "" { secret = "dev" }
	token, _ := auth.GenerateToken(secret, u.ID)
	return &proto.LoginResponse{Token: token, UserId: u.ID}, nil
}

func (s *server) GetProfile(ctx context.Context, in *proto.GetProfileRequest) (*proto.GetProfileResponse, error) {
	var u User; s.db.First(&u, "id = ?", in.UserId)
	return &proto.GetProfileResponse{UserId:u.ID, Email:u.Email, Nickname:u.Nickname, AvatarUrl:u.AvatarURL, Followers:u.Followers, Following:u.Following}, nil
}

func (s *server) Follow(ctx context.Context, in *proto.FollowRequest) (*proto.FollowResponse, error) {
	var u User; s.db.First(&u, "id = ?", in.TargetId); u.Followers++; s.db.Save(&u)
	var me User; s.db.First(&me, "id = ?", in.UserId); me.Following++; s.db.Save(&me)
	return &proto.FollowResponse{Ok:true}, nil
}

func (s *server) Unfollow(ctx context.Context, in *proto.FollowRequest) (*proto.FollowResponse, error) {
	var u User; s.db.First(&u, "id = ?", in.TargetId); if u.Followers>0 { u.Followers-- }; s.db.Save(&u)
	var me User; s.db.First(&me, "id = ?", in.UserId); if me.Following>0 { me.Following-- }; s.db.Save(&me)
	return &proto.FollowResponse{Ok:true}, nil
}

func sha256Hex(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
