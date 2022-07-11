package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Record struct {
	ID			string		`json:"id"`
	Distance	float64		`json:"distance"`
	Start		string		`json:"start"`
	End			string		`json:"end"`
}

type Records struct {
	Records		[]Record
}

func (records *Records) Add(record Record) {
	records.Records = append(records.Records, record)
}

func open(file io.Reader) (*Records, error){
	decoder := json.NewDecoder(file)

	var records Records
	for {
		var record Record

		err := decoder.Decode(&record)
		if err != nil {
			return &records, err
		}
		if err == io.EOF {
			break
		}

		records.Add(record)
	}

	return &records, nil
}

func (records Records) findMaxSpeed() (*Record, error){
	var maxSpeed float64
	var maxSpeedRecord Record
	for _, record := range records.Records {
		speed, err := calculateSpeed(record)
		if err != nil {
			return nil, err
		}
		if speed > maxSpeed {
			maxSpeed = speed
			maxSpeedRecord = record
		}
	}

	log.Printf("Maximum Speed: %.02f", maxSpeed)

	return &maxSpeedRecord, nil
}

func calculateSpeed(record Record) (float64, error) {
	distance := record.Distance

	startTime, err := time.Parse("2006-01-02T15:04", record.Start)
	if err != nil {
		return 0.0, err
	}

	endTime, err := time.Parse("2006-01-02T15:04", record.End)
	if err != nil {
		return 0.0, err
	}

	timeElapsed := endTime.Sub(startTime).Hours()

	speed := distance/timeElapsed
	return speed, nil
}

func main() {
	file, err := os.Open("record.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	records, _ := open(file)
	
	maxSpeedRecord, err := records.findMaxSpeed()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(maxSpeedRecord)
	}
}
