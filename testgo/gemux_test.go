package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/fharding1/gemux"
)

func stringHandler(s string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, s)
	})
}

func pathParametersHandler(t *testing.T, s string, expectedParams []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := PathParameters(r.Context())
		if !reflect.DeepEqual(expectedParams, params) {
			t.Errorf("expected path parameters %v, but got %v", expectedParams, params)
		}

		io.WriteString(w, s)
	})
}

type handlerArgs struct {
	pattern string
	method  string
	handler http.Handler
}

func TestServeMux(t *testing.T) {
	cases := []struct {
		name                    string
		notFoundHandler         http.Handler
		methodNotAllowedHandler http.Handler
		register                []handlerArgs
		requestURL              string
		requestMethod           string
		expectedResponseCode    int
		expectedResponseBody    string
	}{
		{
			name: "root",
			register: []handlerArgs{
				{
					pattern: "/",
					method:  http.MethodGet,
					handler: stringHandler("a"),
				},
			},
			requestURL:           "/",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "a",
		},
		{
			name:                 "root not found",
			register:             []handlerArgs{},
			requestURL:           "/",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: "404 page not found\n",
		},
		{
			name: "root method not allowed",
			register: []handlerArgs{
				{
					pattern: "/",
					method:  http.MethodGet,
					handler: stringHandler("a"),
				},
			},
			requestURL:           "/",
			requestMethod:        "POST",
			expectedResponseCode: http.StatusMethodNotAllowed,
			expectedResponseBody: "405 method not allowed\n",
		},
		{
			name: "root wildcard method",
			register: []handlerArgs{
				{
					pattern: "/",
					method:  "*",
					handler: stringHandler("a"),
				},
			},
			requestURL:           "/",
			requestMethod:        "PUT",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "a",
		},
		{
			name: "wildcard path",
			register: []handlerArgs{
				{
					pattern: "/*",
					method:  "GET",
					handler: pathParametersHandler(t, "a", []string{"foo"}),
				},
			},
			requestURL:           "/foo",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "a",
		},
		{
			name: "multiple children paths",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/bar",
					method:  "GET",
					handler: stringHandler("b"),
				},
				{
					pattern: "/baz",
					method:  "GET",
					handler: stringHandler("c"),
				},
			},
			requestURL:           "/bar",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "b",
		},
		{
			name: "same child different methods",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/foo",
					method:  "PUT",
					handler: stringHandler("b"),
				},
				{
					pattern: "/foo",
					method:  "PATCH",
					handler: stringHandler("c"),
				},
			},
			requestURL:           "/foo",
			requestMethod:        "PATCH",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "c",
		},
		{
			name: "child not found",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/bar",
					method:  "GET",
					handler: stringHandler("b"),
				},
			},
			requestURL:           "/boo",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusNotFound,
			expectedResponseBody: "404 page not found\n",
		},
		{
			name: "child custom not found",
			notFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "custom handler: was not found")
			}),
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/bar",
					method:  "GET",
					handler: stringHandler("b"),
				},
			},
			requestURL:           "/boo",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusBadRequest,
			expectedResponseBody: "custom handler: was not found\n",
		},
		{
			name: "child method not allowed",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/foo",
					method:  "PUT",
					handler: stringHandler("b"),
				},
			},
			requestURL:           "/foo",
			requestMethod:        "PATCH",
			expectedResponseCode: http.StatusMethodNotAllowed,
			expectedResponseBody: "405 method not allowed\n",
		},
		{
			name: "child custom method not allowed",
			methodNotAllowedHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusAccepted)
				io.WriteString(w, "accepted?")
			}),
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: stringHandler("a"),
				},
				{
					pattern: "/foo",
					method:  "PUT",
					handler: stringHandler("b"),
				},
			},
			requestURL:           "/foo",
			requestMethod:        "PATCH",
			expectedResponseCode: http.StatusAccepted,
			expectedResponseBody: "accepted?",
		},
		{
			name: "no path parameters",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  "GET",
					handler: pathParametersHandler(t, "a", nil),
				},
			},
			requestURL:           "/foo",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusOK,
			expectedResponseBody: "a",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mux := new(ServeMux)

			mux.NotFoundHandler = tt.notFoundHandler
			mux.MethodNotAllowedHandler = tt.methodNotAllowedHandler

			for _, route := range tt.register {
				mux.Handle(route.pattern, route.method, route.handler)
			}

			rw := httptest.NewRecorder()
			req, err := http.NewRequest(tt.requestMethod, tt.requestURL, nil)
			if err != nil {
				t.Fatalf("did not expect error setting up test: %v\n", err)
			}

			mux.ServeHTTP(rw, req)

			if rw.Code != tt.expectedResponseCode {
				t.Errorf("expected response code %d, got %d", tt.expectedResponseCode, rw.Code)
			}

			if body := rw.Body.String(); body != tt.expectedResponseBody {
				t.Errorf("expected response body %q, got %q", tt.expectedResponseBody, body)
			}
		})
	}
}

func ExampleServeMux() {
	mux := new(ServeMux)

	mux.Handle("/", http.MethodGet, stringHandler("health check"))
	mux.Handle("/posts", http.MethodGet, stringHandler("create post"))
	mux.Handle("/posts", http.MethodGet, stringHandler("get posts"))
	mux.Handle("/posts/*", http.MethodGet, stringHandler("get post"))
	mux.Handle("/posts/*/comments", http.MethodGet, stringHandler("get post comments"))

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(rw, req)
	fmt.Println(rw.Body.String())

	rw = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/4", nil)
	mux.ServeHTTP(rw, req)
	fmt.Println(rw.Body.String())

	rw = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/posts/5/comments", nil)
	mux.ServeHTTP(rw, req)
	fmt.Println(rw.Body.String())

	// Output:
	// health check
	// get post
	// get post comments
}

func ExamplePathParameters() {
	mux := new(ServeMux)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathParameters := PathParameters(r.Context())
		fmt.Println(pathParameters)
	})

	mux.Handle("/foo/*/bar/*", http.MethodGet, handler)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foo/test/bar/92", nil)
	mux.ServeHTTP(rw, req)
	fmt.Println(rw.Body.String())

	// Output:
	// [test 92]
}

func BenchmarkServeHTTP(b *testing.B) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	cases := []struct {
		name          string
		register      []handlerArgs
		requestURL    string
		requestMethod string
	}{
		{
			name: "one static path",
			register: []handlerArgs{
				{
					pattern: "/foo",
					method:  http.MethodGet,
					handler: handler,
				},
			},
			requestURL:    "/foo",
			requestMethod: http.MethodGet,
		},
		{
			name: "one wildcard path",
			register: []handlerArgs{
				{
					pattern: "/*",
					method:  http.MethodGet,
					handler: handler,
				},
			},
			requestURL:    "/foo",
			requestMethod: http.MethodGet,
		},
		{
			name: "one wildcard path and method",
			register: []handlerArgs{
				{
					pattern: "/*",
					method:  "*",
					handler: handler,
				},
			},
			requestURL:    "/foo",
			requestMethod: http.MethodGet,
		},
		{
			name: "short path with many routes",
			register: []handlerArgs{
				{"/", http.MethodGet, handler},
				{"/openapi.yaml", http.MethodGet, handler},
				{"/users", http.MethodPost, handler},
				{"/users/*", http.MethodGet, handler},
				{"/users/*", http.MethodPatch, handler},
				{"/users/*", http.MethodDelete, handler},
				{"/schemas", http.MethodGet, handler},
				{"/schemas", http.MethodPost, handler},
				{"/schemas/*", http.MethodGet, handler},
				{"/events", http.MethodGet, handler},
				{"/events/*", http.MethodPut, handler},
				{"/events/*", http.MethodGet, handler},
				{"/events/*/stats", http.MethodGet, handler},
				{"/events/*/matches", http.MethodGet, handler},
				{"/events/*/matches", http.MethodGet, handler},
				{"/events/*/matches/*", http.MethodGet, handler},
				{"/events/*/matches/*/reports/*", http.MethodPost, handler},
			},
			requestURL:    "/openapi.yaml",
			requestMethod: "GET",
		},
		{
			name: "very deep static path",
			register: []handlerArgs{
				{
					pattern: "/a/b/c/d/e",
					method:  http.MethodGet,
					handler: handler,
				},
			},
			requestURL:    "/a/b/c/d/e",
			requestMethod: http.MethodGet,
		},
		{
			name: "very deep wildcard path",
			register: []handlerArgs{
				{
					pattern: "/*/*/*/*/*",
					method:  http.MethodGet,
					handler: handler,
				},
			},
			requestURL:    "/a/b/c/d/e",
			requestMethod: http.MethodGet,
		},
	}

	for _, tt := range cases {
		b.Run(tt.name, func(b *testing.B) {
			mux := new(ServeMux)

			for _, route := range tt.register {
				mux.Handle(route.pattern, route.method, route.handler)
			}

			req, _ := http.NewRequest(tt.requestMethod, tt.requestURL, nil)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				mux.ServeHTTP(nil, req)
			}
		})
	}
}
