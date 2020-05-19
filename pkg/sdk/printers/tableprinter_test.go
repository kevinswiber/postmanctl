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

package printers_test

import (
	"bytes"
	"testing"

	"github.com/kevinswiber/postmanctl/pkg/sdk/printers"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources/gen"
)

func TestTablePrinterPrintsCollectionListItems(t *testing.T) {
	options := printers.PrintOptions{
		NoHeaders: false,
	}
	printer := printers.NewTablePrinter(options)

	var collections resources.CollectionListItems = []resources.CollectionListItem{
		{
			ID:    "abcdef",
			Name:  "test collection",
			Owner: "12345",
			UID:   "12345-abcdef",
		},
	}

	var formatter resources.Formatter = collections

	var b bytes.Buffer

	printer.PrintResource(formatter, &b)

	expected := `UID            NAME
12345-abcdef   test collection
`

	actual := b.String()
	if expected != actual {
		t.Errorf("Unexpected output, have: \"%s\", want: \"%s\"", actual, expected)
	}
}

func TestTablePrinterPrintsCollection(t *testing.T) {
	options := printers.PrintOptions{
		NoHeaders: false,
	}
	printer := printers.NewTablePrinter(options)

	collection := resources.Collection{
		Collection: &gen.Collection{
			Info: &gen.Info{
				PostmanID: "abcdef",
				Name:      "test collection",
			},
		},
	}

	var formatter resources.Formatter = collection

	var b bytes.Buffer

	printer.PrintResource(formatter, &b)

	expected := `ID       NAME
abcdef   test collection
`

	actual := b.String()
	if expected != actual {
		t.Errorf("Unexpected output, have: \"%s\", want: \"%s\"", actual, expected)
	}
}
