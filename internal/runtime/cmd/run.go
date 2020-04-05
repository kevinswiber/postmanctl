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
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	var cmd = &cobra.Command{
		Use:   "run",
		Short: "Execute runnable Postman resources.",
	}

	var runMonitorCmd = &cobra.Command{
		Use:     "monitor",
		Aliases: []string{"mon"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg, err := service.RunMonitor(context.Background(), args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			b, err := json.MarshalIndent(&msg, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			fmt.Println(string(b))
		},
	}

	cmd.AddCommand(runMonitorCmd)
	rootCmd.AddCommand(cmd)
}
