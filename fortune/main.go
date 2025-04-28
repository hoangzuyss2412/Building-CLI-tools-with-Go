package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// Run "fortune -f" via WSL because we're on Windows
	fortuneCommand = exec.Command("wsl", "fortune", "-f")

	pipe, err = fortuneCommand.StderrPipe()
	if err != nil {
		panic(err)
	}

	fortuneCommand.Start()

	outputStream := bufio.NewScanner()
	outputStream.Scan()
	fmt.Println(outputStream.Text())

	fmt.Println(string(out))
}