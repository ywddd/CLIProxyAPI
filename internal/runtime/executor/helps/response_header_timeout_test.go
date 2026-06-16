package helps

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWithResponseHeaderTimeoutTimesOutBeforeHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	client := WithResponseHeaderTimeout(server.Client(), 10*time.Millisecond)
	resp, err := client.Get(server.URL)
	if err == nil {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
		t.Fatal("expected response header timeout error")
	}
}

func TestWithResponseHeaderTimeoutDoesNotLimitBodyAfterHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatal("response writer does not implement http.Flusher")
		}
		w.WriteHeader(http.StatusOK)
		flusher.Flush()
		time.Sleep(30 * time.Millisecond)
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	client := WithResponseHeaderTimeout(server.Client(), 10*time.Millisecond)
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected request error after headers: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(body) != "ok" {
		t.Fatalf("body = %q, want ok", string(body))
	}
}
