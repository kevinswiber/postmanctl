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

// ReplaceCollectionFromReader replaces a collection.
func (s *Service) ReplaceCollectionFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.CollectionType, reader, urlParams)
}

// ReplaceEnvironmentFromReader replaces an existing environment.
func (s *Service) ReplaceEnvironmentFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.EnvironmentType, reader, urlParams)
}

// ReplaceMockFromReader replaces an existing mock.
func (s *Service) ReplaceMockFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.MockType, reader, urlParams)
}

// ReplaceMonitorFromReader replaces an existing monitor.
func (s *Service) ReplaceMonitorFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.MonitorType, reader, urlParams)
}

// ReplaceWorkspaceFromReader replaces an existing API.
func (s *Service) ReplaceWorkspaceFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.WorkspaceType, reader, urlParams)
}

// ReplaceAPIFromReader replaces an existing API.
func (s *Service) ReplaceAPIFromReader(ctx context.Context, reader io.Reader, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.ReplaceFromReader(ctx, resources.APIType, reader, urlParams)
}

// ReplaceAPIVersionFromReader replaces an existing API Version.
func (s *Service) ReplaceAPIVersionFromReader(ctx context.Context, reader io.Reader, resourceID, apiID string) (string, error) {
	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new API version")
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID
	urlParams["apiID"] = apiID

	return s.ReplaceFromReader(ctx, resources.APIVersionType, reader, urlParams)
}

// ReplaceSchemaFromReader replaces an existing API Version.
func (s *Service) ReplaceSchemaFromReader(ctx context.Context, reader io.Reader, resourceID, apiID, apiVersionID string) (string, error) {
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

	return s.ReplaceFromReader(ctx, resources.SchemaType, reader, urlParams)
}

// ReplaceFromReader posts a new resource to the Postman API.
func (s *Service) ReplaceFromReader(ctx context.Context, t resources.ResourceType, reader io.Reader, urlParams map[string]string) (string, error) {
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
		requestBody, _ = json.Marshal(c) // already been unmarshalled, no error
	case resources.EnvironmentType:
		path = []string{"environments", urlParams["ID"]}
		responseValueKey = "environment"

		c := struct {
			Environment map[string]interface{} `json:"environment"`
		}{
			Environment: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.MockType:
		path = []string{"mocks", urlParams["ID"]}
		responseValueKey = "mock"

		c := struct {
			Mock map[string]interface{} `json:"mock"`
		}{
			Mock: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.MonitorType:
		path = []string{"monitors", urlParams["ID"]}
		responseValueKey = "monitor"

		c := struct {
			Monitor map[string]interface{} `json:"monitor"`
		}{
			Monitor: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.WorkspaceType:
		path = []string{"workspaces", urlParams["ID"]}
		responseValueKey = "workspace"

		c := struct {
			Workspace map[string]interface{} `json:"workspace"`
		}{
			Workspace: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.APIType:
		path = []string{"apis", urlParams["ID"]}
		responseValueKey = "api"

		c := struct {
			API map[string]interface{} `json:"api"`
		}{
			API: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.APIVersionType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["ID"]}
		responseValueKey = "version"

		c := struct {
			Version map[string]interface{} `json:"version"`
		}{
			Version: v,
		}
		requestBody, _ = json.Marshal(c)
	case resources.SchemaType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["apiVersionID"], "schemas", urlParams["ID"]}
		responseValueKey = "schema"

		c := struct {
			Schema map[string]interface{} `json:"schema"`
		}{
			Schema: v,
		}
		requestBody, _ = json.Marshal(c)
	default:
		return "", fmt.Errorf("unable to replace resource, %+v not supported", t)
	}

	var responseBody interface{}
	if _, err := s.put(ctx, requestBody, &responseBody, path...); err != nil {
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
