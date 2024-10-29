package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/fitness-app/config"
	"github.com/shiibs/fitness-app/db"
	"github.com/shiibs/fitness-app/handlers"
	"github.com/shiibs/fitness-app/middlewares"
	"github.com/shiibs/fitness-app/service"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Fitness-Tracker",
		ServerHeader: "Fiber",
	})

	server := app.Group("/api")

	// Auth
	authService := service.NewAuthService(db)
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	// Protected Routes
	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// User
	userService := service.NewUserService(db)
	handlers.NewUserHandler(privateRoutes, userService)

	// Food Entry
	foodService := service.NewFoodService(db)
	handlers.NewFoodHandler(privateRoutes, foodService)

	// Workout entry
	workoutService := service.NewWorkoutService(db)
	handlers.NewWorkoutHandler(privateRoutes, workoutService)

	app.Listen(fmt.Sprintf(":%s", envConfig.ServerPort))
}
