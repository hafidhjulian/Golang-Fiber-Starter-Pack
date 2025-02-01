package routers

import (
	"golang-fiber-starterpack/handlers"

	"github.com/gofiber/fiber/v2"
)

func ExampleRoute(app fiber.Router) {
	app.Post("/post/example", handlers.PostExample)
	app.Get("/get/example", handlers.GetExample)
}
