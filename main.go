package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://ayo.so/moby")
	})
	
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return.SendString("moby mobyhub apex nocomplymoby moby twitter moby socials")	
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(fmt.Sprintf(":%s", port))
}
