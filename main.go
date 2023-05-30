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
	// ign, _ := readIgnoreFile(".ign")

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != dirPath {
			// Check if the directory contains a ".git" directory
			if strings.HasSuffix(path, ".git") {
				parentDir := filepath.Dir(path)
				fmt.Println("Executing git pull in directory:", parentDir)

				// Execute git pull in the parent directory
				output, err := execGitPull(parentDir)
				if err != nil {
					return err
				}
				fmt.Println(output)
			}

			// Recursively walk through child directories
			if err := walkDirectories(path); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func main() {
	// ignFilePath := ".ign"
	// ignList, err := readIgnoreFile(ignFilePath)
	// if err != nil {
	// 	log.Println(err)
	// }

	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = walkDirectories(dirPath)
	if err != nil {
		fmt.Println("Error walking directories:", err)
	}
}
