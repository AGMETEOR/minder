// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (organization_id, identity_subject) VALUES ($1, $2) RETURNING id, organization_id, identity_subject, created_at, updated_at
`

type CreateUserParams struct {
	OrganizationID  uuid.UUID `json:"organization_id"`
	IdentitySubject string    `json:"identity_subject"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.OrganizationID, arg.IdentitySubject)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdentitySubject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, organization_id, identity_subject, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdentitySubject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserBySubject = `-- name: GetUserBySubject :one
SELECT id, organization_id, identity_subject, created_at, updated_at FROM users WHERE identity_subject = $1
`

func (q *Queries) GetUserBySubject(ctx context.Context, identitySubject string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserBySubject, identitySubject)
	var i User
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.IdentitySubject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, organization_id, identity_subject, created_at, updated_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.IdentitySubject,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByOrganization = `-- name: ListUsersByOrganization :many
SELECT id, organization_id, identity_subject, created_at, updated_at FROM users
WHERE organization_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUsersByOrganizationParams struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	Limit          int32     `json:"limit"`
	Offset         int32     `json:"offset"`
}

func (q *Queries) ListUsersByOrganization(ctx context.Context, arg ListUsersByOrganizationParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByOrganization, arg.OrganizationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.IdentitySubject,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
