项目架构总览
采用微服务架构，包含以下核心组件：
- 网关层：处理 HTTP 请求路由、认证鉴权
- 服务层：用户、视频、社交、消息等微服务（GRPC 通信）
- 存储层：PostgreSQL（主数据）、Redis（缓存 / 会话）、MinIO（视频存储）
- 中间件：RabbitMQ（异步任务）、Consul（服务发现）、Jaeger（链路追踪）
- 监控层：Prometheus + Grafana（指标监控）、Pyroscope（性能分析）
第 0 天：项目基础配置
1.1 项目根目录文件
1. go.mod（依赖版本锁定）
module github.com/yourname/mini-douyin

go 1.20

require (
        github.com/gin-contrib/gzip v0.0.6
        github.com/gin-gonic/gin v1.9.1
        github.com/google/uuid v1.3.0
        github.com/hashicorp/consul/api v1.26.0
        github.com/hashicorp/golang-lru/v2 v2.0.7
        github.com/redis/go-redis/v9 v9.5.1
        github.com/sirupsen/logrus v1.9.3
        github.com/stretchr/testify v1.8.4
        github.com/zsais/go-gin-prometheus v0.1.0
        go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.46.0
        go.opentelemetry.io/otel v1.21.0
        go.opentelemetry.io/otel/exporters/jaeger v1.21.0
        go.opentelemetry.io/otel/sdk v1.21.0
        go.opentelemetry.io/otel/semconv/v1.21.0 v0.1.0
        golang.org/x/crypto v0.14.0
        google.golang.org/grpc v1.59.0
        google.golang.org/protobuf v1.31.0
        gorm.io/driver/postgres v1.5.2
        gorm.io/gorm v1.25.3
)

require (
        // 间接依赖（自动生成，无需手动修改）
        github.com/bytedance/sonic v1.10.2 // indirect
        github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
        github.com/cespare/xxhash/v2 v2.2.0 // indirect
        github.com/davecgh/go-spew v1.1.1 // indirect
        github.com/gin-contrib/sse v0.1.0 // indirect
        github.com/go-logr/logr v1.2.4 // indirect
        github.com/go-logr/stdr v1.2.2 // indirect
        github.com/go-playground/locales v0.14.1 // indirect
        github.com/go-playground/universal-translator v0.18.1 // indirect
        github.com/go-playground/validator/v10 v10.15.0 // indirect
        github.com/gogo/protobuf v1.3.2 // indirect
        github.com/golang/protobuf v1.5.3 // indirect
        github.com/google/pprof v0.0.0-20230825180558-8a8e88d8567c // indirect
        github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
        github.com/hashicorp/go-hclog v1.5.0 // indirect
        github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
        github.com/hashicorp/go-multierror v1.1.1 // indirect
        github.com/hashicorp/go-plugin v1.4.10 // indirect
        github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
        github.com/hashicorp/go-rootcerts v1.0.2 // indirect
        github.com/hashicorp/go-sockaddr v1.0.2 // indirect
        github.com/hashicorp/go-uuid v1.0.3 // indirect
        github.com/hashicorp/go-version v1.6.0 // indirect
        github.com/hashicorp/hcl v1.0.0 // indirect
        github.com/hashicorp/logutils v1.0.0 // indirect
        github.com/hashicorp/serf v0.10.1 // indirect
        github.com/json-iterator/go v1.1.12 // indirect
        github.com/josharian/native v1.1.0 // indirect
        github.com/klauspost/cpuid/v2 v2.2.5 // indirect
        github.com/leodido/go-urn v1.2.4 // indirect
        github.com/magiconair/properties v1.8.7 // indirect
        github.com/mailru/easyjson v0.7.7 // indirect
        github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
        github.com/mitchellh/go-homedir v1.1.0 // indirect
        github.com/mitchellh/go-testing-interface v1.14.1 // indirect
        github.com/mitchellh/mapstructure v1.5.0 // indirect
        github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
        github.com/modern-go/reflect2 v1.0.2 // indirect
        github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
        github.com/pelletier/go-toml/v2 v2.1.1 // indirect
        github.com/pmezard/go-difflib v1.0.0 // indirect
        github.com/prometheus/client_golang v1.17.0 // indirect
        github.com/prometheus/client_model v0.4.1 // indirect
        github.com/prometheus/common v0.44.0 // indirect
        github.com/prometheus/procfs v0.11.1 // indirect
        github.com/rogpeppe/go-internal v1.10.0 // indirect
        github.com/russross/blackfriday/v2 v2.1.0 // indirect
        github.com/shirou/gopsutil/v3 v3.24.4 // indirect
        github.com/sirupsen/logrus v1.9.3 // indirect
        github.com/spf13/afero v1.10.0 // indirect
        github.com/spf13/cast v1.5.1 // indirect
        github.com/spf13/cobra v1.7.0 // indirect
        github.com/spf13/jwalterweatherman v1.1.0 // indirect
        github.com/spf13/pflag v1.0.5 // indirect
        github.com/spf13/viper v1.16.0 // indirect
        github.com/subosito/gotenv v1.4.2 // indirect
        github.com/tklauser/go-sysconf v0.3.12 // indirect
        github.com/tklauser/numcpus v0.6.1 // indirect
        github.com/ugorji/go/codec v1.2.11 // indirect
        github.com/x448/float16 v0.8.4 // indirect
        go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
        go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0 // indirect
        go.opentelemetry.io/otel/metric v1.21.0 // indirect
        go.opentelemetry.io/otel/trace v1.21.0 // indirect
        go.uber.org/multierr v1.11.0 // indirect
        go.uber.org/zap v1.26.0 // indirect
        golang.org/x/net v0.17.0 // indirect
        golang.org/x/sys v0.13.0 // indirect
        golang.org/x/text v0.13.0 // indirect
        golang.org/x/time v0.3.0 // indirect
        google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
        google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
        google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0 // indirect
        gopkg.in/ini.v1 v1.67.0 // indirect
        gopkg.in/yaml.v3 v3.0.1 // indirect
)
2. .env.example（环境变量模板）
# 网关配置
GATEWAY_ADDR=:8080

# JWT配置
JWT_SECRET=mini-douyin-2024-secret-key
PASSWORD_SALT=mini-douyin-salt-123

# Consul配置
CONSUL_ADDR=localhost:8500
RPC_PORT=50051

# PostgreSQL配置（与docker-compose一致）
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=douyin
POSTGRES_PASSWORD=douyin
POSTGRES_DB=douyin

# Redis配置
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# RabbitMQ配置（与docker-compose一致）
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=douyin
RABBITMQ_PASSWORD=douyin
RABBITMQ_VHOST=/

# MinIO配置
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_VIDEO_BUCKET=videos
MINIO_IMAGE_BUCKET=images

# Jaeger配置
JAEGER_ENDPOINT=http://localhost:14268/api/traces
3.  .gitignore（忽略无关文件）
# Go编译产物
*.exe
*.exe~
*.dll
*.so
*.dylib
*.out
*.test
*.prof
bin/
dist/
build/

# 依赖缓存
go.sum
vendor/

# 环境变量文件
.env
.env.local
.env.dev
.env.prod

# 日志文件
logs/
*.log

# 临时文件
tmp/
/tmp/
*.tmp
*.swp
.DS_Store

# IDE配置
.idea/
.vscode/
*.iml
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?

# Docker相关（可选，本地开发可保留）
.dockerignore
!docker-compose.yml
!docker/
1.2 配置文件实现（src/config/config.go）
package config

import (
        "os"
        "strconv"
        "time"

        "github.com/joho/godotenv"
        "github.com/sirupsen/logrus"
)

// 全局配置变量（程序启动后只读）
var (
        // 网关配置
        GatewayAddr string

        // JWT配置
        JWTSecret     string
        PasswordSalt  string
        JWTExpireTime time.Duration

        // Consul配置
        ConsulAddr string
        RPCPort    int

        // PostgreSQL配置
        PostgresHost     string
        PostgresPort     int
        PostgresUser     string
        PostgresPassword string
        PostgresDB       string

        // Redis配置
        RedisAddr     string
        RedisPassword string
        RedisDB       int

        // RabbitMQ配置
        RabbitMQHost     string
        RabbitMQPort     int
        RabbitMQUser     string
        RabbitMQPassword string
        RabbitMQVHost    string

        // MinIO配置
        MinIOEndpoint     string
        MinIOAccessKey    string
        MinIOSecretKey    string
        MinIOVideoBucket  string
        MinIOImageBucket  string

        // Jaeger配置
        JaegerEndpoint string
)

// InitConfig 初始化配置（从.env文件加载）
func InitConfig() {
        // 加载.env文件（优先加载本地.env，无则加载.env.example）
        if err := godotenv.Load(".env"); err != nil {
                logrus.Warnf("加载.env文件失败，尝试加载.env.example: %v", err)
                if err := godotenv.Load(".env.example"); err != nil {
                        logrus.Fatalf("加载.env.example文件失败: %v", err)
                }
        }

        // 网关配置
        GatewayAddr = getEnv("GATEWAY_ADDR", ":8080")

        // JWT配置
        JWTSecret = getEnv("JWT_SECRET", "default-secret-key") // 生产环境必须修改
        PasswordSalt = getEnv("PASSWORD_SALT", "default-salt")
        JWTExpireTime = 7 * 24 * time.Hour // 固定7天过期

        // Consul配置
        ConsulAddr = getEnv("CONSUL_ADDR", "localhost:8500")
        RPCPort = getEnvInt("RPC_PORT", 50051)

        // PostgreSQL配置
        PostgresHost = getEnv("POSTGRES_HOST", "localhost")
        PostgresPort = getEnvInt("POSTGRES_PORT", 5432)
        PostgresUser = getEnv("POSTGRES_USER", "douyin")
        PostgresPassword = getEnv("POSTGRES_PASSWORD", "douyin")
        PostgresDB = getEnv("POSTGRES_DB", "douyin")

        // Redis配置
        RedisAddr = getEnv("REDIS_ADDR", "localhost:6379")
        RedisPassword = getEnv("REDIS_PASSWORD", "")
        RedisDB = getEnvInt("REDIS_DB", 0)

        // RabbitMQ配置
        RabbitMQHost = getEnv("RABBITMQ_HOST", "localhost")
        RabbitMQPort = getEnvInt("RABBITMQ_PORT", 5672)
        RabbitMQUser = getEnv("RABBITMQ_USER", "douyin")
        RabbitMQPassword = getEnv("RABBITMQ_PASSWORD", "douyin")
        RabbitMQVHost = getEnv("RABBITMQ_VHOST", "/")

        // MinIO配置
        MinIOEndpoint = getEnv("MINIO_ENDPOINT", "localhost:9000")
        MinIOAccessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
        MinIOSecretKey = getEnv("MINIO_SECRET_KEY", "minioadmin")
        MinIOVideoBucket = getEnv("MINIO_VIDEO_BUCKET", "videos")
        MinIOImageBucket = getEnv("MINIO_IMAGE_BUCKET", "images")

        // Jaeger配置
        JaegerEndpoint = getEnv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")

        logrus.Info("配置初始化完成")
}

// getEnv 读取环境变量，不存在则返回默认值
func getEnv(key, def string) string {
        val := os.Getenv(key)
        if val == "" {
                return def
        }
        return val
}

// getEnvInt 读取整数类型环境变量，解析失败则返回默认值
func getEnvInt(key string, def int) int {
        val := os.Getenv(key)
        if val == "" {
                return def
        }
        res, err := strconv.Atoi(val)
        if err != nil {
                logrus.Warnf("环境变量%s解析为整数失败，使用默认值%d: %v", key, def, err)
                return def
        }
        return res
}
1.3 公共工具补充
1. 链路追踪工具（src/pkg/tracing/tracer.go）
package tracing

import (
        "context"
        "mini-douyin/src/config"

        "go.opentelemetry.io/otel"
        "go.opentelemetry.io/otel/exporters/jaeger"
        "go.opentelemetry.io/otel/sdk/resource"
        tracesdk "go.opentelemetry.io/otel/sdk/trace"
        semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitTracer 初始化Jaeger链路追踪
func InitTracer(serviceName string) (*tracesdk.TracerProvider, error) {
        // 创建Jaeger exporter（基于HTTP协议）
        exp, err := jaeger.New(
                jaeger.WithCollectorEndpoint(
                        jaeger.WithEndpoint(config.JaegerEndpoint), // 从配置读取Jaeger地址
                ),
        )
        if err != nil {
                return nil, err
        }

        // 创建TracerProvider（配置采样率为100%，开发环境用；生产环境可改为0.5）
        tp := tracesdk.NewTracerProvider(
                tracesdk.WithBatcher(exp), // 批量导出追踪数据
                tracesdk.WithResource(resource.NewWithAttributes(
                        semconv.SchemaURL,
                        semconv.ServiceName(serviceName), // 服务名称（用于Jaeger识别）
                )),
                tracesdk.WithSampler(tracesdk.AlwaysSample()), // 100%采样
        )

        // 设置全局Tracer（后续代码可通过otel.Tracer()获取）
        otel.SetTracerProvider(tp)
        return tp, nil
}

// GetTracer 获取Tracer实例（业务代码中用）
func GetTracer(serviceName string) tracesdk.Tracer {
        return otel.Tracer(serviceName)
}
2. MinIO 工具（src/pkg/minio/minio.go）
package minio

import (
        "context"
        "mini-douyin/src/config"
        "net/url"
        "time"

        "github.com/minio/minio-go/v7"
        "github.com/minio/minio-go/v7/pkg/credentials"
        "github.com/sirupsen/logrus"
)

// 全局MinIO客户端
var Client *minio.Client

// 桶名称（从配置读取）
var (
        VideoBucket = config.MinIOVideoBucket
        ImageBucket = config.MinIOImageBucket
)

// InitMinIO 初始化MinIO客户端
func InitMinIO() error {
        // 解析MinIO地址
        endpoint := config.MinIOEndpoint
        useSSL := false // 本地开发用HTTP，生产环境可改为true

        // 创建客户端
        client, err := minio.New(endpoint, &minio.Options{
                Creds:  credentials.NewStaticV4(config.MinIOAccessKey, config.MinIOSecretKey, ""),
                Secure: useSSL,
        })
        if err != nil {
                return err
        }

        // 检查/创建视频桶
        if err := createBucketIfNotExist(client, VideoBucket); err != nil {
                return err
        }

        // 检查/创建图片桶（封面、头像等）
        if err := createBucketIfNotExist(client, ImageBucket); err != nil {
                return err
        }

        Client = client
        logrus.Info("MinIO客户端初始化完成")
        return nil
}

// createBucketIfNotExist 检查桶是否存在，不存在则创建
func createBucketIfNotExist(client *minio.Client, bucketName string) error {
        exists, err := client.BucketExists(context.Background(), bucketName)
        if err != nil {
                return err
        }
        if !exists {
                // 创建桶（区域填us-east-1，MinIO默认区域）
                if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
                        Region: "us-east-1",
                }); err != nil {
                        return err
                }
                logrus.Infof("MinIO桶%s创建成功", bucketName)
        } else {
                logrus.Infof("MinIO桶%s已存在", bucketName)
        }
        return nil
}

// UploadFile 上传文件到MinIO
func UploadFile(ctx context.Context, bucketName, objectName string, file interface{}, fileSize int64) error {
        // 上传选项（公开读权限，方便前端直接访问）
        opts := minio.PutObjectOptions{
                ContentType: "application/octet-stream", // 通用二进制类型，可根据文件类型修改
        }

        // 执行上传
        _, err := Client.PutObject(ctx, bucketName, objectName, file, fileSize, opts)
        if err != nil {
                logrus.Errorf("MinIO上传文件失败（桶：%s，对象：%s）: %v", bucketName, objectName, err)
                return err
        }
        return nil
}

// GetFileURL 获取文件访问URL（带过期时间，默认7天）
func GetFileURL(bucketName, objectName string) string {
        // 生成带签名的URL（过期时间7天）
        reqParams := make(url.Values)
        presignedURL, err := Client.PresignedGetObject(
                context.Background(),
                bucketName,
                objectName,
                7*24*time.Hour, // 过期时间
                reqParams,
        )
        if err != nil {
                logrus.Errorf("MinIO生成文件URL失败（桶：%s，对象：%s）: %v", bucketName, objectName, err)
                return ""
        }
        return presignedURL.String()
}
3. RabbitMQ 工具（src/pkg/rabbitmq/rabbitmq.go）
package rabbitmq

import (
        "context"
        "encoding/json"
        "mini-douyin/src/config"
        "fmt"

        amqp "github.com/rabbitmq/amqp091-go"
        "github.com/sirupsen/logrus"
)

// 全局RabbitMQ连接和信道
var (
        Conn *amqp.Connection
        Ch   *amqp.Channel
)

// InitRabbitMQ 初始化RabbitMQ连接和信道
func InitRabbitMQ() error {
        // 构建连接地址
        addr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
                config.RabbitMQUser,
                config.RabbitMQPassword,
                config.RabbitMQHost,
                config.RabbitMQPort,
                config.RabbitMQVHost,
        )

        // 建立连接
        conn, err := amqp.Dial(addr)
        if err != nil {
                return fmt.Errorf("连接RabbitMQ失败: %v", err)
        }
        Conn = conn

        // 创建信道（RabbitMQ推荐一个连接对应多个信道）
        ch, err := conn.Channel()
        if err != nil {
                return fmt.Errorf("创建RabbitMQ信道失败: %v", err)
        }
        Ch = ch

        // 声明通用交换机（topic类型，支持路由键匹配）
        if err := declareExchange("video.exchange"); err != nil {
                return err
        }
        if err := declareExchange("message.exchange"); err != nil {
                return err
        }

        logrus.Info("RabbitMQ初始化完成")
        return nil
}

// declareExchange 声明交换机（不存在则创建）
func declareExchange(exchangeName string) error {
        return Ch.ExchangeDeclare(
                exchangeName, // 交换机名称
                "topic",      // 交换机类型（topic支持通配符路由）
                true,         // 持久化（重启后不丢失）
                false,        // 自动删除（无绑定则删除）
                false,        // 排他性（仅当前连接可用）
                false,        // 非阻塞（立即返回）
                nil,          // 额外参数
        )
}

// Publish 发送消息到RabbitMQ
func Publish(exchangeName, routingKey string, data interface{}) error {
        // 序列化消息（JSON格式）
        body, err := json.Marshal(data)
        if err != nil {
                return fmt.Errorf("消息序列化失败: %v", err)
        }

        // 发送消息
        return Ch.Publish(
                exchangeName, // 交换机名称
                routingKey,   // 路由键（匹配队列绑定的键）
                false,        // 强制投递（交换机无法路由则返回错误）
                false,        // 立即投递（无消费者则返回错误）
                amqp.Publishing{
                        ContentType: "application/json", // 消息类型
                        Body:        body,               // 消息体
                        DeliveryMode: amqp.Persistent,   // 持久化消息（重启后不丢失）
                },
        )
}

// Consume 消费消息（业务代码中调用，如封面生成 worker）
func Consume(exchangeName, queueName, routingKey string) (<-chan amqp.Delivery, error) {
        // 声明队列（不存在则创建）
        q, err := Ch.QueueDeclare(
                queueName, // 队列名称
                true,      // 持久化
                false,     // 自动删除
                false,     // 排他性
                false,     // 非阻塞
                nil,
        )
        if err != nil {
                return nil, fmt.Errorf("声明队列失败: %v", err)
        }

        // 绑定队列到交换机
        if err := Ch.QueueBind(
                q.Name,        // 队列名称
                routingKey,    // 路由键（如"video.cover"）
                exchangeName,  // 交换机名称
                false,
                nil,
        ); err != nil {
                return nil, fmt.Errorf("队列绑定交换机失败: %v", err)
        }

        // 消费消息（手动确认模式）
        msgs, err := Ch.Consume(
                q.Name, // 队列名称
                "",     // 消费者标签（自定义）
                false,  // 自动确认（false=手动确认，确保消息处理完成）
                false,  // 排他性
                false,  // 不本地消费（同一连接的消息不消费）
                false,  // 非阻塞
                nil,
        )
        if err != nil {
                return nil, fmt.Errorf("消费消息失败: %v", err)
        }

        return msgs, nil
}

// Close 关闭RabbitMQ连接和信道
func Close() {
        if Ch != nil {
                Ch.Close()
        }
        if Conn != nil {
                Conn.Close()
        }
        logrus.Info("RabbitMQ连接已关闭")
}
4. Redis 工具（src/internal/storage/redis/redis.go）
package redis

import (
        "context"
        "mini-douyin/src/config"

        "github.com/redis/go-redis/v9"
        "github.com/sirupsen/logrus"
)

// 全局Redis客户端
var Client *redis.Client

// InitRedis 初始化Redis客户端
func InitRedis() error {
        // 创建客户端
        client := redis.NewClient(&redis.Options{
                Addr:     config.RedisAddr,     // 地址
                Password: config.RedisPassword, // 密码（无则空）
                DB:       config.RedisDB,       // 数据库编号
                PoolSize: 100,                  // 连接池大小（高并发可调整）
        })

        // 测试连接
        ctx := context.Background()
        if err := client.Ping(ctx).Err(); err != nil {
                return err
        }

        Client = client
        logrus.Info("Redis客户端初始化完成")
        return nil
}

// Close 关闭Redis连接
func Close() error {
        return Client.Close()
}
第 1 天：环境搭建与基础框架
目标
- 搭建开发环境与项目结构
- 实现基础服务注册与发现
- 完成网关初始化
技术栈
- Go 1.20、Gin（网关）、GRPC、Docker、Consul
1. 项目结构
mini-douyin/
├── docker/           # 基础镜像配置
├── scripts/          # 构建脚本
├── src/
│   ├── api/          # 网关API定义
│   ├── cmd/          # 服务入口
│   ├── config/       # 配置文件
│   ├── internal/     # 业务逻辑
│   │   ├── service/  # 微服务实现
│   │   ├── rpc/      # GRPC客户端
│   │   └── storage/  # 存储层
│   └── pkg/          # 公共工具
├── docker-compose.yml # 服务编排
└── go.mod
2. Docker Compose 配置（核心组件）
# docker-compose.yml
version: '3.8'
services:
  consul:
    image: consul:1.15
    ports:
      - "8500:8500"
    command: agent -dev -client=0.0.0.0

  postgres:
    image: postgres:14
    environment:
      POSTGRES_USER: douyin
      POSTGRES_PASSWORD: douyin
      POSTGRES_DB: douyin
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    command: redis-server --appendonly yes

  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      RABBITMQ_DEFAULT_USER: douyin
      RABBITMQ_DEFAULT_PASS: douyin

  jaeger:
    image: jaegertracing/all-in-one:1.47
    ports:
      - "6831:6831/udp"
      - "16686:16686"

volumes:
  pgdata:
3. 网关初始化（src/cmd/gateway/main.go）
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
4. 网关路由实现（src/api/router/router.go）
package router

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
        // 健康检查路由（第1天核心功能）
        r.GET("/health", func(c *gin.Context) {
                c.JSON(200, gin.H{"status": "ok"})
        })
}
5. 服务注册工具（src/pkg/consul/register.go）
package consul

import (
        "context"
        "fmt"
        "github.com/hashicorp/consul/api"
        "mini-douyin/src/config"
        "time"
)

// RegisterService 注册服务到Consul
func RegisterService(serviceName, addr string) error {
        client, err := api.NewClient(&api.Config{
                Address: config.ConsulAddr,
        })
        if err != nil {
                return err
        }

        registration := &api.AgentServiceRegistration{
                Name:    serviceName,
                ID:      fmt.Sprintf("%s-%s", serviceName, addr),
                Address: addr,
                Port:    config.RPCPort,
                Check: &api.AgentServiceCheck{
                        HTTP:                           fmt.Sprintf("http://%s:%d/health", addr, config.RPCPort),
                        Timeout:                        "5s",
                        Interval:                       "10s",
                        DeregisterCriticalServiceAfter: "30s",
                },
        }

        return client.Agent().ServiceRegister(registration)
}
6. main.go
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
测试
1. 环境测试：启动 Docker Compose，验证所有组件正常运行
docker-compose up -d
# 检查Consul服务列表（http://localhost:8500）
2. 网关测试：运行网关服务，使用 curl 测试基础路由
curl http://localhost:8080/health
# 预期返回：{"status":"ok"}
3. 服务注册测试：（src/cmd/consul/test_register.go）
package main

import (
        "mini-douyin/src/config"
        "mini-douyin/src/pkg/consul"
        "mini-douyin/src/utils/logging"
        "time"
)

func main() {
        // 初始化配置和日志
        config.InitConfig()
        logging.InitLog()

        // 注册测试服务到Consul
        serviceName := "test-service"
        serviceAddr := "localhost"
        if err := consul.RegisterService(serviceName, serviceAddr); err != nil {
                logging.Logger.Fatalf("服务注册失败: %v", err)
        }

        logging.Logger.Infof("服务%s（地址%s）注册到Consul成功", serviceName, serviceAddr)
        // 保持进程运行，观察Consul UI（http://localhost:8500）
        time.Sleep(30 * time.Minute)
}
第 2 天：用户服务开发
目标
- 实现用户注册 / 登录功能
- 集成 JWT 认证
- 完成用户信息存储与缓存
技术栈
- GORM（ORM）、JWT、Redis（会话存储）
1.  数据库初始化（src/internal/storage/postgres/postgres.go）
package postgres

import (
        "context"
        "mini-douyin/src/config"
        "mini-douyin/src/internal/storage/model"
        "fmt"

        "gorm.io/driver/postgres"
        "gorm.io/gorm"
        "gorm.io/gorm/logger"
)

// 全局数据库客户端
var DB *gorm.DB

// InitPostgres 初始化PostgreSQL连接
func InitPostgres() error {
        // 构建DSN（数据源名称）
        dsn := fmt.Sprintf(
                "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
                config.PostgresHost,
                config.PostgresPort,
                config.PostgresUser,
                config.PostgresPassword,
                config.PostgresDB,
        )

        // 打开数据库连接
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
                Logger: logger.Default.LogMode(logger.Info), // 开发环境显示SQL日志
        })
        if err != nil {
                return fmt.Errorf("PostgreSQL连接失败: %v", err)
        }

        // 初始化数据库连接池（可选，优化性能）
        sqlDB, err := db.DB()
        if err != nil {
                return fmt.Errorf("获取SQL DB失败: %v", err)
        }
        sqlDB.SetMaxOpenConns(100)    // 最大打开连接数
        sqlDB.SetMaxIdleConns(20)     // 最大空闲连接数
        sqlDB.SetConnMaxLifetime(1*60) // 连接最大存活时间（分钟）

        DB = db
        return nil
}

// Migrate 数据库迁移（创建表）
func Migrate() error {
        // 迁移所有模型
        return DB.AutoMigrate(
                &model.User{},
                &model.Video{},
                &model.Relation{},
                &model.Message{},
                &model.Comment{},
                &model.Favorite{},
        )
}

// Close 关闭数据库连接
func Close() error {
        sqlDB, err := DB.DB()
        if err != nil {
                return err
        }
        return sqlDB.Close()
}
2. 用户 API Handler（src/api/handler/user_handler.go）
package handler

import (
        "context"
        "mini-douyin/src/api/proto/user"
        "mini-douyin/src/config"
        "mini-douyin/src/internal/rpc"
        "mini-douyin/src/internal/storage/redis"
        "net/http"
        "strconv"

        "github.com/gin-gonic/gin"
        "github.com/golang-jwt/jwt/v5"
        "github.com/sirupsen/logrus"
)

// UserRegister 注册接口Handler
func UserRegister(c *gin.Context) {
        // 1. 绑定请求参数
        var req user.RegisterRequest
        if err := c.ShouldBindJSON(&req); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "status_code": 1,
                        "status_msg":  "参数错误: " + err.Error(),
                })
                return
        }

        // 2. 调用用户服务RPC接口
        resp, err := rpc.UserClient.Register(context.Background(), &req)
        if err != nil {
                logrus.Errorf("用户注册RPC调用失败: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{
                        "status_code": 1,
                        "status_msg":  "注册失败: 服务异常",
                })
                return
        }

        // 3. 返回响应
        c.JSON(http.StatusOK, resp)
}

// UserLogin 登录接口Handler
func UserLogin(c *gin.Context) {
        // 1. 绑定请求参数
        var req user.LoginRequest
        if err := c.ShouldBindJSON(&req); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "status_code": 1,
                        "status_msg":  "参数错误: " + err.Error(),
                })
                return
        }

        // 2. 调用用户服务RPC接口
        resp, err := rpc.UserClient.Login(context.Background(), &req)
        if err != nil {
                logrus.Errorf("用户登录RPC调用失败: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{
                        "status_code": 1,
                        "status_msg":  "登录失败: 服务异常",
                })
                return
        }

        // 3. 返回响应
        c.JSON(http.StatusOK, resp)
}

// UserGetInfo 获取用户信息接口Handler
func UserGetInfo(c *gin.Context) {
        // 1. 获取Token并解析用户ID
        tokenStr := c.Query("token")
        if tokenStr == "" {
                c.JSON(http.StatusUnauthorized, gin.H{
                        "status_code": 1,
                        "status_msg":  "未登录: 缺少Token",
                })
                return
        }

        // 解析JWT Token
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
                // 验证签名算法
                if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("不支持的签名算法: %v", t.Header["alg"])
                }
                return []byte(config.JWTSecret), nil
        })
        if err != nil || !token.Valid {
                c.JSON(http.StatusUnauthorized, gin.H{
                        "status_code": 1,
                        "status_msg":  "未登录: Token无效或过期",
                })
                return
        }

        // 提取用户ID
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
                c.JSON(http.StatusUnauthorized, gin.H{
                        "status_code": 1,
                        "status_msg":  "未登录: Token解析失败",
                })
                return
        }
        userID := int64(claims["user_id"].(float64))

        // 2. 先查Redis缓存
        userInfo, err := redis.GetUserInfo(context.Background(), userID)
        if err == nil && userInfo != nil {
                // 缓存命中，直接返回
                c.JSON(http.StatusOK, gin.H{
                        "status_code": 0,
                        "user":        userInfo,
                })
                return
        }

        // 3. 缓存未命中，查数据库（调用用户服务）
        resp, err := rpc.UserClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{
                UserId: userID,
        })
        if err != nil {
                logrus.Errorf("获取用户信息RPC调用失败: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{
                        "status_code": 1,
                        "status_msg":  "获取用户信息失败",
                })
                return
        }

        // 4. 同步到Redis缓存
        redis.SetUserInfo(context.Background(), userID, resp.User)

        // 5. 返回响应
        c.JSON(http.StatusOK, resp)
}
3. 用户服务 RPC 客户端（src/internal/rpc/user_rpc.go）
package rpc

import (
        "mini-douyin/src/api/proto/user"
        "mini-douyin/src/config"
        "mini-douyin/src/pkg/consul"
        "time"

        "github.com/sirupsen/logrus"
        "google.golang.org/grpc"
        "google.golang.org/grpc/credentials/insecure"
)

// UserClient 用户服务RPC客户端
var UserClient user.UserServiceClient

// InitUserRPC 初始化用户服务RPC客户端（从Consul发现服务）
func InitUserRPC() error {
        // 从Consul获取用户服务地址
        serviceAddr, err := consul.DiscoverService("user-service")
        if err != nil {
                return err
        }

        // 连接GRPC服务（开发环境用非安全连接，生产环境需配置TLS）
        conn, err := grpc.Dial(
                serviceAddr,
                grpc.WithTransportCredentials(insecure.NewCredentials()),
                grpc.WithConnectParams(grpc.ConnectParams{
                        Backoff: grpc.DefaultBackoffConfig,
                        Timeout: 5 * time.Second, // 连接超时
                }),
        )
        if err != nil {
                return fmt.Errorf("连接用户服务GRPC失败: %v", err)
        }

        // 创建客户端
        UserClient = user.NewUserServiceClient(conn)
        logrus.Info("用户服务GRPC客户端初始化完成")
        return nil
}
4. Consul 服务发现补充（src/pkg/consul/discover.go）
package consul

import (
        "mini-douyin/src/config"
        "fmt"

        "github.com/hashicorp/consul/api"
        "github.com/sirupsen/logrus"
)

// DiscoverService 从Consul发现服务（返回第一个健康的服务地址）
func DiscoverService(serviceName string) (string, error) {
        // 创建Consul客户端
        client, err := api.NewClient(&api.Config{
                Address: config.ConsulAddr,
        })
        if err != nil {
                return "", err
        }

        // 查找服务（只返回健康的服务）
        services, _, err := client.Health().Service(
                serviceName,
                "",
                true,  // 只返回健康服务
                nil,
        )
        if err != nil {
                return "", err
        }

        // 检查服务是否存在
        if len(services) == 0 {
                return "", fmt.Errorf("未找到服务%s的健康实例", serviceName)
        }

        // 返回第一个服务地址（格式：ip:port）
        service := services[0].Service
        addr := fmt.Sprintf("%s:%d", service.Address, service.Port)
        logrus.Infof("发现服务%s的实例: %s", serviceName, addr)
        return addr, nil
}
5. 用户模型（src/internal/storage/model/user.go）
package model

import (
        "time"
        "gorm.io/gorm"
)

type User struct {
        ID              int64          `gorm:"primaryKey" json:"id"`
        Username        string         `gorm:"uniqueIndex;size:32" json:"username"`
        PasswordHash    string         `gorm:"size:128" json:"-"`
        Avatar          string         `json:"avatar"`
        BackgroundImage string         `json:"background_image"`
        Signature       string         `json:"signature"`
        CreatedAt       time.Time      `json:"-"`
        UpdatedAt       time.Time      `json:"-"`
        DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
6. 用户服务 GRPC 定义（src/api/proto/user.proto）
syntax = "proto3";
package user;

option go_package = "./user";

message RegisterRequest {
  string username = 1;
  string password = 2;
}

message RegisterResponse {
  int32 status_code = 1;
  string status_msg = 2;
  int64 user_id = 3;
  string token = 4;
}

message LoginRequest {
  string username = 1;
  string password = 2;
}

message LoginResponse {
  int32 status_code = 1;
  string status_msg = 2;
  int64 user_id = 3;
  string token = 4;
}

service UserService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
  rpc Login(LoginRequest) returns (LoginResponse);
}
7. 用户服务实现（src/internal/service/user/service.go）
package user

import (
        "context"
        "crypto/sha256"
        "encoding/hex"
        "mini-douyin/src/api/proto/user"
        "mini-douyin/src/config"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/redis"
        "time"

        "github.com/golang-jwt/jwt/v5"
        "gorm.io/gorm"
)

type UserService struct{}

// 生成密码哈希
func hashPassword(password string) string {
        h := sha256.New()
        h.Write([]byte(password + config.PasswordSalt))
        return hex.EncodeToString(h.Sum(nil))
}

// 生成JWT令牌
func generateToken(userID int64) (string, error) {
        claims := jwt.MapClaims{
                "user_id": userID,
                "exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
        }
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        return token.SignedString([]byte(config.JWTSecret))
}

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
        // 检查用户名是否已存在
        var exist model.User
        if err := postgres.DB.WithContext(ctx).Where("username = ?", req.Username).First(&exist).Error; err != gorm.ErrRecordNotFound {
                return &user.RegisterResponse{
                        StatusCode: 1,
                        StatusMsg:  "用户名已存在",
                }, nil
        }

        // 创建用户
        user := model.User{
                Username:     req.Username,
                PasswordHash: hashPassword(req.Password),
        }
        if err := postgres.DB.WithContext(ctx).Create(&user).Error; err != nil {
                return nil, err
        }

        // 生成令牌
        token, err := generateToken(user.ID)
        if err != nil {
                return nil, err
        }

        // 缓存用户信息到Redis
        redis.SetUserInfo(ctx, user.ID, &user)

        return &user.RegisterResponse{
                StatusCode: 0,
                StatusMsg:  "注册成功",
                UserId:     user.ID,
                Token:      token,
        }, nil
}

func (s *UserService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
        var user model.User
        if err := postgres.DB.WithContext(ctx).Where("username = ?", req.Username).First(&user).Error; err != nil {
                return &user.LoginResponse{
                        StatusCode: 1,
                        StatusMsg:  "用户不存在",
                }, nil
        }

        // 验证密码
        if user.PasswordHash != hashPassword(req.Password) {
                return &user.LoginResponse{
                        StatusCode: 1,
                        StatusMsg:  "密码错误",
                }, nil
        }

        // 生成令牌
        token, err := generateToken(user.ID)
        if err != nil {
                return nil, err
        }

        return &user.LoginResponse{
                StatusCode: 0,
                StatusMsg:  "登录成功",
                UserId:     user.ID,
                Token:      token,
        }, nil
}
8. Redis 缓存工具（src/internal/storage/redis/user.go）
package redis

import (
        "context"
        "encoding/json"
        "mini-douyin/src/internal/storage/model"
        "time"

        "github.com/redis/go-redis/v9"
)

// SetUserInfo 缓存用户信息（过期时间24小时）
func SetUserInfo(ctx context.Context, userID int64, user *model.User) error {
        key := fmt.Sprintf("user:info:%d", userID)
        val, err := json.Marshal(user)
        if err != nil {
                return err
        }
        return Client.Set(ctx, key, val, 24*time.Hour).Err()
}

// GetUserInfo 获取缓存的用户信息
func GetUserInfo(ctx context.Context, userID int64) (*model.User, error) {
        key := fmt.Sprintf("user:info:%d", userID)
        val, err := Client.Get(ctx, key).Bytes()
        if err != nil {
                return nil, err
        }
        var user model.User
        if err := json.Unmarshal(val, &user); err != nil {
                return nil, err
        }
        return &user, nil
}
9. main.go
package main

import (
        "context"
        "mini-douyin/src/config"
        "mini-douyin/src/api/router"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/redis"
        "mini-douyin/src/pkg/consul"
        "mini-douyin/src/pkg/tracing"
        "mini-douyin/src/internal/rpc"
        "net"
        "os"
        "os/signal"
        "sync"
        "syscall"
        "time"

        "github.com/gin-contrib/gzip"
        "github.com/gin-gonic/gin"
        "github.com/sirupsen/logrus"
        ginprometheus "github.com/zsais/go-gin-prometheus"
        "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
        "google.golang.org/grpc"
)

func main() {
        // 1. 初始化配置
        config.InitConfig()

        // 2. 初始化链路追踪
        tp, err := tracing.InitTracer("user-system")
        if err != nil {
                logrus.Fatalf("初始化追踪失败: %v", err)
        }
        defer tp.Shutdown(context.Background())

        // 3. 初始化数据库和缓存
        if err := postgres.InitPostgres(); err != nil {
                logrus.Fatalf("PostgreSQL初始化失败: %v", err)
        }
        defer postgres.Close()
        if err := redis.InitRedis(); err != nil {
                logrus.Fatalf("Redis初始化失败: %v", err)
        }
        defer redis.Close()

        // 4. 数据库迁移（创建用户表）
        if err := postgres.Migrate(); err != nil {
                logrus.Fatalf("数据库迁移失败: %v", err)
        }

        // 5. 启动用户GRPC服务（并发启动）
        var wg sync.WaitGroup
        wg.Add(2) // 网关 + 用户服务

        // 5.1 启动用户服务
        go func() {
                defer wg.Done()
                // 注册服务到Consul
                userServiceName := "user-service"
                userAddr := "localhost"
                if err := consul.RegisterService(userServiceName, userAddr); err != nil {
                        logrus.Fatalf("用户服务注册失败: %v", err)
                }

                // 启动GRPC监听
                lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.RPCPort))
                if err != nil {
                        logrus.Fatalf("用户服务监听失败: %v", err)
                }
                grpcServer := grpc.NewServer()
                // 注册用户服务实现（第2天核心：绑定UserService）
                user.RegisterUserServiceServer(grpcServer, &user.UserService{})

                logrus.Infof("用户服务启动，监听端口: %d", config.RPCPort)
                if err := grpcServer.Serve(lis); err != nil {
                        logrus.Fatalf("用户服务启动失败: %v", err)
                }
        }()

        // 5.2 启动网关服务（依赖用户服务RPC客户端）
        go func() {
                defer wg.Done()
                // 初始化用户服务RPC客户端
                if err := rpc.InitUserRPC(); err != nil {
                        logrus.Fatalf("用户RPC客户端初始化失败: %v", err)
                }

                // 启动HTTP网关
                r := gin.Default()
                r.Use(otelgin.Middleware("gateway"))
                r.Use(gzip.Gzip(gzip.DefaultCompression))
                p := ginprometheus.NewPrometheus("douyin_gateway")
                p.Use(r)

                // 注册路由（包含用户注册/登录）
                router.RegisterRoutes(r)

                logrus.Infof("网关服务启动，监听地址: %s", config.GatewayAddr)
                if err := r.Run(config.GatewayAddr); err != nil {
                        logrus.Fatalf("网关启动失败: %v", err)
                }
        }()

        // 6. 优雅关闭
        quit := make(chan os.Signal, 1)
        signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
        <-quit
        logrus.Info("开始关闭服务...")

        wg.Wait()
        logrus.Info("所有服务已关闭")
}
10. 网关路由
package router

import (
        "mini-douyin/src/api/handler"
        "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
        // 健康检查
        r.GET("/health", func(c *gin.Context) {
                c.JSON(200, gin.H{"status": "ok"})
        })

        // 用户相关路由
        douyin := r.Group("/douyin")
        {
                user := douyin.Group("/user")
                {
                        user.POST("/register", handler.UserRegister) // 注册接口
                        user.POST("/login", handler.UserLogin)       // 登录接口
                }
        }
}
测试
1. 单元测试（用户服务逻辑）
// src/internal/service/user/service_test.go
package user

import (
        "context"
        "testing"
        "mini-douyin/src/api/proto/user"
        "github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
        s := &UserService{}
        ctx := context.Background()

        // 测试注册新用户
        req := &user.RegisterRequest{Username: "test", Password: "123456"}
        resp, err := s.Register(ctx, req)
        assert.NoError(t, err)
        assert.Equal(t, int32(0), resp.StatusCode)
        assert.NotZero(t, resp.UserId)

        // 测试重复注册
        resp, err = s.Register(ctx, req)
        assert.NoError(t, err)
        assert.Equal(t, int32(1), resp.StatusCode)
}
2. API 测试（使用 Postman）
- 注册接口：POST http://localhost:8080/douyin/user/register/
  - 请求体：{"username":"test","password":"123456"}
  - 预期响应：{"status_code":0,"user_id":1,"token":"xxx"}
第 3 天：视频服务开发（含异步处理）
目标
- 实现视频上传与存储
- 异步生成视频封面（RabbitMQ）
- 实现视频列表查询
技术栈
- MinIO（对象存储）、RabbitMQ（异步任务）、FFmpeg（视频处理）
1. 视频服务 GRPC 定义（src/api/proto/video.proto）
syntax = "proto3";
package video;

option go_package = "./video";

import "user.proto"; // 导入用户服务的proto，用于关联作者信息

// 创建视频请求
message CreateVideoRequest {
  int64 author_id = 1;  // 作者ID
  string play_url = 2;  // 视频播放URL
  string cover_url = 3; // 封面URL（初始可空，异步生成后更新）
  string title = 4;     // 视频标题
}

// 创建视频响应
message CreateVideoResponse {
  int32 status_code = 1;
  string status_msg = 2;
  int64 video_id = 3; // 生成的视频ID
}

// 获取视频流请求（可选：最后一次获取的时间戳，用于分页）
message GetVideoFeedRequest {
  int64 latest_time = 1; // 毫秒级时间戳，默认0（获取最新视频）
  int32 limit = 2;       // 一次获取数量，默认20
}

// 获取视频流响应
message GetVideoFeedResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated VideoInfo video_list = 3; // 视频列表
  int64 next_time = 4;               // 下次请求的latest_time（用于分页）
}

// 获取用户发布的视频列表请求
message GetUserPublishRequest {
  int64 user_id = 1; // 用户ID
}

// 获取用户发布的视频列表响应
message GetUserPublishResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated VideoInfo video_list = 3; // 视频列表
}

// 视频信息结构体
message VideoInfo {
  int64 id = 1;                // 视频ID
  user.UserInfo author = 2;    // 作者信息（关联用户服务）
  string play_url = 3;         // 播放URL
  string cover_url = 4;        // 封面URL
  int64 favorite_count = 5;    // 点赞数
  int64 comment_count = 6;     // 评论数
  bool is_favorite = 7;        // 当前用户是否点赞（需登录）
  string title = 8;            // 标题
}

// 视频服务接口
service VideoService {
  rpc CreateVideo(CreateVideoRequest) returns (CreateVideoResponse);
  rpc GetVideoFeed(GetVideoFeedRequest) returns (GetVideoFeedResponse);
  rpc GetUserPublish(GetUserPublishRequest) returns (GetUserPublishResponse);
  rpc UpdateVideoCover(UpdateVideoCoverRequest) returns (UpdateVideoCoverResponse); // 用于异步更新封面
}

// 更新视频封面请求（内部接口，由视频处理服务调用）
message UpdateVideoCoverRequest {
  int64 video_id = 1;   // 视频ID
  string cover_url = 2; // 新封面URL
}

// 更新视频封面响应
message UpdateVideoCoverResponse {
  int32 status_code = 1;
  string status_msg = 2;
}
2. 视频处理 Worker 补充（src/cmd/worker/cover_worker.go）
package main

import (
        "context"
        "encoding/json"
        "mini-douyin/src/config"
        "mini-douyin/src/internal/rpc"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/pkg/minio"
        "mini-douyin/src/pkg/rabbitmq"
        "mini-douyin/src/utils/logging"
        "os"
        "os/exec"
        "path/filepath"

        amqp "github.com/rabbitmq/amqp091-go"
        "github.com/sirupsen/logrus"
)

// 消息体结构（与发送端对应）
type VideoCoverMsg struct {
        VideoID  string `json:"video_id"`  // 视频ID（UUID）
        VideoURL string `json:"video_url"` // 视频MinIO URL
}

func main() {
        // 1. 初始化基础组件
        config.InitConfig()
        logging.InitLog()
        if err := postgres.InitPostgres(); err != nil {
                logrus.Fatalf("PostgreSQL初始化失败: %v", err)
        }
        if err := rabbitmq.InitRabbitMQ(); err != nil {
                logrus.Fatalf("RabbitMQ初始化失败: %v", err)
        }
        defer rabbitmq.Close()
        if err := minio.InitMinIO(); err != nil {
                logrus.Fatalf("MinIO初始化失败: %v", err)
        }
        if err := rpc.InitVideoRPC(); err != nil {
                logrus.Fatalf("视频服务RPC客户端初始化失败: %v", err)
        }

        // 2. 消费RabbitMQ消息（视频封面生成队列）
        msgs, err := rabbitmq.Consume(
                "video.exchange", // 交换机名称
                "video.cover",    // 队列名称
                "video.cover",    // 路由键
        )
        if err != nil {
                logrus.Fatalf("消费消息失败: %v", err)
        }

        logrus.Info("视频封面生成Worker启动成功，等待消息...")

        // 3. 处理消息
        for msg := range msgs {
                var data VideoCoverMsg
                if err := json.Unmarshal(msg.Body, &data); err != nil {
                        logrus.Errorf("消息解析失败（消息体：%s）: %v", string(msg.Body), err)
                        msg.Nack(false, false) // 拒绝消息，不重试（格式错误无法修复）
                        continue
                }

                // 生成封面
                coverURL, err := generateCover(data.VideoURL, data.VideoID)
                if err != nil {
                        logrus.Errorf("生成视频封面失败（视频ID：%s）: %v", data.VideoID, err)
                        msg.Nack(false, true) // 拒绝消息，重试（可能是临时错误）
                        continue
                }

                // 调用视频服务更新封面URL
                _, err = rpc.VideoClient.UpdateVideoCover(context.Background(), &rpc.VideoUpdateCoverRequest{
                        VideoId:  data.VideoID,
                        CoverUrl: coverURL,
                })
                if err != nil {
                        logrus.Errorf("更新视频封面URL失败（视频ID：%s）: %v", data.VideoID, err)
                        msg.Nack(false, true) // 重试
                        continue
                }

                logrus.Infof("视频封面生成成功（视频ID：%s，封面URL：%s）", data.VideoID, coverURL)
                msg.Ack(false) // 确认消息处理完成
        }
}

// generateCover 调用FFmpeg生成视频封面
func generateCover(videoURL, videoID string) (string, error) {
        // 1. 创建临时目录（存放下载的视频和生成的封面）
        tmpDir := filepath.Join("/tmp", "mini-douyin", "covers")
        if err := os.MkdirAll(tmpDir, 0755); err != nil {
                return "", fmt.Errorf("创建临时目录失败: %v", err)
        }

        // 2. 下载视频到临时文件
        tmpVideoPath := filepath.Join(tmpDir, fmt.Sprintf("%s.mp4", videoID))
        cmd := exec.Command("curl", "-s", "-o", tmpVideoPath, videoURL) // -s：静默模式
        if err := cmd.Run(); err != nil {
                return "", fmt.Errorf("下载视频失败: %v", err)
        }
        defer os.Remove(tmpVideoPath) // 处理完成后删除临时视频

        // 3. 调用FFmpeg生成封面（取第1秒的帧）
        tmpCoverPath := filepath.Join(tmpDir, fmt.Sprintf("%s.jpg", videoID))
        // FFmpeg参数：-i 输入文件 -ss 时间点 -vframes 取1帧 -q:v 2 画质（1-31，越小越好）
        cmd = exec.Command(
                "ffmpeg",
                "-i", tmpVideoPath,
                "-ss", "1",          // 从第1秒开始
                "-vframes", "1",      // 只取1帧
                "-q:v", "2",          // 画质（2为高质量）
                "-y",                 // 覆盖已存在文件
                tmpCoverPath,
        )
        // 捕获FFmpeg输出（方便调试）
        output, err := cmd.CombinedOutput()
        if err != nil {
                return "", fmt.Errorf("FFmpeg执行失败（输出：%s）: %v", string(output), err)
        }
        defer os.Remove(tmpCoverPath) // 处理完成后删除临时封面

        // 4. 上传封面到MinIO
        coverFile, err := os.Open(tmpCoverPath)
        if err != nil {
                return "", fmt.Errorf("打开封面文件失败: %v", err)
        }
        defer coverFile.Close()

        // 获取文件大小
        fileInfo, err := coverFile.Stat()
        if err != nil {
                return "", fmt.Errorf("获取封面文件大小失败: %v", err)
        }

        // MinIO对象名称（格式：covers/{videoID}.jpg）
        objectName := filepath.Join("covers", fmt.Sprintf("%s.jpg", videoID))
        if err := minio.UploadFile(
                context.Background(),
                minio.ImageBucket,
                objectName,
                coverFile,
                fileInfo.Size(),
        ); err != nil {
                return "", fmt.Errorf("上传封面到MinIO失败: %v", err)
        }

        // 5. 返回封面URL
        return minio.GetFileURL(minio.ImageBucket, objectName), nil
}
3. FFmpeg 安装（Ubuntu）
# 安装FFmpeg（视频处理依赖）
sudo apt update
sudo apt install -y ffmpeg
# 验证安装
ffmpeg -version # 应输出FFmpeg版本信息
4. 视频模型（src/internal/storage/model/video.go）
package model

import (
        "time"
        "gorm.io/gorm"
)

type Video struct {
        ID            int64          `gorm:"primaryKey" json:"id"`
        AuthorID      int64          `gorm:"index" json:"author_id"`
        PlayURL       string         `json:"play_url"`
        CoverURL      string         `json:"cover_url"`
        Title         string         `json:"title"`
        FavoriteCount int64          `json:"favorite_count"`
        CommentCount  int64          `json:"comment_count"`
        CreatedAt     time.Time      `json:"-"`
        UpdatedAt     time.Time      `json:"-"`
        DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
5. 视频上传接口（网关层）
// src/api/handler/video.go
package handler

import (
        "context"
        "mini-douyin/src/api/proto/video"
        "mini-douyin/src/internal/rpc"
        "mini-douyin/src/pkg/minio"
        "mini-douyin/src/pkg/rabbitmq"
        "net/http"
        "path/filepath"

        "github.com/gin-gonic/gin"
        "github.com/google/uuid"
)

func PublishAction(c *gin.Context) {
        // 获取用户ID（从JWT中间件）
        userID, _ := c.Get("user_id")
        
        // 接收视频文件
        file, err := c.FormFile("data")
        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"status_code": 1, "status_msg": "文件上传失败"})
                return
        }

        // 生成唯一文件名
        ext := filepath.Ext(file.Filename)
        videoID := uuid.New().String()
        videoName := fmt.Sprintf("videos/%s%s", videoID, ext)

        // 上传到MinIO
        src, err := file.Open()
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "文件打开失败"})
                return
        }
        defer src.Close()

        if err := minio.UploadFile(context.Background(), minio.VideoBucket, videoName, src, file.Size); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "文件存储失败"})
                return
        }

        // 发送消息到RabbitMQ，触发封面生成
        msg := map[string]interface{}{
                "video_id":  videoID,
                "video_url": minio.GetFileURL(minio.VideoBucket, videoName),
        }
        if err := rabbitmq.Publish("video.exchange", "video.cover", msg); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "消息发送失败"})
                return
        }

        // 保存视频信息到数据库
        resp, err := rpc.VideoClient.CreateVideo(c, &video.CreateVideoRequest{
                AuthorId: userID.(int64),
                PlayUrl:  minio.GetFileURL(minio.VideoBucket, videoName),
                Title:    c.PostForm("title"),
        })
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status_code": 1, "status_msg": "视频保存失败"})
                return
        }

        c.JSON(http.StatusOK, resp)
}
6. RabbitMQ 消费者（封面生成）
// src/cmd/worker/cover_worker.go
package main

import (
        "context"
        "encoding/json"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/pkg/minio"
        "mini-douyin/src/pkg/rabbitmq"
        "os/exec"
        "strings"

        amqp "github.com/rabbitmq/amqp091-go"
        "github.com/sirupsen/logrus"
)

func main() {
        // 连接RabbitMQ
        conn, ch, err := rabbitmq.Connect()
        if err != nil {
                logrus.Fatalf("RabbitMQ连接失败: %v", err)
        }
        defer conn.Close()
        defer ch.Close()

        // 声明队列
        q, err := ch.QueueDeclare(
                "video.cover", // 队列名
                true,          // 持久化
                false,         // 自动删除
                false,         // 排他性
                false,         // 非阻塞
                nil,
        )
        if err != nil {
                logrus.Fatalf("队列声明失败: %v", err)
        }

        // 绑定交换机
        if err := ch.QueueBind(
                q.Name,
                "video.cover", // 路由键
                "video.exchange",
                false,
                nil,
        ); err != nil {
                logrus.Fatalf("队列绑定失败: %v", err)
        }

        // 消费消息
        msgs, err := ch.Consume(
                q.Name,
                "",
                false, // 手动确认
                false,
                false,
                false,
                nil,
        )
        if err != nil {
                logrus.Fatalf("消息消费失败: %v", err)
        }

        // 处理消息
        for msg := range msgs {
                var data map[string]interface{}
                if err := json.Unmarshal(msg.Body, &data); err != nil {
                        logrus.Errorf("消息解析失败: %v", err)
                        msg.Nack(false, false) // 拒绝消息
                        continue
                }

                // 生成封面
                videoID := data["video_id"].(string)
                videoURL := data["video_url"].(string)
                coverURL, err := generateCover(videoURL, videoID)
                if err != nil {
                        logrus.Errorf("封面生成失败: %v", err)
                        msg.Nack(false, true) // 重试
                        continue
                }

                // 更新视频封面URL
                if err := postgres.DB.Model(&model.Video{}).
                        Where("id = ?", videoID).
                        Update("cover_url", coverURL).Error; err != nil {
                        logrus.Errorf("封面URL更新失败: %v", err)
                        msg.Nack(false, true)
                        continue
                }

                msg.Ack(false) // 确认消息处理完成
        }
}

// 调用FFmpeg生成封面
func generateCover(videoURL, videoID string) (string, error) {
        // 下载视频到临时文件
        tmpVideo := fmt.Sprintf("/tmp/%s.mp4", videoID)
        cmd := exec.Command("curl", "-o", tmpVideo, videoURL)
        if err := cmd.Run(); err != nil {
                return "", err
        }

        // 生成封面（第1秒帧）
        tmpCover := fmt.Sprintf("/tmp/%s.jpg", videoID)
        cmd = exec.Command("ffmpeg", "-i", tmpVideo, "-ss", "1", "-vframes", "1", tmpCover)
        if err := cmd.Run(); err != nil {
                return "", err
        }

        // 上传封面到MinIO
        coverName := fmt.Sprintf("covers/%s.jpg", videoID)
        file, err := os.Open(tmpCover)
        if err != nil {
                return "", err
        }
        defer file.Close()

        if err := minio.UploadFile(context.Background(), minio.ImageBucket, coverName, file, 0); err != nil {
                return "", err
        }

        // 清理临时文件
        os.Remove(tmpVideo)
        os.Remove(tmpCover)

        return minio.GetFileURL(minio.ImageBucket, coverName), nil
}
测试
1. 异步任务测试
  - 手动发送消息到 RabbitMQ，验证封面生成逻辑
# 使用RabbitMQ管理界面发送消息：
# 交换机：video.exchange，路由键：video.cover，消息体：{"video_id":"test","video_url":"http://xxx.mp4"}    
2. 压力测试（使用 k6）
// video_upload_test.js
import http from 'k6/http';
import { check, sleep } from 'k6';

export default function() {
  const url = 'http://localhost:8080/douyin/publish/action/';
  const formData = {
    data: http.file('./test.mp4', 'test.mp4'),
    title: 'test video',
  };
  const params = {
    headers: { 'Authorization': 'Bearer xxx' },
  };
  const res = http.post(url, formData, params);
  check(res, { 'status is 200': (r) => r.status === 200 });
  sleep(1);
}

// 运行测试：k6 run --vus 10 --duration 30s video_upload_test.js
第 4 天：社交服务开发（关注 / 粉丝）
目标
- 实现关注 / 取消关注功能
- 粉丝列表 / 关注列表查询
- 高并发场景下的缓存优化
技术栈
- Redis（关系缓存）、PostgreSQL（关系存储）
1. 关注关系模型
// src/internal/storage/model/relation.go
package model

import (
        "time"
        "gorm.io/gorm"
)

type Relation struct {
        ID         int64          `gorm:"primaryKey" json:"id"`
        UserID     int64          `gorm:"index:idx_user" json:"user_id"`     // 关注者
        FollowID   int64          `gorm:"index:idx_follow" json:"follow_id"` // 被关注者
        IsFollow   bool           `json:"is_follow"`
        CreatedAt  time.Time      `json:"-"`
        UpdatedAt  time.Time      `json:"-"`
        DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
2. 关注服务实现
// src/internal/service/relation/service.go
package relation

import (
        "context"
        "mini-douyin/src/api/proto/relation"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/redis"

        "github.com/redis/go-redis/v9"
)

type RelationService struct{}

// 关注用户
func (s *RelationService) Follow(ctx context.Context, req *relation.FollowRequest) (*relation.FollowResponse, error) {
        // 开启事务
        tx := postgres.DB.WithContext(ctx).Begin()
        defer func() {
                if r := recover(); r != nil {
                        tx.Rollback()
                }
        }()

        // 检查是否已关注
        var rel model.Relation
        if err := tx.Where("user_id = ? AND follow_id = ?", req.UserId, req.FollowId).First(&rel).Error; err == nil {
                // 已关注则更新状态
                rel.IsFollow = true
                if err := tx.Save(&rel).Error; err != nil {
                        tx.Rollback()
                        return nil, err
                }
        } else {
                // 未关注则创建关系
                rel = model.Relation{
                        UserID:   req.UserId,
                        FollowID: req.FollowId,
                        IsFollow: true,
                }
                if err := tx.Create(&rel).Error; err != nil {
                        tx.Rollback()
                        return nil, err
                }
        }
        tx.Commit()

        // 更新Redis缓存（关注列表）
        if err := redis.Client.SAdd(ctx, redis.FollowKey(req.UserId), req.FollowId).Err(); err != nil {
                return nil, err
        }
        // 更新粉丝列表
        if err := redis.Client.SAdd(ctx, redis.FollowerKey(req.FollowId), req.UserId).Err(); err != nil {
                return nil, err
        }

        return &relation.FollowResponse{StatusCode: 0}, nil
}

// 获取关注列表
func (s *RelationService) GetFollowList(ctx context.Context, req *relation.FollowListRequest) (*relation.FollowListResponse, error) {
        // 先查Redis缓存
        followIDs, err := redis.Client.SMembers(ctx, redis.FollowKey(req.UserId)).Result()
        if err != nil && err != redis.Nil {
                return nil, err
        }

        // 缓存未命中则查数据库
        if len(followIDs) == 0 || err == redis.Nil {
                var rels []model.Relation
                if err := postgres.DB.WithContext(ctx).Where("user_id = ? AND is_follow = true", req.UserId).Find(&rels).Error; err != nil {
                        return nil, err
                }
                // 同步到Redis
                for _, rel := range rels {
                        followIDs = append(followIDs, fmt.Sprintf("%d", rel.FollowID))
                }
                if len(followIDs) > 0 {
                        redis.Client.SAdd(ctx, redis.FollowKey(req.UserId), followIDs)
                        redis.Client.Expire(ctx, redis.FollowKey(req.UserId), 24*time.Hour)
                }
        }

        // 转换为用户列表（此处省略调用用户服务获取详情）
        return &relation.FollowListResponse{
                StatusCode: 0,
                UserList:   []*relation.UserInfo{}, // 实际应填充用户信息
        }, nil
}
3. 社交服务 GRPC 定义（src/api/proto/relation.proto）
syntax = "proto3";
package relation;

option go_package = "./relation";

import "user.proto";

// 关注/取消关注请求
message FollowRequest {
  int64 user_id = 1;    // 发起关注的用户ID
  int64 follow_id = 2;  // 被关注的用户ID
  int32 action_type = 3; // 1=关注，2=取消关注
}

// 关注/取消关注响应
message FollowResponse {
  int32 status_code = 1;
  string status_msg = 2;
}

// 获取关注列表请求
message GetFollowListRequest {
  int64 user_id = 1; // 用户ID
}

// 获取关注列表响应
message GetFollowListResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated user.UserInfo user_list = 3; // 关注的用户列表
}

// 获取粉丝列表请求
message GetFollowerListRequest {
  int64 user_id = 1; // 用户ID
}

// 获取粉丝列表响应
message GetFollowerListResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated user.UserInfo user_list = 3; // 粉丝列表
}

// 社交服务接口
service RelationService {
  rpc Follow(FollowRequest) returns (FollowResponse);
  rpc GetFollowList(GetFollowListRequest) returns (GetFollowListResponse);
  rpc GetFollowerList(GetFollowerListRequest) returns (GetFollowerListResponse);
}
4. 消息服务 GRPC 定义（src/api/proto/message.proto）
syntax = "proto3";
package message;

option go_package = "./message";

// 发送消息请求
message SendMessageRequest {
  int64 from_user_id = 1; // 发送者ID
  int64 to_user_id = 2;   // 接收者ID
  string content = 3;     // 消息内容
}

// 发送消息响应
message SendMessageResponse {
  int32 status_code = 1;
  string status_msg = 2;
  MessageInfo message = 3; // 发送的消息信息
}

// 获取消息列表请求
message GetMessageListRequest {
  int64 user_id = 1;    // 当前用户ID
  int64 to_user_id = 2; // 对方用户ID
  int64 pre_msg_time = 3; // 最后一条消息的时间戳（用于分页，默认0）
}

// 获取消息列表响应
message GetMessageListResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated MessageInfo message_list = 3; // 消息列表
}

// 获取未读消息计数请求
message GetUnreadCountRequest {
  int64 user_id = 1;    // 当前用户ID
  int64 to_user_id = 2; // 对方用户ID（可选，为空则获取所有未读）
}

// 获取未读消息计数响应
message GetUnreadCountResponse {
  int32 status_code = 1;
  string status_msg = 2;
  int64 unread_count = 3; // 未读消息数
}

// 消息信息结构体
message MessageInfo {
  int64 id = 1;                // 消息ID
  int64 from_user_id = 2;      // 发送者ID
  int64 to_user_id = 3;        // 接收者ID
  string content = 4;          // 消息内容
  string create_time = 5;      // 发送时间（格式：2006-01-02 15:04:05）
  bool is_read = 6;            // 是否已读
}

// 消息服务接口
service MessageService {
  rpc SendMessage(SendMessageRequest) returns (SendMessageResponse);
  rpc GetMessageList(GetMessageListRequest) returns (GetMessageListResponse);
  rpc GetUnreadCount(GetUnreadCountRequest) returns (GetUnreadCountResponse);
}
5. 推荐服务 GRPC 定义（src/api/proto/recommend.proto）
syntax = "proto3";
package recommend;

option go_package = "./recommend";

import "video.proto";

// 获取推荐视频流请求
message GetRecommendFeedRequest {
  int64 user_id = 1;  // 用户ID（可选，用于个性化推荐）
  int32 limit = 2;    // 一次获取数量，默认20
}

// 获取推荐视频流响应
message GetRecommendFeedResponse {
  int32 status_code = 1;
  string status_msg = 2;
  repeated video.VideoInfo video_list = 3; // 推荐视频列表
}

// 推荐服务接口
service RecommendService {
  rpc GetRecommendFeed(GetRecommendFeedRequest) returns (GetRecommendFeedResponse);
}
6. Redis 键定义
// src/internal/storage/redis/key.go
package redis

// 关注列表键：follow:{user_id} -> set{follow_id1, follow_id2...}
func FollowKey(userID int64) string {
        return fmt.Sprintf("follow:%d", userID)
}

// 粉丝列表键：follower:{user_id} -> set{follower_id1, follower_id2...}
func FollowerKey(userID int64) string {
        return fmt.Sprintf("follower:%d", userID)
}
测试
1. 并发关注测试（使用 Go 测试）
// src/internal/service/relation/service_test.go
package relation

import (
        "context"
        "sync"
        "testing"
        "mini-douyin/src/api/proto/relation"
        "github.com/stretchr/testify/assert"
)

func TestConcurrentFollow(t *testing.T) {
        s := &RelationService{}
        ctx := context.Background()
        userID := int64(1)
        followID := int64(2)
        wg := sync.WaitGroup{}
        wg.Add(100) // 100个并发请求

        for i := 0; i < 100; i++ {
                go func() {
                        defer wg.Done()
                        _, err := s.Follow(ctx, &relation.FollowRequest{
                                UserId:   userID,
                                FollowId: followID,
                        })
                        assert.NoError(t, err)
                }()
        }
        wg.Wait()

        // 验证最终状态
        rels := []model.Relation{}
        postgres.DB.Where("user_id = ? AND follow_id = ?", userID, followID).Find(&rels)
        assert.Len(t, rels, 1)
        assert.True(t, rels[0].IsFollow)
}
第 5 天：消息服务开发
目标
- 实现实时消息发送与存储
- 消息列表查询
- 未读消息计数
技术栈
- RabbitMQ（消息异步投递）、Redis（未读计数）
1. 消息模型
// src/internal/storage/model/message.go
package model

import (
        "time"
        "gorm.io/gorm"
)

type Message struct {
        ID           int64          `gorm:"primaryKey" json:"id"`
        FromUserID   int64          `gorm:"index" json:"from_user_id"`
        ToUserID     int64          `gorm:"index" json:"to_user_id"`
        Content      string         `json:"content"`
        CreateTime   time.Time      `json:"create_time"`
        IsRead       bool           `json:"is_read"`
        CreatedAt    time.Time      `json:"-"`
        UpdatedAt    time.Time      `json:"-"`
        DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
2. 消息发送实现
// src/internal/service/message/service.go
package message

import (
        "context"
        "mini-douyin/src/api/proto/message"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/redis"
        "mini-douyin/src/pkg/rabbitmq"
        "time"
)

type MessageService struct{}

// 发送消息
func (s *MessageService) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
        // 保存消息到数据库
        msg := model.Message{
                FromUserID: req.FromUserId,
                ToUserID:   req.ToUserId,
                Content:    req.Content,
                CreateTime: time.Now(),
                IsRead:     false,
        }
        if err := postgres.DB.WithContext(ctx).Create(&msg).Error; err != nil {
                return nil, err
        }

        // 发送消息到RabbitMQ（通知接收方）
        rabbitmq.Publish("message.exchange", "message.notify", map[string]interface{}{
                "message_id":  msg.ID,
                "from_user_id": msg.FromUserID,
                "to_user_id":   msg.ToUserID,
                "content":      msg.Content,
        })

        // 增加未读计数
        redis.Client.Incr(ctx, redis.UnreadKey(req.ToUserId, req.FromUserId))

        return &message.SendMessageResponse{
                StatusCode: 0,
                Message: &message.MessageInfo{
                        Id:         msg.ID,
                        FromUserId: msg.FromUserID,
                        ToUserId:   msg.ToUserID,
                        Content:    msg.Content,
                        CreateTime: msg.CreateTime.Format("2006-01-02 15:04:05"),
                },
        }, nil
}

// 获取消息列表
func (s *MessageService) GetMessageList(ctx context.Context, req *message.MessageListRequest) (*message.MessageListResponse, error) {
        var messages []model.Message
        if err := postgres.DB.WithContext(ctx).
                Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
                        req.UserId, req.ToUserId, req.ToUserId, req.UserId).
                Order("create_time DESC").
                Limit(100).
                Find(&messages).Error; err != nil {
                return nil, err
        }

        // 标记消息为已读
        if err := postgres.DB.WithContext(ctx).
                Model(&model.Message{}).
                Where("from_user_id = ? AND to_user_id = ? AND is_read = false",
                        req.ToUserId, req.UserId).
                Update("is_read", true).Error; err != nil {
                return nil, err
        }

        // 重置未读计数
        redis.Client.Del(ctx, redis.UnreadKey(req.UserId, req.ToUserId))

        // 转换为响应格式（省略具体转换逻辑）
        return &message.MessageListResponse{
                StatusCode: 0,
                MessageList: []*message.MessageInfo{},
        }, nil
}
3. 未读消息计数键
// src/internal/storage/redis/key.go
// 未读消息键：unread:{user_id}:{from_user_id} -> count
func UnreadKey(userID, fromUserID int64) string {
        return fmt.Sprintf("unread:%d:%d", userID, fromUserID)
}
测试
1. 消息可靠性测试
  - 发送 1000 条消息，验证数据库与 RabbitMQ 消息一致性
  - 模拟消费者宕机，验证消息重试机制
2. 性能测试
  - 使用 WRK 测试消息列表查询接口响应时间
wrk -t10 -c100 -d30s "http://localhost:8080/douyin/message/chat/?user_id=1&to_user_id=2"
第 6 天：推荐系统与缓存优化
目标
- 实现基于热门度的视频推荐
- 多级缓存优化（Redis + 本地缓存）
- 定时任务更新推荐列表
技术栈
- Redis（热门视频缓存）、Golang 定时器、本地 LRU 缓存
1. 推荐服务实现
// src/internal/service/recommend/service.go
package recommend

import (
        "context"
        "mini-douyin/src/api/proto/recommend"
        "mini-douyin/src/internal/storage/model"
        "mini-douyin/src/internal/storage/postgres"
        "mini-douyin/src/internal/storage/redis"
        "sync"
        "time"

        "github.com/hashicorp/golang-lru/v2"
)

type RecommendService struct {
        // 本地LRU缓存（容量1000）
        localCache *lru.Cache[int64, []*model.Video]
        mu         sync.RWMutex
}

func NewRecommendService() *RecommendService {
        cache, _ := lru.New[int64, []*model.Video](1000)
        s := &RecommendService{localCache: cache}
        // 启动定时任务（每小时更新热门视频）
        go s.updateHotVideos()
        return s
}

// 获取推荐视频
func (s *RecommendService) GetRecommend(ctx context.Context, req *recommend.RecommendRequest) (*recommend.RecommendResponse, error) {
        // 1. 先查本地缓存
        s.mu.RLock()
        videos, ok := s.localCache.Get(0) // 0表示全局推荐
        s.mu.RUnlock()
        if ok {
                return s.convertToResponse(videos), nil
        }

        // 2. 查Redis缓存
        videoIDs, err := redis.Client.ZRevRange(ctx, "hot_videos", 0, 20).Result()
        if err != nil && err != redis.Nil {
                return nil, err
        }

        // 3. 缓存未命中则查数据库
        if len(videoIDs) == 0 || err == redis.Nil {
                var videos []*model.Video
                if err := postgres.DB.WithContext(ctx).
                        Order("favorite_count + comment_count DESC").
                        Limit(20).
                        Find(&videos).Error; err != nil {
                        return nil, err
                }
                // 同步到Redis和本地缓存
                s.syncCache(videos)
                return s.convertToResponse(videos), nil
        }

        // 4. 从数据库加载视频详情（省略）
        return &recommend.RecommendResponse{}, nil
}

// 定时更新热门视频
func (s *RecommendService) updateHotVideos() {
        ticker := time.NewTicker(1 * time.Hour)
        defer ticker.Stop()

        for range ticker.C {
                var videos []*model.Video
                if err := postgres.DB.
                        Order("favorite_count + comment_count DESC").
                        Limit(100).
                        Find(&videos).Error; err != nil {
                        continue
                }
                s.syncCache(videos)
        }
}

// 同步缓存
func (s *RecommendService) syncCache(videos []*model.Video) {
        // 更新Redis有序集（分数为热度值）
        ctx := context.Background()
        pipe := redis.Client.Pipeline()
        for _, v := range videos {
                score := float64(v.FavoriteCount + v.CommentCount)
                pipe.ZAdd(ctx, "hot_videos", redis.Z{Score: score, Member: v.ID})
        }
        pipe.Expire(ctx, "hot_videos", 2*time.Hour)
        pipe.Exec(ctx)

        // 更新本地缓存
        s.mu.Lock()
        s.localCache.Add(0, videos)
        s.mu.Unlock()
}
测试
1. 缓存一致性测试
  - 修改视频热度值，验证 1 小时后缓存是否更新
  - 对比本地缓存、Redis 缓存与数据库数据一致性
2. 推荐性能测试
  - 测试推荐接口在缓存命中 / 未命中场景下的响应时间
# 首次请求（未命中）
time curl http://localhost:8080/douyin/feed/
# 二次请求（命中）
time curl http://localhost:8080/douyin/feed/
第 7 天：系统集成与监控
目标
- 集成全链路监控
- 性能压测与优化
- 部署文档编写
1. 链路追踪集成（src/pkg/tracing/tracer.go）
package tracing

import (
        "context"
        "go.opentelemetry.io/otel"
        "go.opentelemetry.io/otel/exporters/jaeger"
        "go.opentelemetry.io/otel/sdk/resource"
        tracesdk "go.opentelemetry.io/otel/sdk/trace"
        semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitTracer 初始化Jaeger追踪
func InitTracer(serviceName string) (*tracesdk.TracerProvider, error) {
        // 连接Jaeger
        exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://jaeger:14268/api/traces")))
        if err != nil {
                return nil, err
        }

        // 创建TracerProvider
        tp := tracesdk.NewTracerProvider(
                tracesdk.WithBatcher(exp),
                tracesdk.WithResource(resource.NewWithAttributes(
                        semconv.SchemaURL,
                        semconv.ServiceName(serviceName),
                )),
        )

        // 设置全局Tracer
        otel.SetTracerProvider(tp)
        return tp, nil
}
2. 性能压测（使用 k6）
// 全链路压测脚本
import http from 'k6/http';
import { check, sleep } from 'k6';

export default function() {
  // 登录
  const loginRes = http.post('http://localhost:8080/douyin/user/login/', 
    JSON.stringify({username: 'test', password: '123456'}),
    {headers: {'Content-Type': 'application/json'}}
  );
  const token = JSON.parse(loginRes.body).token;
  check(loginRes, {'login success': (r) => r.status === 200});

  // 刷视频流
  const feedRes = http.get('http://localhost:8080/douyin/feed/', 
    {headers: {'Authorization': `Bearer ${token}`}}
  );
  check(feedRes, {'feed success': (r) => r.status === 200});

  sleep(1);
}

// 运行：k6 run --vus 50 --duration 60s full_link_test.js
3. 部署文档（deploy.md）
# 部署指南

## 环境要求
- Docker 20.10+
- Docker Compose 2.0+
- 内存 >= 4GB

## 部署步骤
1. 克隆代码
   ```bash
   git clone https://github.com/yourname/mini-douyin.git
   cd mini-douyin

2. 配置环境
    cp .env.example .env
    # 编辑.env文件设置关键参数
    
3. 启动服务
    docker-compose up -d

4. 验证部署
    curl http://localhost:8080/health
    # 预期返回：{"status":"ok"}
### 测试
1. **全链路监控验证**
   - 访问Jaeger UI（http://localhost:16686），查看服务调用链路
   - 验证各服务间调用关系与耗时

2. **高并发稳定性测试**
   - 持续压测30分钟，监控系统错误率（目标<0.1%）
   - 检查内存泄漏与数据库连接池状态


## 总结
本项目通过7天实现了一个功能完整的迷你抖音后端，包含用户、视频、社交、消息等核心模块，采用微服务架构确保可扩展性，通过Redis缓存、RabbitMQ异步任务、多级缓存等机制保证高并发场景下的性能。测试环节覆盖单元测试、集成测试、压力测试，确保代码质量与系统稳定性。
[图片]
[图片]

部署与运维补充
网关 Dockerfile（docker/gateway/Dockerfile）
# 构建阶段
FROM golang:1.20-alpine AS builder
WORKDIR /app
# 复制依赖文件
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download
# 复制源代码
COPY . .
# 构建网关二进制文件（指定GOOS和GOARCH，适配Linux）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway ./src/cmd/gateway

# 运行阶段（使用轻量级Alpine镜像）
FROM alpine:3.18
WORKDIR /app
# 复制构建产物
COPY --from=builder /app/gateway .
# 复制环境变量模板
COPY --from=builder /app/.env.example .
# 暴露网关端口
EXPOSE 8080
# 启动命令
CMD ["./gateway"]
用户服务 Dockerfile（docker/user-service/Dockerfile）
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 构建用户服务二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service ./src/cmd/service/user

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/user-service .
COPY --from=builder /app/.env.example .
EXPOSE 50051
CMD ["./user-service"]
docker-compose.yml（完整服务编排）
version: '3.8'

services:
  # 基础组件
  consul:
    image: consul:1.15
    container_name: mini-douyin-consul
    ports:
      - "8500:8500"  # Consul UI端口
      - "8600:8600/udp" # DNS端口
    command: agent -dev -client=0.0.0.0
    restart: always
    networks:
      - mini-douyin-network

  postgres:
    image: postgres:14
    container_name: mini-douyin-postgres
    environment:
      POSTGRES_USER: douyin
      POSTGRES_PASSWORD: douyin
      POSTGRES_DB: douyin
      TZ: Asia/Shanghai
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data
    restart: always
    networks:
      - mini-douyin-network

  redis:
    image: redis:7-alpine
    container_name: mini-douyin-redis
    ports:
      - "6379:6379"
    command: redis-server --appendonly yes --requirepass ""
    volumes:
      - redisdata:/data
    restart: always
    networks:
      - mini-douyin-network

  rabbitmq:
    image: rabbitmq:3-management
    container_name: mini-douyin-rabbitmq
    environment:
      RABBITMQ_DEFAULT_USER: douyin
      RABBITMQ_DEFAULT_PASS: douyin
      RABBITMQ_DEFAULT_VHOST: /
    ports:
      - "5672:5672"  # AMQP端口
      - "15672:15672" # 管理UI端口
    volumes:
      - rabbitmqdata:/var/lib/rabbitmq
    restart: always
    networks:
      - mini-douyin-network

  minio:
    image: minio/minio:latest
    container_name: mini-douyin-minio
    environment:
      MINIO_ROOT_USER: minioadmin
      MINIO_ROOT_PASSWORD: minioadmin
    ports:
      - "9000:9000"  # API端口
      - "9001:9001"  # 管理UI端口
    command: server /data --console-address ":9001"
    volumes:
      - miniostorage:/data
    restart: always
    networks:
      - mini-douyin-network

  jaeger:
    image: jaegertracing/all-in-one:1.47
    container_name: mini-douyin-jaeger
    environment:
      COLLECTOR_OTLP_ENABLED: "true"
    ports:
      - "6831:6831/udp"  # 采样端口
      - "16686:16686"    # UI端口
      - "4317:4317"       # OTLP gRPC端口
    restart: always
    networks:
      - mini-douyin-network

  prometheus:
    image: prom/prometheus:v2.47.0
    container_name: mini-douyin-prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./docker/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
      - prometheusdata:/prometheus
    restart: always
    networks:
      - mini-douyin-network

  grafana:
    image: grafana/grafana:10.1.2
    container_name: mini-douyin-grafana
    ports:
      - "3000:3000"
    volumes:
      - grafanadata:/var/lib/grafana
    environment:
      GF_SECURITY_ADMIN_USER: admin
      GF_SECURITY_ADMIN_PASSWORD: admin
    restart: always
    networks:
      - mini-douyin-network

  # 业务服务
  gateway:
    build:
      context: .
      dockerfile: ./docker/gateway/Dockerfile
    container_name: mini-douyin-gateway
    ports:
      - "8080:8080"
    environment:
      - GATEWAY_ADDR=:8080
      - CONSUL_ADDR=consul:8500
      - JAEGER_ENDPOINT=http://jaeger:14268/api/traces
    depends_on:
      - consul
      - jaeger
    restart: always
    networks:
      - mini-douyin-network

  user-service:
    build:
      context: .
      dockerfile: ./docker/user-service/Dockerfile
    container_name: mini-douyin-user-service
    ports:
      - "50051:50051"
    environment:
      - RPC_PORT=50051
      - CONSUL_ADDR=consul:8500
      - POSTGRES_HOST=postgres
      - REDIS_ADDR=redis:6379
      - JAEGER_ENDPOINT=http://jaeger:14268/api/traces
    depends_on:
      - consul
      - postgres
      - redis
      - jaeger
    restart: always
    networks:
      - mini-douyin-network

  video-service:
    build:
      context: .
      dockerfile: ./docker/video-service/Dockerfile
    container_name: mini-douyin-video-service
    ports:
      - "50052:50052"
    environment:
      - RPC_PORT=50052
      - CONSUL_ADDR=consul:8500
      - POSTGRES_HOST=postgres
      - REDIS_ADDR=redis:6379
      - RABBITMQ_HOST=rabbitmq
      - MINIO_ENDPOINT=minio:9000
      - JAEGER_ENDPOINT=http://jaeger:14268/api/traces
    depends_on:
      - consul
      - postgres
      - redis
      - rabbitmq
      - minio
      - jaeger
    restart: always
    networks:
      - mini-douyin-network

  cover-worker:
    build:
      context: .
      dockerfile: ./docker/cover-worker/Dockerfile
    container_name: mini-douyin-cover-worker
    environment:
      - CONSUL_ADDR=consul:8500
      - RABBITMQ_HOST=rabbitmq
      - MINIO_ENDPOINT=minio:9000
      - VIDEO_SERVICE_RPC=video-service:50052
    depends_on:
      - consul
      - rabbitmq
      - minio
      - video-service
    restart: always
    networks:
      - mini-douyin-network

# 数据卷（持久化存储）
volumes:
  pgdata:
  redisdata:
  rabbitmqdata:
  miniostorage:
  prometheusdata:
  grafanadata:

# 网络（所有服务在同一网络，可通过服务名访问）
networks:
  mini-douyin-network:
    driver: bridge
Prometheus 配置（docker/prometheus/prometheus.yml）
global:
  scrape_interval: 15s # 全局采样间隔

scrape_configs:
  # 监控Prometheus自身
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  # 监控网关服务
  - job_name: 'gateway'
    static_configs:
      - targets: ['gateway:8080']

  # 监控用户服务
  - job_name: 'user-service'
    static_configs:
      - targets: ['user-service:50051']

  # 监控视频服务
  - job_name: 'video-service'
    static_configs:
      - targets: ['video-service:50052']
测试
全链路压测脚本（test/k6/full_link_test.js）
import http from 'k6/http';
import { check, sleep, group } from 'k6';

// 压测配置
export const options = {
  vus: 50,        // 虚拟用户数
  duration: '60s',// 压测时长
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95%请求响应时间<500ms
    http_req_failed: ['rate<0.01'],   // 请求失败率<1%
  },
};

// 全局变量
let token = '';
const baseUrl = 'http://localhost:8080/douyin';

// 压测逻辑
export default function () {
  // 1. 登录（每个虚拟用户首次执行）
  if (token === '') {
    group('用户登录', function () {
      const loginRes = http.post(
        `${baseUrl}/user/login`,
        JSON.stringify({ username: 'test', password: '123456' }),
        { headers: { 'Content-Type': 'application/json' } }
      );
      check(loginRes, {
        '登录成功': (r) => r.status === 200 && JSON.parse(r.body).status_code === 0,
      });
      if (loginRes.status === 200) {
        token = JSON.parse(loginRes.body).token;
      }
    });
    sleep(1);
  }

  // 2. 获取视频流
  group('获取视频流', function () {
    const feedRes = http.get(
      `${baseUrl}/video/feed`,
      { headers: { 'Authorization': `Bearer ${token}` } }
    );
    check(feedRes, {
      '视频流获取成功': (r) => r.status === 200 && JSON.parse(r.body).status_code === 0,
      '视频列表非空': (r) => JSON.parse(r.body).video_list.length > 0,
    });
  });
  sleep(1);

  // 3. 获取用户信息
  group('获取用户信息', function () {
    const userRes = http.get(
      `${baseUrl}/user?token=${token}`,
      { headers: { 'Authorization': `Bearer ${token}` } }
    );
    check(userRes, {
      '用户信息获取成功': (r) => r.status === 200 && JSON.parse(r.body).status_code === 0,
    });
  });
  sleep(1);
}
运行压测命令：
# 安装k6（Ubuntu）
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
echo "deb
问题：
问题1：go.opentelemetry.io/otel/semconv v1.21.0 这个mod包在go mod tidy时一直报错！
问题出在 go.opentelemetry.io/otel/semconv/v1 这个模块路径上。OpenTelemetry 的 semconv 模块实际路径格式与我使用的不一致，导致解析失败。
// 替换掉这一行
go.opentelemetry.io/otel/semconv/v1 v1.21.0

// 改为
go.opentelemetry.io/otel/semconv v1.21.0