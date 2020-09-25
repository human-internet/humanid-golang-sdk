package helper

import (
	"encoding/json"
	"io"
	"net/http"
	"bytes"

	"github.com/asaskevich/govalidator"
)

// DecodeAndValidate makes the response with payload as json format
func DecodeAndValidate(body io.Reader, payload interface{}) (error) {
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return err
	}

	// Validate request
	result, err := govalidator.ValidateStruct(payload)
	if !result {
		return err
	}

	return nil
}

func PostNoBody(url string, respPayload interface{}, serverClientID string, serverClientSecret string) (error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client-id", serverClientID)
	req.Header.Set("client-secret", serverClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = DecodeAndValidate(resp.Body, respPayload)
	if err != nil {
		return err
	}
	return nil
}

func Post(url string, payload interface{}, respPayload interface{}, serverClientID string, serverClientSecret string) error {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client-id", serverClientID)
	req.Header.Set("client-secret", serverClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = DecodeAndValidate(resp.Body, respPayload)
	if err != nil {
		return err
	}
	return nil
}