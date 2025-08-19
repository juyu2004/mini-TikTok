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