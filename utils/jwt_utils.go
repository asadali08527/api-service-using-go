package utils

/**
The utils package provides utility functions to handle JWT (JSON Web Token) generation and validation, as well as extracting user information (such as username and role) from HTTP requests. It also offers a convenient way to manage user context across the application.

*/
import (
	"api-service/config"
	"api-service/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

/* This variable holds the secret key used for signing and verifying JWT tokens. It should ideally be stored in an environment variable for security purposes.*/
var jwtKey = []byte("your_secret_key")

/*
*
This function extracts the JWT token from the Authorization header in the HTTP request, validates it, and retrieves the username from the token claims.
*/
func GetUserIDFromRequest(r *http.Request) (string, error) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return "", errors.New("missing token")
	}

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		print(username)
		return username, nil
	}

	return "", errors.New("invalid token")
}

type contextKey string

const UserKey contextKey = "user_id"
const userContextKey = contextKey("user")
const RoleKey contextKey = "role"

// This function retrieves the role of the user from the request context.
func GetUserRoleFromRequest(r *http.Request) (string, error) {
	role, ok := r.Context().Value(RoleKey).(string)
	if !ok {
		return "", fmt.Errorf("role not found in context")
	}
	return role, nil
}

// This function generates a JWT token for the authenticated user based on their email, role, and username.
func GenerateJWT(user models.User) (string, error) {
	claims := &models.JWTClaims{
		Email:    user.Email,
		Role:     user.Role,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// This function validates a JWT token and returns the user claims embedded in it (email, role).
func ValidateToken(tokenString string) (*models.User, error) {
	claims := &models.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &models.User{
		Email: claims.Email,
		Role:  claims.Role,
	}, nil
}

// This function stores user data in the context of the current HTTP request. This is typically used by middleware to make user details available throughout the request lifecycle.
func ContextWithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// This function retrieves the user details from the request context, where they were previously stored by middleware.
func GetUserFromContext(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return nil, errors.New("no user found in context")
	}
	return user, nil
}
