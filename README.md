# Next-Chat

基于 Eino ADK 的 AI Agent 客服平台，采用简化的微服务架构。

## 架构

```
┌─────────────────────────────────────────────────────────────┐
│                         Frontend                             │
│                   (tgo-web / tgo-widget)                    │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                       apiserver                              │
│              (业务网关、认证、会话、访客)                      │
│                        :8000                                 │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                       aicenter                               │
│           (Agent + RAG + Workflow 统一服务)                  │
│                        :8081                                 │
└─────────────────────────────────────────────────────────────┘
```

## 服务说明

| 服务 | 端口 | 职责 |
|------|------|------|
| **aicenter** | 8081 | AI 全能力中心：Agent、RAG、Workflow |
| **apiserver** | 8000 | 业务网关：认证、会话、访客、转人工 |
| **platform** | 8083 | 多渠道接入（可选） |

## 技术栈

- **语言**：Go 1.23+
- **Web 框架**：Gin
- **ORM**：Gorm
- **AI 框架**：Eino ADK
- **数据库**：PostgreSQL + pgvector
- **配置**：Viper

## 快速开始

### 1. 启动 aicenter

```bash
cd aicenter
cp configs/config.yaml.example configs/config.yaml
# 编辑配置文件
go run cmd/server/main.go
```

### 2. 启动 apiserver（待开发）

```bash
cd apiserver
go run cmd/server/main.go
```

## 目录结构

```
next-chat/
├── aicenter/           # AI 能力中心
│   ├── go.mod
│   ├── cmd/server/
│   └── internal/
│       ├── handler/    # HTTP Handler
│       ├── biz/        # 业务逻辑
│       ├── store/      # 数据访问
│       ├── model/      # 数据模型
│       └── pkg/agent/  # Eino Agent
│
├── apiserver/          # 业务网关（待开发）
│   ├── go.mod
│   └── ...
│
├── platform/           # 多渠道接入（可选）
│   ├── go.mod
│   └── ...
│
├── docker-compose.yml
└── Makefile
```

## 开发计划

- [x] aicenter 基础框架
- [ ] aicenter workflow 模块
- [ ] apiserver 服务
- [ ] platform 服务
- [ ] docker-compose 部署

## 参考

- **Eino ADK**：https://github.com/cloudwego/eino
- **next-show**：原型参考
- **TGO**：业务参考
