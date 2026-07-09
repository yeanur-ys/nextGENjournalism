package server

import "net/http"

func RoleGuard(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requiredRole == "" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/health" || r.URL.Path == "/auth/login" {
			next.ServeHTTP(w, r)
			return
		}

		role := r.Header.Get("X-Role")
		if role == "" {
			http.Error(w, "missing role", http.StatusUnauthorized)
			return
		}
		if role != requiredRole {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
