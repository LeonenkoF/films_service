package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
)

func main() {
	app := fiber.New()

	fmt.Println("The server is running", log.String("port", "8080"))
	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Server startup error: %v", err)
	}
}
