package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//template engine
	engine := html.New("./views", ".html")

	//create fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//static files
	app.Static("/", "./public")

	//get command
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"hello": "world",
		})
	})

	//start server
	app.Listen(":3000")
}
