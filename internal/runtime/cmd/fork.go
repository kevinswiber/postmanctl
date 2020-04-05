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

	"github.com/spf13/cobra"
)

func init() {
	var cmd = &cobra.Command{
		Use:   "fork",
		Short: "Create a fork of a Postman resource.",
	}

	var forkCollectionCmd = &cobra.Command{
		Use:     "collection",
		Aliases: []string{"co"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := service.ForkCollection(context.Background(), args[0], usingWorkspace, forkLabel)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			fmt.Println(id)
		},
	}

	forkCollectionCmd.Flags().StringVarP(&usingWorkspace, "workspace", "w", "", "workspace for fork operation")
	forkCollectionCmd.Flags().StringVarP(&forkLabel, "label", "l", "", "label to associate with the forked collection (required)")
	forkCollectionCmd.MarkFlagRequired("label")

	cmd.AddCommand(forkCollectionCmd)
	rootCmd.AddCommand(cmd)
}
