// Package http 提供 HTTP Handler 层.
package http

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/next-chat/aicenter/internal/biz/auth"
	"github.com/next-chat/aicenter/internal/model"
)

// ContextKey 上下文键类型.
type ContextKey string

const (
	ContextKeyUserID   ContextKey = "user_id"
	ContextKeyTenantID ContextKey = "tenant_id"
	ContextKeyUser     ContextKey = "user"
	ContextKeyRole     ContextKey = "role"
	ContextKeyClaims   ContextKey = "claims"
)

// 无需认证的 API 列表.
var noAuthAPIs = map[string][]string{
	"/health":               {"GET"},
	"/api/v1/auth/register": {"POST"},
	"/api/v1/auth/login":    {"POST"},
}

// isNoAuthAPI 检查请求是否在无需认证的 API 列表中.
func isNoAuthAPI(path string, method string) bool {
	for api, methods := range noAuthAPIs {
		if strings.HasSuffix(api, "*") {
			if strings.HasPrefix(path, strings.TrimSuffix(api, "*")) && slices.Contains(methods, method) {
				return true
			}
		} else if path == api && slices.Contains(methods, method) {
			return true
		}
	}
	return false
}

// AuthMiddleware JWT 认证中间件.
func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// 检查是否在无需认证的 API 列表中
		if isNoAuthAPI(c.Request.URL.Path, c.Request.Method) {
			c.Next()
			return
		}

		// 尝试 JWT Token 认证
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := h.biz.Auth().ValidateToken(c.Request.Context(), token)
			if err == nil && claims != nil {
				// JWT 认证成功，存储信息到上下文
				c.Set(string(ContextKeyUserID), claims.UserID)
				c.Set(string(ContextKeyTenantID), claims.TenantID)
				c.Set(string(ContextKeyRole), claims.Role)
				c.Set(string(ContextKeyClaims), claims)

				// 同时存储到 request context
				ctx := c.Request.Context()
				ctx = context.WithValue(ctx, ContextKeyUserID, claims.UserID)
				ctx = context.WithValue(ctx, ContextKeyTenantID, claims.TenantID)
				ctx = context.WithValue(ctx, ContextKeyRole, claims.Role)
				ctx = context.WithValue(ctx, ContextKeyClaims, claims)
				c.Request = c.Request.WithContext(ctx)

				c.Next()
				return
			}
		}

		// 尝试 API Key 认证
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != "" {
			key, err := h.biz.Tenants().ValidateAPIKey(c.Request.Context(), apiKey)
			if err == nil && key != nil && key.Status == model.APIKeyStatusActive {
				// API Key 认证成功
				c.Set(string(ContextKeyTenantID), key.TenantID)
				c.Set(string(ContextKeyRole), model.UserRoleUser) // API Key 默认为普通用户权限

				ctx := c.Request.Context()
				ctx = context.WithValue(ctx, ContextKeyTenantID, key.TenantID)
				ctx = context.WithValue(ctx, ContextKeyRole, model.UserRoleUser)
				c.Request = c.Request.WithContext(ctx)

				c.Next()
				return
			}
		}

		// 没有提供有效的认证信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: missing or invalid authentication"})
		c.Abort()
	}
}

// RequireAuth 要求认证的中间件（简化版，用于特定路由组）.
func (h *Handler) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试 JWT Token 认证
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: missing token"})
			c.Abort()
			return
		}

		claims, err := h.biz.Auth().ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid token"})
			c.Abort()
			return
		}

		// 存储信息到上下文
		c.Set(string(ContextKeyUserID), claims.UserID)
		c.Set(string(ContextKeyTenantID), claims.TenantID)
		c.Set(string(ContextKeyRole), claims.Role)
		c.Set(string(ContextKeyClaims), claims)

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, ContextKeyUserID, claims.UserID)
		ctx = context.WithValue(ctx, ContextKeyTenantID, claims.TenantID)
		ctx = context.WithValue(ctx, ContextKeyRole, claims.Role)
		ctx = context.WithValue(ctx, ContextKeyClaims, claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// RequireAdmin 要求管理员权限的中间件.
func (h *Handler) RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(string(ContextKeyRole))
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: not authenticated"})
			c.Abort()
			return
		}

		userRole, ok := role.(model.UserRole)
		if !ok || userRole != model.UserRoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: admin role required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRoles 要求指定角色的中间件.
func (h *Handler) RequireRoles(roles ...model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(string(ContextKeyRole))
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: not authenticated"})
			c.Abort()
			return
		}

		userRole, ok := role.(model.UserRole)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: invalid role"})
			c.Abort()
			return
		}

		for _, r := range roles {
			if userRole == r {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
		c.Abort()
	}
}

// GetUserIDFromContext 从上下文获取用户 ID.
func GetUserIDFromContext(c *gin.Context) string {
	if userID, exists := c.Get(string(ContextKeyUserID)); exists {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}

// GetTenantIDFromContext 从上下文获取租户 ID.
func GetTenantIDFromContext(c *gin.Context) string {
	if tenantID, exists := c.Get(string(ContextKeyTenantID)); exists {
		if id, ok := tenantID.(string); ok {
			return id
		}
	}
	return ""
}

// GetClaimsFromContext 从上下文获取 JWT Claims.
func GetClaimsFromContext(c *gin.Context) *auth.Claims {
	if claims, exists := c.Get(string(ContextKeyClaims)); exists {
		if cl, ok := claims.(*auth.Claims); ok {
			return cl
		}
	}
	return nil
}

// GetRoleFromContext 从上下文获取用户角色.
func GetRoleFromContext(c *gin.Context) model.UserRole {
	if role, exists := c.Get(string(ContextKeyRole)); exists {
		if r, ok := role.(model.UserRole); ok {
			return r
		}
	}
	return ""
}
