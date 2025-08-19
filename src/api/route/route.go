package router

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
        // 健康检查路由（第1天核心功能）
        r.GET("/health", func(c *gin.Context) {
                c.JSON(200, gin.H{"status": "ok"})
        })
}