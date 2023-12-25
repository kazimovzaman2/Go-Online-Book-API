package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazimovzaman2/online_book/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Get("/book/:id", controllers.GetBook)
	route.Get("/token/new", controllers.GetNewAccessToken)
}
