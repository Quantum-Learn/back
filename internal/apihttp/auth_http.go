package apihttp

import (
	api "MVP_project/internal/api"
	"MVP_project/internal/auth"
	"MVP_project/internal/handlers"
	"MVP_project/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthHTTP struct {
	H     *handlers.AuthHandler
	Store storage.UserStore
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (a *AuthHTTP) Register(w http.ResponseWriter, r *http.Request) {
	var req api.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	resp, err := a.H.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, resp)
}

func (a *AuthHTTP) Login(w http.ResponseWriter, r *http.Request) {
	var req api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	resp, err := a.H.Login(r.Context(), req)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (a *AuthHTTP) Me(w http.ResponseWriter, r *http.Request) {
	// достаём userID из контекста
	val := r.Context().Value(auth.UserIDKey)
	uidStr, ok := val.(string)
	if !ok || uidStr == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// преобразуем userID в int
	id, err := strconv.Atoi(uidStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	// достаём пользователя из стора
	u, err := a.Store.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// убираем пароль из ответа
	u.PasswordHash = nil

	// отдаём JSON
	writeJSON(w, http.StatusOK, u)
}
