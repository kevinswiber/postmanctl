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

var (
	deleteMux     *http.ServeMux
	deleteService *sdk.Service
)

func setupDeleteTest() func() {
	teardown := setupService(&deleteMux, &deleteService)

	return teardown
}

func TestDeleteCollection(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteCollection(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteEnvironment(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/environments/abcdef"
	subject := "{\"environment\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteEnvironment(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteMock(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/mocks/abcdef"
	subject := "{\"mock\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteMock(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteMonitor(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/monitors/abcdef"
	subject := "{\"monitor\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteMonitor(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteWorkspace(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/workspaces/abcdef"
	subject := "{\"workspace\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteWorkspace(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteAPI(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/apis/abcdef"
	subject := "{\"api\":{\"uid\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteAPI(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteAPIVersion(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/apis/12345/versions/abcdef"
	subject := "{\"version\":{\"id\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteAPIVersion(context.Background(), "abcdef", "12345")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteAPIVersionError(t *testing.T) {
	_, err := deleteService.DeleteAPIVersion(context.Background(), "abcdef", "")
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestDeleteSchema(t *testing.T) {
	teardown := setupDeleteTest()
	defer teardown()

	path := "/apis/12345/versions/4567/schemas/abcdef"
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	deleteMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, deleteMux, path)

	r, err := deleteService.DeleteSchema(context.Background(), "abcdef", "12345", "4567")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestDeleteSchemaMissingAPIIDError(t *testing.T) {
	_, err := deleteService.DeleteSchema(context.Background(), "abcdef", "", "4567")
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestDeleteSchemaMissingAPIVersionIDError(t *testing.T) {
	_, err := deleteService.DeleteSchema(context.Background(), "abcdef", "12345", "")
	if err == nil {
		t.Errorf("Expected error.")
	}
}
