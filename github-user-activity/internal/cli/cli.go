package cli

import (
	"fmt"
	"githubuseractivity/internal/github"
)

type CLI struct {
	requester github.Requester
}

func New(requester github.Requester) *CLI {
	return &CLI{
		requester: requester,
	}
}

func (c *CLI) GithubUserActivities(username string) error {
	events, err := c.requester.GetListEventsUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", events)

	return nil
}
