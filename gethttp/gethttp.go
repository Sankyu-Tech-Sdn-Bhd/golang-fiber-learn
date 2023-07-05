package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//get request without parameter

	// app.Get("/api/list", func(c *fiber.Ctx) error {
	// 	return c.SendString("I'm a GET request!")
	// })

	// get request with parameter

	app.Get("/user/:name/position/:position", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("position"))
		return nil
	})

	app.Listen(":3000")
}
