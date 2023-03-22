package main

import "github.com/gofiber/fiber/v2"

func main() {
	//create fiber app
	app := fiber.New()

	//static files
	app.Static("/", "./public")

	//get command
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is a chat app!")
	})

	//start server
	app.Listen(":3000")
}
