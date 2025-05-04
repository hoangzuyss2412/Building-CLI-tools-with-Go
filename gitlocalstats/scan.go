package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
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

func loadRepoListFile(dotfilePath string) []string {
	f, err := os.Open(dotfilePath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines

}

func updateRepoListFile(dotfilePath string, newRepos []string) {
	existingRepos := loadRepoListFile(dotfilePath)
	repos := mergeUniqueRepos(existingRepos, newRepos)
	saveRepoList(repos, dotfilePath)
}

func mergeUniqueRepos(existing []string, newRepos []string) []string {
	seen := make(map[string]bool)

	for _, repo := range existing {
		seen[repo] = true
	}

	for _, repo := range newRepos {
		if seen[repo] {
			continue
		}
		existing = append(existing, repo)
		seen[repo] = true
	}

	return existing
}

// overwrite the existing content of the dot file, e.g .gitlocalstats
func saveRepoList(repos []string, dotfilePath string) {
	content := strings.Join(repos, "\n")
	err := os.WriteFile(dotfilePath, []byte(content), 0644)
	if err != nil {
		log.Fatal("Failed to write repo list:", err)
	}
}

func scan(folder string) {
	fmt.Println("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	dotfilePath := getDotFilePath()
	updateRepoListFile(dotfilePath, repositories)
	fmt.Println("\n\nSuccessfully added\n\n")
}
