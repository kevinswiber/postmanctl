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
	"errors"
	"fmt"
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

func init() {
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
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get monitors")
			},
		},
		&cobra.Command{
			Use:     "mocks",
			Aliases: []string{"mock"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get mocks")
			},
		},
		&cobra.Command{
			Use:     "workspaces",
			Aliases: []string{"workspace", "ws"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get workspaces")
			},
		},
		&cobra.Command{
			Use: "user",
			RunE: func(cmd *cobra.Command, args []string) error {
				var representation resources.UserResponse
				req := client.NewRequest(apiClient)
				_, err := req.Get().
					Resource("me").
					As(&representation).
					Do()

				if err != nil {
					return handleResponseError(err)
				}

				w := printers.GetNewTabWriter(os.Stdout)

				var r resources.Resource
				r = representation.User
				options := printers.PrintOptions{}
				printer := printers.NewTablePrinter(options)

				printer.Print(r, w)

				return nil
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
		&cobra.Command{
			Use:     "api-versions",
			Aliases: []string{"api-version"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get api-versions")
			},
		},
		&cobra.Command{
			Use:     "api-relations",
			Aliases: []string{"api-relation"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get api-relations")
			},
		},
	)

	getCmd.PersistentFlags().VarP(&outputFormat, "output", "o", "output format (json, yaml, jsonpath)")
	rootCmd.AddCommand(getCmd)
}

func getAllCollections(w *tabwriter.Writer) error {
	var representation resources.CollectionListResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("collections").
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	for i, col := range representation.Collections {
		var options printers.PrintOptions
		if i == 0 {
			options = printers.PrintOptions{
				NoHeaders: false,
			}
		} else {
			options = printers.PrintOptions{
				NoHeaders: true,
			}
		}

		var res resources.Resource
		res = col
		printer := printers.NewTablePrinter(options)
		printer.Print(res, w)
	}

	return nil
}

func getSingleCollection(w *tabwriter.Writer, id string) error {
	var representation resources.CollectionResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("collections", id).
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	options := printers.PrintOptions{
		NoHeaders: false,
	}

	printer := printers.NewTablePrinter(options)

	var r resources.Resource
	r = representation.Collection.Info
	printer.Print(r, w)
	return nil
}

func getAllEnvironments(w *tabwriter.Writer) error {
	var representation resources.EnvironmentListResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("environments").
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	for i, env := range representation.Environments {
		var options printers.PrintOptions
		if i == 0 {
			options = printers.PrintOptions{
				NoHeaders: false,
			}
		} else {
			options = printers.PrintOptions{
				NoHeaders: true,
			}
		}

		var res resources.Resource
		res = env
		printer := printers.NewTablePrinter(options)
		printer.Print(res, w)
	}

	return nil
}

func getSingleEnvironment(w *tabwriter.Writer, id string) error {
	var representation resources.EnvironmentResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("environments", id).
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	options := printers.PrintOptions{
		NoHeaders: false,
	}

	var r resources.Resource
	r = representation.Environment
	printer := printers.NewTablePrinter(options)
	printer.Print(r, w)

	return nil
}

func getAllAPIs(w *tabwriter.Writer) error {
	var representation resources.APIListResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("apis").
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	for i, api := range representation.APIs {
		var options printers.PrintOptions
		if i == 0 {
			options = printers.PrintOptions{
				NoHeaders: false,
			}
		} else {
			options = printers.PrintOptions{
				NoHeaders: true,
			}
		}

		var res resources.Resource
		res = api
		printer := printers.NewTablePrinter(options)
		printer.Print(res, w)
	}

	return nil
}

func getSingleAPI(w *tabwriter.Writer, id string) error {
	var representation resources.APIResponse
	req := client.NewRequest(apiClient)
	_, err := req.Get().
		Resource("apis", id).
		As(&representation).
		Do()

	if err != nil {
		return handleResponseError(err)
	}

	options := printers.PrintOptions{
		NoHeaders: false,
	}

	var r resources.Resource
	r = representation.API
	printer := printers.NewTablePrinter(options)
	printer.Print(r, w)

	return nil
}

func handleResponseError(err error) error {
	if err, ok := err.(*client.RequestError); ok {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}

	return err
}
