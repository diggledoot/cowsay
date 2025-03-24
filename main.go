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

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
