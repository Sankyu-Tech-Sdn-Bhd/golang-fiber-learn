package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Position string `json:"position" xml:"position" form:"position"`
}

func createUser(ctx *fiber.Ctx) error {

	data := new(User)

	if err := ctx.BodyParser(data); err != nil {
		return err
	}
	user := User{
		Name:     data.Name,
		Position: data.Position,
	}

	return ctx.Status(fiber.StatusOK).JSON(user)

}

func main() {
	app := fiber.New()

	//post request without parameter

	// app.Post("/api/register", func(c *fiber.Ctx) error {
	// 	return c.SendString("I'm a POST request!")
	// })

	//post request with parameter

	//example send data

	// {
	// 	"Name" : "warfee",
	// 	"Position" : "Dev"
	// }

	app.Post("/api/register", createUser)

	app.Listen(":3000")
}
