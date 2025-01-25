package main

import (
	"log"
	"moy_proekt/internal/database"
	"moy_proekt/internal/handlers"
	"moy_proekt/internal/taskService"
	"moy_proekt/internal/userService"
	"moy_proekt/internal/web/tasks"
	"moy_proekt/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Task dependencies
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	// User dependencies
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Echo setup
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Task routes
	taskStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	// User routes
	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	// Start server
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
