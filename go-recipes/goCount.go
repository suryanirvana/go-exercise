package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
)

var regex = regexp.MustCompile(`go ([a-z]+)`)

func count(fileName string, term1 string, term2 string) (map[string]int, error) {
	file, err := open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	matches := make(map[string]int)
	for scanner.Scan() {
		match := regex.FindStringSubmatch(scanner.Text())
		if len(match) > 0 {
			splittedText := strings.Split(scanner.Text(), ";")[1]
			command := strings.Split(splittedText, " ")
			combinedCommand := fmt.Sprintf("%s %s", command[0], command[1])
			matches[combinedCommand]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

func open(fileName string) (*os.File, error){
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func main() {
	frequencies, err := count("zsh_history", "go ", ".go")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(frequencies)
	}
}
