package main

import (
	"api-service/controllers"
	"api-service/db"
	"api-service/middleware"
	"api-service/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize DB
	dbConn := db.InitDB()

	// Initialize Services
	userService := &services.UserService{DB: dbConn}
	adminService := &services.AdminService{DB: dbConn}

	// Initialize Controllers
	userController := &controllers.UserController{UserService: userService}
	adminController := &controllers.AdminController{AdminService: adminService}

	// Setup Router
	router := mux.NewRouter()

	// Public Routes
	router.HandleFunc("/register", userController.RegisterAdmin).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")

	// Protected Routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	// User Routes (protected for logged-in users)
	api.HandleFunc("/profile", userController.GetProfile).Methods("GET")
	api.HandleFunc("/profile", userController.UpdateProfile).Methods("PUT")

	// Admin Routes (protected for admin only)
	adminApi := api.PathPrefix("/admin").Subrouter()
	adminApi.Use(middleware.AdminRoleMiddleware)

	adminApi.HandleFunc("/users", adminController.GetAllUsers).Methods("GET")
	adminApi.HandleFunc("/users", adminController.CreateUser).Methods("POST")
	adminApi.HandleFunc("/users/{id}", adminController.DeleteUser).Methods("DELETE")
	adminApi.HandleFunc("/users/{id}/revoke", adminController.RevokeToken).Methods("POST")

	// Start server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
