package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\t", "    ")
		ret = append(ret, line)
	}

	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, line := range lines {
		len := utf8.RuneCountInString(line)
		if len > w {
			w = len
		}
	}

	return w
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | cowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
		lines = append(lines, string(line))
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`
	lines = tabsToSpaces(lines)

}
