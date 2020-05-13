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

package sdk_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/kevinswiber/postmanctl/pkg/sdk"
	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
)

func NewTestService(server *httptest.Server) *sdk.Service {
	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)

	return sdk.NewService(options)
}

func ensurePath(t *testing.T, mux *http.ServeMux, path string) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("API call path is incorrect, have: %s, want: %s", r.URL, path)
	})
}

func setupService(mux **http.ServeMux, service **sdk.Service) func() {
	*mux = http.NewServeMux()
	server := httptest.NewServer(*mux)
	*service = NewTestService(server)

	return func() {
		server.Close()
	}
}

type errReader string

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("errReader")
}
