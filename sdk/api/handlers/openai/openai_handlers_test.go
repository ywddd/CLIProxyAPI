package openai

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestNormalizeOpenAIRequestJSONForRoutingRepairsWindowsPathEscapes(t *testing.T) {
	rawJSON := []byte(`{"model":"kimi-k2.7-code","messages":[{"role":"user","content":"path C:\Users\bad"}],"stream":true}`)

	repaired, err := normalizeOpenAIRequestJSONForRouting(rawJSON)
	if err != nil {
		t.Fatalf("normalizeOpenAIRequestJSONForRouting() error = %v", err)
	}
	if !gjson.ValidBytes(repaired) {
		t.Fatalf("repaired payload is not valid JSON: %s", repaired)
	}
	if got := gjson.GetBytes(repaired, "model").String(); got != "kimi-k2.7-code" {
		t.Fatalf("model after repair = %q, want %q", got, "kimi-k2.7-code")
	}
	if got := gjson.GetBytes(repaired, "messages.0.content").String(); got != `path C:\Users\bad` {
		t.Fatalf("content after repair = %q", got)
	}
}

func TestNormalizeOpenAIRequestJSONForRoutingRejectsUnrecoverableInvalidJSON(t *testing.T) {
	rawJSON := []byte(`{"model":"kimi-k2.7-code","messages":[}`)

	if _, err := normalizeOpenAIRequestJSONForRouting(rawJSON); err == nil {
		t.Fatal("normalizeOpenAIRequestJSONForRouting() error = nil, want invalid JSON error")
	}
}
