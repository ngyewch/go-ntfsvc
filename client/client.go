package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NotificationService struct {
	url    string
	apiKey string
}

type SendNotificationRequest struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func NewNotificationService(url string, apiKey string) *NotificationService {
	return &NotificationService{
		url:    url,
		apiKey: apiKey,
	}
}

func (service *NotificationService) SendNotification(topic string, message string) error {
	requestBody := SendNotificationRequest{
		Topic:   topic,
		Message: message,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", service.url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("X-API-KEY", service.apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("HTTP %d", resp.StatusCode))
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
