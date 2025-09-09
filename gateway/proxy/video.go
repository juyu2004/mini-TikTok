package proxy

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/example/mini-tiktok/proto"
)

func videoClient() (proto.VideoServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("VIDEO_GRPC")|"video:50052", grpc.WithInsecure())
	if err != nil { return nil, nil, err }
	return proto.NewVideoServiceClient(conn), conn, nil
}

func VideoUpload(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("description")
	file, err := c.FormFile("file")
	if err != nil { c.JSON(400, gin.H{"error":"file required"}); return }
	fp, _ := file.Open()
	defer fp.Close()
	buf := make([]byte, file.Size)
	fp.Read(buf)
	cli, conn, err := videoClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Upload(c, &proto.UploadRequest{UserId:"", Title:title, Description:desc, File:buf, Filename:file.Filename})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func VideoFeed(c *gin.Context) {
	cli, conn, err := videoClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Feed(c, &proto.FeedRequest{Page:1, PageSize:10})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func VideoListByUser(c *gin.Context) {
	uid := c.Query("user_id")
	cli, conn, err := videoClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.ListByUser(c, &proto.ListByUserRequest{UserId:uid, Page:1, PageSize:20})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func VideoDelete(c *gin.Context) {
	var body struct{ UserID, VideoID string }
	if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := videoClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	_, err = cli.Delete(c, &proto.DeleteRequest{UserId:body.UserID, VideoId:body.VideoID})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, gin.H{"ok": true})
}
