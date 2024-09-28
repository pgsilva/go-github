package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-github/cmd/route"
	"github.com/pgsilva/go-github/pkg/config"
)

func main() {
	envVar()
	fiberApp()
}

func envVar() {
	slog.Info("Starting the environment variables...")

	if err := config.Env(); err != nil {
		slog.Error("Error in loading the environment variables")
		return
	}
}

func fiberApp() {
	slog.Info("Starting the Fiber app...")

	app := fiber.New()

	route.EnableRoutes(app)

	slog.Info("Fiber app is running on port " + config.Port)

	if err := app.Listen(":" + config.Port); err != nil {
		slog.Error("Failed to start the application", "err", err)
	}
}
