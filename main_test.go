// go test
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/hello":  helloHandler,
		"/data":   dataHandler,
		"/submit": submitHandler,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := routes[r.URL.Path]; ok {
			handler(w, r)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := server.Client()

	testCases := []struct {
		path         string
		method       string
		expectStatus int
	}{
		{"/hello?name=Test", http.MethodGet, http.StatusOK},
		{"/data?id=1", http.MethodGet, http.StatusOK},
		{"/submit", http.MethodPost, http.StatusOK},
		{"/nonexistent", http.MethodGet, http.StatusNotFound},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, server.URL+tc.path, strings.NewReader("{}"))
		if err != nil {
			t.Fatal(err)
		}
		if tc.method == http.MethodPost {
			req.Header.Set("Content-Type", "application/json")
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectStatus {
			t.Errorf("for %s, expected status %v, got %v", tc.path, tc.expectStatus, resp.StatusCode)
		}
	}
}
