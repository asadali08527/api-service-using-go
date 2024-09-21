package config

import (
	"os"
)

// var JWTSecret = os.Getenv("JWT_SECRET")
var JWTSecret = "your_secret_key"
var DBUrl = os.Getenv("DB_URL")
