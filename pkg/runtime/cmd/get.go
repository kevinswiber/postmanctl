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
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
	"github.com/kevinswiber/postmanctl/pkg/sdk/printers"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/liggitt/tabwriter"
)

var outputFormat OutputFormatValue

// OutputFormatValue is a custom Value for the output flag that validates.
type OutputFormatValue struct {
	value string
}

// String returns a string representation of this flag.
func (o *OutputFormatValue) String() string {
	return o.value
}

// Set creates the flag value.
func (o *OutputFormatValue) Set(v string) error {
	if v == "json" || v == "yaml" {
		o.value = v
		return nil
	}

	return errors.New("output format must be json, yaml, or jsonpath")
}

// Type returns the type of this value.
func (o *OutputFormatValue) Type() string {
	return "string"
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must specify the type of resource to get (e.g., collections, environments, monitors, mocks)")
	},
}

var (
	forAPI        string
	forAPIVersion string
)

func init() {
	apiVersionsCmd := &cobra.Command{
		Use:     "api-versions",
		Aliases: []string{"api-version"},
		RunE: func(cmd *cobra.Command, args []string) error {
			w := printers.GetNewTabWriter(os.Stdout)

			if len(args) > 0 {
				return getSingleAPIVersion(w, forAPI, args[0])
			}

			return getAllAPIVersions(w, forAPI)
		},
	}

	apiVersionsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiVersionsCmd.MarkFlagRequired("for-api")

	schemaCmd := &cobra.Command{
		Use:  "schema",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			w := printers.GetNewTabWriter(os.Stdout)

			return getSchema(w, forAPI, forAPIVersion, args[0])
		},
	}

	schemaCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	schemaCmd.MarkFlagRequired("for-api")

	schemaCmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
	schemaCmd.MarkFlagRequired("for-api-version")

	apiRelationsCmd := &cobra.Command{
		Use: "api-relations",
		RunE: func(cmd *cobra.Command, args []string) error {
			w := printers.GetNewTabWriter(os.Stdout)

			return getAPIRelations(w, forAPI, forAPIVersion)
		},
	}

	apiRelationsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiRelationsCmd.MarkFlagRequired("for-api")

	apiRelationsCmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
	apiRelationsCmd.MarkFlagRequired("for-api-version")

	getCmd.AddCommand(
		&cobra.Command{
			Use:     "collections",
			Aliases: []string{"collection", "co"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleCollection(w, args[0])
				}

				return getAllCollections(w)
			},
		},
		&cobra.Command{
			Use:     "environments",
			Aliases: []string{"environment", "env"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleEnvironment(w, args[0])
				}

				return getAllEnvironments(w)
			},
		},
		&cobra.Command{
			Use:     "monitors",
			Aliases: []string{"monitor", "mon"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleMonitor(w, args[0])
				}

				return getAllMonitors(w)
			},
		},
		&cobra.Command{
			Use:     "mocks",
			Aliases: []string{"mock"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleMock(w, args[0])
				}

				return getAllMocks(w)
			},
		},
		&cobra.Command{
			Use:     "workspaces",
			Aliases: []string{"workspace", "ws"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleWorkspace(w, args[0])
				}

				return getAllWorkspaces(w)
			},
		},
		&cobra.Command{
			Use: "user",
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				return getUser(w)
			},
		},
		&cobra.Command{
			Use:     "apis",
			Aliases: []string{"api"},
			RunE: func(cmd *cobra.Command, args []string) error {
				w := printers.GetNewTabWriter(os.Stdout)

				if len(args) > 0 {
					return getSingleAPI(w, args[0])
				}

				return getAllAPIs(w)
			},
		},
		apiVersionsCmd,
		apiRelationsCmd,
		schemaCmd,
	)

	getCmd.PersistentFlags().VarP(&outputFormat, "output", "o", "output format (json, yaml, jsonpath)")
	rootCmd.AddCommand(getCmd)
}

func getAllCollections(w *tabwriter.Writer) error {
	resource, err := service.Collections(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleCollection(w *tabwriter.Writer, id string) error {
	resource, err := service.Collection(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllEnvironments(w *tabwriter.Writer) error {
	resource, err := service.Environments(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleEnvironment(w *tabwriter.Writer, id string) error {
	resource, err := service.Environment(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllAPIs(w *tabwriter.Writer) error {
	resource, err := service.APIs(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleAPI(w *tabwriter.Writer, id string) error {
	resource, err := service.API(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllWorkspaces(w *tabwriter.Writer) error {
	resource, err := service.Workspaces(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleWorkspace(w *tabwriter.Writer, id string) error {
	resource, err := service.Workspace(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllMonitors(w *tabwriter.Writer) error {
	resource, err := service.Monitors(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleMonitor(w *tabwriter.Writer, id string) error {
	resource, err := service.Monitor(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllMocks(w *tabwriter.Writer) error {
	resource, err := service.Mocks(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleMock(w *tabwriter.Writer, id string) error {
	resource, err := service.Mock(context.Background(), id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getUser(w *tabwriter.Writer) error {
	resource, err := service.User(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAllAPIVersions(w *tabwriter.Writer, apiID string) error {
	resource, err := service.APIVersions(context.Background(), apiID)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSingleAPIVersion(w *tabwriter.Writer, apiID, id string) error {
	resource, err := service.APIVersion(context.Background(), apiID, id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getSchema(w *tabwriter.Writer, apiID, apiVersionID, id string) error {
	resource, err := service.Schema(context.Background(), apiID, apiVersionID, id)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func getAPIRelations(w *tabwriter.Writer, apiID, apiVersionID string) error {
	resource, err := service.FormattedAPIRelationItems(context.Background(), apiID, apiVersionID)

	if err != nil {
		return handleResponseError(err)
	}

	print(w, resource)

	return nil
}

func handleResponseError(err error) error {
	if err, ok := err.(*client.RequestError); ok {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}

	return err
}

func print(w io.Writer, f resources.Formatter) {
	if f == nil {
		return
	}
	printer := printers.NewTablePrinter(printers.PrintOptions{})
	printer.PrintResource(f, w)
}
