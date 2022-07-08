package main

import (
	"fmt"
)

func main() {
	const FIZZ string = "fizz"
	const BUZZ string = "buzz"
	
	fizzBuzz(FIZZ, BUZZ)
}

// if divisible by 3, print fizz
// if divisible by 5, print buzz
// if divisible by 3 and 5, print fizz buzz
// else, print the number
func fizzBuzz(fizz string, buzz string) {
	for i := 1; i < 21; i++ {
		if i%3 == 0 {
			if i %5 == 0 {
				fmt.Println(fizz + " " + buzz)
			} else {
				fmt.Println(fizz)
			}
		} else if i%5 == 0 {
			fmt.Println(buzz)
		} else {
			fmt.Println(i)
		}
	}
}