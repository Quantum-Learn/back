package storage

import (
	api "MVP_project/internal/api"
	"context"
	"errors"
)

type UserStore interface {
	Create(ctx context.Context, user api.User) (int, error)
	GetByEmail(ctx context.Context, email string) (*api.User, error)
	GetByID(ctx context.Context, id int) (*api.User, error) // <-- int
}

type InMemoryUserStore struct {
	byEmail map[string]api.User
	lastID  int
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{byEmail: make(map[string]api.User)}
}

func (s *InMemoryUserStore) Create(ctx context.Context, u api.User) (int, error) {
	s.lastID++
	// учти: у oapi-codegen поля, скорее всего, указатели
	id := s.lastID
	u.Id = &id
	if *u.Email == "" {
		return 0, errors.New("empty email")
	}
	if _, exists := s.byEmail[*u.Email]; exists {
		return 0, errors.New("email already exists")
	}
	s.byEmail[*u.Email] = u
	return id, nil
}

func (s *InMemoryUserStore) GetByEmail(ctx context.Context, email string) (*api.User, error) {
	u, ok := s.byEmail[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

func (s *InMemoryUserStore) GetByID(ctx context.Context, id int) (*api.User, error) {
	for _, u := range s.byEmail {
		if u.Id != nil && *u.Id == id {
			uu := u
			return &uu, nil
		}
	}
	return nil, errors.New("user not found")
}
