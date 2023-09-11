package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

// Define a custom context key type
type contextKey string

const (
	userIDContextKey contextKey = "userID"
)

type UserRoutesHandler struct {
	DB *sql.DB
}

func UserRoutes(db *sql.DB) *UserRoutesHandler {
	return &UserRoutesHandler{DB: db}
}

func (h *UserRoutesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		userService := NewUserService(h.DB)
		userController := NewUserController(userService)
		userController.CreateUser(w, r)
	case http.MethodGet:
		if r.URL.Path == "/users" {
			userService := NewUserService(h.DB)
			userController := NewUserController(userService)
			userController.GetAllUsers(w, r)
		} else if len(r.URL.Path) > len("/users/") && r.URL.Path[:len("/users/")] == "/users/" {
			// Extract the user ID from the URL path
			userIDStr := r.URL.Path[len("/users/"):]
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
			userService := NewUserService(h.DB)
			userController := NewUserController(userService)

			// Pass the userID within the request context using the custom context key
			ctx := context.WithValue(r.Context(), userIDContextKey, userID)
			userController.GetUserByID(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Define routes here
func (h *UserRoutesHandler) AddRoutes() {
	http.Handle("/users", h)
	http.Handle("/users/", h)
}
