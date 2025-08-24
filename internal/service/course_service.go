package service

import (
	"MVP_project/internal/models"
	"MVP_project/internal/repository"
)

type CourseService interface {
	ListCourses() ([]models.Course, error)
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) ListCourses() ([]models.Course, error) {
	return s.repo.GetAll()
}
