package main

import (
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	AnthropicAPIURL = "https://api.anthropic.com/"
	DefaultTimeout  = 5 * time.Minute
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  DefaultTimeout,
		WriteTimeout: DefaultTimeout,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	app.All("/anthropic/*", proxyAnthropicRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func proxyAnthropicRequest(c *fiber.Ctx) error {
	url := AnthropicAPIURL + c.Params("*")
	if err := proxy.Do(c, url); err != nil {
		return err
	}

	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
	c.Response().Header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

	return nil
}
