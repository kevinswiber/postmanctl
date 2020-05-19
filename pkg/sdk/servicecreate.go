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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

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

// CreateWorkspaceFromReader creates a new API.
func (s *Service) CreateWorkspaceFromReader(ctx context.Context, reader io.Reader, workspace string) (string, error) {
	return s.CreateFromReader(ctx, resources.WorkspaceType, reader, nil, nil)
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

	return s.CreateFromReader(ctx, resources.SchemaType, reader, queryParams, urlParams)
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
		requestBody, _ = json.Marshal(c) // already been unmarshalled, no error
	case resources.EnvironmentType:
		path = []string{"environments"}
		responseValueKey = "environment"

		c := struct {
			Environment map[string]interface{} `json:"environment"`
		}{
			Environment: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.MockType:
		path = []string{"mocks"}
		responseValueKey = "mock"

		c := struct {
			Mock map[string]interface{} `json:"mock"`
		}{
			Mock: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.MonitorType:
		path = []string{"monitors"}
		responseValueKey = "monitor"

		c := struct {
			Monitor map[string]interface{} `json:"monitor"`
		}{
			Monitor: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.WorkspaceType:
		path = []string{"workspaces"}
		responseValueKey = "workspace"

		c := struct {
			Workspace map[string]interface{} `json:"workspace"`
		}{
			Workspace: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.APIType:
		path = []string{"apis"}
		responseValueKey = "api"

		c := struct {
			API map[string]interface{} `json:"api"`
		}{
			API: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.APIVersionType:
		path = []string{"apis", urlParams["apiID"], "versions"}
		responseValueKey = "version"

		c := struct {
			Version map[string]interface{} `json:"version"`
		}{
			Version: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.SchemaType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["apiVersionID"], "schemas"}
		responseValueKey = "schema"

		c := struct {
			Schema map[string]interface{} `json:"schema"`
		}{
			Schema: v,
		}
		requestBody, _ = json.Marshal(c)
	default:
		return "", fmt.Errorf("unable to create resource, %+v not supported", t)
	}

	var responseBody interface{}
	if _, err := s.post(ctx, requestBody, &responseBody, queryParams, path...); err != nil {
		return "", err
	}

	// Try a best attempt at returning the ID value.
	responseValue := responseBody.(map[string]interface{})
	if v, ok := responseValue[responseValueKey]; ok {
		vMap := v.(map[string]interface{})
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
