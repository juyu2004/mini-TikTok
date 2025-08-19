package main

import (
        "context"
        "mini-douyin/src/config"
        "mini-douyin/src/pkg/tracing"
        "mini-douyin/src/api/router"
        "github.com/gin-contrib/gzip"
        "github.com/gin-gonic/gin"
        "github.com/sirupsen/logrus"
        ginprometheus "github.com/zsais/go-gin-prometheus"
        "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
        // 初始化链路追踪
        tp, err := tracing.InitTracer("gateway")
        if err != nil {
                logrus.Fatalf("初始化追踪失败: %v", err)
        }
        defer tp.Shutdown(context.Background())

        // 初始化Gin
        r := gin.Default()
        r.Use(otelgin.Middleware("gateway"))       // 链路追踪中间件
        r.Use(gzip.Gzip(gzip.DefaultCompression))  // 压缩中间件
        
        // 初始化Prometheus监控
        p := ginprometheus.NewPrometheus("douyin_gateway")
        p.Use(r)

        // 注册路由
        router.RegisterRoutes(r)

        // 启动服务
        if err := r.Run(config.GatewayAddr); err != nil {
                logrus.Fatalf("网关启动失败: %v", err)
        }
}