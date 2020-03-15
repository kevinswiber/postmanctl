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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/kevinswiber/postmanctl/pkg/resources"
)

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
				w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

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
				url := configContext.APIRoot + "/environments"
				method := "GET"

				client := &http.Client{}
				req, err := http.NewRequest(method, url, nil)

				if err != nil {
					fmt.Println(err)
				}
				req.Header.Add("X-Api-Key", configContext.APIKey)

				res, err := client.Do(req)
				if err != nil {
					return err
				}

				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)

				var e resources.EnvironmentList
				json.Unmarshal(body, &e)

				w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
				fmt.Fprintln(w, "ID\tNAME\tOWNER")
				for _, env := range e.Environments {
					fmt.Fprintf(w, "%s\t%s\t%s\n", env.ID, env.Name, env.Owner)
				}
				w.Flush()

				return nil
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
				url := configContext.APIRoot + "/me"
				method := "GET"

				client := &http.Client{}
				req, err := http.NewRequest(method, url, nil)

				if err != nil {
					return err
				}

				req.Header.Add("X-Api-Key", configContext.APIKey)

				res, err := client.Do(req)
				if err != nil {
					return err
				}

				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)

				var u resources.UserResponse
				json.Unmarshal(body, &u)

				w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
				fmt.Fprintln(w, "ID")
				fmt.Fprintf(w, "%d\n", u.User.ID)
				w.Flush()

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
	)
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAllCollections(w *tabwriter.Writer) error {
	url := configContext.APIRoot + "/collections"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("X-Api-Key", configContext.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var c resources.CollectionList
	json.Unmarshal(body, &c)

	fmt.Fprintln(w, "ID\tNAME")
	for _, col := range c.Collections {
		fmt.Fprintf(w, "%s\t%s\n", col.ID, col.Name)
	}
	w.Flush()

	return nil
}

func getSingleCollection(w *tabwriter.Writer, id string) error {
	url := configContext.APIRoot + "/collections/" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("X-Api-Key", configContext.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var cRes resources.CollectionResponse
	json.Unmarshal(body, &cRes)

	fmt.Fprintln(w, "ID\tNAME")
	fmt.Fprintf(w, "%s\t%s\n", cRes.Collection.Info.ID, cRes.Collection.Info.Name)
	w.Flush()

	return nil
}
