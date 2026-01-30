package netutil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drawmoon/xrc/pkg/netutil"
)

func TestPing_Success(t *testing.T) {
	// Create a test server that returns 204 No Content
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := &http.Client{}
	times := 1

	result := netutil.Ping(client, server.URL, times)

	// Should return a non-negative value (average time in ms, 0 if very fast)
	if result < 0 {
		t.Errorf("Expected non-negative ping time, got %d", result)
	}
}

func TestPing_Failure(t *testing.T) {
	// Create a test server that returns 200 OK (not 204)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	client := &http.Client{}
	times := 1

	result := netutil.Ping(client, server.URL, times)

	// Should return -1 on failure
	if result != -1 {
		t.Errorf("Expected -1 on failure, got %d", result)
	}
}
