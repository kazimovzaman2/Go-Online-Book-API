package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazimovzaman2/online_book/app/controllers"
	"github.com/kazimovzaman2/online_book/pkg/middleware"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)  // update one book by ID
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
