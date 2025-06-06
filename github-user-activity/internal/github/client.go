package github

import (
	"encoding/json"
	"fmt"
	"githubuseractivity/config"
	"githubuseractivity/internal/requester"
	"io"
)

const (
	hostname = "https://api.github.com"
)

type Client struct {
	requester requester.Requester
}

func NewClient(requester requester.Requester) *Client {
	return &Client{
		requester: requester,
	}
}

func (c *Client) GetListEventsUser(username string) ([]UserEvent, error) {
	url := fmt.Sprintf("%s/users/%s/events", hostname, username)

	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", config.GetGitHubToken()),
	}

	resp, err := c.requester.Get(url, header)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error getting user events: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var events []UserEvent
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return events, nil
}
