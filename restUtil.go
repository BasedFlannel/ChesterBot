package main

import (
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
