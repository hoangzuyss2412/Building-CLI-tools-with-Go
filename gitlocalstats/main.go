package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, rootFolder string) []string {
	folder := filepath.Clean(rootFolder)
	f, err := os.Open(folder)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	entries, err := f.ReadDir(-1)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			name := entry.Name()

			if name == "node_modules" || name == "vendor" {
				continue
			}

			path := filepath.Join(folder, name)
			if name == ".git" {
				repoPath := filepath.Dir(path) // parent of .git
				fmt.Println(repoPath)
				folders = append(folders, repoPath)
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(usr.HomeDir, ".gitlocalstats")
}

func scan(folder string) {
	print("scan")
}

func stats(email string) {
	print("stats")
}

func main() {
	var folder string

	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")

	if folder != "" {
		scan(folder)
	}

	stats(email)
}
