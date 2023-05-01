package scaffold

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockController struct{}

func (mc *mockController) Index(c *Context)  {}
func (mc *mockController) Create(c *Context) {}
func (mc *mockController) Show(c *Context)   {}
func (mc *mockController) Update(c *Context) {}
func (mc *mockController) Delete(c *Context) {}

func TestNewRouter(t *testing.T) {
	r := NewRouter(nil)
	if r.Internal == nil {
		t.Errorf("expected router to be initialized, got nil")
	}
}

func TestRouterResource(t *testing.T) {
	router := mux.NewRouter()
	r := NewRouter(router)

	controller := &mockController{}

	r.Resource("/users", controller)

	testCases := []struct {
		method string
		path   string
	}{
		{"GET", "/users"},
		{"POST", "/users"},
		{"GET", "/users/123"},
		{"PATCH", "/users/123"},
		{"DELETE", "/users/123"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %s", tc.method, tc.path), func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
			}
		})
	}
}

func TestRouterOPTIONS(t *testing.T) {
	router := mux.NewRouter()
	r := NewRouter(router)

	var called bool
	r.OPTIONS("/users", func(ctx *Context) {
		called = true
	})

	req, err := http.NewRequest(http.MethodOptions, "/users", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	if !called {
		t.Error("expected callback to be called, but it wasn't")
	}
}
