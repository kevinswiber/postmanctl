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
	"time"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources/gen"
)

//go:generate sh -c "schema-generate -p gen ../../../schema/collection.schema.json  | sed 's/Id/ID/g' > ./gen/collection.go"

// Collection represents a Postman Collection.
type Collection struct {
	*gen.Collection
	Items *ItemTree
}

// UnmarshalJSON converts JSON to a struct.
func (c *Collection) UnmarshalJSON(b []byte) error {
	var genC gen.Collection
	if err := json.Unmarshal(b, &genC); err != nil {
		return err
	}

	c.Collection = &genC
	node := ItemTreeNode{}
	if err := populateItemGroup(&node, c.Collection.Item); err != nil {
		return err
	}

	c.Items = &ItemTree{
		Root: node,
	}

	return nil
}

// CollectionListResponse is the top-level struct representation of a collection
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

	return []string{"UID", "Name"}, s
}

// CollectionListItem represents a single item in a CollectionListResponse.
type CollectionListItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	UID   string `json:"uid"`
	Fork  *Fork  `json:"fork,omitempty"`
}

// Fork represents fork metadata for a collection.
type Fork struct {
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"createdAt"`
	From      string    `json:"from"`
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

// Item represents an item (request) in a Collection.
type Item struct {
	*gen.Item
	Events []Event
}

// ItemGroup represents a folder in a Collection.
type ItemGroup struct {
	*gen.ItemGroup
	Events []Event
}

// Event represents an item event.
type Event struct {
	*gen.Event
}

// UnmarshalJSON converts JSON to a struct.
func (item *Item) UnmarshalJSON(b []byte) error {
	var genItem gen.Item
	if err := json.Unmarshal(b, &genItem); err != nil {
		return err
	}

	item.Item = &genItem
	item.Events = make([]Event, len(item.Item.Event))

	for i, genEvent := range item.Item.Event {
		item.Events[i] = Event{Event: genEvent}
	}

	return nil
}

func populateItemGroup(b *ItemTreeNode, item []interface{}) error {
	if len(item) == 0 {
		return nil
	}

	for _, v := range item {
		var m map[string]interface{} = v.(map[string]interface{})

		if v2, ok := m["item"]; ok {
			branch := ItemTreeNode{}

			if _, ok := m["name"]; ok {
				data, err := json.Marshal(m)
				if err != nil {
					return err
				}

				var ig gen.ItemGroup
				if err := json.Unmarshal(data, &ig); err != nil {
					return err
				}

				branch.MakeGroup(ItemGroup{
					ItemGroup: &ig,
				})
			}

			if err := populateItemGroup(&branch, v2.([]interface{})); err != nil {
				return err
			}

			b.AddBranch(branch)
		} else {
			it, err := populateItem(m["name"].(string), m)

			if err != nil {
				return err
			}

			b.AddItem(*it)
		}
	}

	return nil
}

func populateItem(name string, item map[string]interface{}) (*Item, error) {
	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	var it Item
	if err := json.Unmarshal(data, &it); err != nil {
		return nil, err
	}

	return &it, nil
}
