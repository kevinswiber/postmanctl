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
				params := append([]string{forAPI}, args...)
				return getIndividualResources(resources.APIVersionType, params...)
			}

			return getAllResources(resources.APIVersionType, forAPI)
		},
	}

	apiVersionsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiVersionsCmd.MarkFlagRequired("for-api")

	schemaCmd := &cobra.Command{
		Use:  "schema",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			params := append([]string{forAPI, forAPIVersion}, args...)
			return getIndividualResources(resources.SchemaType, params...)
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
					return getIndividualResources(resources.CollectionType, args...)
				}

				return getAllResources(resources.CollectionType)
			},
		},
		&cobra.Command{
			Use:     "environments",
			Aliases: []string{"environment", "env"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getIndividualResources(resources.EnvironmentType, args...)
				}

				return getAllResources(resources.EnvironmentType)
			},
		},
		&cobra.Command{
			Use:     "monitors",
			Aliases: []string{"monitor", "mon"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getIndividualResources(resources.MonitorType, args...)
				}

				return getAllResources(resources.MonitorType)
			},
		},
		&cobra.Command{
			Use:     "mocks",
			Aliases: []string{"mock"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getIndividualResources(resources.MockType, args...)
				}

				return getAllResources(resources.MockType)
			},
		},
		&cobra.Command{
			Use:     "workspaces",
			Aliases: []string{"workspace", "ws"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getIndividualResources(resources.WorkspaceType, args...)
				}

				return getAllResources(resources.WorkspaceType)
			},
		},
		&cobra.Command{
			Use: "user",
			RunE: func(cmd *cobra.Command, args []string) error {
				return getIndividualResources(resources.UserType)
			},
		},
		&cobra.Command{
			Use:     "apis",
			Aliases: []string{"api"},
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return getIndividualResources(resources.APIType, args...)
				}

				return getAllResources(resources.APIType)
			},
		},
		apiVersionsCmd,
		apiRelationsCmd,
		schemaCmd,
	)

	getCmd.PersistentFlags().VarP(&outputFormat, "output", "o", "output format (json, jsonpath)")
	rootCmd.AddCommand(getCmd)
}

func getAllResources(resourceType resources.ResourceType, args ...string) error {
	ctx := context.Background()

	var resource interface{}
	var err error

	switch resourceType {
	case resources.CollectionType:
		resource, err = service.Collections(ctx)
	case resources.EnvironmentType:
		resource, err = service.Environments(ctx)
	case resources.MockType:
		resource, err = service.Mocks(ctx)
	case resources.MonitorType:
		resource, err = service.Monitors(ctx)
	case resources.APIType:
		resource, err = service.APIs(ctx)
	case resources.APIVersionType:
		resource, err = service.APIVersions(ctx, args[0])
	case resources.WorkspaceType:
		resource, err = service.Workspaces(ctx)
	default:
		return fmt.Errorf("invalid resource type: %s", resourceType.String())
	}

	if err != nil {
		return handleResponseError(err)
	}

	print(resource)

	return nil
}

func getIndividualResources(resourceType resources.ResourceType, args ...string) error {
	switch resourceType {
	case resources.CollectionType:
		r := make(resources.CollectionSlice, len(args))
		for i, id := range args {
			resource, err := service.Collection(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.EnvironmentType:
		r := make(resources.EnvironmentSlice, len(args))
		for i, id := range args {
			resource, err := service.Environment(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.MockType:
		r := make(resources.MockSlice, len(args))
		for i, id := range args {
			resource, err := service.Mock(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.MonitorType:
		r := make(resources.MonitorSlice, len(args))
		for i, id := range args {
			resource, err := service.Monitor(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.APIType:
		r := make(resources.APISlice, len(args))
		for i, id := range args {
			resource, err := service.API(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.APIVersionType:
		apiID := args[0]
		ids := args[1:]

		r := make(resources.APIVersionSlice, len(ids))
		for i, id := range ids {
			resource, err := service.APIVersion(context.Background(), apiID, id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.WorkspaceType:
		r := make(resources.WorkspaceSlice, len(args))
		for i, id := range args {
			resource, err := service.Workspace(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		print(r)
	case resources.UserType:
		resource, err := service.User(context.Background())

		if err != nil {
			return handleResponseError(err)
		}

		print(resource)
	case resources.SchemaType:
		apiID := args[0]
		apiVersionID := args[1]
		id := args[2]

		resource, err := service.Schema(context.Background(), apiID, apiVersionID, id)

		if err != nil {
			return handleResponseError(err)
		}

		print(resource)
	default:
		return fmt.Errorf("invalid resource type: %s", resourceType.String())
	}

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
