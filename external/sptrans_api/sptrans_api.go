package sptransapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"julianopedraca/clean_architecture_go/interfaces"
	utils "julianopedraca/clean_architecture_go/utils"
	"log/slog"
	"net/http"
	"os"
)

type ResponseWrapper struct {
	Response []interfaces.SearchLine
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

	respBody := utils.GetBodyFromHttpResponse(resp.Body)
	return respBody, nil
}

func (s *SptransApi) SearchLine(line string) ([]interfaces.SearchLine, error) {
	url := fmt.Sprintf("%s/Linha/Buscar?termosBusca=%s", os.Getenv("API_URL"), line)
	resp := utils.NewGetRequestWithCookies(url, s.cookies)
	defer resp.Body.Close()
	respBody := utils.GetBodyFromHttpResponse(resp.Body)

	var response []interfaces.SearchLine
	utils.ParseJsonBody(respBody, &response)
	return response, nil
}

func (s *SptransApi) SearchLineDirection(line string, direction string) ([]interfaces.SearchLine, error) {
	url := fmt.Sprintf("%s/Linha/BuscarLinhaSentido?termosBusca=%s&sentido=%s", os.Getenv("API_URL"), line, direction)
	resp := utils.NewGetRequestWithCookies(url, s.cookies)
	defer resp.Body.Close()
	respBody := utils.GetBodyFromHttpResponse(resp.Body)

	var response []interfaces.SearchLine
	utils.ParseJsonBody(respBody, &response)
	return response, nil
}

func (s *SptransApi) SearchStops(stop string) ([]interfaces.SearchStop, error) {
	url := fmt.Sprintf("%s/Parada/Buscar?termosBusca=%s", os.Getenv("API_URL"), stop)
	resp := utils.NewGetRequestWithCookies(url, s.cookies)
	defer resp.Body.Close()
	respBody := utils.GetBodyFromHttpResponse(resp.Body)

	var response []interfaces.SearchStop
	utils.ParseJsonBody(respBody, &response)
	return response, nil
}

func (s *SptransApi) SearchStopsByLine(line string) ([]interfaces.SearchStop, error) {
	url := fmt.Sprintf("%s/Parada/BuscarParadasPorLinha?codigoLinha=%s", os.Getenv("API_URL"), line)

	resp := utils.NewGetRequestWithCookies(url, s.cookies)
	defer resp.Body.Close()
	respBody := utils.GetBodyFromHttpResponse(resp.Body)

	var response []interfaces.SearchStop
	utils.ParseJsonBody(respBody, &response)
	return response, nil
}
