package main

import (
	"fmt"
	"os/exec"
)

func main() {
	a, _ := execTh()
	fmt.Print(a)
}

func execTh() (string, error) {
	cmd := exec.Command("git", "pull")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return "", err
	}

	return string(output), nil
}
