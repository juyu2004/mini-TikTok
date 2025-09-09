package main

import (
	"context"
	"log"
	"net"
	"os"
	"math/rand"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"

	"github.com/example/mini-tiktok/pkg/consul"
	"github.com/example/mini-tiktok/pkg/trace"
	"github.com/example/mini-tiktok/proto"
)

type server struct{ proto.UnimplementedRecommendServiceServer; db *gorm.DB; rdb *redis.Client }

func main() {
	shutdown := trace.Init("recommend-service"); defer shutdown(context.Background())
	db, _ := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})

	lis, _ := net.Listen("tcp", ":50054")
	grpcSrv := grpc.NewServer()
	proto.RegisterRecommendServiceServer(grpcSrv, &server{db:db, rdb:rdb})
	go func(){ consul.Register("recommend", "recommend-1", 50054) }()
	log.Println("recommend service on :50054")
	log.Fatal(grpcSrv.Serve(lis))
}

func (s *server) Hot(ctx context.Context, in *proto.HotRequest) (*proto.HotResponse, error) {
	// naive: return random IDs from recent videos table (left as exercise). Here, dummy IDs.
	ids := []string{"v1","v2","v3","v4","v5"}
	rand.Shuffle(len(ids), func(i,j int){ ids[i], ids[j] = ids[j], ids[i] })
	if int(in.Size) < len(ids) { ids = ids[:in.Size] }
	return &proto.HotResponse{VideoIds:ids}, nil
}

func (s *server) ForUser(ctx context.Context, in *proto.ForUserRequest) (*proto.ForUserResponse, error) {
	// naive CF placeholder — recommend hot list as fallback
	h, _ := s.Hot(ctx, &proto.HotRequest{Size: in.Size})
	return &proto.ForUserResponse{VideoIds: h.VideoIds}, nil
}
