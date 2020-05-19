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
	"github.com/kevinswiber/postmanctl/pkg/sdk/resources"
)

var (
	replaceMux     *http.ServeMux
	replaceService *sdk.Service
)

func setupReplaceTest() func() {
	teardown := setupService(&replaceMux, &replaceService)

	return teardown
}

func TestReplaceCollectionFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceCollectionFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceCollectionFromReaderError(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	_, err := replaceService.ReplaceCollectionFromReader(context.Background(), rdr, "abcdef")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestReplaceCollectionFromReaderMissingIDCondition(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := `{"collection":{}}`

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceCollectionFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestReplaceCollectionFromReaderMissingResponseValueCondition(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := `{"blah":{}}`

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceCollectionFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestReplaceEnvironmentFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/environments/abcdef"
	subject := "{\"environment\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceEnvironmentFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceMockFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/mocks/abcdef"
	subject := "{\"mock\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceMockFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceMonitorFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/monitors/abcdef"
	subject := "{\"monitor\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceMonitorFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceWorkspaceFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/workspaces/abcdef"
	subject := "{\"workspace\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceWorkspaceFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceAPIFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/apis/abcdef"
	subject := "{\"api\":{\"uid\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceAPIFromReader(context.Background(), rdr, "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource UID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceAPIVersionFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/apis/1234/versions/abcdef"
	subject := "{\"version\":{\"id\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceAPIVersionFromReader(context.Background(), rdr, "abcdef", "1234")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceAPIVersionFromReaderErrorMissingAPIID(t *testing.T) {
	subject := "{\"version\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := replaceService.ReplaceAPIVersionFromReader(context.Background(), rdr, "abcdef", "")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestReplaceSchemaFromReaderErrorMissingAPIID(t *testing.T) {
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := replaceService.ReplaceSchemaFromReader(context.Background(), rdr, "abcdef", "", "4567")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestReplaceSchemaFromReaderErrorMissingAPIVersionID(t *testing.T) {
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	rdr := strings.NewReader(subject)
	_, err := replaceService.ReplaceSchemaFromReader(context.Background(), rdr, "abcdef", "12345", "")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestReplaceSchemaFromReader(t *testing.T) {
	teardown := setupReplaceTest()
	defer teardown()

	path := "/apis/1234/versions/5678/schemas/abcdef"
	subject := "{\"schema\":{\"id\":\"abcdef\"}}"

	replaceMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, replaceMux, path)

	rdr := strings.NewReader(subject)
	r, err := replaceService.ReplaceSchemaFromReader(context.Background(), rdr, "abcdef", "1234", "5678")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestReplaceFromReaderReadError(t *testing.T) {
	urlParams := make(map[string]string)
	var rdr errReader
	_, err := replaceService.ReplaceFromReader(context.Background(), resources.CollectionType, rdr, urlParams)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestReplaceFromReaderMarshalError(t *testing.T) {
	urlParams := make(map[string]string)
	rdr := strings.NewReader(`{"eRr0r":}`)
	_, err := replaceService.ReplaceFromReader(context.Background(), resources.CollectionType, rdr, urlParams)
	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestReplaceFromReaderInvalidResourceError(t *testing.T) {
	urlParams := make(map[string]string)
	rdr := strings.NewReader(`{"collection":{}}`)
	_, err := replaceService.ReplaceFromReader(context.Background(), 99, rdr, urlParams)
	if err == nil {
		t.Errorf("Expected error.")
	}
}
