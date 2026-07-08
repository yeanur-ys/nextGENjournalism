package services

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yeanur-ys/nextGENjournalism/apps/api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db        *sql.DB
	jwtSecret []byte
}

var ErrInvalidCredentials = errors.New("invalid credentials")

type RegisterInput struct {
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	DisplayName string      `json:"displayName"`
	Role        models.Role `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func NewAuthService(db *sql.DB, jwtSecret string) *AuthService {
	return &AuthService{db: db, jwtSecret: []byte(jwtSecret)}
}

func (s *AuthService) Register(ctx context.Context, in RegisterInput) (models.User, error) {
	if in.Role == "" {
		in.Role = models.RoleReader
	}
	if err := validateRegisterInput(in); err != nil {
		return models.User{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	var u models.User
	err = s.db.QueryRowContext(ctx, `
		INSERT INTO users (email, password_hash, display_name, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, display_name, role, bio, verification, created_at
	`, strings.ToLower(strings.TrimSpace(in.Email)), string(hash), strings.TrimSpace(in.DisplayName), in.Role).Scan(
		&u.ID, &u.Email, &u.DisplayName, &u.Role, &u.Bio, &u.Verification, &u.CreatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (s *AuthService) Login(ctx context.Context, in LoginInput) (AuthResponse, error) {
	if strings.TrimSpace(in.Email) == "" || in.Password == "" {
		return AuthResponse{}, errors.New("email and password are required")
	}

	var (
		u            models.User
		passwordHash string
	)

	err := s.db.QueryRowContext(ctx, `
		SELECT id, email, display_name, role, bio, verification, created_at, password_hash
		FROM users
		WHERE email = $1
	`, strings.ToLower(strings.TrimSpace(in.Email))).Scan(
		&u.ID, &u.Email, &u.DisplayName, &u.Role, &u.Bio, &u.Verification, &u.CreatedAt, &passwordHash,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return AuthResponse{}, ErrInvalidCredentials
		}
		return AuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(in.Password)); err != nil {
		return AuthResponse{}, ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  u.ID,
		"role": string(u.Role),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{Token: tokenString, User: u}, nil
}

func validateRegisterInput(in RegisterInput) error {
	if strings.TrimSpace(in.Email) == "" || in.Password == "" || strings.TrimSpace(in.DisplayName) == "" {
		return errors.New("email, password and displayName are required")
	}
	if len(in.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	switch in.Role {
	case models.RoleJournalist, models.RoleAuditor, models.RoleReader, models.RoleAdmin:
		return nil
	default:
		return errors.New("invalid role")
	}
}
