// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package rule provides the CLI subcommand for managing rules

// Package ingester provides necessary interfaces and implementations for ingesting
// data for rules.
package ingester

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stacklok/minder/internal/db"
	"github.com/stacklok/minder/internal/engine/ingester/artifact"
	"github.com/stacklok/minder/internal/engine/ingester/builtin"
	"github.com/stacklok/minder/internal/engine/ingester/git"
	"github.com/stacklok/minder/internal/engine/ingester/rest"
	"github.com/stacklok/minder/internal/providers"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	provifv1 "github.com/stacklok/minder/pkg/providers/v1"
)

func TestNewRuleDataIngest(t *testing.T) {
	t.Parallel()

	type args struct {
		rt *pb.RuleType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "artifact",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type:     artifact.ArtifactRuleDataIngestType,
							Artifact: &pb.ArtifactType{},
						},
					},
				},
			},
		},
		{
			name: "artifact missing",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: artifact.ArtifactRuleDataIngestType,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "rest",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: rest.RestRuleDataIngestType,
							Rest: &pb.RestType{
								Endpoint: "https://api.github.com/repos/Foo/Bar",
							},
						},
					},
				},
			},
		},
		{
			name: "rest missing",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: rest.RestRuleDataIngestType,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "builtin",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type:    builtin.BuiltinRuleDataIngestType,
							Builtin: &pb.BuiltinType{},
						},
					},
				},
			},
		},
		{
			name: "builtin missing",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: builtin.BuiltinRuleDataIngestType,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "git",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: git.GitRuleDataIngestType,
							Git: &pb.GitType{
								CloneUrl: "https://github.com/staklok/minder.git",
							},
						},
					},
				},
			},
		},
		{
			name: "unsupported",
			args: args{
				rt: &pb.RuleType{
					Def: &pb.RuleType_Definition{
						Ingest: &pb.RuleType_Definition_Ingest{
							Type: "unsupported",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewRuleDataIngest(tt.args.rt, providers.NewProviderBuilder(
				&db.Provider{
					Name:    "github",
					Version: provifv1.V1,
					Implements: []db.ProviderType{
						"rest",
						"git",
						"github",
					},
					Definition: json.RawMessage(`{
	"rest": {
		"endpoint": "https://api.github.com/repos/Foo/Bar"
	},
	"git": {},
	"github": {}
}`),
				},
				db.ProviderAccessToken{},
				"token",
			))
			if tt.wantErr {
				require.Error(t, err, "Expected error")
				require.Nil(t, got, "Expected nil")
				return
			}

			require.NoError(t, err, "Unexpected error")
			require.NotNil(t, got, "Expected non-nil")
		})
	}
}
