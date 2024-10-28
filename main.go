package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/config"
	"github.com/shiibs/fitness-app/db"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Fitness-Tracker",
		ServerHeader: "Fiber",
	})

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
