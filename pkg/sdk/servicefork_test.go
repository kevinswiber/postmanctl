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
	forkMux     *http.ServeMux
	forkService *sdk.Service
)

func setupForkTest() func() {
	teardown := setupService(&forkMux, &forkService)

	return teardown
}

func TestForkCollection(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/fork/abcdef"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.ForkCollection(context.Background(), "abcdef", "12345", "forkd")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestForkCollectionMissingIDCondition(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/fork/abcdef"
	subject := `{"collection": {}}`

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.ForkCollection(context.Background(), "abcdef", "12345", "forkd")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestForkCollectionMissingResponseValueKeyCondition(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/fork/abcdef"
	subject := `{"blah":{"uid":"abcdef"}}`

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.ForkCollection(context.Background(), "abcdef", "12345", "forkd")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestForkCollectionError(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/fork/abcdef"
	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, forkMux, path)

	_, err := forkService.ForkCollection(context.Background(), "abcdef", "12345", "forkd")
	if err == nil {
		t.Error("Expected error.")
	}
}

func TestMergeCollection(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/merge"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.MergeCollection(context.Background(), "abcdef", "ghijkl", "")
	if err != nil {
		t.Fatal(err)
	}

	if r != "abcdef" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "abcdef")
	}
}

func TestMergeCollectionMissingIDCondition(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/merge"
	subject := `{"collection":{}}`

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.MergeCollection(context.Background(), "abcdef", "ghijkl", "")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestMergeCollectionMissingResponseValueKeyCondition(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/merge"
	subject := `{"blah":{}}`

	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(subject)); err != nil {
			t.Error(err)
		}
	})

	ensurePath(t, forkMux, path)

	r, err := forkService.MergeCollection(context.Background(), "abcdef", "ghijkl", "")
	if err != nil {
		t.Fatal(err)
	}

	if r != "" {
		t.Errorf("Resource ID is incorrect, have: %s, want: %s", r, "")
	}
}

func TestMergeCollectionError(t *testing.T) {
	teardown := setupForkTest()
	defer teardown()

	path := "/collections/merge"
	forkMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, forkMux, path)

	_, err := forkService.MergeCollection(context.Background(), "abcdef", "ghijkl", "")
	if err == nil {
		t.Error("Expected error.")
	}
}
