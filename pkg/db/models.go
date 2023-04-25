// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type AccessToken struct {
	ID             int32     `json:"id"`
	OrganisationID int32     `json:"organisation_id"`
	EncryptedToken string    `json:"encrypted_token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Group struct {
	ID             int32     `json:"id"`
	OrganisationID int32     `json:"organisation_id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GroupRole struct {
	ID        int32     `json:"id"`
	GroupID   int32     `json:"group_id"`
	RoleID    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Organisation struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID             int32     `json:"id"`
	OrganisationID int32     `json:"organisation_id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type User struct {
	ID             int32         `json:"id"`
	OrganisationID sql.NullInt32 `json:"organisation_id"`
	GroupID        sql.NullInt32 `json:"group_id"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	IsAdmin        bool          `json:"is_admin"`
	IsSuperAdmin   bool          `json:"is_super_admin"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type UserRole struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	RoleID    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}