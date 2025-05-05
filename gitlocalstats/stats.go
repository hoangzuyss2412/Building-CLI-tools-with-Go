package main

import (
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const daysInLastSixMonths = 183

// An "infinite" number to represent any date that is too far in the past
// Here, older is than 6 months is considered as too far in the past
const outOfRange = 99999

func stats(email string) {
	commits := processRepositories(email)
	printCommitStats(commits)
}

// processRepositories returns a map of day-of-year -> number of ,
// representing commits made in the last 6 months of the user with their given email
func processRepositories(email string) map[int]int {
	daysInMap := daysInLastSixMonths
	commits := make(map[int]int)
	dotFilePath := getDotFilePath()
	repos := loadRepoListFile(dotFilePath)

	for i := daysInMap; i >= 0; i-- {
		commits[i] = 0
	}

	for _, repo := range repos {
		commits = fillCommits(email, repo, commits)
	}

	return commits
}

// populate the commits map based on the given repo and user email
func fillCommits(email string, repoPath string, commits map[int]int) map[int]int {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	// Get all commits history starting from head
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	err = iterator.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}

		daysAgo := countDaysSinceDate(c.Author.When)

		// if day is not too old
		if daysAgo != outOfRange {
			commits[daysAgo]++
		}

		return nil

	})

	if err != nil {
		panic(err)
	}

	return commits
}

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	loc := t.Location()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}
func countDaysSinceDate(date time.Time) int {
	today := getBeginningOfDay(time.Now())
	start := getBeginningOfDay(date)

	duration := today.Sub(start)
	days := int(duration.Hours() / 24)

	if days > daysInLastSixMonths {
		return outOfRange
	}

	return days
}

func printCommitStats(commits map[int]int) {

}
