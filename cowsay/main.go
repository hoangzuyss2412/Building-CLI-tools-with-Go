package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | cowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}

}
