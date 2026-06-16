package management

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/router-for-me/CLIProxyAPI/v7/internal/config"
	"github.com/router-for-me/CLIProxyAPI/v7/internal/runtime/executor/helps"
)

func TestAuthFileTestExecutionContextUsesFirstConfiguredClientAPIKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/v0/management/auth-files/test", nil)

	h := &Handler{cfg: &config.Config{SDKConfig: config.SDKConfig{APIKeys: []string{"  first-key  ", "second-key"}}}}
	ctx := h.authFileTestExecutionContext(c)

	if got := helps.APIKeyFromContext(ctx); got != "first-key" {
		t.Fatalf("APIKeyFromContext() = %q, want %q", got, "first-key")
	}
	rawProvider, ok := c.Get("accessProvider")
	if !ok || rawProvider != "management-auth-test" {
		t.Fatalf("accessProvider = %v, %v; want management-auth-test", rawProvider, ok)
	}
}

func TestAuthFileTestExecutionContextPreservesExistingClientAPIKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/v0/management/auth-files/test", nil)
	c.Set("userApiKey", "caller-key")

	h := &Handler{cfg: &config.Config{SDKConfig: config.SDKConfig{APIKeys: []string{"first-key"}}}}
	ctx := h.authFileTestExecutionContext(c)

	if got := helps.APIKeyFromContext(ctx); got != "caller-key" {
		t.Fatalf("APIKeyFromContext() = %q, want %q", got, "caller-key")
	}
}
