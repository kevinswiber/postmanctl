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

package resources

import "time"

// SchemaResponse represents the top-level schema response in the Postman API.
type SchemaResponse struct {
	Schema Schema `json:"schema"`
}

// Schema represents an API schema from the Postman API
type Schema struct {
	APIVersion string    `json:"apiVersion"`
	CreatedBy  string    `json:"createdBy"`
	UpdatedBy  string    `json:"updatedBy"`
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Language   string    `json:"language"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Schema     string    `json:"schema"`
}

// Format returns column headers and values for the resource.
func (r Schema) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r

	return []string{"ID", "Type", "Language"}, s
}
