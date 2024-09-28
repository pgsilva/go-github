package config

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClientService struct {
	Client *http.Client
}

func NewHttpClientService() *HttpClientService {
	return &HttpClientService{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func GetHttpClient() *HttpClientService {
	return NewHttpClientService()
}

func (s *HttpClientService) Do(req *http.Request) ([]byte, error) {
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a solicitação: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %v", err)
	}

	return respBody, nil
}
