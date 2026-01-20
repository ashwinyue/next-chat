// Package http TGO 兼容 Handler 实现.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ==================== TGO-AI Handlers ====================

// RunAgent 运行 Agent（SSE 流式）.
func (h *Handler) RunAgent(c *gin.Context) {
	// TODO: 实现 Agent 运行，支持 SSE 流式输出
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ToggleAgent 启用/禁用 Agent.
func (h *Handler) ToggleAgent(c *gin.Context) {
	// TODO: 实现 Agent 启用/禁用
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListAgentCollections 获取 Agent 关联的 Collections.
func (h *Handler) ListAgentCollections(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// AddAgentCollection 添加 Agent Collection 关联.
func (h *Handler) AddAgentCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// RemoveAgentCollection 移除 Agent Collection 关联.
func (h *Handler) RemoveAgentCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListTeams 获取 Team 列表.
func (h *Handler) ListTeams(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateTeam 创建 Team.
func (h *Handler) CreateTeam(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetTeam 获取 Team 详情.
func (h *Handler) GetTeam(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateTeam 更新 Team.
func (h *Handler) UpdateTeam(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteTeam 删除 Team.
func (h *Handler) DeleteTeam(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListTools 获取 Tool 列表.
func (h *Handler) ListTools(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateTool 创建 Tool.
func (h *Handler) CreateTool(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetTool 获取 Tool 详情.
func (h *Handler) GetTool(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateTool 更新 Tool.
func (h *Handler) UpdateTool(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteTool 删除 Tool.
func (h *Handler) DeleteTool(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetProjectAIConfig 获取项目 AI 配置.
func (h *Handler) GetProjectAIConfig(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateProjectAIConfig 更新项目 AI 配置.
func (h *Handler) UpdateProjectAIConfig(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ==================== TGO-RAG Handlers ====================

// ListCollections 获取 Collection 列表.
func (h *Handler) ListCollections(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateCollection 创建 Collection.
func (h *Handler) CreateCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// BatchGetCollections 批量获取 Collections.
func (h *Handler) BatchGetCollections(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetCollection 获取 Collection 详情.
func (h *Handler) GetCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateCollection 更新 Collection.
func (h *Handler) UpdateCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteCollection 删除 Collection.
func (h *Handler) DeleteCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// SearchCollection 搜索 Collection.
func (h *Handler) SearchCollection(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListFiles 获取文件列表.
func (h *Handler) ListFiles(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UploadFile 上传文件.
func (h *Handler) UploadFile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetFile 获取文件详情.
func (h *Handler) GetFile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteFile 删除文件.
func (h *Handler) DeleteFile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ReprocessFile 重新处理文件.
func (h *Handler) ReprocessFile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListWebsites 获取网站列表.
func (h *Handler) ListWebsites(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateWebsite 创建网站.
func (h *Handler) CreateWebsite(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetWebsite 获取网站详情.
func (h *Handler) GetWebsite(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteWebsite 删除网站.
func (h *Handler) DeleteWebsite(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CrawlWebsite 爬取网站.
func (h *Handler) CrawlWebsite(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListQA 获取 QA 列表.
func (h *Handler) ListQA(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateQA 创建 QA.
func (h *Handler) CreateQA(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// BatchCreateQA 批量创建 QA.
func (h *Handler) BatchCreateQA(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateQA 更新 QA.
func (h *Handler) UpdateQA(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteQA 删除 QA.
func (h *Handler) DeleteQA(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetEmbeddingConfig 获取嵌入配置.
func (h *Handler) GetEmbeddingConfig(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateEmbeddingConfig 更新嵌入配置.
func (h *Handler) UpdateEmbeddingConfig(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ==================== TGO-Workflow Handlers ====================

// ListWorkflows 获取 Workflow 列表.
func (h *Handler) ListWorkflows(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateWorkflow 创建 Workflow.
func (h *Handler) CreateWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// BatchGetWorkflows 批量获取 Workflows.
func (h *Handler) BatchGetWorkflows(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ValidateWorkflow 验证 Workflow.
func (h *Handler) ValidateWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetWorkflow 获取 Workflow 详情.
func (h *Handler) GetWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateWorkflow 更新 Workflow.
func (h *Handler) UpdateWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteWorkflow 删除 Workflow.
func (h *Handler) DeleteWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DuplicateWorkflow 复制 Workflow.
func (h *Handler) DuplicateWorkflow(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetWorkflowVariables 获取 Workflow 可用变量.
func (h *Handler) GetWorkflowVariables(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// ListExecutions 获取执行记录列表.
func (h *Handler) ListExecutions(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CreateExecution 创建执行（运行 Workflow）.
func (h *Handler) CreateExecution(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetExecution 获取执行详情.
func (h *Handler) GetExecution(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// CancelExecution 取消执行.
func (h *Handler) CancelExecution(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// RetryExecution 重试执行.
func (h *Handler) RetryExecution(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// StreamExecution SSE 流式获取执行状态.
func (h *Handler) StreamExecution(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
