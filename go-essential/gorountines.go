package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.linkedin.com",
		"https://www.google.com",
		"https://www.github.com",
	}

	channel := make(chan string)

	for _, url := range urls {
		go retrieveContentType(url, channel)
	}

	for value := range channel {
		log.Printf("Receiving %s", value)
	}
}

func retrieveContentType(url string, channel chan string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprintf("%s -> error: %s", url, err)
		return "", err
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	channel <- fmt.Sprintf("%s -> %s", url, contentType)
	return contentType, err
}
