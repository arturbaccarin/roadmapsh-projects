package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	githubToken string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetGitHubToken() string {
	if githubToken == "" {
		githubToken = os.Getenv("GITHUB_TOKEN")
	}

	return githubToken
}
