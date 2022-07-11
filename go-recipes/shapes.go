package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"math"
)

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

type Device interface {
	Set(int, int, color.Color)
}

type Drawer interface {
	Draw(device Device)
}

type Shape struct {
	X 			float64
	Y			float64
	Color		color.Color
}

type Circle struct {
	Shape		Shape
	Radius		float64
}

func NewCircle(x float64, y float64, radius float64, color color.Color) (*Circle) {
	if radius <= 0 {
		return nil
	}

	circle := Circle{
		Shape: Shape{
			X: x,
			Y: y,
			Color: color,
		},
		Radius: radius,
	}
	return &circle
}

func (circle *Circle) Draw(device Device) {
	minX, minY := circle.Shape.X-circle.Radius, circle.Shape.Y-circle.Radius
	maxX, maxY := circle.Shape.X+circle.Radius, circle.Shape.Y+circle.Radius

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dx, dy := x-circle.Shape.X, y-circle.Shape.Y
			if math.Sqrt(float64(dx*dx+dy*dy)) <= circle.Radius {
				device.Set(int(x), int(y), circle.Shape.Color)
			}
		}
	}
}

type Rectangle struct {
	Shape		Shape
	Length		float64
	Width		float64
}

func NewRectangle(x float64, y float64, length float64, width float64, color color.Color) (*Rectangle) {
	if length <= 0 || width <= 0 {
		return nil
	}

	rectangle := Rectangle{
		Shape: Shape{
			X: x,
			Y: y,
			Color: color,
		},
		Length: length,
		Width: width,
	}
	return &rectangle
}

func (rectangle *Rectangle) Draw(device Device) {
	minX, minY := rectangle.Shape.X-rectangle.Width, rectangle.Shape.Y-rectangle.Length
	maxX, maxY := rectangle.Shape.X+rectangle.Width, rectangle.Shape.Y+rectangle.Length

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			device.Set(int(x), int(y), rectangle.Shape.Color)
		}
	}
}

type ImageCanvas struct {
	Length		float64
	Width		float64
	Shapes		[]Drawer
}

func NewImageCanvas(length float64, width float64) (*ImageCanvas, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length can't be zero")
	}
	
	if width <= 0 {
		return nil, fmt.Errorf("Width can't be zero")
	}

	imageCanvas := ImageCanvas{
		Length: length,
		Width: width,
	}
	return &imageCanvas, nil
}

func (imageCanvas *ImageCanvas) Add(drawer Drawer) {
	imageCanvas.Shapes = append(imageCanvas.Shapes, drawer)
}

func (imageCanvas *ImageCanvas) Draw(writer io.Writer) error {
	image := image.NewRGBA(image.Rect(0, 0, int(imageCanvas.Length), int(imageCanvas.Width)))
	for _, shape := range imageCanvas.Shapes {
		shape.Draw(image)
	}
	return png.Encode(writer, image)
}

func main() {
	imageCanvas, err := NewImageCanvas(200, 200)
	if err != nil {
		log.Fatal(err)
	}

	imageCanvas.Add(NewCircle(100, 100, 80, Green))
	imageCanvas.Add(NewCircle(60, 60, 10, Blue))
	imageCanvas.Add(NewCircle(140, 60, 10, Blue))
	imageCanvas.Add(NewRectangle(100, 130, 10, 50, Red))
	file, err := os.Create("face.png")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err := imageCanvas.Draw(file); err != nil {
		log.Fatal(err)
	}
}
