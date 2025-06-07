package github

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockRequester struct {
	get func(url string, headers map[string]string) (*http.Response, error)
}

func (m *mockRequester) Get(url string, headers map[string]string) (*http.Response, error) {
	return m.get(url, headers)
}

func NewMockClient(doFunc func(url string, headers map[string]string) (*http.Response, error)) *Client {
	return NewClient(&mockRequester{get: doFunc})
}

type errorReadCloser struct{}

func (e errorReadCloser) Read([]byte) (int, error) {
	return 0, errors.New("error")
}
func (e errorReadCloser) Close() error {
	return nil
}

func TestGetListEventsUser(t *testing.T) {
	t.Run("if requester returns error then return error", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return nil, errors.New("error")
		})

		events, err := client.GetListEventsUser("test")
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}

		if events != nil {
			t.Fatalf("Expected nil events, got %v", events)
		}
	})

	t.Run("if requester returns status code != 200 then return error getting user events and status", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return &http.Response{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error content"))}, nil
		})

		events, err := client.GetListEventsUser("test")
		if err.Error() != "error getting user events: 500" {
			t.Fatalf("Expected error getting user events: 500, got %s", err.Error())
		}

		if events != nil {
			t.Fatalf("Expected nil events, got %v", events)
		}
	})

	t.Run("if requester returns response body with error then return error", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       errorReadCloser{},
			}, nil
		})

		events, err := client.GetListEventsUser("test")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		if events != nil {
			t.Fatalf("Expected nil events, got %v", events)
		}
	})

	t.Run("if requester returns invalid json then return error", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("invalid json"))}, nil
		})

		events, err := client.GetListEventsUser("test")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		if events != nil {
			t.Fatalf("Expected nil events, got %v", events)
		}
	})

	t.Run("if requester returns success then return no error", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`[{"type": "PushEvent", "id": 12345}]`))}, nil
		})

		events, err := client.GetListEventsUser("test")
		if err != nil {
			t.Fatalf("Expected not error, got %s", err.Error())
		}

		if len(events) != 1 || events[0].Type != "PushEvent" {
			t.Fatalf("Unexpected events: %+v", events)
		}
	})
}
