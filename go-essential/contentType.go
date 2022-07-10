package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://www.linkedin.comm"
	contentType, err := retrieveContentType(url)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(contentType)
	}
}

func retrieveContentType(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	return contentType, err
}
