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
	"strings"
	"testing"

	"github.com/kevinswiber/postmanctl/pkg/sdk"
)

var (
	createMux     *http.ServeMux
	createService *sdk.Service
)

func setupCreateTest() func() {
	teardown := setupService(&createMux, &createService)

	return teardown
}

func TestCreateCollectionFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/collections"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateCollectionFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateEnvironmentFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/environments"
	subject := "{\"environment\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateEnvironmentFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateMockFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/mocks"
	subject := "{\"mock\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateMockFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateMonitorFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/monitors"
	subject := "{\"monitor\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateMonitorFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateWorkspaceFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/workspaces"
	subject := "{\"workspace\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateWorkspaceFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateAPIFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/apis"
	subject := "{\"api\":{\"uid\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateAPIFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateAPIVersionFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/apis/1234/versions"
	subject := "{\"version\":{\"id\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateAPIVersionFromReader(context.Background(), rdr, "abcdef", "1234")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestCreateAPIVersionFromReaderErrorMissingAPIID(t *testing.T) {
	subject := "{\"version\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := createService.CreateAPIVersionFromReader(context.Background(), rdr, "abcdef", "")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestCreateSchemaFromReaderErrorMissingAPIID(t *testing.T) {
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := createService.CreateSchemaFromReader(context.Background(), rdr, "abcdef", "", "4567")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestCreateSchemaFromReaderErrorMissingAPIVersionID(t *testing.T) {
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := createService.CreateSchemaFromReader(context.Background(), rdr, "abcdef", "12345", "")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestCreateSchemaFromReader(t *testing.T) {
	teardown := setupCreateTest()
	defer teardown()

	path := "/apis/1234/versions/5678/schemas"
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	createMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, createMux, path)

	rdr := strings.NewReader(subject)
	r, err := createService.CreateSchemaFromReader(context.Background(), rdr, "", "1234", "5678")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}
