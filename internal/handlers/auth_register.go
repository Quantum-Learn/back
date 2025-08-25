package handlers

import (
	api "MVP_project/internal/api"
	"MVP_project/internal/auth"
	"context"
	"errors"
	"fmt"
)

func (h *AuthHandler) Register(ctx context.Context, req api.RegisterRequest) (api.User, error) {
	if req.Email == "" || req.Password == "" {
		return api.User{}, errors.New("email and password required")
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		return api.User{}, fmt.Errorf("hash error: %w", err)
	}

	u := api.User{
		Email:        &req.Email,
		Name:         &req.Name,
		PasswordHash: &hash, // важно: храним ХЭШ
	}
	id, err := h.store.Create(ctx, u)
	if err != nil {
		return api.User{}, err
	}
	idClone := int(id)
	u.Id = &idClone

	// Не возвращаем PasswordHash наружу
	u.PasswordHash = nil
	return u, nil
}
