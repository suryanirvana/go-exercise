package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"regexp"
)

func main() {
	text := readInput()
	textMap := processText(text)
	for key, value := range textMap {
		fmt.Println(key, value)
	}
}

func readInput() (string) {
	fmt.Print("Enter a text: ")
	reader := bufio.NewReader(os.Stdin)
	
	input, _ := reader.ReadString('\n')
	
	re, _ := regexp.Compile(`[^\w]`)
	return re.ReplaceAllString(strings.ToLower(strings.ReplaceAll(input, "\n", "")), " ")
}

func processText(text string) (map[string]int) {
	textSlices := strings.Fields(text)
	textMap := map[string]int{}
	for _, slice := range textSlices {
		textMap[slice]++
	}

	return textMap
}
