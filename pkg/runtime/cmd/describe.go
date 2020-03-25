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
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an entity in the Postman API",
}

func init() {
	describeCmd.AddCommand(
		generateDescribeSubcommand(resources.CollectionType, "collections", []string{"collection", "co"}),
	)
	rootCmd.AddCommand(describeCmd)
}

func generateDescribeSubcommand(t resources.ResourceType, use string, aliases []string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Aliases: aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return describeResource(t, args...)
		},
	}
}

func describeResource(resourceType resources.ResourceType, args ...string) error {
	switch resourceType {
	case resources.CollectionType:
		r := make(resources.CollectionSlice, len(args))
		for i, id := range args {
			resource, err := service.Collection(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		out, err := describeCollections(r)
		if err != nil {
			return err
		}

		fmt.Println(out)
	case resources.EnvironmentType:
		r := make(resources.EnvironmentSlice, len(args))
		for i, id := range args {
			resource, err := service.Environment(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.MockType:
		r := make(resources.MockSlice, len(args))
		for i, id := range args {
			resource, err := service.Mock(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.MonitorType:
		r := make(resources.MonitorSlice, len(args))
		for i, id := range args {
			resource, err := service.Monitor(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.APIType:
		r := make(resources.APISlice, len(args))
		for i, id := range args {
			resource, err := service.API(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.APIVersionType:
		apiID := args[0]
		ids := args[1:]

		r := make(resources.APIVersionSlice, len(ids))
		for i, id := range ids {
			resource, err := service.APIVersion(context.Background(), apiID, id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.WorkspaceType:
		r := make(resources.WorkspaceSlice, len(args))
		for i, id := range args {
			resource, err := service.Workspace(context.Background(), id)

			if err != nil {
				return handleResponseError(err)
			}

			r[i] = resource
		}

		printGetOutput(r)
	case resources.UserType:
		resource, err := service.User(context.Background())

		if err != nil {
			return handleResponseError(err)
		}

		printGetOutput(resource)
	case resources.SchemaType:
		apiID := args[0]
		apiVersionID := args[1]
		id := args[2]

		resource, err := service.Schema(context.Background(), apiID, apiVersionID, id)

		if err != nil {
			return handleResponseError(err)
		}

		printGetOutput(resource)
	default:
		return fmt.Errorf("invalid resource type: %s", resourceType.String())
	}

	return nil
}

func describeCollections(r resources.CollectionSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, c := range r {
			out.Write([]byte("Info:\n"))
			out.Write([]byte(fmt.Sprintf("  ID:\t%s\n", c.Info.PostmanId)))
			out.Write([]byte(fmt.Sprintf("  Name:\t%s\n", c.Info.Name)))
			out.Write([]byte(fmt.Sprintf("  Schema:\t%s\n", c.Info.Schema)))
			out.Write([]byte(fmt.Sprint("Items:\n")))
			tree := treeprint.New()
			writeCollectionItemOrItemGroup(out, c.Item, tree)
			for _, s := range strings.Split(tree.String(), "\n") {
				out.Write([]byte(fmt.Sprintf("  %s\n", s)))
			}
		}
		return nil
	})
}

func writeCollectionItemOrItemGroup(out io.Writer, c []interface{}, branch treeprint.Tree) {
	if len(c) == 0 {
		return
	}

	for _, v := range c {
		var m map[string]interface{}
		m = v.(map[string]interface{})
		if _, ok := m["item"]; ok {
			b := branch.AddBranch(m["name"].(string) + " (folder)")
			writeCollectionItemOrItemGroup(out, m["item"].([]interface{}), b)
		} else {
			branch.AddNode(m["name"].(string) + " (request)")
		}
	}
}

func tabbedString(f func(io.Writer) error) (string, error) {
	out := new(tabwriter.Writer)
	buf := &bytes.Buffer{}
	out.Init(buf, 0, 8, 2, ' ', 0)

	err := f(out)
	if err != nil {
		return "", err
	}

	out.Flush()
	str := string(buf.String())
	return str, nil
}
