package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func genK8s(serverURL string, resource string, inputFilePath string) error {
	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Errorf("error reading input data from file: %v", err)
		log.Panic(err)
	}
	contentType := determineContentType(inputFilePath)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	body := bytes.NewBuffer(inputData)
	resp, err := client.Post(serverURL, contentType, body)
	if err != nil {
		log.Errorf("error sending POST request: %v", err)
		return fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error reading response body: %v", err)
		return fmt.Errorf("%v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Errorf("server returned a non-OK status code: %d. Response Body: %s", resp.StatusCode, responseBody)
		return fmt.Errorf("%v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Errorf("error parsing server response: %v", err)
		return fmt.Errorf("%v", err)
	}

	return nil
}
