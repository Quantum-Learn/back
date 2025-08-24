package storage

import (
	api "MVP_project/internal/api"
	"context"
)

// UserStore описывает работу с пользователями
type UserStore interface {
	Create(ctx context.Context, user api.User) (int64, error)
	GetByEmail(ctx context.Context, email string) (*api.User, error)
	GetByID(ctx context.Context, id int64) (*api.User, error)
}
