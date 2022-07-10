package main

import (
	"fmt"
	"log"
	"math"
)

type Square struct {
	X 		float64
	Y 		float64
	Length 	float64
}

func NewSquare(x float64, y float64, length float64) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("Length must not be negative")
	}

	square := Square{
		X: x,
		Y: y,
		Length: length,
	}
	return &square, nil
}

func (square *Square) Move(dx float64, dy float64) {
	square.X += dx
	square.Y += dy
}

func (square *Square) Area() (float64) {
	return math.Pow(square.Length, 2)
}

func main() {
	square, err := NewSquare(1.0, 1.0, 10.0)
	if err != nil {
		log.Println(err)
	}

	square.Move(1.0, 1.0)
	log.Printf("X: %v", square.X)
	log.Printf("Y: %v", square.Y)

	area := square.Area()
	log.Printf("Area: %v", area)
}
