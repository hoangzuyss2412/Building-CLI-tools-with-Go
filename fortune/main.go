package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run "fortune -f" via WSL because we're on Windows
	out, err := exec.Command("wsl", "fortune", "-f").CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}