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