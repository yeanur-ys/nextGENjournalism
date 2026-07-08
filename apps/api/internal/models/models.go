package models

import "time"

type Role string

const (
	RoleJournalist Role = "journalist"
	RoleAuditor    Role = "auditor"
	RoleReader     Role = "reader"
	RoleAdmin      Role = "admin"
)

type User struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	DisplayName   string    `json:"displayName"`
	Role          Role      `json:"role"`
	Bio           string    `json:"bio"`
	CredentialURL string    `json:"credentialUrl,omitempty"`
	Verification  string    `json:"verification"`
	CreatedAt     time.Time `json:"createdAt"`
}

type Article struct {
	ID            string    `json:"id"`
	AuthorID      string    `json:"authorId"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Status        string    `json:"status"`
	SyncedToGraph bool      `json:"syncedToGraph"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
