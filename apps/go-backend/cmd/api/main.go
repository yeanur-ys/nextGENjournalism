package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yeanur-ys/nextGENjournalism/apps/go-backend/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := server.NewRouter()
	log.Printf("go-backend listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
