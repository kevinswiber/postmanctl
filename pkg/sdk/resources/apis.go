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

// GetResourceKind returns a string representation of the resource type.
func (ali APIListItem) GetResourceKind() ResourceKind {
	return "APIListItem"
}

// GetPrintColumns returns a list of fields to print for this resource output.
func (ali APIListItem) GetPrintColumns() []string {
	return []string{"ID", "Name", "Summary"}
}

// APIResponse is a single API representation in the Postman API.
type APIResponse struct {
	API API `json:"api"`
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

// GetResourceKind returns a string representation of the resource type.
func (ali API) GetResourceKind() ResourceKind {
	return "API"
}

// GetPrintColumns returns a list of fields to print for this resource output.
func (ali API) GetPrintColumns() []string {
	return []string{"ID", "Name", "Summary"}
}
