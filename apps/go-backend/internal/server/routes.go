package server

import (
	"net/http"

	"github.com/yeanur-ys/nextGENjournalism/apps/go-backend/internal/articles"
	"github.com/yeanur-ys/nextGENjournalism/apps/go-backend/internal/auth"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("POST /auth/login", auth.LoginHandler)
	mux.HandleFunc("GET /articles", articles.ListHandler)
	mux.HandleFunc("POST /articles", articles.CreateHandler)

	return RoleGuard("journalist", mux)
}
