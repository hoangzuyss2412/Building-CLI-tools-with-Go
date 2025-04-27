package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("fortune", "-f").CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}