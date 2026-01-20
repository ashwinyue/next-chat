// Package biz 提供业务逻辑层.
package biz

import (
	"github.com/next-chat/aicenter/internal/biz/agent"
	"github.com/next-chat/aicenter/internal/biz/auth"
	"github.com/next-chat/aicenter/internal/biz/knowledge"
	"github.com/next-chat/aicenter/internal/biz/mcp"
	"github.com/next-chat/aicenter/internal/biz/provider"
	"github.com/next-chat/aicenter/internal/biz/session"
	"github.com/next-chat/aicenter/internal/biz/settings"
	"github.com/next-chat/aicenter/internal/biz/tenant"
	"github.com/next-chat/aicenter/internal/biz/websearch"
	"github.com/next-chat/aicenter/internal/pkg/agent/factory"
	"github.com/next-chat/aicenter/internal/store"
	"github.com/cloudwego/eino/components/embedding"
)

// Biz 业务层聚合接口.
type Biz interface {
	Agents() agent.AgentBiz
	AgentConfig() agent.ConfigBiz
	Providers() provider.Biz
	MCP() mcp.Biz
	WebSearch() websearch.Biz
	Settings() settings.Biz
	Sessions() session.SessionBiz
	Knowledge() knowledge.Biz
	Tenants() tenant.Biz
	Auth() auth.Biz
}

type biz struct {
	agentBiz       agent.AgentBiz
	agentConfigBiz agent.ConfigBiz
	providerBiz    provider.Biz
	mcpBiz         mcp.Biz
	webSearchBiz   websearch.Biz
	settingsBiz    settings.Biz
	sessionBiz     session.SessionBiz
	knowledgeBiz   knowledge.Biz
	tenantBiz      tenant.Biz
	authBiz        auth.Biz
}

// NewBiz 创建业务层实例.
func NewBiz(store store.Store, agentFactory *factory.AgentFactory, embedder embedding.Embedder) Biz {
	return &biz{
		agentBiz:       agent.NewAgentBiz(store, agentFactory),
		agentConfigBiz: agent.NewConfigBiz(store),
		providerBiz:    provider.NewBiz(store),
		mcpBiz:         mcp.NewBiz(store),
		webSearchBiz:   websearch.NewBiz(store),
		settingsBiz:    settings.NewBiz(store),
		sessionBiz:     session.NewSessionBiz(store),
		knowledgeBiz:   knowledge.NewBiz(store, embedder),
		tenantBiz:      tenant.NewBiz(store),
		authBiz:        auth.NewBiz(store, nil),
	}
}

func (b *biz) Agents() agent.AgentBiz {
	return b.agentBiz
}

func (b *biz) Sessions() session.SessionBiz {
	return b.sessionBiz
}

func (b *biz) AgentConfig() agent.ConfigBiz {
	return b.agentConfigBiz
}

func (b *biz) Providers() provider.Biz {
	return b.providerBiz
}

func (b *biz) MCP() mcp.Biz {
	return b.mcpBiz
}

func (b *biz) WebSearch() websearch.Biz {
	return b.webSearchBiz
}

func (b *biz) Settings() settings.Biz {
	return b.settingsBiz
}

func (b *biz) Knowledge() knowledge.Biz {
	return b.knowledgeBiz
}

func (b *biz) Tenants() tenant.Biz {
	return b.tenantBiz
}

func (b *biz) Auth() auth.Biz {
	return b.authBiz
}
