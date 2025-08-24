package handlers

import (
	api "MVP_project/internal/api"
	"MVP_project/internal/auth"
	"MVP_project/internal/storage"
	"context"
	"fmt"
)

// Обработчик аутентификации
type AuthHandler struct {
	store storage.UserStore
}

// Логин пользователя. Проверяет учетные данные и возвращает JWT.
func (h *AuthHandler) Login(ctx context.Context, req api.LoginRequest) (api.AuthResponse, error) {
	user, err := h.store.GetByEmail(ctx, req.Email)
	if err != nil {
		return api.AuthResponse{}, err
	}

	//Проверить как происходит хэширование
	if !auth.CheckPassword(req.Password, *user.PasswordHash) {
		return api.AuthResponse{}, fmt.Errorf("invalid credentials")
	}

	token, err := auth.GenerateToken(*user.Name)
	if err != nil {
		return api.AuthResponse{}, err
	}

	return api.AuthResponse{Token: &token}, nil
}
