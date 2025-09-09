package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/example/mini-tiktok/pkg/consul"
	"github.com/example/mini-tiktok/pkg/trace"
	"github.com/example/mini-tiktok/proto"
)

type server struct{ proto.UnimplementedUserServiceServer; db *gorm.DB }

type User struct {
	ID string `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex"`
	PasswordHash string
	Nickname string
	AvatarURL string
	Followers int64
	Following int64
}

func main() {
	shutdown := trace.Init("user-service"); defer shutdown(context.Background())
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil { log.Fatal(err) }
	db.AutoMigrate(&User{})

	lis, _ := net.Listen("tcp", ":50051")
	grpcSrv := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcSrv, &server{db:db})
	go func(){ consul.Register("user", "user-1", 50051) }()
	log.Println("user service on :50051")
	log.Fatal(grpcSrv.Serve(lis))
}

