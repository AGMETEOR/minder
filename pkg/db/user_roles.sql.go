// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user_roles.sql

package db

import (
	"context"
)

const assignRoleToUser = `-- name: AssignRoleToUser :one
INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) RETURNING id, user_id, role_id, created_at, updated_at
`

type AssignRoleToUserParams struct {
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

func (q *Queries) AssignRoleToUser(ctx context.Context, arg AssignRoleToUserParams) (UserRole, error) {
	row := q.db.QueryRowContext(ctx, assignRoleToUser, arg.UserID, arg.RoleID)
	var i UserRole
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserRoles = `-- name: GetUserRoles :many
SELECT id, user_id, role_id, created_at, updated_at FROM user_roles WHERE user_id = $1
`

func (q *Queries) GetUserRoles(ctx context.Context, userID int32) ([]UserRole, error) {
	rows, err := q.db.QueryContext(ctx, getUserRoles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserRole{}
	for rows.Next() {
		var i UserRole
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.RoleID,
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

const revokeRoleFromUser = `-- name: RevokeRoleFromUser :exec
DELETE FROM user_roles WHERE user_id = $1 AND role_id = $2
`

type RevokeRoleFromUserParams struct {
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

func (q *Queries) RevokeRoleFromUser(ctx context.Context, arg RevokeRoleFromUserParams) error {
	_, err := q.db.ExecContext(ctx, revokeRoleFromUser, arg.UserID, arg.RoleID)
	return err
}