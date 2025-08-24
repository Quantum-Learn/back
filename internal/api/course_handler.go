package api

import (
	"encoding/json"
	"net/http"

	"MVP_project/internal/service"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(s service.CourseService) *CourseHandler {
	return &CourseHandler{service: s}
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.service.ListCourses()
	if err != nil {
		http.Error(w, "Unable to fetch courses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(courses)
}
