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

// APIVersionListResponse represents the top-level API Versions response from the Postman API.
type APIVersionListResponse struct {
	APIVersions APIVersionListItems `json:"versions"`
}

// APIVersionListItems is a slice of APIVersionListItem
type APIVersionListItems []APIVersionListItem

// Format returns column headers and values for the resource.
func (r APIVersionListItems) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// APIVersionListItem represents a single item in an APIVersionListResponse.
type APIVersionListItem struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	TransactionID string    `json:"transactionId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	API           string    `json:"api"`
	CreatedBy     string    `json:"createdBy"`
	UpdatedBy     string    `json:"updatedBy"`
	LastRevision  int64     `json:"lastRevision"`
}

// APIVersionResponse is a single APIVersion representation in the Postman APIVersion.
type APIVersionResponse struct {
	APIVersion APIVersion `json:"version"`
}

// APIVersion represents a single item in an APIVersionListResponse.
type APIVersion struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	TransactionID string    `json:"transactionId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	API           string    `json:"api"`
	CreatedBy     string    `json:"createdBy"`
	UpdatedBy     string    `json:"updatedBy"`
	LastRevision  int64     `json:"lastRevision"`
	Schema        []string  `json:"schema"`
}

// Format returns column headers and values for the resource.
func (r APIVersion) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r

	return []string{"ID", "Name"}, s
}

// APIVersionSlice is a slice of APIVersion.
type APIVersionSlice []*APIVersion

// Format returns column headers and values for the resource.
func (r APIVersionSlice) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}
