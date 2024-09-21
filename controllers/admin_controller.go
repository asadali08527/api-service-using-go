package controllers

/**The AdminController is responsible for handling HTTP requests related to user management from an admin perspective. It interacts with the AdminService to perform actions such as creating users, retrieving all users, deleting users, and revoking user tokens.
 */
import (
	"api-service/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminController struct {
	/**
	AdminService: The AdminService is a service layer that handles the business logic related to user management. It is injected into the controller to perform database operations.
	*/
	AdminService *services.AdminService
}

/*
*
This endpoint allows an admin to create a new user by passing the required user details (username, password, role, and email) in JSON format.
Request:

Method: POST
Endpoint: /admin/users
Body (JSON format):
json

	{
	  "username": "user1",
	  "password": "password123",
	  "role": "user",
	  "email": "user1@example.com"
	}
*/
func (ac *AdminController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := ac.AdminService.CreateUser(data.Username, data.Password, data.Role, data.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (ac *AdminController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ac.AdminService.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (ac *AdminController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	if err := ac.AdminService.DeleteUser(uint(userID)); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}

func (ac *AdminController) RevokeToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	if err := ac.AdminService.RevokeToken(uint(userID)); err != nil {
		http.Error(w, "Failed to revoke token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User's token revoked"})
}
