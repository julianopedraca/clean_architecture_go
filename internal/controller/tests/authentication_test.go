package tests

import (
	"julianopedraca/clean_architecture_go/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAuthentication(t *testing.T) {
	router := router.Setup()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/authentication", nil)
	assert.NilError(t, err, "Failed to create request")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP status 200")
	assert.Equal(t, `{"response":"true"}`, w.Body.String(), `Expected result:"true"`)
}
