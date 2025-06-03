package nethttp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNetHTTP_Get(t *testing.T) {
	// Setup test HTTP server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assert headers
		if r.Header.Get("Authorization") != "Bearer testtoken" {
			t.Errorf("Expected Authorization header 'Bearer testtoken', got '%s'", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Create NetHTTP instance with test server client
	n := New()

	// Call Get method
	headers := map[string]string{
		"Authorization": "Bearer testtoken",
	}
	resp, err := n.Get(testServer.URL, headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	// Validate response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "ok") {
		t.Errorf("Expected body to contain 'ok', got '%s'", string(body))
	}
}
