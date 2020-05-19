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

package client_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/kevinswiber/postmanctl/pkg/sdk/client"
)

func TestGetMethod(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodGet)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestPost(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPost)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Post().
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestPut(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPut)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Put().
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodDelete)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Delete().
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestMethod(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Method is incorrect, have: %s, want: %s", r.Method, http.MethodPatch)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Method("PATCH").
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestPath(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/one/two/three/four" {
			t.Errorf("Path is incorrect, have: %s, want: %s", r.URL.Path, "/one/two/three/four")
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Path("one", "two", "three", "four").
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestAddHeader(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Expected Content-Type to equal application/json")
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		AddHeader("Content-Type", "application/json").
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestParam(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("hello")
		if param != "world" {
			t.Errorf("Unexpected param, have: %s, want: %s", param, "world")
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Param("hello", "world").
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestParams(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("hello")
		if param != "world" {
			t.Errorf("Unexpected param, have: %s, want: %s", param, "world")
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Params(map[string]string{"hello": "world"}).
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestBody(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	subject := `{"hello":"world"}`
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != subject {
			t.Errorf("Incorrect request body, have: %s, want: %s", string(body), subject)
		}
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	rdr := strings.NewReader(subject)
	_, err := req.
		Get().
		Body(rdr).
		Do()

	if err != nil {
		t.Error(err)
	}
}

func TestInto(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	type subj struct {
		Hello string `json:"hello"`
	}

	subject := `{"hello":"world"}`
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(subject))

		if err != nil {
			t.Error(err)
		}
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	var s subj
	_, err := req.
		Get().
		Into(&s).
		Do()

	if err != nil {
		t.Error(err)
	}

	if s.Hello != "world" {
		t.Errorf("Unexpected value, have: %s, want: %s", s.Hello, "world")
	}
}

func TestHTTPError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Do()
	if err == nil {
		t.Error("Expected error.")
	} else {
		switch err.(type) {
		case *client.RequestError:
			s := err.Error()
			if !strings.Contains(s, "status code: 500") {
				t.Error("Expected error to contain status code: 500")
			}
		default:
			t.Errorf("Incorrect error, expected RequestError, got: %s", err)
		}
	}
}

func TestRequestError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	_, err := req.
		Method("@@").
		Do()

	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestRequestCreatesDefaultClient(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", nil)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Do()

	if err != nil {
		t.Fatal(err)
	}
}

func TestPropagatesClientError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	u, _ := url.Parse("http://domain.false$")
	options := client.NewOptions(u, "", nil)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Do()

	if err == nil {
		t.Error("Expected error.")
	}
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

type errReader string

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("errReader")
}

func (errReader) Close() error {
	return nil
}

func TestPropagateBodyReadErrorOnNon200StatusCode(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	u, _ := url.Parse(server.URL)
	c := &http.Client{
		Transport: roundTripFunc(func(*http.Request) *http.Response {
			var rdr errReader
			return &http.Response{
				StatusCode: 500,
				Body:       rdr,
			}
		}),
	}
	options := client.NewOptions(u, "", c)
	req := client.NewRequest(options)

	_, err := req.
		Get().
		Do()

	if err == nil {
		t.Error("Expected error.")
	}
}

func TestResponseBodyReadErrorOn200StatusCode(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	type subj struct {
		Hello string `json:"hello"`
	}

	c := &http.Client{
		Transport: roundTripFunc(func(*http.Request) *http.Response {
			var rdr errReader
			return &http.Response{
				StatusCode: 200,
				Body:       rdr,
			}
		}),
	}

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", c)
	req := client.NewRequest(options)

	var s subj
	_, err := req.
		Get().
		Into(&s).
		Do()

	if err == nil {
		t.Error("Expected error.")
	}
}

func TestUnmarshalErrorOn200StatusCode(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	type subj struct {
		Hello string `json:"hello"`
	}

	subject := `{"hello":}`
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(subject))

		if err != nil {
			t.Error(err)
		}
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	var s subj
	_, err := req.
		Get().
		Into(&s).
		Do()

	if err == nil {
		t.Error("Expected error.")
	}
}

func TestUnmarshalErrorOnNon200StatusCode(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	type subj struct {
		Hello string `json:"hello"`
	}

	subject := `{"hello":}`
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(subject))

		if err != nil {
			t.Error(err)
		}
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)
	req := client.NewRequest(options)

	var s subj
	_, err := req.
		Get().
		Into(&s).
		Do()

	if err == nil {
		t.Error("Expected error.")
	} else if !strings.Contains(err.Error(), "invalid character '}' looking for beginning of value") {
		t.Error("Expected unmarshal error.")
	}
}

func TestContext(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})

	u, _ := url.Parse(server.URL)
	options := client.NewOptions(u, "", http.DefaultClient)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	req := client.NewRequestWithContext(ctx, options)

	_, err := req.
		Get().
		Do()

	if err == nil {
		t.Error("Expected error.")
	} else if !strings.HasSuffix(err.Error(), "context deadline exceeded") {
		t.Errorf("Unexpected error, have: %s, want: context deadline exceeded", err)
	}
}
