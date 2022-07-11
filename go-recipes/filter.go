package main

import (
	"log"
)

type NumberSlices struct {
	nums	[]int
}

func filter(pred func(int) bool, values []int) []int {
	if len(values) == 0 {
		return nil
	}

	filterMap := make(map[string][]int)
	for _, value := range values {
		if pred(value) {
			currentValues := filterMap["Odd"]
			currentValues = append(currentValues, value)
			filterMap["Odd"] = currentValues
		} else {
			currentValues := filterMap["Even"]
			currentValues = append(currentValues, value)
			filterMap["Even"] = currentValues
		}
	}

	log.Println(filterMap)
	return filterMap["Odd"]
}

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	values := []int{1,2,3,4,5,6,7,8}
	filteredValues := filter(isOdd, values)
	log.Println(filteredValues)
}