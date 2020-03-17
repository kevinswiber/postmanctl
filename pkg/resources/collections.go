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
	Collections []CollectionListItem `json:"collections"`
}

// CollectionListItem represents a single item in a CollectionListResponse.
type CollectionListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	UID   string `json:"uid"`
}

// GetResourceKind returns a string representation of the resource type.
func (c CollectionListItem) GetResourceKind() ResourceKind {
	return "CollectionListItem"
}

// GetPrintColumns returns a list of fields to print for this resource output.
func (c CollectionListItem) GetPrintColumns() []string {
	return []string{"ID", "Name"}
}

// CollectionResponse is the top-level struct representation of a collection
// response from the Postman API.
type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

// Collection is a single item representation of the CollectionResponse.
type Collection struct {
	Info      CollectionInfo   `json:"info"`
	Items     []CollectionItem `json:"item"`
	Events    []Event          `json:"event,omitempty"`
	Variables []Variable       `json:"variable,omitempty"`
}

// CollectionInfo contains metadata associated with a Collection.
type CollectionInfo struct {
	ID     string `json:"_postman_id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

// GetResourceKind returns a string representation of the resource type.
func (c CollectionInfo) GetResourceKind() ResourceKind {
	return "CollectionInfo"
}

// GetPrintColumns returns a list of fields to print for this resource output.
func (c CollectionInfo) GetPrintColumns() []string {
	return []string{"ID", "Name"}
}

// CollectionItem is a single unit of a collection entity.
type CollectionItem struct {
	ID                      string           `json:"_postman_id"`
	Name                    string           `json:"name"`
	Items                   []CollectionItem `json:"item"`
	Events                  []Event          `json:"event"`
	ProtocolProfileBehavior struct{}         `json:"protocolProfileBehavior"`
	Request                 Request          `json:"request"`
	Responses               []Response       `json:"response"`
}

// Event represents a pre-request or test script.
type Event struct {
	Listen string `json:"listen"` /* prerequest, test */
	Script Script `json:"script"`
}

// Variable is a representation of a Postman variable.
type Variable struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// Script is a representation of a pre-request or test script in the
// Postman API.
type Script struct {
	ID   string   `json:"id"`
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}

// Request contains HTTP request info associated with a CollectionItem.
type Request struct{}

// Response contains HTTP response info associated with a CollectionItem.
type Response struct{}
