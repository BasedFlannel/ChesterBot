package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

func restPutImg(imgUrl string) {
	// Replace with the API ID provided by imgchest
	apiID := "XXXXXX"

	// HTTP endpoint for uploading new posts
	createResourceURL := "https://api.imgchest.com/v1/post"

	// Declare these here so they can be used in the below two if blocks and
	// still carry over to the end of this function.
	var base64img string

	// pull image from URL and convert into a base64 encoding of bytes to send to imgchest
	if imgUrl != "" {

		resp, err := http.Get(imgUrl)
		if err != nil {
			fmt.Println("Error retrieving the file, ", err)
			return
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		img, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading the response, ", err)
			return
		}

		//contentType = http.DetectContentType(img)
		base64img = base64.StdEncoding.EncodeToString(img)
	}

	//title string based on date and random number
	titleString := "Chester Upload " + time.Now().Format("RFC3339Nano")

	// JSON payload containing resource information
	requestBody, _ := json.Marshal(map[string]string{
		"title":     titleString,
		"privacy":   "secret",
		"anonymous": "false",
		"nsfw":      "false",
		"images":    "[" + base64img + "]",
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
	//Adding API Key as a header
	request.Header.Add("Authorization", apiID)

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
