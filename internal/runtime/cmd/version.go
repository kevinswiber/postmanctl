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
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	commit  = "snapshot"
	date    = "<unknown>"
)

type versionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

var versionOutputFormat VersionOutputFormatValue

// VersionOutputFormatValue is a custom Value for the output flag that validates.
type VersionOutputFormatValue struct {
	value string
}

// String returns a string representation of this flag.
func (o *VersionOutputFormatValue) String() string {
	return o.value
}

// Set creates the flag value.
func (o *VersionOutputFormatValue) Set(v string) error {
	if v == "json" || v == "short" {
		o.value = v
		return nil
	}

	return errors.New("output format must be json or short")
}

// Type returns the type of this value.
func (o *VersionOutputFormatValue) Type() string {
	return "string"
}

func init() {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information for postmanctl.",
		Run: func(cmd *cobra.Command, args []string) {
			v := versionInfo{
				Version: version,
				Commit:  commit,
				Date:    date,
			}

			f := versionOutputFormat.value
			if f == "short" {
				fmt.Printf("Version: %s\n", v.Version)
				fmt.Printf("Commit: %s\n", v.Commit)
				fmt.Printf("Date: %s\n", v.Date)
			} else if f == "json" {
				p, err := json.MarshalIndent(&v, "", "  ")

				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %s\n", err)
					os.Exit(1)
					return
				}

				fmt.Println(string(p))
			} else {
				fmt.Printf("%#v\n", v)
			}
		},
	}

	cmd.Flags().VarP(&versionOutputFormat, "output", "o", "output format (json, short)")
	rootCmd.AddCommand(cmd)
}
