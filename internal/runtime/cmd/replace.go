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
	replaceCmd := &cobra.Command{
		Use:   "replace",
		Short: "Replace existing Postman resources.",
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
	replaceCmd.PersistentFlags().StringVarP(&inputFile, "filename", "f", "", "the filename used to replace the resource (required when not using data from stdin)")

	replaceCmd.AddCommand(
		generateReplaceSubcommand(resources.CollectionType, "collection", []string{"co"}),
		generateReplaceSubcommand(resources.EnvironmentType, "environment", []string{"env"}),
		generateReplaceSubcommand(resources.MonitorType, "monitor", []string{"mon"}),
		generateReplaceSubcommand(resources.MockType, "mock", []string{}),
		generateReplaceSubcommand(resources.WorkspaceType, "workspace", []string{"ws"}),
		generateReplaceSubcommand(resources.APIType, "api", []string{}),
		generateReplaceSubcommand(resources.APIVersionType, "api-version", []string{}),
		generateReplaceSubcommand(resources.SchemaType, "schema", []string{}),
	)

	rootCmd.AddCommand(replaceCmd)
}

func generateReplaceSubcommand(t resources.ResourceType, use string, aliases []string) *cobra.Command {
	cmd := cobra.Command{
		Use:     use,
		Aliases: aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return replaceResource(t, args[0])
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

func replaceResource(t resources.ResourceType, resourceID string) error {
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
		id, err = service.ReplaceCollectionFromReader(ctx, inputReader, resourceID)
	case resources.EnvironmentType:
		id, err = service.ReplaceEnvironmentFromReader(ctx, inputReader, resourceID)
	case resources.MockType:
		id, err = service.ReplaceMockFromReader(ctx, inputReader, resourceID)
	case resources.MonitorType:
		id, err = service.ReplaceMonitorFromReader(ctx, inputReader, resourceID)
	case resources.WorkspaceType:
		id, err = service.ReplaceWorkspaceFromReader(ctx, inputReader, resourceID)
	case resources.APIType:
		id, err = service.ReplaceAPIFromReader(ctx, inputReader, resourceID)
	case resources.APIVersionType:
		id, err = service.ReplaceAPIVersionFromReader(ctx, inputReader, forAPI, resourceID)
	case resources.SchemaType:
		id, err = service.ReplaceSchemaFromReader(ctx, inputReader, resourceID, forAPI, forAPIVersion)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(id)

	return nil
}
