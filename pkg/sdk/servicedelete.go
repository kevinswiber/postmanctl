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
	"errors"
	"fmt"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

// DeleteCollection deletes a collection.
func (s *Service) DeleteCollection(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.CollectionType, urlParams)
}

// DeleteEnvironment deletes a environment.
func (s *Service) DeleteEnvironment(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.EnvironmentType, urlParams)
}

// DeleteMock deletes a mock.
func (s *Service) DeleteMock(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.MockType, urlParams)
}

// DeleteMonitor deletes a monitor.
func (s *Service) DeleteMonitor(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.MonitorType, urlParams)
}

// DeleteWorkspace deletes a API.
func (s *Service) DeleteWorkspace(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.WorkspaceType, urlParams)
}

// DeleteAPI deletes a API.
func (s *Service) DeleteAPI(ctx context.Context, resourceID string) (string, error) {
	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID

	return s.Delete(ctx, resources.APIType, urlParams)
}

// DeleteAPIVersion deletes a API Version.
func (s *Service) DeleteAPIVersion(ctx context.Context, resourceID, apiID string) (string, error) {
	if apiID == "" {
		return "", errors.New("an API ID is required for creating a new API version")
	}

	urlParams := make(map[string]string)
	urlParams["ID"] = resourceID
	urlParams["apiID"] = apiID

	return s.Delete(ctx, resources.APIVersionType, urlParams)
}

// DeleteSchema deletes a API Version.
func (s *Service) DeleteSchema(ctx context.Context, resourceID, apiID, apiVersionID string) (string, error) {
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

	return s.Delete(ctx, resources.APIVersionType, urlParams)
}

// Delete posts a new resource to the Postman API.
func (s *Service) Delete(ctx context.Context, t resources.ResourceType, urlParams map[string]string) (string, error) {
	var (
		path             []string
		responseValueKey string
	)

	switch t {
	case resources.CollectionType:
		path = []string{"collections", urlParams["ID"]}
		responseValueKey = "collection"
	case resources.EnvironmentType:
		path = []string{"environments", urlParams["ID"]}
		responseValueKey = "environment"
	case resources.MockType:
		path = []string{"mocks", urlParams["ID"]}
		responseValueKey = "mock"
	case resources.MonitorType:
		path = []string{"monitors", urlParams["ID"]}
		responseValueKey = "monitor"
	case resources.APIType:
		path = []string{"apis", urlParams["ID"]}
		responseValueKey = "api"
	case resources.APIVersionType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["ID"]}
		responseValueKey = "version"
	case resources.SchemaType:
		path = []string{"apis", urlParams["apiID"], "versions", urlParams["apiVersionID"], "schemas", urlParams["ID"]}
		responseValueKey = "version"
	default:
		return "", fmt.Errorf("unable to delete resource, %+v not supported", t)
	}

	var responseBody interface{}
	if _, err := s.delete(ctx, &responseBody, path...); err != nil {
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
