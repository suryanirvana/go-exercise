package main

import (
	"fmt"
)

func main() {
	nums := []int{321,541,332,434,543,462,413,753,34,546,234,645}
	max := findMax(nums)
	fmt.Println("Maximum Number:", max)
}

func findMax(nums []int) (int) {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
