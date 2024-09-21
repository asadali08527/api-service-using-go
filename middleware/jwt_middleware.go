package middleware

/**
The JWTMiddleware is responsible for validating the JSON Web Token (JWT) provided by the user in the Authorization header. It ensures that only authenticated users can access protected routes by verifying the token and adding user information to the request context for downstream use in the application.
*/
import (
	"api-service/utils"
	"net/http"
)

/*
*
JWTMiddleware

func JWTMiddleware(next http.Handler) http.Handler
Description: This middleware intercepts incoming HTTP requests, checks if the request contains a valid JWT token in the Authorization header, and validates it. If the token is valid, it adds the user information to the request context and allows the request to proceed. If the token is missing or invalid, it returns an unauthorized error (401 Unauthorized).
*/
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//The token is retrieved from the Authorization header. If no token is present, the middleware responds with an error.
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			//If the token is missing, it sends a 401 Unauthorized response:
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}
		// print(tokenString)
		//  The token is passed to the ValidateToken function in the utils package, where the JWT token is decrypted and validated. The ValidateToken function returns the user information if the token is valid.
		user, err := utils.ValidateToken(tokenString)
		if err != nil {
			// If the token is invalid or expired, a 401 Unauthorized error is returned:
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, the user information (extracted from the token) is stored in the request context using the ContextWithUser function. This allows downstream handlers to access the authenticated user's information via the context.
		ctx := utils.ContextWithUser(r.Context(), user)
		//The middleware calls the next handler in the chain, passing the modified request with the user information in the context. This ensures that only authenticated requests can proceed to the protected endpoint.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
