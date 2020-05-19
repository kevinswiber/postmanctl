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
)

// ForkCollection makes a fork of an existing collection.
func (s *Service) ForkCollection(ctx context.Context, id, workspace, label string) (string, error) {
	responseValueKey := "collection"
	queryParams := make(map[string]string)
	queryParams["workspace"] = workspace

	input := struct {
		Label string `json:"label"`
	}{
		Label: label,
	}

	// swallow error here, strings will always marshal
	requestBody, _ := json.Marshal(input)

	var responseBody interface{}
	if _, err := s.post(ctx, requestBody, &responseBody, queryParams, "collections", "fork", id); err != nil {
		return "", err
	}

	// Try a best attempt at returning the ID value.
	responseValue := responseBody.(map[string]interface{})
	if v, ok := responseValue[responseValueKey]; ok {
		vMap := v.(map[string]interface{})
		if v2, ok := vMap["uid"]; ok {
			return v2.(string), nil
		}

		return "", nil
	}
	return "", nil
}

// MergeCollection makes a fork of an existing collection.
func (s *Service) MergeCollection(ctx context.Context, id, destination, strategy string) (string, error) {
	responseValueKey := "collection"

	input := struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
		Strategy    string `json:"strategy"`
	}{
		Source:      id,
		Destination: destination,
		Strategy:    strategy,
	}

	// swallow error here, strings will always marshal
	requestBody, _ := json.Marshal(input)

	var responseBody interface{}
	if _, err := s.post(ctx, requestBody, &responseBody, nil, "collections", "merge"); err != nil {
		return "", err
	}

	// Try a best attempt at returning the ID value.
	responseValue := responseBody.(map[string]interface{})
	if v, ok := responseValue[responseValueKey]; ok {
		vMap := v.(map[string]interface{})
		if v2, ok := vMap["uid"]; ok {
			return v2.(string), nil
		}

		return "", nil
	}
	return "", nil
}
