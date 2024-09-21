package middleware

import (
	"api-service/utils"
	"net/http"
)

func AdminRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.GetUserFromContext(r.Context())
		if err != nil || user.Role != "admin" {
			http.Error(w, "Forbidden - Admins only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
