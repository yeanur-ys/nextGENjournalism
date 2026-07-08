package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/services"
)

type AuthController struct {
	auth *services.AuthService
}

func NewAuthController(auth *services.AuthService) *AuthController {
	return &AuthController{auth: auth}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var in services.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	user, err := c.auth.Register(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{"user": user})
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var in services.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	res, err := c.auth.Login(r.Context(), in)
	if err != nil {
		code := http.StatusBadRequest
		if errors.Is(err, services.ErrInvalidCredentials) {
			code = http.StatusUnauthorized
		}
		writeError(w, code, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, res)
}
