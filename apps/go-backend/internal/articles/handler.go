package articles

import (
	"encoding/json"
	"net/http"
	"time"
)

func ListHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode([]Article{})
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	if article.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	article.ID = "draft"
	article.CreatedAt = time.Now().UTC()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(article)
}
