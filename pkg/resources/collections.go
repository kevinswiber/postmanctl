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

type CollectionList struct {
	Collections []CollectionListItem `json:"collections"`
}

type CollectionListItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Uid   string `json:"uid"`
}

type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

type Collection struct {
	Info      CollectionInfo   `json:"info"`
	Items     []CollectionItem `json:"item"`
	Events    []Event          `json:"event,omitempty"`
	Variables []Variable       `json:"variable,omitempty"`
}

type CollectionInfo struct {
	Id     string `json:"_postman_id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type CollectionItem struct {
	Id                      string           `json:"_postman_id"`
	Name                    string           `json:"name"`
	Items                   []CollectionItem `json:"item`
	Events                  []Event          `json:"event"`
	ProtocolProfileBehavior struct{}         `json:"protocolProfileBehavior"`
	Request                 Request          `json:"request"`
	Responses               []Response       `json:"response"`
}

type Event struct {
	Listen string `json:"listen"` /* prerequest, test */
	Script Script `json:"script"`
}

type Variable struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Script struct {
	Id   string   `json:"id"`
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}

type Request struct{}
type Response struct{}
