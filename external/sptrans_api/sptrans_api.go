package sptransapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"julianopedraca/clean_architecture_go/interfaces"
	"log/slog"
	"net/http"
	"os"
)

type ResponseWrapper struct {
	Response []interfaces.BusLine
}

type SptransApi struct {
	cookies []*http.Cookie
}

func (s *SptransApi) Authentication() ([]byte, error) {
	postBody, _ := json.Marshal(map[string]string{"": ""})
	body := bytes.NewReader(postBody)
	url := fmt.Sprintf("%s/Login/Autenticar?token=%s", os.Getenv("API_URL"), os.Getenv("API_KEY"))

	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		slog.Error("Failed to make POST request", "error", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	s.cookies = resp.Cookies()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to parse io.ReadAll", "error", err.Error())
		return nil, err
	}
	return respBody, nil
}

func (s *SptransApi) SearchLine(line string) ([]interfaces.BusLine, error) {
	url := fmt.Sprintf("%s/Linha/Buscar?termosBusca=%s", os.Getenv("API_URL"), line)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Failed to create request", "error", err.Error())
		return nil, err
	}

	for _, cookie := range s.cookies {
		req.AddCookie(cookie)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to make GET request", "error", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to parse io.ReadAll", "error", err.Error())
		return nil, err
	}

	var busLines []interfaces.BusLine
	err = json.Unmarshal(respBody, &busLines)
	if err != nil {
		slog.Error("Failed to unmarshal respBody", "error", err.Error())
		return nil, err
	}
	return busLines, nil
}
