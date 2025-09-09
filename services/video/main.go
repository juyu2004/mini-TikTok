package main

import (
	"context"
	"log"
	"net"
	"os"
	"fmt"

	"google.golang.org/grpc"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/example/mini-tiktok/pkg/consul"
	"github.com/example/mini-tiktok/pkg/trace"
	"github.com/example/mini-tiktok/proto"
)

type server struct{ proto.UnimplementedVideoServiceServer; db *gorm.DB; minio *minio.Client }

type Video struct { ID string `gorm:"primaryKey"`; Title, Description, URL string; Likes int64; Comments int64; UserID string }

func main() {
	shutdown := trace.Init("video-service"); defer shutdown(context.Background())
	db, _ := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	db.AutoMigrate(&Video{})

	mc, err := minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{Creds:credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""), Secure:false})
	if err != nil { log.Fatal(err) }
	bucket := "videos"; exists, _ := mc.BucketExists(context.Background(), bucket); if !exists { mc.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{}) }

	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()
	proto.RegisterVideoServiceServer(grpcSrv, &server{db:db, minio:mc})
	go func(){ consul.Register("video", "video-1", 50052) }()
	log.Println("video service on :50052")
	log.Fatal(grpcSrv.Serve(lis))
}

func (s *server) Upload(ctx context.Context, in *proto.UploadRequest) (*proto.UploadResponse, error) {
	name := in.Filename
	_, err := s.minio.PutObject(ctx, "videos", name,  bytes.NewReader(in.File), int64(len(in.File)), minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil { return nil, err }
	url := fmt.Sprintf("http://localhost:9000/videos/%s", name)
	v := &Video{ID: uuid.NewString(), Title: in.Title, Description: in.Description, URL: url, UserID: in.UserId}
	s.db.Create(v)
	return &proto.UploadResponse{VideoId: v.ID, PlaybackUrl: url}, nil
}

func (s *server) ListByUser(ctx context.Context, in *proto.ListByUserRequest) (*proto.ListByUserResponse, error) {
	var list []Video; s.db.Where("user_id = ?", in.UserId).Order("id desc").Limit(int(in.PageSize)).Find(&list)
	resp := &proto.ListByUserResponse{}; for _, v := range list { resp.Items = append(resp.Items, &proto.Video{VideoId:v.ID, Title:v.Title, Description:v.Description, PlaybackUrl:v.URL, Likes:v.Likes, Comments:v.Comments}) }
	return resp, nil
}

func (s *server) Feed(ctx context.Context, in *proto.FeedRequest) (*proto.FeedResponse, error) {
	var list []Video; s.db.Order("id desc").Limit(int(in.PageSize)).Find(&list)
	resp := &proto.FeedResponse{}; for _, v := range list { resp.Items = append(resp.Items, &proto.Video{VideoId:v.ID, Title:v.Title, Description:v.Description, PlaybackUrl:v.URL, Likes:v.Likes, Comments:v.Comments}) }
	return resp, nil
}

func (s *server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	s.db.Delete(&Video{}, "id = ? AND user_id = ?", in.VideoId, in.UserId)
	return &proto.DeleteResponse{Ok:true}, nil
}
