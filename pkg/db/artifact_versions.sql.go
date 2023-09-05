// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: artifact_versions.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createArtifactVersion = `-- name: CreateArtifactVersion :one
INSERT INTO artifact_versions (
    artifact_id,
    version,
    tags,
    sha,
    signature_verification,
    github_workflow, created_at) VALUES ($1, $2, $3, $4,
    $6::jsonb,
    $7::jsonb,
    $5) RETURNING id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at
`

type CreateArtifactVersionParams struct {
	ArtifactID            int32           `json:"artifact_id"`
	Version               int64           `json:"version"`
	Tags                  sql.NullString  `json:"tags"`
	Sha                   string          `json:"sha"`
	CreatedAt             time.Time       `json:"created_at"`
	SignatureVerification json.RawMessage `json:"signature_verification"`
	GithubWorkflow        json.RawMessage `json:"github_workflow"`
}

func (q *Queries) CreateArtifactVersion(ctx context.Context, arg CreateArtifactVersionParams) (ArtifactVersion, error) {
	row := q.db.QueryRowContext(ctx, createArtifactVersion,
		arg.ArtifactID,
		arg.Version,
		arg.Tags,
		arg.Sha,
		arg.CreatedAt,
		arg.SignatureVerification,
		arg.GithubWorkflow,
	)
	var i ArtifactVersion
	err := row.Scan(
		&i.ID,
		&i.ArtifactID,
		&i.Version,
		&i.Tags,
		&i.Sha,
		&i.SignatureVerification,
		&i.GithubWorkflow,
		&i.CreatedAt,
	)
	return i, err
}

const deleteArtifactVersion = `-- name: DeleteArtifactVersion :exec
DELETE FROM artifact_versions
WHERE id = $1
`

func (q *Queries) DeleteArtifactVersion(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteArtifactVersion, id)
	return err
}

const deleteOldArtifactVersions = `-- name: DeleteOldArtifactVersions :exec
DELETE FROM artifact_versions
WHERE artifact_id = $1 AND created_at <= $2
`

type DeleteOldArtifactVersionsParams struct {
	ArtifactID int32     `json:"artifact_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (q *Queries) DeleteOldArtifactVersions(ctx context.Context, arg DeleteOldArtifactVersionsParams) error {
	_, err := q.db.ExecContext(ctx, deleteOldArtifactVersions, arg.ArtifactID, arg.CreatedAt)
	return err
}

const getArtifactVersionByID = `-- name: GetArtifactVersionByID :one
SELECT id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at FROM artifact_versions WHERE id = $1
`

func (q *Queries) GetArtifactVersionByID(ctx context.Context, id int32) (ArtifactVersion, error) {
	row := q.db.QueryRowContext(ctx, getArtifactVersionByID, id)
	var i ArtifactVersion
	err := row.Scan(
		&i.ID,
		&i.ArtifactID,
		&i.Version,
		&i.Tags,
		&i.Sha,
		&i.SignatureVerification,
		&i.GithubWorkflow,
		&i.CreatedAt,
	)
	return i, err
}

const getArtifactVersionBySha = `-- name: GetArtifactVersionBySha :one
SELECT id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at FROM artifact_versions WHERE artifact_id = $1 AND sha = $2
`

type GetArtifactVersionByShaParams struct {
	ArtifactID int32  `json:"artifact_id"`
	Sha        string `json:"sha"`
}

func (q *Queries) GetArtifactVersionBySha(ctx context.Context, arg GetArtifactVersionByShaParams) (ArtifactVersion, error) {
	row := q.db.QueryRowContext(ctx, getArtifactVersionBySha, arg.ArtifactID, arg.Sha)
	var i ArtifactVersion
	err := row.Scan(
		&i.ID,
		&i.ArtifactID,
		&i.Version,
		&i.Tags,
		&i.Sha,
		&i.SignatureVerification,
		&i.GithubWorkflow,
		&i.CreatedAt,
	)
	return i, err
}

const listArtifactVersionsByArtifactID = `-- name: ListArtifactVersionsByArtifactID :many
SELECT id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at FROM artifact_versions
WHERE artifact_id = $1
ORDER BY created_at DESC
LIMIT COALESCE($2::int, 2147483647)
`

type ListArtifactVersionsByArtifactIDParams struct {
	ArtifactID int32         `json:"artifact_id"`
	Limit      sql.NullInt32 `json:"limit"`
}

func (q *Queries) ListArtifactVersionsByArtifactID(ctx context.Context, arg ListArtifactVersionsByArtifactIDParams) ([]ArtifactVersion, error) {
	rows, err := q.db.QueryContext(ctx, listArtifactVersionsByArtifactID, arg.ArtifactID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ArtifactVersion{}
	for rows.Next() {
		var i ArtifactVersion
		if err := rows.Scan(
			&i.ID,
			&i.ArtifactID,
			&i.Version,
			&i.Tags,
			&i.Sha,
			&i.SignatureVerification,
			&i.GithubWorkflow,
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

const listArtifactVersionsByArtifactIDAndTag = `-- name: ListArtifactVersionsByArtifactIDAndTag :many
SELECT id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at FROM artifact_versions
WHERE artifact_id = $1
AND $2=ANY(STRING_TO_ARRAY(tags, ','))
ORDER BY created_at DESC
LIMIT COALESCE($3::int, 2147483647)
`

type ListArtifactVersionsByArtifactIDAndTagParams struct {
	ArtifactID int32          `json:"artifact_id"`
	Tags       sql.NullString `json:"tags"`
	Limit      sql.NullInt32  `json:"limit"`
}

func (q *Queries) ListArtifactVersionsByArtifactIDAndTag(ctx context.Context, arg ListArtifactVersionsByArtifactIDAndTagParams) ([]ArtifactVersion, error) {
	rows, err := q.db.QueryContext(ctx, listArtifactVersionsByArtifactIDAndTag, arg.ArtifactID, arg.Tags, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ArtifactVersion{}
	for rows.Next() {
		var i ArtifactVersion
		if err := rows.Scan(
			&i.ID,
			&i.ArtifactID,
			&i.Version,
			&i.Tags,
			&i.Sha,
			&i.SignatureVerification,
			&i.GithubWorkflow,
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

const upsertArtifactVersion = `-- name: UpsertArtifactVersion :one
INSERT INTO artifact_versions (
    artifact_id,
    version,
    tags,
    sha,
    signature_verification,
    github_workflow,
    created_at
) VALUES ($1, $2, $3, $4,
    $6::jsonb,
    $7::jsonb,
    $5)
ON CONFLICT (artifact_id, sha)
DO UPDATE SET
    version = $2,
    tags = $3,
    signature_verification = $6::jsonb,
    github_workflow = $7::jsonb,
    created_at = $5
WHERE artifact_versions.artifact_id = $1 AND artifact_versions.sha = $4
RETURNING id, artifact_id, version, tags, sha, signature_verification, github_workflow, created_at
`

type UpsertArtifactVersionParams struct {
	ArtifactID            int32           `json:"artifact_id"`
	Version               int64           `json:"version"`
	Tags                  sql.NullString  `json:"tags"`
	Sha                   string          `json:"sha"`
	CreatedAt             time.Time       `json:"created_at"`
	SignatureVerification json.RawMessage `json:"signature_verification"`
	GithubWorkflow        json.RawMessage `json:"github_workflow"`
}

func (q *Queries) UpsertArtifactVersion(ctx context.Context, arg UpsertArtifactVersionParams) (ArtifactVersion, error) {
	row := q.db.QueryRowContext(ctx, upsertArtifactVersion,
		arg.ArtifactID,
		arg.Version,
		arg.Tags,
		arg.Sha,
		arg.CreatedAt,
		arg.SignatureVerification,
		arg.GithubWorkflow,
	)
	var i ArtifactVersion
	err := row.Scan(
		&i.ID,
		&i.ArtifactID,
		&i.Version,
		&i.Tags,
		&i.Sha,
		&i.SignatureVerification,
		&i.GithubWorkflow,
		&i.CreatedAt,
	)
	return i, err
}
