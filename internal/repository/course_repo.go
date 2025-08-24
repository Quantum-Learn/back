package repository

import "MVP_project/internal/models"

type CourseRepository interface {
	GetAll() ([]models.Course, error)
}

type courseRepo struct{}

func NewCourseRepository() CourseRepository {
	return &courseRepo{}
}

func (r *courseRepo) GetAll() ([]models.Course, error) {
	// Тут будет обращение к БД
	return []models.Course{
		{ID: 1, Title: "Go Basics"},
		{ID: 2, Title: "React for Beginners"},
	}, nil
}
