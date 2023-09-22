// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.role/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package vulncheck provides the vulnerability check evaluator
package vulncheck

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pb "github.com/stacklok/mediator/pkg/generated/protobuf/go/mediator/v1"
)

func TestSendRecvRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mockHandler http.HandlerFunc
		depName     string
		expectError bool
		expectReply *packageJson
	}{
		{
			name: "ValidResponse",
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				data := packageJson{
					Name:    "my-package",
					Version: "1.0.0",
					Dist: struct {
						Integrity string `json:"integrity"`
						Tarball   string `json:"tarball"`
					}{
						Integrity: "sha512-...",
						Tarball:   "https://example.com/path/to/tarball.tgz",
					},
				}
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(data)
				if err != nil {
					t.Fatal(err)
				}
			},
			depName: "my-package",
			expectReply: &packageJson{
				Name:    "my-package",
				Version: "1.0.0",
				Dist: struct {
					Integrity string `json:"integrity"`
					Tarball   string `json:"tarball"`
				}{
					Integrity: "sha512-...",
					Tarball:   "https://example.com/path/to/tarball.tgz",
				},
			},
		},
		{
			name: "Non200Response",
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
				_, err := w.Write([]byte("Not Found"))
				if err != nil {
					t.Fatal(err)
				}
			},
			depName:     "non-existing-package",
			expectError: true,
		},
		{
			name: "InvalidJSON",
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte("{ invalid json }"))
				if err != nil {
					t.Fatal(err)
				}
			},
			depName:     "package-with-invalid-json",
			expectError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(tt.mockHandler)
			defer server.Close()

			repo := newNpmRepository(server.URL)

			dep := &pb.Dependency{
				Name: tt.depName,
			}
			req, err := repo.NewRequest(context.Background(), dep)
			require.NoError(t, err, "Failed to create request")

			reply, err := repo.SendRecvRequest(req)
			if tt.expectError {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
				require.Equal(t, tt.expectReply.IndentedString(0), reply.IndentedString(0), "expected reply to match mock data")
			}
		})
	}
}

func TestRepoCache(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		ecoConfig      *ecosystemConfig
		expectError    bool
		expectSameRepo bool
	}{
		{
			name: "ValidEcosystemCachesConnections",
			ecoConfig: &ecosystemConfig{
				Name: "npm",
				PackageRepository: packageRepository{
					Url: "http://mock.url",
				},
			},
			expectError: false,
		},
		{
			name: "ErrorWithUnknownEcosystem",
			ecoConfig: &ecosystemConfig{
				Name: "unknown-ecosystem",
				PackageRepository: packageRepository{
					Url: "http://mock.url",
				},
			},
			expectError: true,
		},
	}

	cache := newRepoCache()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			newRepo, err := cache.newRepository(tt.ecoConfig)
			if tt.expectError {
				assert.Error(t, err, "Expected error")
				return
			}

			assert.NoError(t, err, "Failed to fetch repository for the first time")
			cachedRepo, err := cache.newRepository(tt.ecoConfig)
			assert.NoError(t, err, "Failed to fetch repository for the second time")
			assert.Equal(t, newRepo, cachedRepo, "Repositories from cache should be the same instance")
		})
	}
}