package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/middleware"
)

type ProfileController struct {
	db *sql.DB
}

func NewProfileController(db *sql.DB) *ProfileController {
	return &ProfileController{db: db}
}

func (c *ProfileController) Me(w http.ResponseWriter, r *http.Request) {
	subject, ok := middleware.SubjectFromContext(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "missing auth subject")
		return
	}
	var payload []byte
	err := c.db.QueryRowContext(r.Context(), `
		SELECT json_build_object(
			'id', id,
			'email', email,
			'displayName', display_name,
			'role', role,
			'bio', bio,
			'verification', verification,
			'credentialUrl', credential_url,
			'createdAt', created_at
		)
		FROM users WHERE id = $1
	`, subject).Scan(&payload)
	if err != nil {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	var profile map[string]any
	if err := json.Unmarshal(payload, &profile); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to parse profile")
		return
	}
	writeJSON(w, http.StatusOK, profile)
}
