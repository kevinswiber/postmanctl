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

// EnvironmentListResponse represents the top-level environments response from the
// Postman API.
type EnvironmentListResponse struct {
	Environments []EnvironmentListItem `json:"environments"`
}

// Format returns column headers and values for the resource.
func (r EnvironmentListResponse) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r.Environments))
	for i, v := range r.Environments {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// EnvironmentListItem represents a single item in an EnvironmentListResponse.
type EnvironmentListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	UID   string `json:"uid"`
}

// EnvironmentResponse is the top-level environment response from the
// Postman API.
type EnvironmentResponse struct {
	Environment Environment `json:"environment"`
}

// Format returns column headers and values for the resource.
func (r EnvironmentResponse) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r.Environment

	return []string{"ID", "Name"}, s
}

// Environment represents the single environment response from the
// Postman API
type Environment struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Values []KeyValuePair `json:"values"`
}

// KeyValuePair represents a key and value in the Postman API.
type KeyValuePair struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}
