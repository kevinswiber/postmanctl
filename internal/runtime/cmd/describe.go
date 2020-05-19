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
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
)

func init() {
	describeCmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe an entity in the Postman API",
	}

	userCmd := &cobra.Command{
		Use: "user",
		RunE: func(cmd *cobra.Command, args []string) error {
			resource, err := service.User(context.Background())

			if err != nil {
				return handleResponseError(err)
			}

			out, err := describeUser(resource)
			if err != nil {
				return err
			}
			fmt.Println(out)

			return nil
		},
	}

	apiVersionsCmd := &cobra.Command{
		Use:     "api-versions",
		Aliases: []string{"api-version"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return fetchAPIVersions(args)
		},
	}

	apiVersionsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiVersionsCmd.MarkFlagRequired("for-api")

	apiRelationsCmd := &cobra.Command{
		Use: "api-relations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fetchAPIRelations(args)
		},
	}

	apiRelationsCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	apiRelationsCmd.MarkFlagRequired("for-api")

	apiRelationsCmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
	apiRelationsCmd.MarkFlagRequired("for-api-version")

	schemaCmd := &cobra.Command{
		Use: "schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fetchSchema(args)
		},
	}

	schemaCmd.Flags().StringVar(&forAPI, "for-api", "", "the associated API ID (required)")
	schemaCmd.MarkFlagRequired("for-api")

	schemaCmd.Flags().StringVar(&forAPIVersion, "for-api-version", "", "the associated API Version ID (required)")
	schemaCmd.MarkFlagRequired("for-api-version")

	describeCmd.AddCommand(
		generateDescribeSubcommand("collections", []string{"collection", "co"}, fetchCollections),
		generateDescribeSubcommand("environments", []string{"environment", "env"}, fetchEnvironments),
		generateDescribeSubcommand("monitors", []string{"monitor", "mon"}, fetchMonitors),
		generateDescribeSubcommand("mocks", []string{"mock"}, fetchMocks),
		generateDescribeSubcommand("workspaces", []string{"workspace", "ws"}, fetchWorkspaces),
		userCmd,
		generateDescribeSubcommand("apis", []string{"api"}, fetchAPIs),
		apiVersionsCmd,
		apiRelationsCmd,
		schemaCmd,
	)
	rootCmd.AddCommand(describeCmd)
}

func generateDescribeSubcommand(use string, aliases []string, fn func(args []string) error) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Aliases: aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fn(args)
		},
	}
}

func fetchCollections(args []string) error {
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

	return nil
}

func fetchEnvironments(args []string) error {
	r := make(resources.EnvironmentSlice, len(args))
	for i, id := range args {
		resource, err := service.Environment(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeEnvironments(r)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func fetchMocks(args []string) error {
	r := make(resources.MockSlice, len(args))
	for i, id := range args {
		resource, err := service.Mock(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeMocks(r)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchMonitors(args []string) error {
	r := make(resources.MonitorSlice, len(args))
	for i, id := range args {
		resource, err := service.Monitor(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeMonitors(r)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchWorkspaces(args []string) error {
	r := make(resources.WorkspaceSlice, len(args))
	for i, id := range args {
		resource, err := service.Workspace(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeWorkspaces(r)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchAPIs(args []string) error {
	r := make(resources.APISlice, len(args))
	for i, id := range args {
		resource, err := service.API(context.Background(), id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeAPIs(r)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchAPIVersions(args []string) error {
	ids := args[0:]

	r := make(resources.APIVersionSlice, len(ids))
	for i, id := range ids {
		resource, err := service.APIVersion(context.Background(), forAPI, id)

		if err != nil {
			return handleResponseError(err)
		}

		r[i] = resource
	}

	out, err := describeAPIVersions(r)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchAPIRelations(args []string) error {
	resource, err := service.APIRelations(context.Background(), forAPI, forAPIVersion)

	if err != nil {
		return handleResponseError(err)
	}

	out, err := describeAPIRelations(resource)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func fetchSchema(args []string) error {
	var id string
	if len(args) == 0 {
		version, err := service.APIVersion(context.Background(), forAPI, forAPIVersion)

		if err != nil {
			return handleResponseError(err)
		}

		if len(version.Schema) > 0 {
			id = version.Schema[0]
		}
	} else {
		id = args[0]
	}

	resource, err := service.Schema(context.Background(), forAPI, forAPIVersion, id)

	if err != nil {
		return handleResponseError(err)
	}

	out, err := describeSchema(resource)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}

func describeCollections(r resources.CollectionSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, c := range r {
			buf := new(bytes.Buffer)
			buf.WriteString("Info:\n")
			buf.WriteString(fmt.Sprintf("  ID:\t%s\n", c.Info.PostmanID))
			buf.WriteString(fmt.Sprintf("  Name:\t%s\n", c.Info.Name))
			buf.WriteString(fmt.Sprintf("  Schema:\t%s\n", c.Info.Schema))

			var hasPreRequest, hasTest bool
			for _, e := range c.Event {
				if hasPreRequest && hasTest {
					break
				}

				if !hasPreRequest && e.Listen == "prerequest" {
					hasPreRequest = true
				}

				if !hasTest && e.Listen == "test" {
					hasTest = true
				}
			}
			buf.WriteString("Scripts:\n")
			buf.WriteString(fmt.Sprintf("  PreRequest:\t%t\n", hasPreRequest))
			buf.WriteString(fmt.Sprintf("  Test:\t%t\n", hasTest))

			var variables []string = make([]string, len(c.Variable))
			for i, v := range c.Variable {
				variables[i] = v.Key
			}

			buf.WriteString(fmt.Sprintf("Variables:\t%s\n", strings.Join(variables, ", ")))

			buf.WriteString(fmt.Sprint("Items:\n"))
			tree := treeprint.New()
			writeCollectionItemOrItemGroup(out, c.Items.Root, tree)
			for _, s := range strings.Split(tree.String(), "\n") {
				buf.WriteString(fmt.Sprintf("  %s\n", s))
			}
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeEnvironments(r resources.EnvironmentSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, e := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("ID:\t%s\n", e.ID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", e.Name))
			buf.WriteString(fmt.Sprintf("Variables:\n"))
			for _, v := range e.Values {
				buf.WriteString(fmt.Sprintf("  %s\n", v.Key))
			}
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeMonitors(r resources.MonitorSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, m := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("UID:\t%s\n", m.UID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", m.Name))
			buf.WriteString(fmt.Sprintf("Owner ID:\t%s\n", m.Owner))
			buf.WriteString(fmt.Sprintf("Collection UID:\t%s\n", m.CollectionUID))
			buf.WriteString(fmt.Sprintf("Environment UID:\t%s\n", m.EnvironmentUID))
			buf.WriteString(fmt.Sprintf("Options:\n"))
			buf.WriteString(fmt.Sprintf("  Strict SSL:\t%t\n", m.Options.StrictSSL))
			buf.WriteString(fmt.Sprintf("  Follow Redirects:\t%t\n", m.Options.FollowRedirects))
			requestTimeout := "<default>"
			if m.Options.RequestTimeout != nil {
				requestTimeout = string(*m.Options.RequestTimeout)
			}
			buf.WriteString(fmt.Sprintf("  Request Timeout:\t%s\n", requestTimeout))
			buf.WriteString(fmt.Sprintf("  Request Delay:\t%d\n", m.Options.RequestDelay))
			buf.WriteString(fmt.Sprintf("Schedule:\n"))
			buf.WriteString(fmt.Sprintf("  Cron:\t%s\n", m.Schedule.Cron))
			buf.WriteString(fmt.Sprintf("  Time Zone:\t%s\n", m.Schedule.Timezone))
			buf.WriteString(fmt.Sprintf("  Next Run:\t%s\n", m.Schedule.NextRun))
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeMocks(r resources.MockSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, m := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("UID:\t%s\n", m.UID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", m.Name))
			buf.WriteString(fmt.Sprintf("Owner ID:\t%s\n", m.Owner))
			buf.WriteString(fmt.Sprintf("Collection UID:\t%s\n", m.Collection))
			buf.WriteString(fmt.Sprintf("Environment UID:\t%s\n", m.Environment))
			buf.WriteString(fmt.Sprintf("Mock URL:\t%s\n", m.MockURL))
			buf.WriteString(fmt.Sprintf("Config:\n"))
			buf.WriteString(fmt.Sprintf("  Match Body:\t%t\n", m.Config.MatchBody))
			buf.WriteString(fmt.Sprintf("  Match Query Params:\t%t\n", m.Config.MatchQueryParams))
			buf.WriteString(fmt.Sprintf("  Match Wildcards:\t%t\n", m.Config.MatchWildcards))
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeWorkspaces(r resources.WorkspaceSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, m := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("ID:\t%s\n", m.ID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", m.Name))
			buf.WriteString(fmt.Sprintf("Type:\t%s\n", m.Type))
			buf.WriteString(fmt.Sprintf("Collections:\n"))
			for _, c := range m.Collections {
				buf.WriteString(fmt.Sprintf("  %s (%s)\n", c.Name, c.ID))
			}
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeUser(r *resources.User) (string, error) {
	return tabbedString(func(out io.Writer) error {
		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("ID:\t%s\n", r.ID))
		buf.WriteTo(out)
		return nil
	})
}

func describeAPIs(r resources.APISlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, m := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("ID:\t%s\n", m.ID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", m.Name))
			buf.WriteString(fmt.Sprintf("Created By:\t%s\n", m.CreatedBy))
			buf.WriteString(fmt.Sprintf("Created At:\t%s\n", m.CreatedAt))
			buf.WriteString(fmt.Sprintf("Updated By:\t%s\n", m.UpdatedBy))
			buf.WriteString(fmt.Sprintf("Updated By:\t%s\n", m.UpdatedAt))
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeAPIVersions(r resources.APIVersionSlice) (string, error) {
	return tabbedString(func(out io.Writer) error {
		for _, m := range r {
			buf := new(bytes.Buffer)
			buf.WriteString(fmt.Sprintf("ID:\t%s\n", m.ID))
			buf.WriteString(fmt.Sprintf("Name:\t%s\n", m.Name))
			buf.WriteString(fmt.Sprintf("API ID:\t%s\n", m.API))
			buf.WriteString(fmt.Sprintf("Schema ID:\t%s\n", strings.Join(m.Schema, ", ")))
			buf.WriteString(fmt.Sprintf("Created By:\t%s\n", m.CreatedBy))
			buf.WriteString(fmt.Sprintf("Created At:\t%s\n", m.CreatedAt))
			buf.WriteString(fmt.Sprintf("Updated By:\t%s\n", m.UpdatedBy))
			buf.WriteString(fmt.Sprintf("Updated By:\t%s\n", m.UpdatedAt))
			buf.WriteTo(out)
		}
		return nil
	})
}

func describeAPIRelations(r *resources.APIRelations) (string, error) {
	return tabbedString(func(out io.Writer) error {
		buf := new(bytes.Buffer)

		if len(r.Mock) > 0 {
			buf.WriteString(fmt.Sprintf("Mock Servers:\n"))
			for _, s := range r.Mock {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\tURL\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\t---\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt, s.URL))
			}
		}

		if len(r.Documentation) > 0 {
			buf.WriteString(fmt.Sprintf("Documentation:\n"))
			for _, s := range r.Documentation {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		if len(r.Environment) > 0 {
			buf.WriteString(fmt.Sprintf("Environments:\n"))
			for _, s := range r.Environment {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		if len(r.TestSuite) > 0 {
			buf.WriteString(fmt.Sprintf("Test Suites:\n"))
			for _, s := range r.TestSuite {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		if len(r.IntegrationTest) > 0 {
			buf.WriteString(fmt.Sprintf("Integration Tests:\n"))
			for _, s := range r.IntegrationTest {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		if len(r.ContractTest) > 0 {
			buf.WriteString(fmt.Sprintf("Contract Tests:\n"))
			for _, s := range r.ContractTest {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		if len(r.Monitor) > 0 {
			buf.WriteString(fmt.Sprintf("Monitors:\n"))
			for _, s := range r.Monitor {
				buf.WriteString(fmt.Sprintf("  UID\tName\tCreated At\tUpdated At\n"))
				buf.WriteString(fmt.Sprintf("  ---\t----\t----------\t----------\n"))
				buf.WriteString(fmt.Sprintf("  %s\t%s\t%s\t%s\n", s.ID, s.Name, s.CreatedAt, s.UpdatedAt))
			}
		}

		buf.WriteTo(out)
		return nil
	})
}

func describeSchema(r *resources.Schema) (string, error) {
	return tabbedString(func(out io.Writer) error {
		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("ID:\t%s\n", r.ID))
		buf.WriteString(fmt.Sprintf("API Version:\t%s\n", r.APIVersion))
		buf.WriteString(fmt.Sprintf("Created By:\t%s\n", r.CreatedBy))
		buf.WriteString(fmt.Sprintf("Created At:\t%s\n", r.CreatedAt))
		buf.WriteString(fmt.Sprintf("Updated By:\t%s\n", r.UpdatedBy))
		buf.WriteString(fmt.Sprintf("Updated At:\t%s\n", r.UpdatedAt))
		buf.WriteString(fmt.Sprintf("Type:\t%s\n", r.Type))
		buf.WriteString(fmt.Sprintf("Language:\t%s\n", r.Language))
		buf.WriteString(fmt.Sprintf("\nSchema:\n\n%s\n", r.Schema))
		buf.WriteTo(out)
		return nil
	})
}

func writeCollectionItemOrItemGroup(out io.Writer, c resources.ItemTreeNode, printBranch treeprint.Tree) {
	if c.ItemGroup != nil {
		var evs []string
		if c.ItemGroup.Events != nil {
			for _, e := range c.ItemGroup.Events {
				evs = append(evs, e.Listen)
			}
		}

		metaString := ""
		if len(evs) > 0 {
			metaString += fmt.Sprintf(" (scripts: %s)", strings.Join(evs, ","))
		}
		printBranch = printBranch.AddBranch(c.ItemGroup.Name + metaString)
	}

	if c.Branches != nil {
		if c.Branches != nil {
			for _, br := range *c.Branches {
				writeCollectionItemOrItemGroup(out, br, printBranch)
			}
		}
	}

	if c.Items != nil {
		for _, it := range *c.Items {
			var evs []string
			if it.Events != nil {
				for _, e := range it.Events {
					evs = append(evs, e.Listen)
				}
			}

			metaString := ""
			if len(evs) > 0 {
				sort.Strings(evs)
				metaString += fmt.Sprintf(" (scripts: %s)", strings.Join(evs, ","))
			}
			printBranch.AddNode(it.Name + metaString)
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
