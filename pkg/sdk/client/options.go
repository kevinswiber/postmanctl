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

package client

import (
	"net/http"
	"net/url"
	"strings"
)

// Options allows for storing a base URL and containing common functionality.
type Options struct {
	base   *url.URL
	APIKey string
	Client *http.Client
}

// NewOptions creates a new instance of the Postman API client options.
func NewOptions(baseURL *url.URL, apiKey string, client *http.Client) *Options {
	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &Options{
		base:   &base,
		APIKey: apiKey,
		Client: client,
	}
}
