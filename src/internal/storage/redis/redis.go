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