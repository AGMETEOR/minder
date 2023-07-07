// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Group struct {
	ID             int32          `json:"id"`
	OrganizationID int32          `json:"organization_id"`
	Name           string         `json:"name"`
	Description    sql.NullString `json:"description"`
	IsProtected    bool           `json:"is_protected"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type Organization struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Company   string    `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Policy struct {
	ID               int32           `json:"id"`
	Provider         string          `json:"provider"`
	GroupID          int32           `json:"group_id"`
	PolicyType       int32           `json:"policy_type"`
	PolicyDefinition json.RawMessage `json:"policy_definition"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type PolicyType struct {
	ID          int32          `json:"id"`
	Provider    string         `json:"provider"`
	PolicyType  string         `json:"policy_type"`
	Description sql.NullString `json:"description"`
	Version     string         `json:"version"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ProviderAccessToken struct {
	ID             int32     `json:"id"`
	Provider       string    `json:"provider"`
	GroupID        int32     `json:"group_id"`
	EncryptedToken string    `json:"encrypted_token"`
	ExpirationTime time.Time `json:"expiration_time"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Repository struct {
	ID         int32         `json:"id"`
	Provider   string        `json:"provider"`
	GroupID    int32         `json:"group_id"`
	RepoOwner  string        `json:"repo_owner"`
	RepoName   string        `json:"repo_name"`
	RepoID     int32         `json:"repo_id"`
	IsPrivate  bool          `json:"is_private"`
	IsFork     bool          `json:"is_fork"`
	WebhookID  sql.NullInt32 `json:"webhook_id"`
	WebhookUrl string        `json:"webhook_url"`
	DeployUrl  string        `json:"deploy_url"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

type Role struct {
	ID             int32         `json:"id"`
	OrganizationID int32         `json:"organization_id"`
	GroupID        sql.NullInt32 `json:"group_id"`
	Name           string        `json:"name"`
	IsAdmin        bool          `json:"is_admin"`
	IsProtected    bool          `json:"is_protected"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type SessionStore struct {
	ID           int32         `json:"id"`
	Provider     string        `json:"provider"`
	GrpID        sql.NullInt32 `json:"grp_id"`
	Port         sql.NullInt32 `json:"port"`
	SessionState string        `json:"session_state"`
	CreatedAt    time.Time     `json:"created_at"`
}

type User struct {
	ID                  int32          `json:"id"`
	OrganizationID      int32          `json:"organization_id"`
	Email               sql.NullString `json:"email"`
	Username            string         `json:"username"`
	Password            string         `json:"password"`
	NeedsPasswordChange bool           `json:"needs_password_change"`
	FirstName           sql.NullString `json:"first_name"`
	LastName            sql.NullString `json:"last_name"`
	IsProtected         bool           `json:"is_protected"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	MinTokenIssuedTime  sql.NullTime   `json:"min_token_issued_time"`
}

type UserGroup struct {
	ID      int32 `json:"id"`
	UserID  int32 `json:"user_id"`
	GroupID int32 `json:"group_id"`
}

type UserRole struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}
