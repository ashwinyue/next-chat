// Package http TGO 兼容路由注册.
package http

import "github.com/gin-gonic/gin"

// ==================== TGO-AI 兼容路由 ====================

// registerAIAgentRoutes 注册 AI Agent 路由（对应 tgo-ai /v1/agents）.
func (h *Handler) registerAIAgentRoutes(r *gin.RouterGroup) {
	agents := r.Group("/agents")
	{
		agents.GET("", h.ListAgents)
		agents.POST("", h.CreateAgent)
		agents.GET("/:agent_id", h.GetAgent)
		agents.PUT("/:agent_id", h.UpdateAgent)
		agents.DELETE("/:agent_id", h.DeleteAgent)
		agents.POST("/run", h.RunAgent)
		agents.POST("/:agent_id/toggle", h.ToggleAgent)
		agents.GET("/:agent_id/tools", h.ListAgentTools)
		agents.POST("/:agent_id/tools", h.AddAgentTool)
		agents.DELETE("/:agent_id/tools/:tool_id", h.RemoveAgentTool)
		agents.GET("/:agent_id/collections", h.ListAgentCollections)
		agents.POST("/:agent_id/collections", h.AddAgentCollection)
		agents.DELETE("/:agent_id/collections/:collection_id", h.RemoveAgentCollection)
	}
}

// registerAITeamRoutes 注册 AI Team 路由（对应 tgo-ai /v1/teams）.
func (h *Handler) registerAITeamRoutes(r *gin.RouterGroup) {
	teams := r.Group("/teams")
	{
		teams.GET("", h.ListTeams)
		teams.POST("", h.CreateTeam)
		teams.GET("/:team_id", h.GetTeam)
		teams.PUT("/:team_id", h.UpdateTeam)
		teams.DELETE("/:team_id", h.DeleteTeam)
	}
}

// registerAIProviderRoutes 注册 LLM Provider 路由（对应 tgo-ai /v1/llm-providers）.
func (h *Handler) registerAIProviderRoutes(r *gin.RouterGroup) {
	providers := r.Group("/llm-providers")
	{
		providers.GET("", h.ListProviders)
		providers.POST("", h.CreateProvider)
		providers.GET("/:provider_id", h.GetProvider)
		providers.PUT("/:provider_id", h.UpdateProvider)
		providers.DELETE("/:provider_id", h.DeleteProvider)
	}
}

// registerAIToolRoutes 注册 AI Tool 路由（对应 tgo-ai /v1/tools）.
func (h *Handler) registerAIToolRoutes(r *gin.RouterGroup) {
	tools := r.Group("/tools")
	{
		tools.GET("", h.ListTools)
		tools.POST("", h.CreateTool)
		tools.GET("/:tool_id", h.GetTool)
		tools.PUT("/:tool_id", h.UpdateTool)
		tools.DELETE("/:tool_id", h.DeleteTool)
	}
}

// registerAIConfigRoutes 注册项目 AI 配置路由（对应 tgo-ai /v1/project-ai-configs）.
func (h *Handler) registerAIConfigRoutes(r *gin.RouterGroup) {
	config := r.Group("/project-ai-configs")
	{
		config.GET("", h.GetProjectAIConfig)
		config.PUT("", h.UpdateProjectAIConfig)
	}
}

// ==================== TGO-RAG 兼容路由 ====================

// registerRAGCollectionRoutes 注册 Collection 路由（对应 tgo-rag /collections）.
func (h *Handler) registerRAGCollectionRoutes(r *gin.RouterGroup) {
	collections := r.Group("/collections")
	{
		collections.GET("", h.ListCollections)
		collections.POST("", h.CreateCollection)
		collections.POST("/batch", h.BatchGetCollections)
		collections.GET("/:collection_id", h.GetCollection)
		collections.PUT("/:collection_id", h.UpdateCollection)
		collections.DELETE("/:collection_id", h.DeleteCollection)
		collections.POST("/:collection_id/search", h.SearchCollection)
	}
}

// registerRAGFileRoutes 注册 File 路由（对应 tgo-rag /files）.
func (h *Handler) registerRAGFileRoutes(r *gin.RouterGroup) {
	files := r.Group("/files")
	{
		files.GET("", h.ListFiles)
		files.POST("", h.UploadFile)
		files.GET("/:file_id", h.GetFile)
		files.DELETE("/:file_id", h.DeleteFile)
		files.POST("/:file_id/reprocess", h.ReprocessFile)
	}
}

// registerRAGWebsiteRoutes 注册 Website 路由（对应 tgo-rag /websites）.
func (h *Handler) registerRAGWebsiteRoutes(r *gin.RouterGroup) {
	websites := r.Group("/websites")
	{
		websites.GET("", h.ListWebsites)
		websites.POST("", h.CreateWebsite)
		websites.GET("/:website_id", h.GetWebsite)
		websites.DELETE("/:website_id", h.DeleteWebsite)
		websites.POST("/:website_id/crawl", h.CrawlWebsite)
	}
}

// registerRAGQARoutes 注册 QA 路由（对应 tgo-rag /qa）.
func (h *Handler) registerRAGQARoutes(r *gin.RouterGroup) {
	qa := r.Group("/qa")
	{
		qa.GET("", h.ListQA)
		qa.POST("", h.CreateQA)
		qa.POST("/batch", h.BatchCreateQA)
		qa.PUT("/:qa_id", h.UpdateQA)
		qa.DELETE("/:qa_id", h.DeleteQA)
	}
}

// registerRAGEmbeddingConfigRoutes 注册嵌入配置路由（对应 tgo-rag /embedding-config）.
func (h *Handler) registerRAGEmbeddingConfigRoutes(r *gin.RouterGroup) {
	embedding := r.Group("/embedding-config")
	{
		embedding.GET("", h.GetEmbeddingConfig)
		embedding.PUT("", h.UpdateEmbeddingConfig)
	}
}

// ==================== TGO-Workflow 兼容路由 ====================

// registerWorkflowRoutes 注册 Workflow 路由（对应 tgo-workflow /workflows）.
func (h *Handler) registerWorkflowRoutes(r *gin.RouterGroup) {
	workflows := r.Group("/workflows")
	{
		workflows.GET("", h.ListWorkflows)
		workflows.POST("", h.CreateWorkflow)
		workflows.GET("/batch", h.BatchGetWorkflows)
		workflows.POST("/validate", h.ValidateWorkflow)
		workflows.GET("/:workflow_id", h.GetWorkflow)
		workflows.PUT("/:workflow_id", h.UpdateWorkflow)
		workflows.DELETE("/:workflow_id", h.DeleteWorkflow)
		workflows.POST("/:workflow_id/duplicate", h.DuplicateWorkflow)
		workflows.GET("/:workflow_id/variables", h.GetWorkflowVariables)
	}
}

// registerExecutionRoutes 注册 Execution 路由（对应 tgo-workflow /executions）.
func (h *Handler) registerExecutionRoutes(r *gin.RouterGroup) {
	executions := r.Group("/executions")
	{
		executions.GET("", h.ListExecutions)
		executions.POST("", h.CreateExecution)
		executions.GET("/:execution_id", h.GetExecution)
		executions.POST("/:execution_id/cancel", h.CancelExecution)
		executions.POST("/:execution_id/retry", h.RetryExecution)
		executions.GET("/:execution_id/stream", h.StreamExecution)
	}
}
