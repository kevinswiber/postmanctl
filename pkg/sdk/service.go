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
	if _, err := s.get(ctx, &resource, "collections"); err != nil {
		return nil, err
	}

	return &resource.Collections, nil
}

// Collection returns a single collection.
func (s *Service) Collection(ctx context.Context, id string) (*resources.Collection, error) {
	var resource resources.CollectionResponse
	if _, err := s.get(ctx, &resource, "collections", id); err != nil {
		return nil, err
	}

	return &resource.Collection, nil
}

// Environments returns all environments.
func (s *Service) Environments(ctx context.Context) (*resources.EnvironmentListItems, error) {
	var resource resources.EnvironmentListResponse
	if _, err := s.get(ctx, &resource, "environments"); err != nil {
		return nil, err
	}

	return &resource.Environments, nil
}

// Environment returns a single environment.
func (s *Service) Environment(ctx context.Context, id string) (*resources.Environment, error) {
	var resource resources.EnvironmentResponse
	if _, err := s.get(ctx, &resource, "environments", id); err != nil {
		return nil, err
	}

	return &resource.Environment, nil
}

// APIs returns all APIs.
func (s *Service) APIs(ctx context.Context) (*resources.APIListItems, error) {
	var resource resources.APIListResponse
	if _, err := s.get(ctx, &resource, "apis"); err != nil {
		return nil, err
	}

	return &resource.APIs, nil
}

// API returns a single API.
func (s *Service) API(ctx context.Context, id string) (*resources.API, error) {
	var resource resources.APIResponse
	if _, err := s.get(ctx, &resource, "apis", id); err != nil {
		return nil, err
	}

	return &resource.API, nil
}

// User returns the current user.
func (s *Service) User(ctx context.Context) (*resources.User, error) {
	var resource resources.UserResponse
	if _, err := s.get(ctx, &resource, "me"); err != nil {
		return nil, err
	}

	return &resource.User, nil
}

func (s *Service) get(ctx context.Context, r interface{}, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Get().
		Path(path...).
		Body(&r).
		Do()

	return res, err
}
