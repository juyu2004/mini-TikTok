package main

import (
	"context"
	"time"
	"github.com/example/mini-tiktok/proto"
)

func (s *server) Like(ctx context.Context, in *proto.LikeRequest) (*proto.LikeResponse, error) {
	s.db.Create(&Like{UserID: in.UserId, VideoID: in.VideoId})
	var count int64; s.db.Model(&Like{}).Where("video_id = ?", in.VideoId).Count(&count)
	return &proto.LikeResponse{Ok:true, LikeCount: count}, nil
}

func (s *server) Unlike(ctx context.Context, in *proto.LikeRequest) (*proto.LikeResponse, error) {
	s.db.Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).Delete(&Like{})
	var count int64; s.db.Model(&Like{}).Where("video_id = ?", in.VideoId).Count(&count)
	return &proto.LikeResponse{Ok:true, LikeCount: count}, nil
}

func (s *server) Comment(ctx context.Context, in *proto.CommentRequest) (*proto.CommentResponse, error) {
	s.db.Create(&Comment{UserID: in.UserId, VideoID: in.VideoId, Content: in.Content, CreatedAt: time.Now().Unix()})
	return &proto.CommentResponse{Ok:true}, nil
}

func (s *server) ListComments(ctx context.Context, in *proto.ListCommentsRequest) (*proto.ListCommentsResponse, error) {
	var items []Comment; s.db.Where("video_id = ?", in.VideoId).Order("id desc").Limit(int(in.PageSize)).Find(&items)
	resp := &proto.ListCommentsResponse{}
	for _, it := range items { resp.Items = append(resp.Items, &proto.Comment{UserId:it.UserID, Content:it.Content, CreatedAt: it.CreatedAt}) }
	return resp, nil
}
