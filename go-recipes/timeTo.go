package main

import (
	"log"
	"time"
)

func convert(timestamp string, from string, to string) (string, error) {
	fromTimezone, err := time.LoadLocation(from)
	if err != nil {
		return "", err
	}

	originalTimestamp, err := time.ParseInLocation("2006-01-02T15:04", timestamp, fromTimezone)
	if err != nil {
		return "", err
	}
	
	toTimezone, err := time.LoadLocation(to)
	if err != nil {
		return "", nil
	}

	convertedTimestamp := originalTimestamp.In(toTimezone)
	return convertedTimestamp.String(), nil
}

func main() {
	timestamp := "2022-07-10T16:44"
	from := "Australia/Brisbane"
	to := "America/Los_Angeles"

	convertedTimestamp, err := convert(timestamp, from, to)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(convertedTimestamp)
	}
}