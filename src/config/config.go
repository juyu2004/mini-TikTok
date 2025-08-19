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