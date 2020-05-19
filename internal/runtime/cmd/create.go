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
	"errors"
	"fmt"
	"os"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/spf13/cobra"
)

func init() {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create new Postman resources.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				inputReader = os.Stdin
			} else {
				if inputFile == "" {
					return errors.New("flag \"filename\" not set, use \"--filename\" or stdin")
				}
			}

			return nil
		},
	}
	createCmd.PersistentFlags().StringVarP(&inputFile, "filename", "f", "", "the filename used to create the resource (required when not using data from stdin)")

	createCmd.AddCommand(
		generateCreateSubcommand(resources.CollectionType, "collection", []string{"co"}),
		generateCreateSubcommand(resources.EnvironmentType, "environment", []string{"env"}),
		generateCreateSubcommand(resources.MonitorType, "monitor", []string{"mon"}),
		generateCreateSubcommand(resources.MockType, "mock", []string{}),
		generateCreateSubcommand(resources.WorkspaceType, "workspace", []string{"ws"}),
		generateCreateSubcommand(resources.APIType, "api", []string{}),
		generateCreateSubcommand(resources.APIVersionType, "api-version", []string{}),
		generateCreateSubcommand(resources.SchemaType, "schema", []string{}),
	)

	rootCmd.AddCommand(createCmd)
}

func generateCreateSubcommand(t resources.ResourceType, use string, aliases []string) *cobra.Command {
	cmd := cobra.Command{
		Use:     use,
		Aliases: aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createResource(t)
		},
	}

	cmd.Flags().StringVarP(&usingWorkspace, "workspace", "w", "", "workspace for create operation")

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

func createResource(t resources.ResourceType) error {
	if inputReader == nil {
		r, err := os.Open(inputFile)

		if err != nil {
			return err
		}

		defer r.Close()

		inputReader = r
	}

	var (
		id  string
		err error
	)

	ctx := context.Background()
	switch t {
	case resources.CollectionType:
		id, err = service.CreateCollectionFromReader(ctx, inputReader, usingWorkspace)
	case resources.EnvironmentType:
		id, err = service.CreateEnvironmentFromReader(ctx, inputReader, usingWorkspace)
	case resources.MockType:
		id, err = service.CreateMockFromReader(ctx, inputReader, usingWorkspace)
	case resources.MonitorType:
		id, err = service.CreateMonitorFromReader(ctx, inputReader, usingWorkspace)
	case resources.WorkspaceType:
		id, err = service.CreateWorkspaceFromReader(ctx, inputReader, usingWorkspace)
	case resources.APIType:
		id, err = service.CreateAPIFromReader(ctx, inputReader, usingWorkspace)
	case resources.APIVersionType:
		id, err = service.CreateAPIVersionFromReader(ctx, inputReader, usingWorkspace, forAPI)
	case resources.SchemaType:
		id, err = service.CreateSchemaFromReader(ctx, inputReader, usingWorkspace, forAPI, forAPIVersion)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(id)

	return nil
}
