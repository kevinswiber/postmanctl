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

package printers

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
	"github.com/liggitt/tabwriter"
)

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = tabwriter.RememberWidths
)

// TablePrinter prints human-readable table structures.
type TablePrinter struct {
	options PrintOptions
}

// NewTablePrinter creates a new human-readable table printer.
func NewTablePrinter(o PrintOptions) *TablePrinter {
	return &TablePrinter{
		options: o,
	}
}

// PrintResource executes the printer and creates an output.
func (p *TablePrinter) PrintResource(r resources.Formatter, output io.Writer) {
	w, found := output.(*tabwriter.Writer)
	if !found {
		w = GetNewTabWriter(output)
	}

	defer w.Flush()

	cols, objs := r.Format()

	if !p.options.NoHeaders {
		printCols := make([]string, len(cols))
		for i, c := range cols {
			if c == "PostmanID" {
				printCols[i] = "ID"
			} else {
				printCols[i] = strings.ToUpper(c)
			}
		}
		fmt.Fprintln(w, strings.Join(printCols, "\t"))
	}

	for _, obj := range objs {
		vals := make([]string, len(cols))
		for i, c := range cols {
			rVal := reflect.ValueOf(obj)
			f := reflect.Indirect(rVal).FieldByName(c).String()
			vals[i] = f
		}

		fmt.Fprintln(w, strings.Join(vals, "\t"))
	}
}

// GetNewTabWriter returns a new formatted tabwriter.Writer.
func GetNewTabWriter(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, tabwriterMinWidth, tabwriterWidth,
		tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}
