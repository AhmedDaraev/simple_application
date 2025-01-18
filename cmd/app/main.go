package main

import (
	"log"
	"moy_proekt/internal/database"
	"moy_proekt/internal/handlers"
	"moy_proekt/internal/taskService"
	"moy_proekt/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	if err := database.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Настройка репозитория, сервиса и хендлера
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация обработчиков
	strictHandler := tasks.NewStrictHandler(handler, nil) // Нужно добавить middleware, если они понадобятся
	tasks.RegisterHandlers(e, strictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
