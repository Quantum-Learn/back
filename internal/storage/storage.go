package storage

import (
	api "MVP_project/internal/api"
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, email, passwordHash string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*api.User, error)

	CreateCourse(ctx context.Context, course api.Course) (int64, error)
	GetCourseByID(ctx context.Context, id int64) (*api.Course, error)

	EnrollUser(ctx context.Context, userID, courseID int64) error
	GetEnrollments(ctx context.Context, userID int64) ([]api.Enrollment, error)

	SaveProgress(ctx context.Context, progress api.Progress) error
	GetProgress(ctx context.Context, userID, lessonID int64) (*api.Progress, error)
}
