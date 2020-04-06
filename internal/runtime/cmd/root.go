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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kevinswiber/postmanctl/internal/runtime/config"
	"github.com/kevinswiber/postmanctl/pkg/sdk"
	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	cfg              *config.Config
	configContext    config.Context
	configContextKey string
	options          *client.Options
	service          *sdk.Service
	forAPI           string
	forAPIVersion    string
	inputFile        string
	inputReader      io.Reader
	usingWorkspace   string
	forkLabel        string
	mergeStrategy    string
	mergeCollection  string
)

var configContextFound = true
var configFileFound = true
var configContextSet = true

var rootCmd = &cobra.Command{
	Use:   "postmanctl",
	Short: "Controls the Postman API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !configFileFound || !configContextFound {
			processArgs := os.Args
			if len(processArgs) > 2 {
				command := processArgs[1]
				subcommand := processArgs[2]

				if command == "config" && subcommand != "current-context" && subcommand != "get-contexts" {
					return
				}
			}

			if !configContextSet {
				fmt.Fprintln(os.Stderr, "error: context is not set, run: postmanctl config use-context --help")
			} else if !configContextFound {
				fmt.Fprintf(os.Stderr, "error: context '%s' is not configured, run: postmanctl config set-context --help\n", configContextKey)
			} else if !configFileFound {
				fmt.Fprintln(os.Stderr, "error: config file not found at $HOME/.postmanctl.yaml, run: postmanctl config set-context --help")
			}

			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GenMarkdownTree generates Markdown documentation for postmanctl.
func GenMarkdownTree(path string) error {
	return doc.GenMarkdownTree(rootCmd, path)
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initAPIClientConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.postmanctl.yaml)")
	rootCmd.PersistentFlags().StringVar(&configContextKey, "context", "", "context to use, overrides the current context in the config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".postmanctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".postmanctl")
	}

	viper.SetEnvPrefix("POSTMANCTL_")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			configFileFound = false
		} else {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}

	cfg = &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	if !configFileFound {
		return
	}

	if configContextKey == "" {
		configContextKey = cfg.CurrentContext
	}

	// viper keys are case-insensitive
	if val, ok := cfg.Contexts[strings.ToLower(configContextKey)]; ok {
		configContext = val
		if len(configContext.APIRoot) == 0 {
			configContext.APIRoot = "https://api.postman.com"
		}
	} else {
		context := cfg.CurrentContext
		configContextFound = false

		if context == "" {
			configContextSet = false
			return
		}
	}
}

func initAPIClientConfig() {
	u, err := url.Parse(configContext.APIRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	options = client.NewOptions(u, configContext.APIKey, http.DefaultClient)
	service = sdk.NewService(options)
}
