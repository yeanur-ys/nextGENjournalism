package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	subjectKey contextKey = "subject"
	roleKey    contextKey = "role"
)

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid claims", http.StatusUnauthorized)
				return
			}
			sub, _ := claims["sub"].(string)
			role, _ := claims["role"].(string)
			if sub == "" || role == "" {
				http.Error(w, "missing token claims", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), subjectKey, sub)
			ctx = context.WithValue(ctx, roleKey, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRoles(allowed ...string) func(http.Handler) http.Handler {
	set := make(map[string]struct{}, len(allowed))
	for _, role := range allowed {
		set[role] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := RoleFromContext(r.Context())
			if !ok {
				http.Error(w, "missing auth role", http.StatusForbidden)
				return
			}
			if _, ok := set[role]; !ok {
				http.Error(w, "insufficient role", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func SubjectFromContext(ctx context.Context) (string, bool) {
	sub, ok := ctx.Value(subjectKey).(string)
	return sub, ok
}

func RoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(roleKey).(string)
	return role, ok
}
