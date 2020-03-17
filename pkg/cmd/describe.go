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
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/kevinswiber/postmanctl/pkg/resources"
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("describe called")
	},
}

func init() {
	describeCmd.AddCommand(
		&cobra.Command{
			Use:     "collections",
			Aliases: []string{"collection", "co"},
			Args:    cobra.MinimumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) < 1 {
					return errors.New("Please specify a collection ID")
				}

				url := configContext.APIRoot + "/collections/" + args[0]
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

				w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
				fmt.Fprintln(w, "ID\tNAME")
				fmt.Fprintf(w, "%s\t%s\n", cRes.Collection.Info.ID, cRes.Collection.Info.Name)
				w.Flush()

				return nil
			},
		},
	)
	rootCmd.AddCommand(describeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// describeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// describeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
