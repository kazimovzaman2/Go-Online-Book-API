package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazimovzaman2/online_book/pkg/configs"
	"github.com/kazimovzaman2/online_book/pkg/middleware"
	"github.com/kazimovzaman2/online_book/pkg/routes"
	"github.com/kazimovzaman2/online_book/pkg/utils"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
