package handlers

import (
	"context"
	"net/http"

	"moy_proekt/internal/userService"
	"moy_proekt/internal/web/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Получение всех пользователей
func (h *UserHandler) GetUsers(ctx context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := make(users.GetUsers200JSONResponse, len(usersList))
	for i, user := range usersList {
		response[i] = users.User{
			Id:       toIntPtr(int(user.ID)),
			Email:    toStringPtr(user.Email),
			Password: toStringPtr(user.Password),
		}
	}
	return response, nil
}

// Создание пользователя
func (h *UserHandler) PostUsers(ctx context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Проверяем, что поля email и password переданы в запросе
	if req.Body.Email == "" || req.Body.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "Email and Password are required",
		})
	}

	newUser := userService.User{
		Email:    req.Body.Email,
		Password: req.Body.Password,
	}

	createdUser, err := h.service.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return users.PostUsers201JSONResponse{
		Id:       toIntPtr(int(createdUser.ID)),
		Email:    toStringPtr(createdUser.Email),
		Password: toStringPtr(createdUser.Password),
	}, nil
}

// Обновление пользователя
func (h *UserHandler) PatchUsersId(ctx context.Context, req users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	updatedUser := userService.User{}

	// Проверяем, есть ли email и password в запросе перед разыменованием
	if req.Body.Email != nil {
		updatedUser.Email = *req.Body.Email
	}
	if req.Body.Password != nil {
		updatedUser.Password = *req.Body.Password
	}

	// Проверяем, что хотя бы одно поле для обновления передано
	if updatedUser.Email == "" && updatedUser.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "At least one field (Email or Password) must be provided",
		})
	}

	// Обновляем пользователя
	result, err := h.service.UpdateUserByID(uint(req.Id), updatedUser)
	if err != nil {
		return nil, err
	}

	return users.PatchUsersId200JSONResponse{
		Id:       toIntPtr(int(result.ID)),
		Email:    toStringPtr(result.Email),
		Password: toStringPtr(result.Password),
	}, nil
}

// Удаление пользователя
func (h *UserHandler) DeleteUsersId(ctx context.Context, req users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUserByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
