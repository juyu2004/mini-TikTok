# GuGoTik

_✨ 迷你抖音后端微服务项目，基于Go语言实现的微服务架构实践 ✨_

## 项目简介

GuGoTik 是一个基于微服务架构的迷你抖音后端项目，支持用户管理、视频上传/处理、评论、消息互动等核心功能。项目采用 Go 语言开发，结合 gRPC 实现服务间通信，使用 Docker 容器化部署，适配 Kubernetes 集群环境，旨在提供高可扩展性和可维护性的后端解决方案。

## 核心功能

- 用户模块：注册、登录、个人信息管理
- 视频模块：上传、转码、封面生成、流加载
- 互动模块：评论、点赞、关注/粉丝关系
- 消息模块：实时聊天、消息通知
- 推荐模块：基于用户行为的视频推荐

## 技术栈

- **开发语言**：Go 1.20+
- **通信协议**：gRPC（服务间）、HTTP（网关层）
- **数据存储**：PostgreSQL（主数据库）、Redis（缓存/会话）
- **消息队列**：RabbitMQ（异步任务处理）
- **容器化**：Docker、Docker Compose
- **CI/CD**：GitHub Actions
- **监控**：OpenTelemetry、Grafana、Jaeger

## 项目结构

GuGoTik/
├── docker/ # 基础镜像配置
├── scripts/ # 构建与部署脚本
├── src/ # 源代码
│ ├── constant/ # 项目常量定义
│ ├── idl/ # gRPC 接口定义
│ ├── models/ # 数据模型
│ ├── rpc/ # gRPC 客户端 / 服务端代码
│ ├── services/ # 微服务实现
│ ├── storage/ # 存储层封装
│ ├── utils/ # 工具函数
│ └── web/ # 网关层代码
├── test/ # 单元测试与集成测试
└── docker-compose.yml # 本地部署配置


## 快速开始

### 本地部署（Docker Compose）

1. 克隆仓库：
   ```bash
   git clone https://github.com/juyu2004/mini-TikTok.git
   cd GuGoTik

2. 启动服务
    docker compose up -d

3. 启动服务
    访问网关：默认地址 http://localhost:8080

### 手动构建

1. 安装依赖：
    go mod download

2. 构建服务
    cd scripts
    ./build-all.sh  # 选择对应平台的脚本

3. 运行服务
    ./run-all.sh

## 贡献指南

提交 Issue 描述功能需求或 Bug
Fork 仓库并创建特性分支（feature/xxx 或 fix/xxx）
提交代码前确保通过 lint 和单元测试
发起 Pull Request 到 dev 分支

## 许可证

本项目基于 GNU General Public License v3.0 开源，详见 LICENSE
