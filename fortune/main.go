package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var files []string

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if strings.Contains(path, "/off/") {
		return nil
	}

	if filepath.Ext(path) == ".dat" {
		return nil
	}

	if strings.HasSuffix(path, ".u8") {
		return nil
	}

	if f.IsDir() {
		return nil
	}

	files = append(files, path)
	return nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var fortuneCommand *exec.Cmd

	if runtime.GOOS == "windows" {
		// Run "fortune -f" via WSL if we're on Windows
		fortuneCommand = exec.Command("wsl", "fortune", "-f")
	} else {
		// Assume that we're only using Linux/MacOS
		fortuneCommand = exec.Command("fortune", "-f")
	}

	pipe, err := fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err := fortuneCommand.Start(); err != nil {
		panic(err)
	}

	outputStream := bufio.NewScanner(pipe)
	outputStream.Scan()
	line := outputStream.Text()
	root := line[strings.Index(line, "/"):]

	// Convert Linux path to Windows path if running on Window machine
	if runtime.GOOS == "windows" {
		translate := exec.Command("wsl", "wslpath", "-w", root)
		translatedBytes, translateErr := translate.Output()
		if translateErr != nil {
			panic(translateErr)
		}
		windowsPath := strings.TrimSpace(string(translatedBytes))
		err = filepath.Walk(windowsPath, visit)
	} else {
		err = filepath.Walk(root, visit)
	}

	if err != nil {
		panic(err)
	}

	i := randomInt(1, len(files))
	randomFile := files[i]

	file, err := os.Open(randomFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(b)
	quotesSlice := strings.Split(quotes, "%")
	j := randomInt(1, len(quotesSlice))
	fmt.Println(quotesSlice[j])

}
