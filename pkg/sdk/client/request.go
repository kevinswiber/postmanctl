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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

// RequestError represents an error from the Postman API.
type RequestError struct {
	StatusCode int
	Name       string
	Message    string
	Details    map[string]interface{}
}

// NewRequestError creates a new RequestError for Postman API responses.
func NewRequestError(code int, name string, message string, details map[string]interface{}) *RequestError {
	return &RequestError{
		StatusCode: code,
		Name:       name,
		Message:    message,
		Details:    details,
	}
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("status code: %d, name: %s, message: %s, details %s", e.StatusCode,
		e.Name, e.Message, e.Details)
}

// Request holds state for a Postman API request.
type Request struct {
	ctx           context.Context
	options       *Options
	method        string
	path          string
	requestReader io.Reader
	result        interface{}
	headers       http.Header
	params        url.Values
	err           error
}

// NewRequest initializes a Postman API Request.
func NewRequest(c *Options) *Request {
	return NewRequestWithContext(context.Background(), c)
}

// NewRequestWithContext intiializes a Postman API Request with a
// given context.
func NewRequestWithContext(ctx context.Context, c *Options) *Request {
	r := &Request{
		ctx:     ctx,
		options: c,
	}

	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Add("X-API-Key", c.APIKey)

	return r
}

// AddHeader adds a header to the request.
func (r *Request) AddHeader(key string, value string) *Request {
	r.headers.Add(key, value)
	return r
}

// Param sets a query parameter.
func (r *Request) Param(k, v string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}

	r.params.Add(k, v)

	return r
}

// Params sets all query parameters.
func (r *Request) Params(params map[string]string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}

	for k, v := range params {
		r.params.Add(k, v)
	}

	return r
}

// Method sets the HTTP method of the request.
func (r *Request) Method(m string) *Request {
	r.method = m
	return r
}

// Get sets the HTTP method to GET
func (r *Request) Get() *Request {
	r.method = "GET"
	return r
}

// Post sets the HTTP method to POST
func (r *Request) Post() *Request {
	r.method = "POST"
	return r
}

// Put sets the HTTP method to PUT
func (r *Request) Put() *Request {
	r.method = "PUT"
	return r
}

// Delete sets the HTTP method to DELETE
func (r *Request) Delete() *Request {
	r.method = "DELETE"
	return r
}

// Path sets the path of the HTTP request.
func (r *Request) Path(p ...string) *Request {
	r.path = path.Join(p...)
	return r
}

// Body sets an input resource for the request
func (r *Request) Body(reader io.Reader) *Request {
	r.requestReader = reader
	return r
}

// Into sets a destination resource for the output response
func (r *Request) Into(o interface{}) *Request {
	r.result = o
	return r
}

// URL returns a complete URL for the current request.
func (r *Request) URL() *url.URL {
	finalURL := &url.URL{}
	if r.options.base != nil {
		*finalURL = *r.options.base
	}
	finalURL.Path = r.path
	finalURL.RawQuery = r.params.Encode()

	return finalURL
}

// Do executes the HTTP request.
func (r *Request) Do() (*http.Response, error) {
	url := r.URL().String()

	req, err := http.NewRequestWithContext(r.ctx, r.method, url, r.requestReader)
	if err != nil {
		return nil, err
	}
	req.Header = r.headers
	client := r.options.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return resp, err
		}

		var e resources.ErrorResponse
		json.Unmarshal(body, &e)
		errorMessage := NewRequestError(resp.StatusCode, e.Error.Name, e.Error.Message, e.Error.Details)
		r.err = errorMessage
		return nil, errorMessage
	}

	if r.result != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return resp, err
		}

		if err := json.Unmarshal(body, &r.result); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
