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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
	"github.com/kevinswiber/postmanctl/pkg/sdk/printers"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"k8s.io/client-go/util/jsonpath"
)

func handleResponseError(err error) error {
	if err, ok := err.(*client.RequestError); ok {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil
	}

	return err
}

func printGetOutput(r interface{}) {
	if r == nil {
		return
	}

	if outputFormat.value == "json" {
		t, err := json.MarshalIndent(&r, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		fmt.Println(string(t))
	} else if strings.HasPrefix(outputFormat.value, "jsonpath=") {
		tmpl := outputFormat.value[9:]
		j := jsonpath.New("out")
		if err := j.Parse(tmpl); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		t, err := json.Marshal(&r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		var queryObj interface{}
		queryObj = map[string]interface{}{}
		if err := json.Unmarshal(t, &queryObj); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		buf := bytes.NewBuffer(nil)
		j.Execute(buf, queryObj)

		fmt.Println(buf)
	} else {
		var f resources.Formatter
		f = r.(resources.Formatter)
		printTable(f)
	}
}

func printTable(f resources.Formatter) {
	w := printers.GetNewTabWriter(os.Stdout)
	printer := printers.NewTablePrinter(printers.PrintOptions{})
	printer.PrintResource(f, w)
}
