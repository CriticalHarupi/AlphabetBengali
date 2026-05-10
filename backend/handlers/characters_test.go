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

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/characters", handlers.GetCharacters)
	r.GET("/api/characters/:id", handlers.GetCharacterByID)
	return r
}

func TestGetCharacters_ReturnsOK(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/characters", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var chars []models.Character
	if err := json.Unmarshal(w.Body.Bytes(), &chars); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}
	if len(chars) == 0 {
		t.Error("expected non-empty character list")
	}
}

func TestGetCharacterByID_ReturnsCorrectCharacter(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/characters/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var ch models.Character
	if err := json.Unmarshal(w.Body.Bytes(), &ch); err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}
	if ch.ID != 1 {
		t.Errorf("expected id 1, got %d", ch.ID)
	}
}

func TestGetCharacterByID_NotFound(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/characters/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected JSON error body, got: %s", w.Body.String())
	}
	if body["error"] == "" {
		t.Error("expected non-empty error message in body")
	}
}

func TestGetCharacterByID_InvalidID(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/characters/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected JSON error body, got: %s", w.Body.String())
	}
	if body["error"] == "" {
		t.Error("expected non-empty error message in body")
	}
}
