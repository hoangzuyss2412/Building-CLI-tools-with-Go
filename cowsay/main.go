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

func normalizeStringLength(lines []string, maxWidth int) []string {
	var ret []string
	for _, line := range lines {
		line = line + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(line))
		ret = append(ret, line)
	}

	return ret
}

func buildBalloon(lines []string, maxWidth int) string {
	var ret []string

	top := strings.Repeat("-", maxWidth+2)
	bottom := strings.Repeat("-", maxWidth+2)
	count := len(lines)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("< %-*s >", maxWidth, lines[0])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf("/ %-*s \\", maxWidth, lines[0])
		ret = append(ret, s)

		for i := 1; i < count-1; i++ {
			s = fmt.Sprintf("| %-*s |", maxWidth, lines[i])
			ret = append(ret, s)
		}

		s = fmt.Sprintf("\\ %-*s /", maxWidth, lines[count-1])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")

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
	maxWidth := calculateMaxWidth(lines)
	lines = normalizeStringLength(lines, maxWidth)
	balloon := buildBalloon(lines, maxWidth)

	fmt.Println(balloon)
	fmt.Println(cow)

}
