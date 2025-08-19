package main

import (
        "context"
        "mini-douyin/src/config"
        "mini-douyin/src/api/router"
        "mini-douyin/src/pkg/consul"
        "mini-douyin/src/pkg/tracing"
        "os"
        "os/signal"
        "syscall"
        "time"

        "github.com/gin-contrib/gzip"
        "github.com/gin-gonic/gin"
        "github.com/sirupsen/logrus"
        ginprometheus "github.com/zsais/go-gin-prometheus"
        "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
		"go.opentelemetry.io/otel/semconv"
)

func main() {
        // 1. 初始化配置
        config.InitConfig()

        // 2. 初始化链路追踪
        tp, err := tracing.InitTracer("gateway")
        if err != nil {
                logrus.Fatalf("初始化追踪失败: %v", err)
        }
        defer tp.Shutdown(context.Background())

        // 3. 注册网关服务到Consul
        gatewayServiceName := "gateway-service"
        gatewayAddr := "localhost" // 本地部署地址
        if err := consul.RegisterService(gatewayServiceName, gatewayAddr); err != nil {
                logrus.Fatalf("网关服务注册失败: %v", err)
        }
        logrus.Infof("网关服务已注册到Consul: %s", gatewayServiceName)

        // 4. 初始化Gin框架
        r := gin.Default()
        r.Use(otelgin.Middleware("gateway"))       // 链路追踪中间件
        r.Use(gzip.Gzip(gzip.DefaultCompression))  // 响应压缩中间件

        // 5. 初始化Prometheus监控
        p := ginprometheus.NewPrometheus("douyin_gateway")
        p.Use(r)

        // 6. 注册路由（第1天仅实现健康检查）
        router.RegisterRoutes(r)

        // 7. 启动HTTP服务（带优雅关闭）
        go func() {
                logrus.Infof("网关服务启动，监听地址: %s", config.GatewayAddr)
                if err := r.Run(config.GatewayAddr); err != nil && err != gin.ErrServerClosed {
                        logrus.Fatalf("网关启动失败: %v", err)
                }
        }()

        // 8. 监听退出信号（Ctrl+C）
        quit := make(chan os.Signal, 1)
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
        <-quit
        logrus.Info("开始优雅关闭网关...")

        // 9. 关闭服务
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        if err := r.Shutdown(ctx); err != nil {
                logrus.Fatalf("网关关闭失败: %v", err)
        }
        logrus.Info("网关已关闭")
}