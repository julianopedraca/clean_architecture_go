package utils

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func GetBodyFromHttpResponse(body io.ReadCloser) []byte {
	parsedBody, err := io.ReadAll(body)
	if err != nil {
		slog.Error("Failed to parse io.ReadAll", "error", err.Error())
		return nil
	}
	return parsedBody
}

func ParseJsonBody(body []byte, v any) {
	err := json.Unmarshal(body, v)
	if err != nil {
		slog.Error("Failed to unmarshal body", "error", err.Error())
	}
}

func NewGetRequestWithCookies(url string, s []*http.Cookie) *http.Response {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		slog.Error("Failed to create request", "error", err.Error())
		return nil
	}
	for _, cookie := range s {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to make GET request", "error", err.Error())
		return nil
	}
	return resp
}
