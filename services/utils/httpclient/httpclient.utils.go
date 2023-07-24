package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	GetMethod  = "GET"
	PostMethod = "POST"
)

func Request(
	url string,
	method string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	client := &http.Client{}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}
