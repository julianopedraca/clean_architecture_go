package tests

import (
	"fmt"
	"julianopedraca/clean_architecture_go/internal/router"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestAuthentication(t *testing.T) {
	router := router.Setup()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/authentication", nil)
	assert.NilError(t, err, "Failed to create request")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP status 200")
	assert.Equal(t, `{"response":"true"}`, w.Body.String(), `Expected result:"true"`)
}

// TODO: this test isn't working because it is not saving the cookies that come as response from authentication
func TestSearchLine(t *testing.T) {
	os.Getenv("API_KEY")
	os.Getenv("API_URL")
	router := router.Setup()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/authentication", nil)
	assert.NilError(t, err, "Failed to create new request")

	client := &ClientMock{}

	resp, err := client.Do(req)
	fmt.Println("🚀 ~ file: routes_test.go ~ line 41 ~ funcTestSearchLine ~ resp : ", resp)
	assert.NilError(t, err, "Failed to create request")
	// router.ServeHTTP(w, req)

	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatal("No cookies found in authentication response")
	}

	w2 := httptest.NewRecorder()
	req2, err := http.NewRequest("GET", "/searchline?line=8000", nil)
	assert.NilError(t, err, "Failed to create searchline request")
	for _, cookie := range cookies {
		req2.AddCookie(cookie)
	}
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code, "Expected HTTP status 200")
}
