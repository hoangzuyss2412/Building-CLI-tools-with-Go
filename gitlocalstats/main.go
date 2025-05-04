package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, rootFolder string) []string {
	folder := strings.TrimSuffix(rootFolder, "/")
	f, err := os.Open(folder)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	files, err := f.ReadDir(-1)

	for _, file := range files {
		if file.IsDir() {
			if file.Name() == "node_modules" || file.Name() == "vendor" {
				continue
			}

			path := folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				folders = append(folders, path)
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}

	return folders

}

func main() {
	var folder string

	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")

	if folder != "" {
		scan(folder)
	}
}
