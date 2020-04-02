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

func init() {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create Postman resources.",
	}
	createCmd.PersistentFlags().StringVarP(&inputFile, "filename", "f", "", "the filename used to create the resource (required)")
	createCmd.MarkPersistentFlagRequired("filename")

	createCmd.AddCommand(
		generateCreateSubcommand(resources.CollectionType, "collection", []string{"co"}),
	)

	rootCmd.AddCommand(createCmd)
}

func generateCreateSubcommand(t resources.ResourceType, use string, aliases []string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Aliases: aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createResource(t)
		},
	}
}

func createResource(t resources.ResourceType) error {
	r, err := os.Open(inputFile)
	defer r.Close()

	if err != nil {
		return err
	}

	id := ""
	switch t {
	case resources.CollectionType:
		if id, err = service.CreateCollectionFromReader(context.Background(), r); err != nil {
			return err
		}
	}

	fmt.Println(id)

	return nil
}
