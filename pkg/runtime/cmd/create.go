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
		Short: "Create Postman resources.",
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
		generateCreateSubcommand(resources.MockType, "mock", []string{}),
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

	return &cmd
}

func createResource(t resources.ResourceType) error {
	if inputReader == nil {
		r, err := os.Open(inputFile)
		defer r.Close()

		if err != nil {
			return err
		}

		inputReader = r
	}

	var (
		id  string
		err error
	)

	switch t {
	case resources.CollectionType:
		id, err = service.CreateCollectionFromReader(context.Background(), inputReader, usingWorkspace)
	case resources.EnvironmentType:
		id, err = service.CreateEnvironmentFromReader(context.Background(), inputReader, usingWorkspace)
	case resources.MockType:
		id, err = service.CreateMockFromReader(context.Background(), inputReader, usingWorkspace)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(id)

	return nil
}
