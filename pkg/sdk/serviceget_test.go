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
	getMux     *http.ServeMux
	getService *sdk.Service
)

func setupGetTest() func() {
	teardown := setupService(&getMux, &getService)

	return teardown
}

func TestCollectionsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/collections"
	subject := "{\"collections\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Collections(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestCollectionsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/collections"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Collections(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestCollectionsItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/collections/abcdef"
	subject := "{\"collection\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Collection(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestCollectionsItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/collections/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Collection(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestEnvironmentsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/environments"
	subject := "{\"environments\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Environments(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestEnvironmentsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/environments"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Environments(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestEnvironmentsItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/environments/abcdef"
	subject := "{\"environment\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Environment(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestEnvironmentsItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/environments/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Environment(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestMocksList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/mocks"
	subject := "{\"mocks\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Mocks(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestMocksListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/mocks"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Mocks(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestMocksItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/mocks/abcdef"
	subject := "{\"mock\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Mock(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestMocksItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/mocks/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Mock(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestMonitorsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/monitors"
	subject := "{\"monitors\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Monitors(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestMonitorsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/monitors"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Monitors(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestMonitorsItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/monitors/abcdef"
	subject := "{\"monitor\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Monitor(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestMonitorsItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/monitors/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Monitor(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestWorkspacesList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/workspaces"
	subject := "{\"workspaces\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Workspaces(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestWorkspacesListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/workspaces"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Workspaces(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestWorkspacesItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/workspaces/abcdef"
	subject := "{\"workspace\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Workspace(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestWorkspacesItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/workspaces/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Workspace(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestUser(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/me"
	subject := "{\"user\":{\"id\":12345}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.User(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestUserError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/me"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.User(context.Background())
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestAPIsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis"
	subject := "{\"apis\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.APIs(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIsListWithWorkspace(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis"
	subject := "{\"apis\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		workspace := r.URL.Query().Get("workspace")
		if workspace != "12345" {
			t.Errorf("Expected workspace ID, have: %s, want: %s", workspace, "12345")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.APIs(context.Background(), "12345")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.APIs(context.Background(), "")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestAPIsItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/abcdef"
	subject := "{\"api\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.API(context.Background(), "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIsItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.API(context.Background(), "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestAPIVersionList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions"
	subject := "{\"versions\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.APIVersions(context.Background(), "12345")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIVersionListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.APIVersions(context.Background(), "12345")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestAPIVersionsItem(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/abcdef"
	subject := "{\"version\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.APIVersion(context.Background(), "12345", "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIVersionsItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.APIVersion(context.Background(), "12345", "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestSchema(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/schemas/abcdef"
	subject := "{\"schema\":{\"uid\":\"abcdef\"}}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.Schema(context.Background(), "12345", "4567", "abcdef")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestSchemasItemError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/schemas/abcdef"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.Schema(context.Background(), "12345", "4567", "abcdef")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}

func TestAPIRelationsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/relations"
	subject := "{\"relations\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.APIRelations(context.Background(), "12345", "4567")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestAPIRelationsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/relations"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.APIRelations(context.Background(), "12345", "4567")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}
func TestFormattedAPIRelationsList(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/relations"
	subject := "{\"relations\":[{\"uid\":\"abcdef\"}]}"

	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subject))
	})

	ensurePath(t, getMux, path)

	r, err := getService.FormattedAPIRelationItems(context.Background(), "12345", "4567")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Errorf("Should return a response.")
	}
}

func TestFormattedAPIRelationsListError(t *testing.T) {
	teardown := setupGetTest()
	defer teardown()

	path := "/apis/12345/versions/4567/relations"
	getMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusInternalServerError)
	})

	ensurePath(t, getMux, path)

	_, err := getService.FormattedAPIRelationItems(context.Background(), "12345", "4567")
	if err == nil {
		t.Errorf("Should return an error.")
	}
}
