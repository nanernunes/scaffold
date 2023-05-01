package scaffold

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContext_RenderHTML(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	context := NewContext(w, r)

	context.RenderHTML(http.StatusOK, "<h1>Hello, World!</h1>")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	expectedContentType := "text/html"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected content type %s but got %s", expectedContentType, w.Header().Get("Content-Type"))
	}

	expectedResponse := "<h1>Hello, World!</h1>"
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s but got %s", expectedResponse, w.Body.String())
	}
}

func TestContext_RenderText(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	context := NewContext(w, r)

	context.RenderText(http.StatusOK, "Hello, World!")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	expectedContentType := "plain/text"
	if w.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected content type %s but got %s", expectedContentType, w.Header().Get("Content-Type"))
	}

	expectedResponse := "Hello, World!"
	if w.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s but got %s", expectedResponse, w.Body.String())
	}
}

func TestContext_UseCORS(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("OPTIONS", "/", nil)

	context := NewContext(w, r)

	context.UseCORS()

	expectedHeaders := map[string]string{
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Origin":  "*",
	}

	for key, value := range expectedHeaders {
		if w.Header().Get(key) != value {
			t.Errorf("Expected header %s to have value %s but got %s", key, value, w.Header().Get(key))
		}
	}
}
