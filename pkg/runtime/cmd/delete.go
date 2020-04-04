/*
Copyright © 2020 Kevin Swiber <kswiber@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a Postman resource.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete existing Postman resources.",
	}

	deleteCmd.AddCommand(
		generateDeleteSubcommand(resources.CollectionType, "collection", []string{"co"}),
		generateDeleteSubcommand(resources.EnvironmentType, "environment", []string{"env"}),
		generateDeleteSubcommand(resources.MonitorType, "monitor", []string{"mon"}),
		generateDeleteSubcommand(resources.MockType, "mock", []string{}),
		generateDeleteSubcommand(resources.WorkspaceType, "workspace", []string{"ws"}),
		generateDeleteSubcommand(resources.APIType, "api", []string{}),
		generateDeleteSubcommand(resources.APIVersionType, "api-version", []string{}),
		generateDeleteSubcommand(resources.SchemaType, "schema", []string{}),
	)

	rootCmd.AddCommand(deleteCmd)
}

func generateDeleteSubcommand(t resources.ResourceType, use string, aliases []string) *cobra.Command {
	cmd := cobra.Command{
		Use:     use,
		Aliases: aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteResource(t, args[0])
		},
	}

	if t == resources.APIVersionType || t == resources.SchemaType {
		cmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
		cmd.MarkFlagRequired("for-api")
	}

	if t == resources.SchemaType {
		cmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
		cmd.MarkFlagRequired("for-api-version")
	}

	return &cmd
}

func deleteResource(t resources.ResourceType, resourceID string) error {
	var (
		id  string
		err error
	)

	ctx := context.Background()
	switch t {
	case resources.CollectionType:
		id, err = service.DeleteCollection(ctx, resourceID)
	case resources.EnvironmentType:
		id, err = service.DeleteEnvironment(ctx, resourceID)
	case resources.MockType:
		id, err = service.DeleteMock(ctx, resourceID)
	case resources.MonitorType:
		id, err = service.DeleteMonitor(ctx, resourceID)
	case resources.APIType:
		id, err = service.DeleteAPI(ctx, resourceID)
	case resources.APIVersionType:
		id, err = service.DeleteAPIVersion(ctx, forAPI, resourceID)
	case resources.SchemaType:
		id, err = service.DeleteSchema(ctx, forAPI, forAPIVersion, resourceID)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(id)

	return nil
}
