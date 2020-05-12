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
	"context"
	"net/http"
	"testing"

	"github.com/kevinswiber/postmanctl/pkg/sdk"
)

func TestServiceRunMonitor(t *testing.T) {
	var (
		mux     *http.ServeMux
		service *sdk.Service
	)

	teardown := setupService(&mux, &service)
	defer teardown()

	path := "/monitors/3/run"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"run\":{}}"))
	})

	ensurePath(t, mux, path)

	msg, err := service.RunMonitor(context.Background(), "3")
	if err != nil {
		t.Fatal(err)
	}

	b, err := msg.MarshalJSON()

	if err != nil {
		t.Fatal(err)
	}

	subject := "{\"run\":{}}"

	if string(b) != subject {
		t.Errorf("Response body is incorrect, have: %s, want: %s", string(b), subject)
	}
}

func TestServiceRunMonitorError(t *testing.T) {
	var (
		mux     *http.ServeMux
		service *sdk.Service
	)

	teardown := setupService(&mux, &service)
	defer teardown()

	path := "/monitors/3/run"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, mux, path)

	_, err := service.RunMonitor(context.Background(), "3")
	if err == nil {
		t.Error("Expected error")
	}
}
