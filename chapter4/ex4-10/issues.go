package main

/*
Modify issues to report the results in age categories, say less than a month
old, less than a year old, and more than a year old.
*/

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

// time constants
// assume a month have 28 days
const (
	OneMonth = time.Hour * 24 * 28
	OneYear  = time.Hour * 24 * 28 * 12
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var month, year, more []*github.Issue
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		passed := time.Since(item.CreatedAt)

		switch {
		case passed < OneMonth:
			month = append(month, item)
		case passed < OneYear:
			year = append(year, item)
		case passed >= OneYear:
			more = append(more, item)
		}
	}

	fmt.Printf("\nIssues less than a month old\n")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nIssues less than a year old\n")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Printf("\nIssues more than a year old\n")
	for _, item := range more {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
