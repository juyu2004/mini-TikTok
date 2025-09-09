# Mini TikTok (迷你抖音)

基于微服务架构的短视频社交平台示例（教学/演示用）。
包含：Go gRPC 微服务、Gin API Gateway、Vue 3 前端、RabbitMQ/Redis/MySQL/MinIO、
Consul 服务注册、Jaeger 链路追踪、Prometheus + Grafana 监控、Pyroscope 性能分析、
Docker Compose 一键启动（Kubernetes 清单可选）。

> 该仓库为可运行的最小实现与脚手架，便于进一步扩展。

## 快速开始（Docker Compose）
1. 安装并启动 Docker 与 Docker Compose
2. 在项目根目录执行：
   ```bash
   docker compose up -d --build
   ```
3. 访问：
   - 前端（Vue）: http://localhost:5173
   - API Gateway: http://localhost:8080
   - Consul: http://localhost:8500
   - Jaeger: http://localhost:16686
   - RabbitMQ: http://localhost:15672  用户名/密码: guest/guest
   - MinIO: http://localhost:9001  用户名/密码: minioadmin/minioadmin
   - Grafana: http://localhost:3000  用户名/密码: admin/admin
   - Prometheus: http://localhost:9090
   - Pyroscope: http://localhost:4040

> 首次启动会执行数据库迁移、创建最小表结构和示例数据（见各服务 init 逻辑）。

## 项目结构
```
mini-tiktok/
  gateway/                # Gin API Gateway
  services/
    user/                 # 用户服务（gRPC）
    video/                # 视频服务（gRPC）
    social/               # 社交服务（gRPC）
    recommend/            # 推荐服务（gRPC）
  proto/                  # gRPC protobuf 定义
  pkg/                    # 公共库（JWT、Consul、Jaeger、日志等）
  webapp/                 # 前端 Vue 3 + Vite + Element Plus
  deploy/                 # 监控、配置、K8s 等
  .github/workflows/      # CI/CD
  docker-compose.yml
```
## 默认账号
- 测试账号: `test@example.com` / `Passw0rd!`

## 开发命令（本地）
- 运行单个服务：在对应目录执行
  ```bash
  go run ./
  ```
- 生成 gRPC 代码：
  ```bash
  (在容器或本机安装 protoc 和 protoc-gen-go / protoc-gen-go-grpc)
  protoc -I=proto --go_out=. --go-grpc_out=. proto/*.proto
  ```

## 注意
- 该工程为教学示例，安全、鲁棒性与生产最佳实践需结合实际完善。
