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
	"os"

	"github.com/kevinswiber/postmanctl/internal/runtime/config"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

var setContextAPIRoot string

func init() {
	var cmd = &cobra.Command{
		Use:   "config",
		Short: "Configure access to the Postman API.",
	}

	var setContextCmd = &cobra.Command{
		Use:   "set-context",
		Short: "Create a context for accessing the Postman API.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Postman API Key: ")
			apiKey, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			fmt.Printf("\n")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			cfg = &config.Config{}
			if configFileFound {
				if err := viper.Unmarshal(cfg); err != nil {
					fmt.Fprintf(os.Stderr, "error: %s\n", err)
					os.Exit(1)
				}
			}

			if len(cfg.Contexts) == 0 {
				cfg.Contexts = make(map[string]config.Context)
			}

			newContext := config.Context{}
			if setContextAPIRoot != "" {
				newContext.APIRoot = setContextAPIRoot
			}

			newContext.APIKey = string(apiKey)

			cfg.Contexts[args[0]] = newContext
			cfg.CurrentContext = args[0]

			mergeConfig(cfg)
		},
	}

	setContextCmd.Flags().StringVar(&setContextAPIRoot, "api-root", "https://api.postman.com", "API root URL for accessing the Postman API.")

	useContextCmd := &cobra.Command{
		Use:   "use-context",
		Short: "Use an existing context for postmanctl commands.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cfg = &config.Config{}
			if err := viper.Unmarshal(cfg); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}

			cfg.CurrentContext = args[0]

			mergeConfig(cfg)
		},
	}

	currentContextCmd := &cobra.Command{
		Use:   "current-context",
		Short: "Get the currently set context for postmanctl commands.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cfg.CurrentContext)
		},
	}

	cmd.AddCommand(setContextCmd, useContextCmd, currentContextCmd)
	rootCmd.AddCommand(cmd)
}

func prepareConfig(cfg *config.Config, result *map[string]interface{}) {
	if err := mapstructure.Decode(cfg, &result); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	contexts := make(map[string]interface{})
	for k, v := range cfg.Contexts {
		var vi map[string]interface{}
		if err := mapstructure.Decode(v, &vi); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		contexts[k] = vi
	}

	(*result)["contexts"] = contexts
}

func mergeConfig(cfg *config.Config) {
	var result map[string]interface{}
	prepareConfig(cfg, &result)

	if err := viper.MergeConfigMap(result); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	viper.SetConfigType("yaml")

	if err := viper.WriteConfig(); err != nil {
		if err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				if err := viper.SafeWriteConfig(); err != nil {
					fmt.Fprintf(os.Stderr, "error: %s\n", err)
					os.Exit(1)
				}
			default:
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
		}
	}

	fmt.Println("config file written to $HOME/.postmanctl.yaml")
}
