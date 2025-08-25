package handlers

import (
	api "MVP_project/internal/api"
	"MVP_project/internal/auth"
	"MVP_project/internal/storage"
	"context"
	"errors"
	"fmt"
)

type AuthHandler struct {
	store storage.UserStore
}

func NewAuthHandler(s storage.UserStore) *AuthHandler {
	return &AuthHandler{store: s}
}

func (h *AuthHandler) Login(ctx context.Context, req api.LoginRequest) (api.AuthResponse, error) {
	// 1) найти пользователя по email
	user, err := h.store.GetByEmail(ctx, req.Email)
	if err != nil {
		return api.AuthResponse{}, fmt.Errorf("user not found: %w", err)
	}

	// 2) проверить, что хэш есть
	if user.PasswordHash == nil {
		return api.AuthResponse{}, errors.New("user has no password hash")
	}

	// 3) сравнить пароль с хэшем
	if !auth.CheckPassword(req.Password, *user.PasswordHash) {
		return api.AuthResponse{}, fmt.Errorf("invalid credentials")
	}

	// 4) сгенерировать JWT — лучше по userID
	// проверь точное имя поля ID в сгенерённой модели (часто это Id или ID)
	var uid int
	switch {
	case user.Id != nil:
		uid = *user.Id
	default:
		return api.AuthResponse{}, errors.New("user id is nil")
	}

	token, err := auth.GenerateToken(uid)
	if err != nil {
		return api.AuthResponse{}, fmt.Errorf("token error: %w", err)
	}

	// 5) вернуть ответ
	return api.AuthResponse{Token: &token}, nil
}
