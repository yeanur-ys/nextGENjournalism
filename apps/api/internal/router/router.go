package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/config"
	controllers "github.com/yeanur-ys/nextGENjournalism/apps/api/internal/controllers"
	authmw "github.com/yeanur-ys/nextGENjournalism/apps/api/internal/middleware"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/models"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/services"
)

func New(cfg config.Config, db *sql.DB, graph neo4j.DriverWithContext) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authService := services.NewAuthService(db, cfg.JWTSecret)
	articleService := services.NewArticleService(db, graph)
	authController := controllers.NewAuthController(authService)
	articleController := controllers.NewArticleController(articleService)
	profileController := controllers.NewProfileController(db)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/auth/register", authController.Register)
		api.Post("/auth/login", authController.Login)

		api.Group(func(protected chi.Router) {
			protected.Use(authmw.JWTAuth(cfg.JWTSecret))
			protected.Get("/profile", profileController.Me)

			protected.With(authmw.RequireRoles(string(models.RoleJournalist), string(models.RoleAdmin))).Post("/articles", articleController.Create)
			protected.With(authmw.RequireRoles(string(models.RoleJournalist), string(models.RoleAdmin))).Put("/articles/{articleID}", articleController.Update)
		})

		api.Get("/articles/{articleID}", articleController.GetByID)
	})

	return r
}
