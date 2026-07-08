package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/config"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/router"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := store.NewPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer db.Close()

	graph, err := store.NewNeo4j(context.Background(), cfg.Neo4jURI, cfg.Neo4jUser, cfg.Neo4jPassword)
	if err != nil {
		log.Fatalf("connect neo4j: %v", err)
	}
	defer graph.Close(context.Background())

	srv := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      router.New(cfg, db, graph),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	log.Printf("api listening on :%s", cfg.HTTPPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
