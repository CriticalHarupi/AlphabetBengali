package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"alphabetbengali/handlers"
	"alphabetbengali/models"
)

func setupConjunctRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/conjuncts", handlers.GetConjuncts)
	r.GET("/api/conjuncts/:id", handlers.GetConjunctByID)
	return r
}

func TestGetConjuncts_ReturnsOK(t *testing.T) {
	r := setupConjunctRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/conjuncts", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var list []models.Conjunct
	if err := json.Unmarshal(w.Body.Bytes(), &list); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if len(list) == 0 {
		t.Error("expected non-empty conjunct list")
	}
}

func TestGetConjunctByID_ReturnsCorrect(t *testing.T) {
	r := setupConjunctRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/conjuncts/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var cj models.Conjunct
	if err := json.Unmarshal(w.Body.Bytes(), &cj); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if cj.ID != 1 {
		t.Errorf("expected id 1, got %d", cj.ID)
	}
	if len(cj.Components) == 0 {
		t.Error("expected non-empty components")
	}
}

func TestGetConjunctByID_NotFound(t *testing.T) {
	r := setupConjunctRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/conjuncts/9999", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected JSON error body: %s", w.Body.String())
	}
	if body["error"] == "" {
		t.Error("expected non-empty error message")
	}
}

func TestGetConjunctByID_InvalidID(t *testing.T) {
	r := setupConjunctRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/conjuncts/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}
