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

// CollectionListResponse is the top-level struct represenation of a collection
// list response in the Postman API.
type CollectionListResponse struct {
	Collections CollectionListItems `json:"collections"`
}

// CollectionListItems is a slice of CollectionListItem
type CollectionListItems []CollectionListItem

// Format returns column headers and values for the resource.
func (r CollectionListItems) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// CollectionListItem represents a single item in a CollectionListResponse.
type CollectionListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	UID   string `json:"uid"`
}

// CollectionResponse is the top-level struct representation of a collection
// response from the Postman API.
type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

// Format returns column headers and values for the resource.
func (r Collection) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r.Info

	return []string{"ID", "Name"}, s
}

// CollectionSlice is a slice of Collection.
type CollectionSlice []*Collection

// Format returns column headers and values for the resource.
func (r CollectionSlice) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v.Info
	}

	return []string{"ID", "Name"}, s
}
