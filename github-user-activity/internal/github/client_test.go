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

func TestGetListEventsUser(t *testing.T) {
	t.Run("if requester returns error then return error", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return nil, errors.New("error")
		})

		_, err := client.GetListEventsUser("test")
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	t.Run("if requester returns status code != 200 then return error getting user events and status", func(t *testing.T) {
		client := NewMockClient(func(url string, headers map[string]string) (*http.Response, error) {
			return &http.Response{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error content"))}, nil
		})

		_, err := client.GetListEventsUser("test")
		if err.Error() != "error getting user events: 500" {
			t.Fatalf("Expected error getting user events: 500, got %s", err.Error())
		}
	})
}
