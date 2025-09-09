package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"github.com/hashicorp/consul/api"

	"gateway/proxy"
	"gateway/mw"
	"gateway/types"
	"github.com/example/mini-tiktok/pkg/trace"
)

func main() {
	shutdown := trace.Init("gateway")
	defer shutdown(context.Background())

	r := gin.Default()
	r.Use(mw.CORSMiddleware())

	v1 := r.Group("/api/v1")

	v1.POST("/auth/register", proxy.UserRegister)
	v1.POST("/auth/login", proxy.UserLogin)
	v1.GET("/user/profile", Auth(), proxy.GetProfile)
	v1.POST("/user/follow", Auth(), proxy.Follow)
	v1.POST("/user/unfollow", Auth(), proxy.Unfollow)

	v1.POST("/video/upload", Auth(), proxy.VideoUpload)
	v1.GET("/video/feed", proxy.VideoFeed)
	v1.GET("/video/list", proxy.VideoListByUser)
	v1.POST("/video/delete", Auth(), proxy.VideoDelete)

	v1.POST("/social/like", Auth(), proxy.Like)
	v1.POST("/social/unlike", Auth(), proxy.Unlike)
	v1.POST("/social/comment", Auth(), proxy.Comment)
	v1.GET("/social/comments", proxy.ListComments)

	r.Run(":8080")
}

func Auth() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"missing token"}); return
		}
		token := strings.TrimPrefix(h, "Bearer ")
		_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
		if err != nil { c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"}); return }
		c.Next()
	}
}

// Example consul-based dialer for gRPC (not used directly in this file, but for reference)
func grpcConn(service string) (*grpc.ClientConn, error) {
	cfg := api.DefaultConfig()
	if addr := os.Getenv("CONSUL_ADDR"); addr != "" { cfg.Address = addr }
	cli, _ := api.NewClient(cfg)
	entries, _, _ := cli.Health().Service(service, "", true, nil)
	addr := entries[0].Service.Address
	port := entries[0].Service.Port
	return grpc.DialContext(context.Background(),  fmt.Sprintf("%s:%d", addr, port), grpc.WithInsecure())
}

