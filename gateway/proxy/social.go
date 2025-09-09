package proxy

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/example/mini-tiktok/proto"
)

func socialClient() (proto.SocialServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("SOCIAL_GRPC")|"social:50053", grpc.WithInsecure())
	if err != nil { return nil, nil, err }
	return proto.NewSocialServiceClient(conn), conn, nil
}

func Like(c *gin.Context) {
	var body struct{ UserID, VideoID string }
	if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := socialClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Like(c, &proto.LikeRequest{UserId:body.UserID, VideoId:body.VideoID})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func Unlike(c *gin.Context) {
	var body struct{ UserID, VideoID string }
	if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := socialClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Unlike(c, &proto.LikeRequest{UserId:body.UserID, VideoId:body.VideoID})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func Comment(c *gin.Context) {
	var body struct{ UserID, VideoID, Content string }
	if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := socialClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	_, err = cli.Comment(c, &proto.CommentRequest{UserId:body.UserID, VideoId:body.VideoID, Content:body.Content})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, gin.H{"ok": true})
}

func ListComments(c *gin.Context) {
	vid := c.Query("video_id")
	cli, conn, err := socialClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.ListComments(c, &proto.ListCommentsRequest{VideoId:vid, Page:1, PageSize:20})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}
