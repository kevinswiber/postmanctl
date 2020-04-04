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
	"net/http"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
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

func (s *Service) put(ctx context.Context, input []byte, output interface{}, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Put().
		Path(path...).
		AddHeader("Content-Type", "application/json").
		Body(bytes.NewReader(input)).
		Into(&output).
		Do()

	return res, err
}

func (s *Service) delete(ctx context.Context, output interface{}, path ...string) (*http.Response, error) {
	req := client.NewRequestWithContext(ctx, s.Options)
	res, err := req.Put().
		Path(path...).
		AddHeader("Content-Type", "application/json").
		Into(&output).
		Do()

	return res, err
}
