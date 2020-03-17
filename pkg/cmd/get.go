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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/kevinswiber/postmanctl/pkg/client"
	"github.com/kevinswiber/postmanctl/pkg/printers"
	"github.com/kevinswiber/postmanctl/pkg/resources"
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
				req := client.NewRequest(apiClient)
				res, err := req.Get().
					Resource("me").
					Do()

				if err != nil {
					return err
				}

				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)

				if err != nil {
					return err
				}

				var u resources.UserResponse
				json.Unmarshal(body, &u)

				w := printers.GetNewTabWriter(os.Stdout)
				var r resources.Resource
				r = u.User
				options := printers.PrintOptions{}
				printer := printers.NewTablePrinter(options)
				printer.Print(r, w)
				return nil
			},
		},
		&cobra.Command{
			Use:     "apis",
			Aliases: []string{"api"},
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("get apis")
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
	req := client.NewRequest(apiClient)
	res, err := req.Get().
		Resource("collections").
		Do()

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var c resources.CollectionListResponse
	json.Unmarshal(body, &c)

	for i, col := range c.Collections {
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
	req := client.NewRequest(apiClient)
	res, err := req.Get().
		Resource("collections", id).
		Do()

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var cRes resources.CollectionResponse
	json.Unmarshal(body, &cRes)

	options := printers.PrintOptions{
		NoHeaders: false,
	}

	var r resources.Resource
	r = cRes.Collection.Info
	printer := printers.NewTablePrinter(options)
	printer.Print(r, w)
	return nil
}

func getAllEnvironments(w *tabwriter.Writer) error {
	req := client.NewRequest(apiClient)
	res, err := req.Get().
		Resource("environments").
		Do()

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var e resources.EnvironmentListResponse
	json.Unmarshal(body, &e)

	for i, env := range e.Environments {
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
	req := client.NewRequest(apiClient)
	res, err := req.Get().
		Resource("environments", id).
		Do()

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var e resources.EnvironmentResponse
	json.Unmarshal(body, &e)

	options := printers.PrintOptions{
		NoHeaders: false,
	}

	var r resources.Resource
	r = e.Environment
	printer := printers.NewTablePrinter(options)
	printer.Print(r, w)

	return nil
}
