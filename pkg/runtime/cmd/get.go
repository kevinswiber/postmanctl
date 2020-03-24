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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
	"github.com/kevinswiber/postmanctl/pkg/sdk/printers"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/jsonpath"
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
	if v == "json" || strings.HasPrefix(v, "jsonpath=") {
		o.value = v
		return nil
	}

	return errors.New("output format must be json or jsonpath")
}

// Type returns the type of this value.
func (o *OutputFormatValue) Type() string {
	return "string"
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve Postman resources.",
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
			if len(args) > 0 {
				return getSingleAPIVersion(forAPI, args...)
			}

			return getAllAPIVersions(forAPI)
		},
	}

	apiVersionsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiVersionsCmd.MarkFlagRequired("for-api")

	schemaCmd := &cobra.Command{
		Use:  "schema",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getSchema(forAPI, forAPIVersion, args[0])
		},
	}

	schemaCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	schemaCmd.MarkFlagRequired("for-api")

	schemaCmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
	schemaCmd.MarkFlagRequired("for-api-version")

	apiRelationsCmd := &cobra.Command{
		Use: "api-relations",
		RunE: func(cmd *cobra.Command, args []string) error {
			if outputFormat.value == "" {
				return getFormattedAPIRelations(forAPI, forAPIVersion)
			}
			return getAPIRelations(forAPI, forAPIVersion)
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
				if len(args) > 0 {
					return getSingleCollection(args...)
				}

				return getAllCollections()
			},
		},
		&cobra.Command{
			Use:     "environments",
			Aliases: []string{"environment", "env"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getSingleEnvironment(args...)
				}

				return getAllEnvironments()
			},
		},
		&cobra.Command{
			Use:     "monitors",
			Aliases: []string{"monitor", "mon"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getSingleMonitor(args...)
				}

				return getAllMonitors()
			},
		},
		&cobra.Command{
			Use:     "mocks",
			Aliases: []string{"mock"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getSingleMock(args...)
				}

				return getAllMocks()
			},
		},
		&cobra.Command{
			Use:     "workspaces",
			Aliases: []string{"workspace", "ws"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getSingleWorkspace(args...)
				}

				return getAllWorkspaces()
			},
		},
		&cobra.Command{
			Use: "user",
			RunE: func(cmd *cobra.Command, args []string) error {
				return getUser()
			},
		},
		&cobra.Command{
			Use:     "apis",
			Aliases: []string{"api"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getSingleAPI(args...)
				}

				return getAllAPIs()
			},
		},
		apiVersionsCmd,
		apiRelationsCmd,
		schemaCmd,
	)

	getCmd.PersistentFlags().VarP(&outputFormat, "output", "o", "output format (json, jsonpath)")
	rootCmd.AddCommand(getCmd)
}

func getAllCollections() error {
	resource, err := service.Collections(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleCollection(ids ...string) error {
	r := make(resources.CollectionSlice, len(ids))
	for i, id := range ids {
		resource, err := service.Collection(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getAllEnvironments() error {
	resource, err := service.Environments(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleEnvironment(ids ...string) error {
	r := make(resources.EnvironmentSlice, len(ids))
	for i, id := range ids {
		resource, err := service.Environment(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getAllAPIs() error {
	resource, err := service.APIs(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleAPI(ids ...string) error {
	r := make(resources.APISlice, len(ids))
	for i, id := range ids {
		resource, err := service.API(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getAllWorkspaces() error {
	resource, err := service.Workspaces(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleWorkspace(ids ...string) error {
	r := make(resources.WorkspaceSlice, len(ids))
	for i, id := range ids {
		resource, err := service.Workspace(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getAllMonitors() error {
	resource, err := service.Monitors(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleMonitor(ids ...string) error {
	r := make(resources.MonitorSlice, len(ids))
	for i, id := range ids {
		resource, err := service.Monitor(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getAllMocks() error {
	resource, err := service.Mocks(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleMock(ids ...string) error {
	r := make(resources.MockSlice, len(ids))
	for i, id := range ids {
		resource, err := service.Mock(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getUser() error {
	resource, err := service.User(context.Background())

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getAllAPIVersions(apiID string) error {
	resource, err := service.APIVersions(context.Background(), apiID)

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getSingleAPIVersion(apiID string, ids ...string) error {
	r := make(resources.APIVersionSlice, len(ids))
	for i, id := range ids {
		resource, err := service.APIVersion(context.Background(), apiID, id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	print(r)

	return nil
}

func getSchema(apiID, apiVersionID, id string) error {
	resource, err := service.Schema(context.Background(), apiID, apiVersionID, id)

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getAPIRelations(apiID, apiVersionID string) error {
	resource, err := service.APIRelations(context.Background(), apiID, apiVersionID)

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getFormattedAPIRelations(apiID, apiVersionID string) error {
	resource, err := service.FormattedAPIRelationItems(context.Background(), apiID, apiVersionID)

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func handleResponseError(err error) error {
	if err, ok := err.(*client.RequestError); ok {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}

	return err
}

func print(r interface{}) {
	if r == nil {
		return
	}

	if outputFormat.value == "json" {
		t, err := json.MarshalIndent(&r, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		fmt.Println(string(t))
	} else if strings.HasPrefix(outputFormat.value, "jsonpath=") {
		tmpl := outputFormat.value[9:]
		j := jsonpath.New("out")
		if err := j.Parse(tmpl); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		t, err := json.Marshal(&r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		var queryObj interface{}
		queryObj = map[string]interface{}{}
		if err := json.Unmarshal(t, &queryObj); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		buf := bytes.NewBuffer(nil)
		j.Execute(buf, queryObj)

		fmt.Println(buf)
	} else {
		var f resources.Formatter
		f = r.(resources.Formatter)
		printTable(f)
	}
}

func printTable(f resources.Formatter) {
	w := printers.GetNewTabWriter(os.Stdout)
	printer := printers.NewTablePrinter(printers.PrintOptions{})
	printer.PrintResource(f, w)
}
