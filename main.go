package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("i'm middleware before processing request")
		err := c.Next()
		fmt.Println("i'm middleware after processing request")
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if fiber.IsChild() {
		fmt.Println("i'm child process")
	} else {
		fmt.Println("i'm parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
