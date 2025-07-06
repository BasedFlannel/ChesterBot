package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func restGet(url string) string {
	resp, err := http.Get(url)
	errorCheck(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panicf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	errorCheck(err)

	return string(data)
}

func restPut() {
	// Replace with the API ID provided by crudcrud.com
	apiID := "XXXXXX"

	// HTTP endpoint with path of resource (vehicles)
	createResourceURL := fmt.Sprintf("https://crudcrud.com/api/%s/vehicles",
		apiID)

	// JSON payload containing resource information
	requestBody, _ := json.Marshal(map[string]string{
		"color":     "White",
		"license":   "ABC123",
		"numWheels": "4",
		"type":      "Car",
	})

	// Converting request body into bytes buffer
	requestBodyBytes := bytes.NewBuffer(requestBody)

	// Creating a POST request on createResourceURL
	request, err := http.NewRequest("POST", createResourceURL, requestBodyBytes)
	if err != nil {
		panic(err)
	}

	// Adding application/json as payload type
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	// Performing POST request
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Reading the response to the request
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error received in response:", err)
	}

	responseBodyString := string(responseBody)
	fmt.Println(responseBodyString)
}
