package github

import "net/http"

func NewClient(request *http.Request) *Client {
	return &Client{
		request: request,
	}
}

func (c *Client) GetListEventsUser(username string) ([]UserEvent, error) {
	return nil, nil
}
