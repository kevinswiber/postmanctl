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
	"time"
)

// FormattedAPIRelationItems is a slice of FormattedAPIRelationItem
type FormattedAPIRelationItems []FormattedAPIRelationItem

// FormattedAPIRelationItem returns a list item version of an APIRelation.
type FormattedAPIRelationItem struct {
	ID   string
	Name string
	Type string
}

// NewFormattedAPIRelationItems returns a new human-readable view of API Relations.
func NewFormattedAPIRelationItems(r *APIRelations) FormattedAPIRelationItems {
	totalLen := len(r.Documentation) + len(r.ContractTest) + len(r.TestSuite) +
		len(r.IntegrationTest) + len(r.Mock) + len(r.Monitor)

	i := 0
	ret := make([]FormattedAPIRelationItem, totalLen)

	for _, s := range r.Documentation {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "documentation",
		}
		i++
	}

	for _, s := range r.Environment {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "environment",
		}
		i++
	}

	for _, s := range r.ContractTest {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "contracttest",
		}
		i++
	}

	for _, s := range r.TestSuite {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "testsuite",
		}
		i++
	}

	for _, s := range r.IntegrationTest {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "integrationtest",
		}
		i++
	}

	for _, s := range r.Mock {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "mock",
		}
		i++
	}

	for _, s := range r.Monitor {
		ret[i] = FormattedAPIRelationItem{
			ID:   s.ID,
			Name: s.Name,
			Type: "monitor",
		}
		i++
	}

	return ret
}

// Format returns column headers and values for the resource.
func (r FormattedAPIRelationItems) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name", "Type"}, s
}

// APIRelationsResource provides the top-level wrapper for API Relations.
type APIRelationsResource struct {
	Relations APIRelations `json:"relations"`
}

// APIRelations provides the top-level relations representation for APIs.
type APIRelations struct {
	Documentation   map[string]LinkedCollection  `json:"documentation,omitempty"`
	Environment     map[string]LinkedEnvironment `json:"environment,omitempty"`
	ContractTest    map[string]LinkedCollection  `json:"contracttest,omitempty"`
	TestSuite       map[string]LinkedCollection  `json:"testsuite,omitempty"`
	IntegrationTest map[string]LinkedCollection  `json:"integrationtest,omitempty"`
	Mock            map[string]LinkedMock        `json:"mock,omitempty"`
	Monitor         map[string]LinkedMonitor     `json:"monitor,omitempty"`
}

// LinkedCollection describes a single linked collection representation.
type LinkedCollection struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// LinkedEnvironment describes a single linked collection representation.
type LinkedEnvironment struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// LinkedMock describes a single linked collection representation.
type LinkedMock struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	URL       string    `json:"url"`
}

// LinkedMonitor describes a single linked collection representation.
type LinkedMonitor struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
