package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/middleware"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/services"
)

type ArticleController struct {
	articles *services.ArticleService
}

func NewArticleController(articles *services.ArticleService) *ArticleController {
	return &ArticleController{articles: articles}
}

func (c *ArticleController) Create(w http.ResponseWriter, r *http.Request) {
	authorID, ok := middleware.SubjectFromContext(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "missing auth subject")
		return
	}

	var in services.ArticleInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	article, err := c.articles.Create(r.Context(), authorID, in)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, article)
}

func (c *ArticleController) GetByID(w http.ResponseWriter, r *http.Request) {
	article, err := c.articles.GetByID(r.Context(), chi.URLParam(r, "articleID"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "article not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to fetch article")
		return
	}
	writeJSON(w, http.StatusOK, article)
}

func (c *ArticleController) Update(w http.ResponseWriter, r *http.Request) {
	var in services.ArticleInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	article, err := c.articles.Update(r.Context(), chi.URLParam(r, "articleID"), in)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "article not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, article)
}
