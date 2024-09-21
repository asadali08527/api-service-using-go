package models

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
	Role     string `json:"role"`            // Admin or User
	Token    string `json:"token,omitempty"` // Optional, stores JWT token for revocation
}

// LoginCredentials for login
type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWTClaims stores the claims for JWT
type JWTClaims struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}
