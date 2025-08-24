package storage

import (
	api "MVP_project/internal/api"
	"context"
	"database/sql"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateUser(ctx context.Context, email, passwordHash string) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx,
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id",
		email, passwordHash,
	).Scan(&id)
	return id, err
}

func (s *PostgresStorage) GetUserByEmail(ctx context.Context, email string) (*api.User, error) {
	row := s.db.QueryRowContext(ctx,
		"SELECT id, email, password_hash FROM users WHERE email = $1",
		email,
	)
	u := &api.User{}
	err := row.Scan(&u.Email, &u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}
