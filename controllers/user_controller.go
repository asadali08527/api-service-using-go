package controllers

/***
The UserController manages user-related operations such as user registration, login, profile viewing, profile updates, and JWT token management. It interacts with the UserService to perform business logic, including authentication and profile management.
*/
import (
	"api-service/models"
	"api-service/services"
	"api-service/utils"
	"encoding/json"
	"net/http"
)

type UserController struct {
	/**
	The UserService handles the core business logic for user management, including creating users, authenticating logins, and managing user profiles.
	*/
	UserService *services.UserService
}

/*
GetProfile

func (uc *UserController) GetProfile(w http.ResponseWriter, r *http.Request)
Description: This endpoint allows a user to retrieve their profile information.

Request:

Method: GET
Endpoint: /api/profile
Headers: Must contain a valid JWT token in the Authorization header.
Logic:

The JWT token is extracted from the request using utils.GetUserIDFromRequest.
The username is retrieved from the token and used to fetch the user profile from the UserService.
If the profile is found, the user data is returned with a 200 OK status.
If the user is not found, a 404 Not Found error is returned.

Response:

On success:

	{
	  "username": "user1",
	  "email": "user1@example.com",
	  "mobile": "123456789",
	  "address": "123 Street, City"
	}

On error: 404 Not Found
*/
func (uc *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	username, _ := utils.GetUserIDFromRequest(r)

	profile, err := uc.UserService.GetProfile(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

/*
*
UpdateProfile

func (uc *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request)
Description: This endpoint allows users to update their profile details such as mobile number and address.

Request:

Method: PUT
Endpoint: /api/profile
Headers: Must contain a valid JWT token in the Authorization header.
Body (JSON format):

	{
	  "mobile": "987654321",
	  "address": "New Address, City"
	}

Logic:

Extracts the username from the JWT token.
Decodes the request body to get the updated mobile and address.
The UpdateProfile function in UserService updates the user's profile with the new data.
If the update is successful, the updated profile is returned with a 200 OK status.
If the user is not found, a 404 Not Found error is returned.

Response:

On success:

	{
	  "username": "user1",
	  "email": "user1@example.com",
	  "mobile": "987654321",
	  "address": "New Address, City"
	}

On error: 404 Not Found
*/
func (uc *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	username, _ := utils.GetUserIDFromRequest(r)

	var updateData struct {
		Mobile  string `json:"mobile"`
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	profile, err := uc.UserService.UpdateProfile(username, updateData.Mobile, updateData.Address)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

/*
* Register

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request)
Description: This endpoint allows a new user to register with a default role of "user".

Request:

Method: POST
Endpoint: /register
Body (JSON format):

	{
	  "username": "user1",
	  "password": "password123",
	  "email": "user1@example.com"
	}

Logic:

Decodes the request body into a User model.
Sets the user's role to user by default.
Calls the CreateUser function in UserService to register the user.
If successful, the newly created user data is returned with a 201 Created status.
If an error occurs, a 500 Internal Server Error is returned.
Response:

On success:

	{
	  "username": "user1",
	  "email": "user1@example.com",
	  "role": "user"
	}

On error: 500 Internal Server Error
*/
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Role = "user" // Default role is 'user'

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

/*
*
RegisterAdmin

func (uc *UserController) RegisterAdmin(w http.ResponseWriter, r *http.Request)
Description: This endpoint allows an admin to register another admin by setting the role to "admin".

Request:

Method: POST
Endpoint: /register/admin
Body (JSON format):

	{
	  "username": "admin1",
	  "password": "admin123",
	  "email": "admin1@example.com"
	}

Logic:

Decodes the request body into a User model.
Sets the user's role to admin.
Calls the CreateUser function in UserService to register the admin.
If successful, the newly created admin user data is returned with a 201 Created status.
If an error occurs, a 500 Internal Server Error is returned.
Response:

On success:

	{
	  "username": "admin1",
	  "email": "admin1@example.com",
	  "role": "admin"
	}

On error: 500 Internal Server Error
*/
func (uc *UserController) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.User
	json.NewDecoder(r.Body).Decode(&admin)
	admin.Role = "admin" // Set role to 'admin'

	err := uc.UserService.CreateUser(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

/*
*
Login

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request)
Description: This endpoint allows users and admins to log in by providing their username and password. If successful, it returns a JWT token for authentication.

Request:

Method: POST
Endpoint: /login
Body (JSON format):

	{
	  "username": "user1",
	  "password": "password123"
	}

Logic:

The request body is decoded into a LoginCredentials structure containing the username and password.
The Authenticate function in UserService is called to verify the credentials.
If the credentials are valid, a JWT token is generated using utils.GenerateJWT.
The JWT token is returned in the response with a 200 OK status.
If authentication fails, a 401 Unauthorized error is returned.
Response:

On success:

	{
	  "token": "your_jwt_token_here"
	}

On error: 401 Unauthorized
*/
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials
	json.NewDecoder(r.Body).Decode(&credentials)

	user, err := uc.UserService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(*user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	user.Token = token
	// print(token)
	uc.UserService.DB.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
