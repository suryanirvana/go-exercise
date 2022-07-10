package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	Writer io.Writer
}

func (capper *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	output := make([]byte, len(p))

	for index, char := range p {
		if char >= 'a' && char <= 'z' {
			char -= diff
		}
		output[index] = char
	}

	return capper.Writer.Write(output)
}

func main() {
	capper := &Capper{
		Writer: os.Stdout,
	}
	fmt.Fprintf(capper, "Hello There")
}
