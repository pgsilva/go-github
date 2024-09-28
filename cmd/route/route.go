package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-github/internal/handler"
)

const (
	ApiV1BasePath  = "/v1"
	GithubBasePath = "/gohub"
)

func EnableRoutes(app *fiber.App) {
	v1 := app.Group(ApiV1BasePath)
	{
		v1.Get("health", handler.HealthCheck)
	}

	gohubRoutes := v1.Group(GithubBasePath)
	{
		gohubRoutes.Get("users", handler.SearchUsers)
	}

}
