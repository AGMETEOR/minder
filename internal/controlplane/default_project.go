//
// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stacklok/minder/internal/authz"
	"github.com/stacklok/minder/internal/db"
	github "github.com/stacklok/minder/internal/providers/github"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

// OrgMeta is the metadata associated with an organization
type OrgMeta struct {
	Company string `json:"company"`
}

// ProjectMeta is the metadata associated with a project
type ProjectMeta struct {
	Description string `json:"description"`
}

// CreateDefaultRecordsForOrg creates the default records, such as projects, roles and provider for the organization
func (s *Server) CreateDefaultRecordsForOrg(ctx context.Context, qtx db.Querier,
	org db.Project, projectName string, userSub string) (outproj *pb.Project, projerr error) {
	projectmeta := &ProjectMeta{
		Description: fmt.Sprintf("Default admin project for %s", org.Name),
	}

	jsonmeta, err := json.Marshal(projectmeta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal meta: %v", err)
	}

	projectID := uuid.New()

	// Create authorization tuple
	// NOTE: This is only creating a tuple for the project, not the organization
	//       We currently have no use for the organization and it might be
	//       removed in the future.
	if err := s.authzClient.Write(ctx, userSub, authz.AuthzRoleAdmin, projectID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create authorization tuple: %v", err)
	}
	defer func() {
		if outproj == nil && projerr != nil {
			if err := s.authzClient.Delete(ctx, userSub, authz.AuthzRoleAdmin, projectID); err != nil {
				log.Ctx(ctx).Error().Err(err).Msg("failed to delete authorization tuple")
			}
		}
	}()

	// we need to create the default records for the organization
	project, err := qtx.CreateProjectWithID(ctx, db.CreateProjectWithIDParams{
		ParentID: uuid.NullUUID{
			UUID:  org.ID,
			Valid: true,
		},
		ID:       projectID,
		Name:     projectName,
		Metadata: jsonmeta,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create default project: %v", err)
	}

	prj := pb.Project{
		ProjectId:   project.ID.String(),
		Name:        project.Name,
		Description: projectmeta.Description,
		CreatedAt:   timestamppb.New(project.CreatedAt),
		UpdatedAt:   timestamppb.New(project.UpdatedAt),
	}

	if err != nil {
		if err := s.authzClient.Delete(ctx, userSub, authz.AuthzRoleAdmin, projectID); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("failed to delete authorization tuple")
		}
		return nil, status.Errorf(codes.Internal, "failed to create default project role: %v", err)
	}

	// Create GitHub provider
	_, err = qtx.CreateProvider(ctx, db.CreateProviderParams{
		Name:       github.Github,
		ProjectID:  project.ID,
		Implements: github.Implements,
		Definition: json.RawMessage(`{"github": {}}`),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create provider: %v", err)
	}
	return &prj, nil
}
