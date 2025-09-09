package proxy

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/example/mini-tiktok/proto"
)

func userClient() (proto.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(os.Getenv("USER_GRPC")|"user:50051", grpc.WithInsecure())
	if err != nil { return nil, nil, err }
	return proto.NewUserServiceClient(conn), conn, nil
}

func UserRegister(c *gin.Context) {
	var req struct{ Email, Password, Nickname string }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()}); return }
	cli, conn, err := userClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Register(c, &proto.RegisterRequest{Email:req.Email, Password:req.Password, Nickname:req.Nickname})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, gin.H{"user_id":resp.UserId})
}

func UserLogin(c *gin.Context) {
	var req struct{ Email, Password string }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := userClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.Login(c, &proto.LoginRequest{Email:req.Email, Password:req.Password})
	if err != nil { c.JSON(401, gin.H{"error":"invalid credentials"}); return }
	c.JSON(200, gin.H{"token":resp.Token, "user_id":resp.UserId})
}

func GetProfile(c *gin.Context) {
	uid := c.Query("user_id")
	cli, conn, err := userClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	resp, err := cli.GetProfile(c, &proto.GetProfileRequest{UserId: uid})
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, resp)
}

func Follow(c *gin.Context)   { followOp(c, true) }
func Unfollow(c *gin.Context) { followOp(c, false) }

func followOp(c *gin.Context, on bool) {
	var body struct{ UserID, TargetID string `json:"user_id","target_id"` }
	if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":err.Error()}); return }
	cli, conn, err := userClient(); if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	defer conn.Close()
	if on {
		_, err = cli.Follow(c, &proto.FollowRequest{UserId:body.UserID, TargetId:body.TargetID})
	} else {
		_, err = cli.Unfollow(c, &proto.FollowRequest{UserId:body.UserID, TargetId:body.TargetID})
	}
	if err != nil { c.JSON(500, gin.H{"error":err.Error()}); return }
	c.JSON(200, gin.H{"ok": true})
}
