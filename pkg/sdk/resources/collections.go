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

import (
	"encoding/json"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources/raw"
)

//go:generate sh -c "schema-generate -p raw ../../../collection.json  | sed 's/Id/ID/g' > ./raw/collection.go"

// Collection represents a Postman Collection.
type Collection struct {
	*raw.Collection
	Items *ItemTree
}

// UnmarshalJSON converts JSON to a struct.
func (c *Collection) UnmarshalJSON(b []byte) error {
	var rawC raw.Collection
	if err := json.Unmarshal(b, &rawC); err != nil {
		return err
	}

	c.Collection = &rawC
	c.Items = newItemTree(c.Collection.Item)

	return nil
}

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
func (c Collection) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = c.Info

	return []string{"PostmanID", "Name"}, s
}

// CollectionSlice is a slice of Collection.
type CollectionSlice []*Collection

// Format returns column headers and values for the resource.
func (r CollectionSlice) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v.Info
	}

	return []string{"PostmanID", "Name"}, s
}

// ItemTree represents a folder/request structure in a Collection
type ItemTree struct {
	Root *itemBranch
}

func (b *itemBranch) populate(item []interface{}) {
	if len(item) == 0 {
		return
	}

	for _, v := range item {
		var m map[string]interface{}
		m = v.(map[string]interface{})

		if v2, ok := m["item"]; ok {
			b.
				addBranch(m["name"].(string)).
				populate(v2.([]interface{}))
		} else {
			b.addNode(m["name"].(string), m)
		}
	}
}

func (b *itemBranch) addBranch(name string) *itemBranch {
	return &itemBranch{
		Name: name,
	}
}

func (b *itemBranch) addNode(name string, item map[string]interface{}) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	var it raw.Item
	if err := json.Unmarshal(data, &it); err != nil {
		return err
	}

	b.Branch = nil
	b.Item = &[]Item{Item{&it}}

	return nil
}

func newItemTree(item []interface{}) *ItemTree {
	tree := &ItemTree{Root: &itemBranch{}}
	tree.Root.populate(item)
	return tree
}

type itemBranch struct {
	Name   string
	Branch *[]itemBranch
	Item   *[]Item
}

// Item represents an item (request) in a Collection.
type Item struct {
	*raw.Item
}
