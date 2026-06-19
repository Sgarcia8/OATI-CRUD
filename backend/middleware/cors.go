package middleware

import (
	"net/http"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func CorsFilter(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	allowOrigin := resolveAllowOrigin(origin)

	if allowOrigin != "" {
		ctx.Output.Header("Access-Control-Allow-Origin", allowOrigin)
	}
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	ctx.Output.Header("Access-Control-Max-Age", "86400")

	if ctx.Input.Method() == http.MethodOptions {
		ctx.Output.SetStatus(http.StatusNoContent)
		_ = ctx.Output.Body([]byte{})
		return
	}
}

func resolveAllowOrigin(origin string) string {
	if beego.BConfig.RunMode == "dev" {
		return "*"
	}

	if origin == "" {
		return ""
	}

	for _, allowed := range allowedOrigins() {
		if origin == allowed {
			return origin
		}
	}
	return ""
}

func allowedOrigins() []string {
	raw := os.Getenv("CORS_ALLOWED_ORIGINS")
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			origins = append(origins, trimmed)
		}
	}
	return origins
}
