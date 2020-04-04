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
		Use:   "merge",
		Short: "Merge a fork of a Postman resource.",
	}

	var mergeCollectionCmd = &cobra.Command{
		Use:     "collection",
		Aliases: []string{"co"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := service.MergeCollection(context.Background(), args[0], mergeCollection, mergeStrategy)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println(id)
		},
	}

	mergeCollectionCmd.Flags().StringVar(&mergeCollection, "to", "", "the destination collection to receive the merged changes")
	mergeCollectionCmd.MarkFlagRequired("to")

	mergeCollectionCmd.Flags().StringVarP(&mergeStrategy, "strategy", "s", "", "strategy for merging fork (optional, values: deleteSource, updateSourceWithDestination)")

	cmd.AddCommand(mergeCollectionCmd)
	rootCmd.AddCommand(cmd)
}
