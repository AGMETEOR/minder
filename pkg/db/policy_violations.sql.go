// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: policy_violations.sql

package db

import (
	"context"
	"encoding/json"
	"time"
)

const createPolicyViolation = `-- name: CreatePolicyViolation :one
INSERT INTO policy_violations (  
    repository_id,
    policy_id,
    metadata,
    violation) VALUES ($1, $2, $3::jsonb, $4::jsonb) RETURNING id, repository_id, policy_id, metadata, violation, created_at
`

type CreatePolicyViolationParams struct {
	RepositoryID int32           `json:"repository_id"`
	PolicyID     int32           `json:"policy_id"`
	Metadata     json.RawMessage `json:"metadata"`
	Violation    json.RawMessage `json:"violation"`
}

func (q *Queries) CreatePolicyViolation(ctx context.Context, arg CreatePolicyViolationParams) (PolicyViolation, error) {
	row := q.db.QueryRowContext(ctx, createPolicyViolation,
		arg.RepositoryID,
		arg.PolicyID,
		arg.Metadata,
		arg.Violation,
	)
	var i PolicyViolation
	err := row.Scan(
		&i.ID,
		&i.RepositoryID,
		&i.PolicyID,
		&i.Metadata,
		&i.Violation,
		&i.CreatedAt,
	)
	return i, err
}

const getPolicyViolationsByGroup = `-- name: GetPolicyViolationsByGroup :many
SELECT pt.policy_type, r.id as repo_id, r.repo_owner, r.repo_name,
pv.metadata, pv.violation, pv.created_at FROM policy_violations pv
INNER JOIN policies p ON p.id = pv.policy_id
INNER JOIN repositories r ON r.id = pv.repository_id
INNER JOIN policy_types pt ON pt.id = p.policy_type
WHERE p.provider=$1 AND p.group_id=$2 ORDER BY pv.created_at DESC LIMIT $3 OFFSET $4
`

type GetPolicyViolationsByGroupParams struct {
	Provider string `json:"provider"`
	GroupID  int32  `json:"group_id"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type GetPolicyViolationsByGroupRow struct {
	PolicyType string          `json:"policy_type"`
	RepoID     int32           `json:"repo_id"`
	RepoOwner  string          `json:"repo_owner"`
	RepoName   string          `json:"repo_name"`
	Metadata   json.RawMessage `json:"metadata"`
	Violation  json.RawMessage `json:"violation"`
	CreatedAt  time.Time       `json:"created_at"`
}

func (q *Queries) GetPolicyViolationsByGroup(ctx context.Context, arg GetPolicyViolationsByGroupParams) ([]GetPolicyViolationsByGroupRow, error) {
	rows, err := q.db.QueryContext(ctx, getPolicyViolationsByGroup,
		arg.Provider,
		arg.GroupID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPolicyViolationsByGroupRow{}
	for rows.Next() {
		var i GetPolicyViolationsByGroupRow
		if err := rows.Scan(
			&i.PolicyType,
			&i.RepoID,
			&i.RepoOwner,
			&i.RepoName,
			&i.Metadata,
			&i.Violation,
			&i.CreatedAt,
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

const getPolicyViolationsById = `-- name: GetPolicyViolationsById :many
SELECT pt.policy_type, r.id as repo_id, r.repo_owner, r.repo_name,
pv.metadata, pv.violation, pv.created_at FROM policy_violations pv
INNER JOIN policies p ON p.id = pv.policy_id
INNER JOIN repositories r ON r.id = pv.repository_id
INNER JOIN policy_types pt ON pt.id = p.policy_type
WHERE p.id = $1 ORDER BY pv.created_at DESC LIMIT $2 OFFSET $3
`

type GetPolicyViolationsByIdParams struct {
	ID     int32 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPolicyViolationsByIdRow struct {
	PolicyType string          `json:"policy_type"`
	RepoID     int32           `json:"repo_id"`
	RepoOwner  string          `json:"repo_owner"`
	RepoName   string          `json:"repo_name"`
	Metadata   json.RawMessage `json:"metadata"`
	Violation  json.RawMessage `json:"violation"`
	CreatedAt  time.Time       `json:"created_at"`
}

func (q *Queries) GetPolicyViolationsById(ctx context.Context, arg GetPolicyViolationsByIdParams) ([]GetPolicyViolationsByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getPolicyViolationsById, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPolicyViolationsByIdRow{}
	for rows.Next() {
		var i GetPolicyViolationsByIdRow
		if err := rows.Scan(
			&i.PolicyType,
			&i.RepoID,
			&i.RepoOwner,
			&i.RepoName,
			&i.Metadata,
			&i.Violation,
			&i.CreatedAt,
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

const getPolicyViolationsByRepositoryId = `-- name: GetPolicyViolationsByRepositoryId :many
SELECT pt.policy_type, r.id as repo_id, r.repo_owner, r.repo_name,
pv.metadata, pv.violation, pv.created_at FROM policy_violations pv
INNER JOIN policies p ON p.id = pv.policy_id
INNER JOIN repositories r ON r.id = pv.repository_id
INNER JOIN policy_types pt ON pt.id = p.policy_type
WHERE r.id = $1 ORDER BY pv.created_at DESC LIMIT $2 OFFSET $3
`

type GetPolicyViolationsByRepositoryIdParams struct {
	ID     int32 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPolicyViolationsByRepositoryIdRow struct {
	PolicyType string          `json:"policy_type"`
	RepoID     int32           `json:"repo_id"`
	RepoOwner  string          `json:"repo_owner"`
	RepoName   string          `json:"repo_name"`
	Metadata   json.RawMessage `json:"metadata"`
	Violation  json.RawMessage `json:"violation"`
	CreatedAt  time.Time       `json:"created_at"`
}

func (q *Queries) GetPolicyViolationsByRepositoryId(ctx context.Context, arg GetPolicyViolationsByRepositoryIdParams) ([]GetPolicyViolationsByRepositoryIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getPolicyViolationsByRepositoryId, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPolicyViolationsByRepositoryIdRow{}
	for rows.Next() {
		var i GetPolicyViolationsByRepositoryIdRow
		if err := rows.Scan(
			&i.PolicyType,
			&i.RepoID,
			&i.RepoOwner,
			&i.RepoName,
			&i.Metadata,
			&i.Violation,
			&i.CreatedAt,
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
