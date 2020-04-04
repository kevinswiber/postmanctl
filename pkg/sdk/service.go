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

package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

// Service is used by Postman API consumers.
type Service struct {
	Options *client.Options
}

// NewService returns a new instance of the Postman API service client.
func NewService(options *client.Options) *Service {
	return &Service{
		Options: options,
	}
}

// Collections returns all collections.
func (s *Service) Collections(ctx context.Context) (*resources.CollectionListItems, error) {
	var resource resources.CollectionListResponse
	if _, err := s.get(ctx, &resource, nil, "collections"); err != nil {
		return nil, err
	}

	return &resource.Collections, nil
}

// Collection returns a single collection.
func (s *Service) Collection(ctx context.Context, id string) (*resources.Collection, error) {
	var resource resources.CollectionResponse
	if _, err := s.get(ctx, &resource, nil, "collections", id); err != nil {
		return nil, err
	}

	return &resource.Collection, nil
}

// CreateCollectionFromReader creates a new collection.
func (s *Service) CreateCollectionFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	return s.CreateFromReader(ctx, resources.CollectionType, reader, params, nil)
}

// CreateEnvironmentFromReader creates a new environment.
func (s *Service) CreateEnvironmentFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	return s.CreateFromReader(ctx, resources.EnvironmentType, reader, params, nil)
}

// CreateMockFromReader creates a new mock.
func (s *Service) CreateMockFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	return s.CreateFromReader(ctx, resources.MockType, reader, params, nil)
}

// CreateMonitorFromReader creates a new monitor.
func (s *Service) CreateMonitorFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	return s.CreateFromReader(ctx, resources.MonitorType, reader, params, nil)
}

// CreateAPIFromReader creates a new API.
func (s *Service) CreateAPIFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	return s.CreateFromReader(ctx, resources.APIType, reader, params, nil)
}

// CreateAPIVersionFromReader creates a new API Version.
func (s *Service) CreateAPIVersionFromReader(ctx context.Context, reader io.Reader, workspace, apiID string) (string, error) {
	var queryParams map[string]string
	if workspace != "" {
		queryParams = make(map[string]string)
		queryParams["workspace"] = workspace
	}

	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new API version")
	}

	urlParams := make(map[string]string)
	urlParams["apiID"] = apiID

	return s.CreateFromReader(ctx, resources.APIVersionType, reader, queryParams, urlParams)
}

// CreateSchemaFromReader creates a new API Version.
func (s *Service) CreateSchemaFromReader(ctx context.Context, reader io.Reader, workspace, apiID, apiVersionID string) (string, error) {
	var queryParams map[string]string
	if workspace != "" {
		queryParams = make(map[string]string)
		queryParams["workspace"] = workspace
	}

	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new schema")
	}

	if apiVersionID == "" {
		return "", errors.New("an API Version ID is required for creating a new schema")
	}

	urlParams := make(map[string]string)
	urlParams["apiID"] = apiID
	urlParams["apiVersionID"] = apiVersionID

	return s.CreateFromReader(ctx, resources.APIVersionType, reader, queryParams, urlParams)
}

// CreateFromReader posts a new resource to the Postman API.
func (s *Service) CreateFromReader(ctx context.Context, t resources.ResourceType, reader io.Reader, queryParams, urlParams map[string]string) (string, error) {
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	var v map[string]interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return "", err
	}

	var (
		path             []string
		requestBody      []byte
		responseValueKey string
	)

	switch t {
	case resources.CollectionType:
		path = []string{"collections"}
		responseValueKey = "collection"

		c := struct {
			Collection map[string]interface{} `json:"collection"`
		}{
			Collection: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.EnvironmentType:
		path = []string{"environments"}
		responseValueKey = "environment"

		c := struct {
			Environment map[string]interface{} `json:"environment"`
		}{
			Environment: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.MockType:
		path = []string{"mocks"}
		responseValueKey = "mock"

		c := struct {
			Mock map[string]interface{} `json:"mock"`
		}{
			Mock: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.MonitorType:
		path = []string{"monitors"}
		responseValueKey = "monitor"

		c := struct {
			Monitor map[string]interface{} `json:"monitor"`
		}{
			Monitor: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.APIType:
		path = []string{"apis"}
		responseValueKey = "api"

		c := struct {
			API map[string]interface{} `json:"api"`
		}{
			API: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.APIVersionType:
		path = []string{"apis", urlParams["apiID"], "versions"}
		responseValueKey = "version"

		c := struct {
			Version map[string]interface{} `json:"version"`
		}{
			Version: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.SchemaType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["apiVersionID"], "schemas"}
		responseValueKey = "version"

		c := struct {
			Schema map[string]interface{} `json:"schema"`
		}{
			Schema: v,
		}
		requestBody, err = json.Marshal(c)
	default:
		return "", fmt.Errorf("unable to create resource, %+v not supported", t)
	}

	if err != nil {
		return "", err
	}

	var responseBody interface{}
	if _, err := s.post(ctx, requestBody, &responseBody, queryParams, path...); err != nil {
		return "", err
	}

	// Try a best attempt at returning the ID value.
	var responseValue map[string]interface{}
	responseValue = responseBody.(map[string]interface{})
	if v, ok := responseValue[responseValueKey]; ok {
		var vMap map[string]interface{}
		vMap = v.(map[string]interface{})
		if v2, ok := vMap["uid"]; ok {
			return v2.(string), nil
		}
		if v2, ok := vMap["id"]; ok {
			return v2.(string), nil
		}
		return "", nil
	}
	return "", nil
}

// ReplaceCollectionFromReader creates a new collection.
func (s *Service) ReplaceCollectionFromReader(ctx context.Context, reader io.Reader, workspace, resourceID string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.CollectionType, reader, params, urlParams)
}

// ReplaceEnvironmentFromReader creates a new environment.
func (s *Service) ReplaceEnvironmentFromReader(ctx context.Context, reader io.Reader, workspace, resourceID string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.EnvironmentType, reader, params, urlParams)
}

// ReplaceMockFromReader creates a new mock.
func (s *Service) ReplaceMockFromReader(ctx context.Context, reader io.Reader, workspace, resourceID string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.MockType, reader, params, urlParams)
}

// ReplaceMonitorFromReader creates a new monitor.
func (s *Service) ReplaceMonitorFromReader(ctx context.Context, reader io.Reader, workspace, resourceID string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.MonitorType, reader, params, urlParams)
}

// ReplaceAPIFromReader creates a new API.
func (s *Service) ReplaceAPIFromReader(ctx context.Context, reader io.Reader, workspace, resourceID string) (string, error) {
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.APIType, reader, params, urlParams)
}

// ReplaceAPIVersionFromReader creates a new API Version.
func (s *Service) ReplaceAPIVersionFromReader(ctx context.Context, reader io.Reader, workspace, resourceID, apiID string) (string, error) {
	var queryParams map[string]string
	if workspace != "" {
		queryParams = make(map[string]string)
		queryParams["workspace"] = workspace
	}

	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new API version")
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID
	urlParams["apiID"] = apiID

	return s.ReplaceFromReader(ctx, resources.APIVersionType, reader, queryParams, urlParams)
}

// ReplaceSchemaFromReader creates a new API Version.
func (s *Service) ReplaceSchemaFromReader(ctx context.Context, reader io.Reader, workspace, resourceID, apiID, apiVersionID string) (string, error) {
	var queryParams map[string]string
	if workspace != "" {
		queryParams = make(map[string]string)
		queryParams["workspace"] = workspace
	}

	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new schema")
	}

	if apiVersionID == "" {
		return "", errors.New("an API Version ID is required for creating a new schema")
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID
	urlParams["apiID"] = apiID
	urlParams["apiVersionID"] = apiVersionID

	return s.ReplaceFromReader(ctx, resources.APIVersionType, reader, queryParams, urlParams)
}

// ReplaceFromReader posts a new resource to the Postman API.
func (s *Service) ReplaceFromReader(ctx context.Context, t resources.ResourceType, reader io.Reader, queryParams, urlParams map[string]string) (string, error) {
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	var v map[string]interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return "", err
	}

	var (
		path             []string
		requestBody      []byte
		responseValueKey string
	)

	switch t {
	case resources.CollectionType:
		path = []string{"collections", urlParams["ID"]}
		responseValueKey = "collection"

		c := struct {
			Collection map[string]interface{} `json:"collection"`
		}{
			Collection: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.EnvironmentType:
		path = []string{"environments", urlParams["ID"]}
		responseValueKey = "environment"

		c := struct {
			Environment map[string]interface{} `json:"environment"`
		}{
			Environment: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.MockType:
		path = []string{"mocks", urlParams["ID"]}
		responseValueKey = "mock"

		c := struct {
			Mock map[string]interface{} `json:"mock"`
		}{
			Mock: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.MonitorType:
		path = []string{"monitors", urlParams["ID"]}
		responseValueKey = "monitor"

		c := struct {
			Monitor map[string]interface{} `json:"monitor"`
		}{
			Monitor: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.APIType:
		path = []string{"apis", urlParams["ID"]}
		responseValueKey = "api"

		c := struct {
			API map[string]interface{} `json:"api"`
		}{
			API: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.APIVersionType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["ID"]}
		responseValueKey = "version"

		c := struct {
			Version map[string]interface{} `json:"version"`
		}{
			Version: v,
		}
		requestBody, err = json.Marshal(c)
	case resources.SchemaType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["apiVersionID"], "schemas", urlParams["ID"]}
		responseValueKey = "version"

		c := struct {
			Schema map[string]interface{} `json:"schema"`
		}{
			Schema: v,
		}
		requestBody, err = json.Marshal(c)
	default:
		return "", fmt.Errorf("unable to create resource, %+v not supported", t)
	}

	if err != nil {
		return "", err
	}

	var responseBody interface{}
	if _, err := s.put(ctx, requestBody, &responseBody, queryParams, path...); err != nil {
		return "", err
	}

	// Try a best attempt at returning the ID value.
	var responseValue map[string]interface{}
	responseValue = responseBody.(map[string]interface{})
	if v, ok := responseValue[responseValueKey]; ok {
		var vMap map[string]interface{}
		vMap = v.(map[string]interface{})
		if v2, ok := vMap["uid"]; ok {
			return v2.(string), nil
		}
		if v2, ok := vMap["id"]; ok {
			return v2.(string), nil
		}
		return "", nil
	}
	return "", nil
}

// Environments returns all environments.
func (s *Service) Environments(ctx context.Context) (*resources.EnvironmentListItems, error) {
	var resource resources.EnvironmentListResponse
	if _, err := s.get(ctx, &resource, nil, "environments"); err != nil {
		return nil, err
	}

	return &resource.Environments, nil
}

// Environment returns a single environment.
func (s *Service) Environment(ctx context.Context, id string) (*resources.Environment, error) {
	var resource resources.EnvironmentResponse
	if _, err := s.get(ctx, &resource, nil, "environments", id); err != nil {
		return nil, err
	}

	return &resource.Environment, nil
}

// APIs returns all APIs.
func (s *Service) APIs(ctx context.Context, workspace string) (*resources.APIListItems, error) {
	var resource resources.APIListResponse
	var params map[string]string
	if workspace != "" {
		params = make(map[string]string)
		params["workspace"] = workspace
	}

	if _, err := s.get(ctx, &resource, params, "apis"); err != nil {
		return nil, err
	}

	return &resource.APIs, nil
}

// API returns a single API.
func (s *Service) API(ctx context.Context, id string) (*resources.API, error) {
	var resource resources.APIResponse
	if _, err := s.get(ctx, &resource, nil, "apis", id); err != nil {
		return nil, err
	}

	return &resource.API, nil
}

// APIVersions returns all API Versions.
func (s *Service) APIVersions(ctx context.Context, apiID string) (*resources.APIVersionListItems, error) {
	var resource resources.APIVersionListResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions"); err != nil {
		return nil, err
	}

	return &resource.APIVersions, nil
}

// APIVersion returns a single API Version.
func (s *Service) APIVersion(ctx context.Context, apiID, id string) (*resources.APIVersion, error) {
	var resource resources.APIVersionResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", id); err != nil {
		return nil, err
	}

	return &resource.APIVersion, nil
}

// Schema returns a single schema for an API version.
func (s *Service) Schema(ctx context.Context, apiID, apiVersionID, id string) (*resources.Schema, error) {
	var resource resources.SchemaResponse
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", apiVersionID, "schemas", id); err != nil {
		return nil, err
	}

	return &resource.Schema, nil
}

// APIRelations returns the linked relations of an API
func (s *Service) APIRelations(ctx context.Context, apiID, apiVersionID string) (*resources.APIRelations, error) {
	var resource resources.APIRelationsResource
	if _, err := s.get(ctx, &resource, nil, "apis", apiID, "versions", apiVersionID, "relations"); err != nil {
		return nil, err
	}

	return &resource.Relations, nil
}

// FormattedAPIRelationItems returns the formatted linked relations of an API
func (s *Service) FormattedAPIRelationItems(ctx context.Context, apiID, apiVersionID string) (*resources.FormattedAPIRelationItems, error) {
	r, err := s.APIRelations(ctx, apiID, apiVersionID)
	if err != nil {
		return nil, err
	}

	f := resources.NewFormattedAPIRelationItems(r)
	return &f, nil
}

// User returns the current user.
func (s *Service) User(ctx context.Context) (*resources.User, error) {
	var resource resources.UserResponse
	if _, err := s.get(ctx, &resource, nil, "me"); err != nil {
		return nil, err
	}

	return &resource.User, nil
}

// Workspaces returns the workspaces for the current user.
func (s *Service) Workspaces(ctx context.Context) (*resources.WorkspaceListItems, error) {
	var resource resources.WorkspaceListResponse
	if _, err := s.get(ctx, &resource, nil, "workspaces"); err != nil {
		return nil, err
	}

	return &resource.Workspaces, nil
}

// Workspace returns a single workspace for the current user.
func (s *Service) Workspace(ctx context.Context, id string) (*resources.Workspace, error) {
	var resource resources.WorkspaceResponse
	if _, err := s.get(ctx, &resource, nil, "workspaces", id); err != nil {
		return nil, err
	}

	return &resource.Workspace, nil
}

// Monitors returns the monitors for the current user.
func (s *Service) Monitors(ctx context.Context) (*resources.MonitorListItems, error) {
	var resource resources.MonitorListResponse
	if _, err := s.get(ctx, &resource, nil, "monitors"); err != nil {
		return nil, err
	}

	return &resource.Monitors, nil
}

// Monitor returns a single monitor for the current user.
func (s *Service) Monitor(ctx context.Context, id string) (*resources.Monitor, error) {
	var resource resources.MonitorResponse
	if _, err := s.get(ctx, &resource, nil, "monitors", id); err != nil {
		return nil, err
	}

	return &resource.Monitor, nil
}

// Mocks returns the mocks for the current user.
func (s *Service) Mocks(ctx context.Context) (*resources.MockListItems, error) {
	var resource resources.MockListResponse
	if _, err := s.get(ctx, &resource, nil, "mocks"); err != nil {
		return nil, err
	}

	return &resource.Mocks, nil
}

// Mock returns a single mock for the current user.
func (s *Service) Mock(ctx context.Context, id string) (*resources.Mock, error) {
	var resource resources.MockResponse
	if _, err := s.get(ctx, &resource, nil, "mocks", id); err != nil {
		return nil, err
	}

	return &resource.Mock, nil
}

func (s *Service) get(ctx context.Context, r interface{}, queryParams map[string]string, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Get().
		Path(path...).
		Params(queryParams).
		Into(&r).
		Do()

	return res, err
}

func (s *Service) post(ctx context.Context, input []byte, output interface{}, queryParams map[string]string, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Post().
		Path(path...).
		Params(queryParams).
		AddHeader("Content-Type", "application/json").
		Body(bytes.NewReader(input)).
		Into(&output).
		Do()

	return res, err
}

func (s *Service) put(ctx context.Context, input []byte, output interface{}, queryParams map[string]string, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Put().
		Path(path...).
		Params(queryParams).
		AddHeader("Content-Type", "application/json").
		Body(bytes.NewReader(input)).
		Into(&output).
		Do()

	return res, err
}
