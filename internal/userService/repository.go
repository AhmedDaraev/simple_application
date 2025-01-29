package userService

import (
	"fmt"

	"gorm.io/gorm"
)

// UserRepository определяет интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, updatedUser User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) UpdateUserByID(id uint, updatedUser User) (User, error) {
	var user User

	// Проверка, существует ли такой email у другого пользователя
	if updatedUser.Email != "" {
		var existingUser User
		err := r.db.Where("email = ? AND id != ?", updatedUser.Email, id).First(&existingUser).Error
		if err == nil {
			return User{}, fmt.Errorf("email %s уже используется другим пользователем", updatedUser.Email)
		}
	}

	// Выполняем обновление
	result := r.db.Model(&User{}).Where("id = ?", id).Updates(updatedUser)
	if result.Error != nil {
		return User{}, result.Error
	}

	// Загружаем обновлённый объект из базы данных
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
