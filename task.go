package main

import "gorm.io/gorm"

// Структура для записи в БД
type Task struct {
	gorm.Model        // Добавляет ID, CreatedAt, UpdatedAt
	Message    string `json:"message"` // Поле "message" из JSON
	IsDone     bool   `json:"is_done"` // Поле "is_done" из JSON
}
