package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func executeGitPull(repositoryDir string) (string, error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = repositoryDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v\nCommand output: %s", err, string(output))
	}

	return string(output), nil
}

func walkRepositories(rootDir string) error {
	err := filepath.WalkDir(rootDir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %v", err)
		}

		if info.IsDir() && info.Name() != rootDir {
			if info.Name() == ".git" {
				repoDir := filepath.Dir(path)
				fmt.Println("Executing 'git pull' in directory:", repoDir)

				output, err := executeGitPull(repoDir)
				if err != nil {
					return fmt.Errorf("failed to execute 'git pull' in directory %s: %v", repoDir, err)
				}
				fmt.Println(output)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk repositories: %v", err)
	}

	return nil
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %v", err))
	}

	if err := walkRepositories(rootDir); err != nil {
		fmt.Println("Error walking repositories:", err)
	}
}
