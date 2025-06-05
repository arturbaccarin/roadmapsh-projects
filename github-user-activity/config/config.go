package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	githubToken string
)

func LoadEnvs() error {
	return godotenv.Load()
}

func GetGitHubToken() string {
	if githubToken == "" {
		githubToken = os.Getenv("GITHUB_TOKEN")
	}

	return githubToken
}
