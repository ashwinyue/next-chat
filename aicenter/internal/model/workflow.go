// Package model 定义数据模型.
package model

import (
	"time"
)

// WorkflowStatus 工作流状态.
type WorkflowStatus string

const (
	WorkflowStatusDraft    WorkflowStatus = "draft"
	WorkflowStatusActive   WorkflowStatus = "active"
	WorkflowStatusInactive WorkflowStatus = "inactive"
)

// Workflow 工作流定义（对应 tgo-workflow）.
type Workflow struct {
	ID              string         `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ProjectID       string         `json:"project_id" gorm:"type:uuid;not null;index"`
	Name            string         `json:"name" gorm:"size:255;not null"`
	Description     string         `json:"description,omitempty" gorm:"type:text"`
	GraphDefinition JSONMap        `json:"graph_definition" gorm:"type:jsonb;not null"`
	Status          WorkflowStatus `json:"status" gorm:"size:20;not null;default:draft"`
	Version         int            `json:"version" gorm:"not null;default:1"`
	IsDefault       bool           `json:"is_default" gorm:"default:false"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at,omitempty" gorm:"index"`

	// 关联
	Nodes      []WorkflowNode      `json:"nodes,omitempty" gorm:"foreignKey:WorkflowID"`
	Executions []WorkflowExecution `json:"executions,omitempty" gorm:"foreignKey:WorkflowID"`
}

func (Workflow) TableName() string {
	return "workflows"
}

// NodeType 节点类型（对应 tgo-workflow 的节点类型）.
type NodeType string

const (
	NodeTypeStart      NodeType = "start"
	NodeTypeEnd        NodeType = "end"
	NodeTypeLLM        NodeType = "llm"
	NodeTypeAPI        NodeType = "api"
	NodeTypeCondition  NodeType = "condition"
	NodeTypeClassifier NodeType = "classifier"
	NodeTypeAgent      NodeType = "agent"
	NodeTypeTool       NodeType = "tool"
	NodeTypeRAG        NodeType = "rag"
	NodeTypeTimer      NodeType = "timer"
	NodeTypeWebhook    NodeType = "webhook"
	NodeTypeAnswer     NodeType = "answer"
)

// WorkflowNode 工作流节点.
type WorkflowNode struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WorkflowID string    `json:"workflow_id" gorm:"type:uuid;not null;index"`
	NodeID     string    `json:"node_id" gorm:"size:100;not null"`
	NodeType   NodeType  `json:"node_type" gorm:"size:50;not null"`
	Name       string    `json:"name" gorm:"size:255"`
	Config     JSONMap   `json:"config,omitempty" gorm:"type:jsonb"`
	Position   JSONMap   `json:"position,omitempty" gorm:"type:jsonb"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// 关联
	Workflow *Workflow      `json:"workflow,omitempty" gorm:"foreignKey:WorkflowID"`
	Edges    []WorkflowEdge `json:"edges,omitempty" gorm:"foreignKey:SourceNodeID"`
}

func (WorkflowNode) TableName() string {
	return "workflow_nodes"
}

// WorkflowEdge 工作流边（连接）.
type WorkflowEdge struct {
	ID           string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WorkflowID   string    `json:"workflow_id" gorm:"type:uuid;not null;index"`
	SourceNodeID string    `json:"source_node_id" gorm:"type:uuid;not null;index"`
	TargetNodeID string    `json:"target_node_id" gorm:"type:uuid;not null;index"`
	Condition    string    `json:"condition,omitempty" gorm:"type:text"`
	Config       JSONMap   `json:"config,omitempty" gorm:"type:jsonb"`
	CreatedAt    time.Time `json:"created_at"`

	// 关联
	Workflow   *Workflow     `json:"workflow,omitempty" gorm:"foreignKey:WorkflowID"`
	SourceNode *WorkflowNode `json:"source_node,omitempty" gorm:"foreignKey:SourceNodeID"`
	TargetNode *WorkflowNode `json:"target_node,omitempty" gorm:"foreignKey:TargetNodeID"`
}

func (WorkflowEdge) TableName() string {
	return "workflow_edges"
}

// ExecutionStatus 执行状态.
type ExecutionStatus string

const (
	ExecutionStatusPending   ExecutionStatus = "pending"
	ExecutionStatusRunning   ExecutionStatus = "running"
	ExecutionStatusCompleted ExecutionStatus = "completed"
	ExecutionStatusFailed    ExecutionStatus = "failed"
	ExecutionStatusCancelled ExecutionStatus = "cancelled"
	ExecutionStatusPaused    ExecutionStatus = "paused"
)

// WorkflowExecution 工作流执行记录.
type WorkflowExecution struct {
	ID          string          `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WorkflowID  string          `json:"workflow_id" gorm:"type:uuid;not null;index"`
	SessionID   string          `json:"session_id,omitempty" gorm:"type:uuid;index"`
	Status      ExecutionStatus `json:"status" gorm:"size:20;not null;default:pending"`
	Input       JSONMap         `json:"input,omitempty" gorm:"type:jsonb"`
	Output      JSONMap         `json:"output,omitempty" gorm:"type:jsonb"`
	Error       string          `json:"error,omitempty" gorm:"type:text"`
	StartedAt   *time.Time      `json:"started_at,omitempty"`
	CompletedAt *time.Time      `json:"completed_at,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`

	// 关联
	Workflow   *Workflow               `json:"workflow,omitempty" gorm:"foreignKey:WorkflowID"`
	NodeStates []WorkflowNodeExecution `json:"node_states,omitempty" gorm:"foreignKey:ExecutionID"`
}

func (WorkflowExecution) TableName() string {
	return "workflow_executions"
}

// WorkflowNodeExecution 节点执行状态.
type WorkflowNodeExecution struct {
	ID          string          `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ExecutionID string          `json:"execution_id" gorm:"type:uuid;not null;index"`
	NodeID      string          `json:"node_id" gorm:"size:100;not null"`
	Status      ExecutionStatus `json:"status" gorm:"size:20;not null;default:pending"`
	Input       JSONMap         `json:"input,omitempty" gorm:"type:jsonb"`
	Output      JSONMap         `json:"output,omitempty" gorm:"type:jsonb"`
	Error       string          `json:"error,omitempty" gorm:"type:text"`
	StartedAt   *time.Time      `json:"started_at,omitempty"`
	CompletedAt *time.Time      `json:"completed_at,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`

	// 关联
	Execution *WorkflowExecution `json:"execution,omitempty" gorm:"foreignKey:ExecutionID"`
}

func (WorkflowNodeExecution) TableName() string {
	return "workflow_node_executions"
}

// WorkflowVariable 工作流变量（对应 tgo-workflow 的变量系统）.
type WorkflowVariable struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WorkflowID string    `json:"workflow_id" gorm:"type:uuid;not null;index"`
	Name       string    `json:"name" gorm:"size:100;not null"`
	Type       string    `json:"type" gorm:"size:50;not null"` // string, number, boolean, object, array
	Default    JSONMap   `json:"default,omitempty" gorm:"type:jsonb"`
	Required   bool      `json:"required" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`

	// 关联
	Workflow *Workflow `json:"workflow,omitempty" gorm:"foreignKey:WorkflowID"`
}

func (WorkflowVariable) TableName() string {
	return "workflow_variables"
}
