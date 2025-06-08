package cli

import "githubuseractivity/internal/github"

type CLI struct {
	requester github.Requester
}

func New(requester github.Requester) *CLI {
	return &CLI{
		requester: requester,
	}
}

func (c *CLI) GithubUserActivities