package taskService

import "gorm.io/gorm"

// Task представляет задачу с ID, описанием и статусом выполнения
type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"index"`
}

// Migrate выполняет автоматическую миграцию структуры Task в базе данных
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
