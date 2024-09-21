API Service Using Go with JWT Authentication and Authorization
This project demonstrates how to implement authentication and authorization in a Go web application using PostgreSQL for data storage and web tokens (JWT) to secure API endpoints. It covers key concepts like password hashing, role-based access control (RBAC), JWT generation and validation, and secure coding practices.

Table of Contents
Introduction
Prerequisites
Project Structure
Setup and Installation
Configuration
Running the Application
API Endpoints
Public Routes
Protected Routes
File Explanations
config.go
main.go
Controllers
admin_controller.go
user_controller.go
Middleware
jwt_middleware.go
role_middleware.go
Models
user.go
Services
admin_service.go
user_service.go
Utilities
jwt_utils.go
Database
db.go
Security Best Practices
Postman Demo
Conclusion
Introduction
This application is a simple API service that demonstrates how to implement secure authentication and authorization mechanisms in Go. It allows users to register, log in, and access protected endpoints based on their roles (admin or user). The admin can manage users, while regular users can view and update their profiles.

Prerequisites
Go (version 1.16 or higher)
PostgreSQL (version 12 or higher)
go get to install dependencies
Postman (optional, for API testing)
Project Structure
go
 ├── main.go
├── config
│   └── config.go
├── controllers
│   ├── admin_controller.go
│   └── user_controller.go
├── middleware
│   ├── jwt_middleware.go
│   └── role_middleware.go
├── models
│   └── user.go
├── services
│   ├── admin_service.go
│   └── user_service.go
├── utils
│   └── jwt_utils.go
├── db
│   └── db.go
├── go.mod
└── go.sum
Setup and Installation
Clone the repository:

bash
 git clone https://github.com/yourusername/api-service.git
cd api-service
Install dependencies:

bash
 go mod download
Set up the PostgreSQL database:

Create a database named api_service.
Update the database credentials in db/db.go or set them using environment variables.
Run database migrations:

The application uses GORM's AutoMigrate feature, which automatically creates the necessary tables when the application starts.

Configuration
Environment Variables
It's recommended to store sensitive information like database credentials and JWT secret keys in environment variables.

DB_URL: Your PostgreSQL connection string.
JWT_SECRET: A secret key for signing JWT tokens.
Example of setting environment variables on Unix systems:

bash
 export DB_URL="host=localhost user=postgres password=yourpassword dbname=api_service port=5432 sslmode=disable"
export JWT_SECRET="your_secret_key"
Running the Application
To start the server, run:

bash
 go run main.go
The server will start on port 8080.

API Endpoints
Public Routes
POST /register: Register a new admin user.
POST /login: Log in as a user or admin.
Protected Routes
User Routes (Requires JWT)
GET /api/profile: Get the authenticated user's profile.
PUT /api/profile: Update the authenticated user's profile.
Admin Routes (Requires JWT and Admin Role)
GET /api/admin/users: Get all users.
POST /api/admin/users: Create a new user.
DELETE /api/admin/users/{id}: Delete a user by ID.
POST /api/admin/users/{id}/revoke: Revoke a user's token.
File Explanations
config.go
Location: config/config.go

Purpose: Stores configuration variables such as the JWT secret key and database URL.

go
 package config

import (
    "os"
)

var JWTSecret = os.Getenv("JWT_SECRET")
var DBUrl = os.Getenv("DB_URL")
Security Note: Always store sensitive information in environment variables rather than hardcoding them.
main.go
Purpose: Entry point of the application. Sets up routes, initializes services, and starts the server.

Key Components:

Route Initialization: Sets up public and protected routes using Gorilla Mux.
Middleware: Applies JWTMiddleware and AdminRoleMiddleware to protect routes.
Services and Controllers: Initializes services and controllers for users and admins.
Controllers
admin_controller.go
Location: controllers/admin_controller.go

Purpose: Handles admin-related HTTP requests.

CreateUser: Allows an admin to create a new user.
GetAllUsers: Retrieves all users.
DeleteUser: Deletes a user by ID.
RevokeToken: Revokes a user's JWT token.
user_controller.go
Location: controllers/user_controller.go

Purpose: Handles user-related HTTP requests.

Register: Registers a new user.
RegisterAdmin: Registers a new admin.
Login: Authenticates a user and returns a JWT token.
GetProfile: Retrieves the authenticated user's profile.
UpdateProfile: Updates the authenticated user's profile.
Middleware
jwt_middleware.go
Location: middleware/jwt_middleware.go

Purpose: Validates JWT tokens on protected routes.

Function: Checks the Authorization header for a valid JWT.
Operation: If the token is valid, the request proceeds; otherwise, it returns a 401 Unauthorized response.
role_middleware.go
Location: middleware/role_middleware.go

Purpose: Ensures that only users with the admin role can access certain routes.

Function: Checks the user's role from the context.
Operation: If the role is 'admin', the request proceeds; otherwise, it returns a 403 Forbidden response.
Models
user.go
Location: models/user.go

Purpose: Defines the User model and related structures.

go
 type User struct {
    ID       uint   `gorm:"primaryKey" :"id"`
    Name     string `:"name"`
    Email    string `:"email" gorm:"unique"`
    Username string `gorm:"unique" :"username"`
    Password string `:"password"`
    Mobile   string `:"mobile"`
    Address  string `:"address"`
    Role     string `:"role"`            // Admin or User
    Token    string `:"token,omitempty"` // Optional, stores JWT token for revocation
}
LoginCredentials: Used for handling login requests.
JWTClaims: Custom claims for JWT tokens, including email, role, and username.
Services
admin_service.go
Location: services/admin_service.go

Purpose: Contains business logic related to admin operations.

CreateUser: Hashes the password and creates a new user.
GetAllUsers: Retrieves all users from the database.
DeleteUser: Deletes a user by ID.
RevokeToken: Clears the user's token to revoke access.
user_service.go
Location: services/user_service.go

Purpose: Contains business logic related to user operations.

CreateUser: Hashes the password and creates a new user.
Authenticate: Validates user credentials.
GetProfile: Retrieves the user's profile.
UpdateProfile: Updates the user's profile details.
Utilities
jwt_utils.go
Location: utils/jwt_utils.go

Purpose: Handles JWT token generation and validation.

GenerateJWT: Generates a JWT token with user claims.
ValidateToken: Validates a JWT token and extracts user information.
Context Management:
ContextWithUser: Adds user data to the request context.
GetUserFromContext: Retrieves user data from the request context.
Database
db.go
Location: db/db.go

Purpose: Initializes the database connection and performs migrations.

Database Connection: Connects to PostgreSQL using GORM.
Auto Migration: Automatically creates the users table based on the User model.
Note: Update your database credentials or use environment variables to avoid hardcoding sensitive information.

Security Best Practices
Password Hashing: Always hash passwords using a strong algorithm like bcrypt before storing them.
JWT Secrets: Store your JWT secret keys securely, preferably in environment variables.
Role-Based Access Control: Implement RBAC to manage user permissions effectively.
Token Expiration: Set appropriate expiration times for JWT tokens.
Input Validation: Always validate and sanitize user input to prevent injection attacks.
HTTPS: Use HTTPS to encrypt data in transit.
Postman Demo
You can use Postman to test the API endpoints.

Register an Admin

Request: POST /register

Body:


 {
  "username": "admin",
  "password": "admin123",
  "email": "admin@example.com"
}
Login

Request: POST /login

Body:


 {
  "username": "admin",
  "password": "admin123"
}
Response:


 {
  "token": "your_jwt_token_here"
}
Create a User (Admin Only)

Request: POST /api/admin/users

Headers:

 
 Authorization: Bearer your_jwt_token_here
Body:


 {
  "username": "user1",
  "password": "password123",
  "role": "user",
  "email": "user1@example.com"
}
Get User Profile

Request: GET /api/profile

Headers:

 
 Authorization: Bearer your_jwt_token_here
Update User Profile

Request: PUT /api/profile

Headers:

 
 Authorization: Bearer your_jwt_token_here
Body:


 {
  "mobile": "123456789",
  "address": "123 Main St"
}
Delete a User (Admin Only)

Request: DELETE /api/admin/users/{id}

Headers:

 
 Authorization: Bearer your_jwt_token_here
Conclusion
This project provides a foundational understanding of how to implement authentication and authorization in a Go web application using PostgreSQL and JWTs. By following best practices and utilizing secure coding techniques, you can build robust applications that protect user data and ensure only authorized access to resources.

Feel free to contribute or raise issues if you find any problems or have suggestions for improvements.

Author: Asad Ali

Date: September 17, 2024
