package main

import (
	"bufio"
	"fmt"
	"runtime"
	"math/rand"
	"os"
	"os/exec" 
	"path/filepath"
	"strings"
	"time"
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

	if f.IsDir() { 
		return nil 
	}

	files = append(files, path)
	return nil
}

func randomInt(min, max int) int { 
	return min + rand.Intn(max - min)
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
		translatedBytes, err := translate.Output()
		if err != nil {
			panic(err)
		}
		windowsPath := strings.TrimSpace(string(translatedBytes))
		err = filepath.Walk(windowsPath, visit)
	} else { 
		err = filepath.Walk(root, visit)
	}
	
	if err != nil { 
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := randomInt(1, len(files))
	randomFile := files[i]
	fmt.Println(randomFile)


}	