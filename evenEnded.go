package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	num, length := readInput()
	processNum(num, length)
}

func readInput() (int, int) {
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.ReplaceAll(input, "\n", ""))
	length := len(input) - 2

	return num, length
}

func processNum(num int, length int) {
	firstDigit := num/(int(math.Pow10(length)))
	secondDigit := num%10

	if firstDigit == secondDigit {
		fmt.Println("Even-Ended Number")
	} else {
		fmt.Println("Not an Even-Ended Number")
	}
}