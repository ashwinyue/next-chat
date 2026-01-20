# AICenter

next-chat 的 AI 能力中心，基于 Eino ADK 构建，提供 Agent、RAG、Workflow 等核心 AI 能力。

## 核心特性

- **AI 全能力**：Agent + RAG + Workflow 统一服务
- **MVC 分层架构**：Handler → Biz → Store → Model
- **Eino ADK 集成**：ReAct Agent、Supervisor 多 Agent、compose 编排
- **SSE 流式输出**：实时返回思考过程、工具调用、回复内容
- **知识库 RAG**：文档解析、向量检索、混合检索、重排序
- **工作流引擎**：基于 Eino compose 的 DAG 编排

## 项目结构（MVC 架构）

```
aicenter/
├── cmd/                           # 命令行入口
│   └── server/                   # API Server
│       └── main.go
├── internal/                      # 内部实现（MVC 核心）
│   ├── handler/                  # Handler 层（Controller）
│   │   └── http/                # HTTP 处理器
│   │       ├── router.go        # 路由注册
│   │       ├── chat.go          # Chat API Handler
│   │       └── sse_handler.go   # SSE 流式处理
│   ├── biz/                      # Biz 层（Service/业务逻辑）
│   │   ├── biz.go               # 业务层入口
│   │   ├── agent/               # Agent 业务
│   │   │   └── agent.go
│   │   └── session/             # Session 业务
│   │       └── session.go
│   ├── store/                    # Store 层（Repository/数据访问）
│   │   ├── store.go             # 存储层入口
│   │   ├── session.go           # Session 存储
│   │   └── message.go           # Message 存储
│   ├── model/                    # Model 层（数据模型）
│   │   ├── session.go           # Session 模型
│   │   └── message.go           # Message 模型
│   └── pkg/                      # 内部公共包
│       ├── agent/               # Agent 相关
│       │   └── event/           # 事件发送器（Envelope）
│       │       ├── types.go     # 事件类型定义
│       │       └── sender.go    # 事件发送器
│       └── sse/                 # SSE 协议
│           ├── types.go         # SSE 事件类型
│           ├── writer.go        # SSE 写入器
│           └── processor.go     # 事件处理器
├── pkg/                          # 公共包（可被外部引用）
│   └── api/                     # API 定义
│       └── v1/                  # v1 版本
├── configs/                      # 配置文件
│   └── config.yaml
└── go.mod
```

## MVC 分层说明

| 层级 | 目录 | 职责 |
|------|------|------|
| **Handler** | `internal/handler/` | 处理 HTTP 请求，参数校验，调用 Biz 层，返回响应 |
| **Biz** | `internal/biz/` | 业务逻辑，Agent 执行，事件处理，调用 Store 层 |
| **Store** | `internal/store/` | 数据访问，CRUD 操作，数据库交互 |
| **Model** | `internal/model/` | 数据模型定义，ORM 映射 |

## 快速开始

```bash
# 安装依赖
go mod tidy

# 运行服务
go run cmd/server/main.go
```

## 事件系统

### 统一事件信封（Envelope）

所有自定义事件通过 `Envelope` 结构传递：

```go
type Envelope struct {
    Type    EventType       `json:"type"`
    Content string          `json:"content,omitempty"`
    Payload json.RawMessage `json:"payload,omitempty"`
    Data    map[string]any  `json:"data,omitempty"`
}
```

### SSE 协议

| response_type | 描述 |
|---------------|------|
| `agent_query` | 查询开始 |
| `thinking`    | Agent 思考过程 |
| `tool_call`   | 工具调用 |
| `tool_result` | 工具结果 |
| `references`  | 知识引用 |
| `answer`      | 最终回答 |
| `reflection`  | 反思内容 |
| `stop`        | 完成 |
| `error`       | 错误 |

## License

MIT
