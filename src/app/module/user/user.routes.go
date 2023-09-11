package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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
		h.handleCreateUser(w, r)
	case http.MethodGet:
		if r.URL.Path == "/users" {
			h.handleGetAllUsers(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/users/") {
			h.handleGetUserByID(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	case http.MethodPut:
		if strings.HasPrefix(r.URL.Path, "/users/") {
			h.handleUpdateUser(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	case http.MethodDelete:
		if strings.HasPrefix(r.URL.Path, "/users/") {
			h.handleDeleteUserByID(w, r)
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

func (h *UserRoutesHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	userService := NewUserService(h.DB)
	userController := NewUserController(userService)
	userController.CreateUser(w, r)
}

func (h *UserRoutesHandler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	userService := NewUserService(h.DB)
	userController := NewUserController(userService)
	userController.GetAllUsers(w, r)
}

func (h *UserRoutesHandler) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
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
}

func (h *UserRoutesHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	userIDStr := r.URL.Path[len("/users/"):]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the controller to update the user
	userController := NewUserController(NewUserService(h.DB))
	userController.UpdateUser(w, r, userID)
}

// func (h *UserRoutesHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
// 	// Extract the user ID from the URL parameter
// 	userIDStr := r.URL.Path[len("/users/"):]
// 	userID, err := strconv.Atoi(userIDStr)
// 	if err != nil {
// 		http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Parse the JSON request body to get the updated user data
// 	var updatedUser User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&updatedUser); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Call the service to update the user
// 	userService := NewUserService(h.DB)
// 	if err := userService.UpdateUser(userID, &updatedUser); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "User updated successfully")
// }

func (h *UserRoutesHandler) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {

}

/*package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

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
*/
