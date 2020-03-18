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

// APIListResponse represents the top-level APIs response from the Postman API.
type APIListResponse struct {
	APIs []APIListItem `json:"apis"`
}

// Format returns column headers and values for the resource.
func (r APIListResponse) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r.APIs))
	for i, v := range r.APIs {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// APIListItem represents a single item in an APIListResponse.
type APIListItem struct {
	CreatedBy   string `json:"createdBy"`
	UpdatedBy   string `json:"updatedBy"`
	Team        string `json:"team"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// APIResponse is a single API representation in the Postman API.
type APIResponse struct {
	API API `json:"api"`
}

// Format returns column headers and values for the resource.
func (r APIResponse) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r.API

	return []string{"ID", "Name"}, s
}

// API represents a single item in an APIListResponse.
type API struct {
	CreatedBy   string `json:"createdBy"`
	UpdatedBy   string `json:"updatedBy"`
	Team        string `json:"team"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
