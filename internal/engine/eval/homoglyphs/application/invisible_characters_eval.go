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

package application

import (
	"context"
	"fmt"

	"github.com/stacklok/minder/internal/engine/eval/homoglyphs/communication"
	"github.com/stacklok/minder/internal/engine/eval/homoglyphs/domain"
	engif "github.com/stacklok/minder/internal/engine/interfaces"
	"github.com/stacklok/minder/internal/providers"
)

// InvisibleCharactersEvaluator is an evaluator for the invisible characters rule type
type InvisibleCharactersEvaluator struct {
	processor     domain.HomoglyphProcessor
	reviewHandler *communication.GhReviewPrHandler
}

// NewInvisibleCharactersEvaluator creates a new invisible characters evaluator
func NewInvisibleCharactersEvaluator(pbuild *providers.ProviderBuilder) (*InvisibleCharactersEvaluator, error) {
	if pbuild == nil {
		return nil, fmt.Errorf("provider builder is nil")
	}

	ghClient, err := pbuild.GetGitHub(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not fetch GitHub client: %w", err)
	}

	return &InvisibleCharactersEvaluator{
		processor:     domain.NewInvisibleCharactersProcessor(),
		reviewHandler: communication.NewGhReviewPrHandler(ghClient),
	}, nil
}

// Eval evaluates the invisible characters rule type
func (ice *InvisibleCharactersEvaluator) Eval(ctx context.Context, _ map[string]any, res *engif.Result) error {
	return evaluateHomoglyphs(ctx, ice.processor, res, ice.reviewHandler)
}
