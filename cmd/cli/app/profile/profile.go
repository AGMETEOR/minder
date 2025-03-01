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

// Package profile provides the CLI subcommand for managing profiles
package profile

import (
	"github.com/spf13/cobra"

	"github.com/stacklok/minder/cmd/cli/app"
	ghclient "github.com/stacklok/minder/internal/providers/github"
)

// ProfileCmd is the root command for the profile subcommands
var ProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
	Long:  `The profile subcommands allows the management of profiles within Minder.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Usage()
	},
}

func init() {
	app.RootCmd.AddCommand(ProfileCmd)
	// Flags for all subcommands
	ProfileCmd.PersistentFlags().StringP("provider", "p", ghclient.Github, "Name of the provider, i.e. github")
	ProfileCmd.PersistentFlags().StringP("project", "j", "", "ID of the project")
}
