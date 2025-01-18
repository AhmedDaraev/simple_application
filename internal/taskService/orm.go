package taskService

import (
	"gorm.io/gorm"
)

// Task представляет задачу с ID, описанием и статусом выполнения
type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"` // Уникальный идентификатор задачи
	Task   string `json:"task"`                 // Описание задачи
	IsDone bool   `json:"is_done"`              // Статус выполнения
}

// Migrate выполняет автоматическую миграцию структуры Task в базе данных
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
