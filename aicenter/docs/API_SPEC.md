# AICenter API 规范

本文档定义 aicenter 的 API 接口，**与 TGO 保持完全兼容**。

## 接口来源

| 原服务 | 端口 | aicenter 路由前缀 |
|--------|------|------------------|
| tgo-ai | 8081 | `/v1/ai/` |
| tgo-rag | 8082 | `/v1/rag/` |
| tgo-workflow | 8004 | `/v1/workflow/` |

---

## 1. tgo-ai 接口 (`/v1/ai/`)

### 1.1 Agents

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/agents` | 获取 Agent 列表 |
| `POST` | `/agents` | 创建 Agent |
| `GET` | `/agents/{agent_id}` | 获取 Agent 详情 |
| `PUT` | `/agents/{agent_id}` | 更新 Agent |
| `DELETE` | `/agents/{agent_id}` | 删除 Agent |
| `POST` | `/agents/run` | 运行 Agent（支持 SSE 流式） |
| `POST` | `/agents/{agent_id}/toggle` | 启用/禁用 Agent |

### 1.2 Teams

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/teams` | 获取 Team 列表 |
| `POST` | `/teams` | 创建 Team |
| `GET` | `/teams/{team_id}` | 获取 Team 详情 |
| `PUT` | `/teams/{team_id}` | 更新 Team |
| `DELETE` | `/teams/{team_id}` | 删除 Team |

### 1.3 LLM Providers

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/llm-providers` | 获取 Provider 列表 |
| `POST` | `/llm-providers` | 创建 Provider |
| `GET` | `/llm-providers/{provider_id}` | 获取 Provider 详情 |
| `PUT` | `/llm-providers/{provider_id}` | 更新 Provider |
| `DELETE` | `/llm-providers/{provider_id}` | 删除 Provider |

### 1.4 Tools

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/tools` | 获取 Tool 列表 |
| `POST` | `/tools` | 创建 Tool |
| `GET` | `/tools/{tool_id}` | 获取 Tool 详情 |
| `PUT` | `/tools/{tool_id}` | 更新 Tool |
| `DELETE` | `/tools/{tool_id}` | 删除 Tool |

### 1.5 Chat

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/chat` | 对话（支持 SSE 流式） |

### 1.6 Project AI Configs

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/project-ai-configs` | 获取项目 AI 配置 |
| `PUT` | `/project-ai-configs` | 更新项目 AI 配置 |

---

## 2. tgo-rag 接口 (`/v1/rag/`)

### 2.1 Collections

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/collections` | 获取 Collection 列表 |
| `POST` | `/collections` | 创建 Collection |
| `GET` | `/collections/{collection_id}` | 获取 Collection 详情 |
| `PUT` | `/collections/{collection_id}` | 更新 Collection |
| `DELETE` | `/collections/{collection_id}` | 删除 Collection |
| `POST` | `/collections/{collection_id}/search` | 搜索 Collection |
| `POST` | `/collections/batch` | 批量获取 Collection |

### 2.2 Files

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/files` | 获取文件列表 |
| `POST` | `/files` | 上传文件 |
| `GET` | `/files/{file_id}` | 获取文件详情 |
| `DELETE` | `/files/{file_id}` | 删除文件 |
| `POST` | `/files/{file_id}/reprocess` | 重新处理文件 |

### 2.3 Websites

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/websites` | 获取网站列表 |
| `POST` | `/websites` | 添加网站 |
| `GET` | `/websites/{website_id}` | 获取网站详情 |
| `DELETE` | `/websites/{website_id}` | 删除网站 |
| `POST` | `/websites/{website_id}/crawl` | 爬取网站 |

### 2.4 QA

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/qa` | 获取 QA 列表 |
| `POST` | `/qa` | 创建 QA |
| `PUT` | `/qa/{qa_id}` | 更新 QA |
| `DELETE` | `/qa/{qa_id}` | 删除 QA |
| `POST` | `/qa/batch` | 批量创建 QA |

### 2.5 Embedding Config

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/embedding-config` | 获取嵌入配置 |
| `PUT` | `/embedding-config` | 更新嵌入配置 |

---

## 3. tgo-workflow 接口 (`/v1/workflow/`)

### 3.1 Workflows

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/workflows` | 获取 Workflow 列表 |
| `POST` | `/workflows` | 创建 Workflow |
| `GET` | `/workflows/batch` | 批量获取 Workflow |
| `GET` | `/workflows/{workflow_id}` | 获取 Workflow 详情 |
| `PUT` | `/workflows/{workflow_id}` | 更新 Workflow |
| `DELETE` | `/workflows/{workflow_id}` | 删除 Workflow |
| `POST` | `/workflows/{workflow_id}/duplicate` | 复制 Workflow |
| `POST` | `/workflows/validate` | 验证 Workflow |
| `GET` | `/workflows/{workflow_id}/variables` | 获取可用变量 |

### 3.2 Executions

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/executions` | 获取执行记录列表 |
| `POST` | `/executions` | 创建执行（运行 Workflow） |
| `GET` | `/executions/{execution_id}` | 获取执行详情 |
| `POST` | `/executions/{execution_id}/cancel` | 取消执行 |
| `POST` | `/executions/{execution_id}/retry` | 重试执行 |
| `GET` | `/executions/{execution_id}/stream` | SSE 流式获取执行状态 |

---

## 4. 通用参数

### Query Parameters

| 参数 | 类型 | 说明 |
|------|------|------|
| `project_id` | UUID | 项目 ID（必需） |
| `limit` | int | 分页大小（默认 20） |
| `offset` | int | 分页偏移（默认 0） |
| `skip` | int | 分页偏移（兼容 tgo-workflow） |

### Response Format

```json
{
  "data": [...],
  "pagination": {
    "total": 100,
    "limit": 20,
    "offset": 0,
    "has_next": true,
    "has_prev": false
  }
}
```

---

## 5. 实现进度

| 模块 | 接口 | 状态 |
|------|------|------|
| **tgo-ai** | Agents | 🟡 部分 |
| **tgo-ai** | Teams | 🔴 待实现 |
| **tgo-ai** | LLM Providers | 🟡 部分 |
| **tgo-ai** | Tools | 🟡 部分 |
| **tgo-ai** | Chat | 🟢 已实现 |
| **tgo-rag** | Collections | 🔴 待实现 |
| **tgo-rag** | Files | 🔴 待实现 |
| **tgo-rag** | Websites | 🔴 待实现 |
| **tgo-rag** | QA | 🔴 待实现 |
| **tgo-workflow** | Workflows | 🔴 待实现 |
| **tgo-workflow** | Executions | 🔴 待实现 |
