package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	fmt.Println("user created successfully")
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the userID from the URL path
	userID := r.URL.Path[len("/users/"):]

	// Parse the userID as an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service to retrieve the user by ID
	user, err := c.userService.GetUser(userIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the user data to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}






