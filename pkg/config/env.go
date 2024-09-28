package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port         string
	GitHubApiUrl string
)

func Env() error {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "err", err)
	}

	GitHubApiUrl = os.Getenv("API_GITHUB_URL")
	Port = os.Getenv("PORT")

	return nil
}
