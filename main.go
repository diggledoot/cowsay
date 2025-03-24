package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Get file information about standard input
	info, _ := os.Stdin.Stat()

	// Check if input is from a terminal (character device) using a bitwise operation
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command  is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	// Create new buffered reader
	reader := bufio.NewReader(os.Stdin)

	// Initialise empty rune slice
	var output []rune

	// Infinite loop
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	// Print each unicode character
	// source: https://pkg.go.dev/fmt
	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
