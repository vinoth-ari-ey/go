package userHandler

import (
	"encoding/json"
	"go-crud/models"
	userService "go-crud/services"
	"net/http"
	"strconv"
)

// UserHandler handles requests for user-related operations.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case http.MethodGet:
		idStr := r.URL.Path[len("/users/"):]
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			users, err := userService.GetUsers()
			if err != nil {
				http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(users)
		} else {
			user, err := userService.GetUser(userId)
			if err != nil {
				http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(user)

		}

	case http.MethodPost:
		var userPayload models.User
		if err := json.NewDecoder(r.Body).Decode(&userPayload); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		createdUser, err := userService.CreateUser(&userPayload)
		if err != nil {
			http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(createdUser)

	case http.MethodPut:
		idStr := r.URL.Path[len("/users/"):]
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		var updatedUser models.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		updatedUserPtr, err := userService.UpdateUser(userId, &updatedUser)
		if err != nil {
			http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedUserPtr)

	case http.MethodDelete:
		idStr := r.URL.Path[len("/users/"):]
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		if err := userService.DeleteUser(userId); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User deleted successfully"))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
