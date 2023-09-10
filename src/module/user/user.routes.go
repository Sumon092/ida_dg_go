package user

import (
	"database/sql"
	"net/http"
)

type UserRoutesHandler struct {
	DB *sql.DB
}

func UserRoutes(db *sql.DB) *UserRoutesHandler {
	return &UserRoutesHandler{DB: db}
}

func (h *UserRoutesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userService := NewUserService(h.DB)
		userController := NewUserController(userService)
		userController.CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
