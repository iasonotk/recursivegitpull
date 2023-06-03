package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func execGitPull(workingDir string) (string, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = workingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		fmt.Println("Command output:", string(output))
		return "", err
	}

	return string(output), nil
}

func walkDirectories(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != dirPath {
			if strings.HasSuffix(path, ".git") {
				parentDir := filepath.Dir(path)
				fmt.Println("Executing git pull in directory:", parentDir)

				output, err := execGitPull(parentDir)
				if err != nil {
					return err
				}
				fmt.Println(output)
			}
		}

		return nil
	})

	return err
}

func main() {

	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = walkDirectories(dirPath)
	if err != nil {
		fmt.Println("Error walking directories:", err)
	}

}
