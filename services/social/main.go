package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"

	"github.com/example/mini-tiktok/pkg/consul"
	"github.com/example/mini-tiktok/pkg/trace"
	"github.com/example/mini-tiktok/proto"
)

type server struct{ proto.UnimplementedSocialServiceServer; db *gorm.DB; rdb *redis.Client }

type Like struct { ID uint `gorm:"primaryKey"`; UserID string; VideoID string }
type Comment struct { ID uint `gorm:"primaryKey"`; UserID string; VideoID string; Content string; CreatedAt int64 }

func main() {
	shutdown := trace.Init("social-service"); defer shutdown(context.Background())
	db, _ := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	db.AutoMigrate(&Like{}, &Comment{})
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})

	lis, _ := net.Listen("tcp", ":50053")
	grpcSrv := grpc.NewServer()
	proto.RegisterSocialServiceServer(grpcSrv, &server{db:db, rdb:rdb})
	go func(){ consul.Register("social", "social-1", 50053) }()
	log.Println("social service on :50053")
	log.Fatal(grpcSrv.Serve(lis))
}
