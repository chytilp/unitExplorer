package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Sender struct {
	url string
}

func NewSender(url string) *Sender {
	return &Sender{url: url}
}

func (s *Sender) sendRequest(request Request) ([]byte, error) {
	jsonBody, err := json.Marshal(request)
	fmt.Printf("json: %s\n", jsonBody)
	if err != nil {
		fmt.Printf("marshalling to json error: %s\n", err)
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", s.url, bodyReader)
	if err != nil {
		fmt.Printf("create request error: %s\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}
	fmt.Printf("Response status code: %d, reason: %s\n", res.StatusCode, res.Status)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	return resBody, nil

}

func (s *Sender) GetDomains(request Request) (*DomainPayload, error) {
	data, err := s.sendRequest(request)
	if err != nil {
		return nil, err
	}
	payload := DomainPayload{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (s *Sender) GetEvents(request Request) (*EventPayload, error) {
	data, err := s.sendRequest(request)
	if err != nil {
		return nil, err
	}
	payload := EventPayload{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (s *Sender) GetMarkets(request Request) (*MarketPayload, error) {
	data, err := s.sendRequest(request)
	if err != nil {
		return nil, err
	}
	payload := MarketPayload{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
