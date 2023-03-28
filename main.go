package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

const (
	owner           = "Rlyehan"
	repo            = "api_polling_service"
	desiredLabel    = "test-label"
	pollingInterval = 30 * time.Second
)

func main() {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("Please set GITHUB_ACCESS_TOKEN environment variable")
		return
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for {
		err := checkAndCloseIssues(ctx, client)
		if err != nil {
			fmt.Printf("Error checking and closing issues: %v\n", err)
		}
		time.Sleep(pollingInterval)
	}
}

func checkAndCloseIssues(ctx context.Context, client *github.Client) error {
	issues, _, err := client.Issues.ListByRepo(ctx, owner, repo, &github.IssueListByRepoOptions{
		State: "open",
	})

	if err != nil {
		return err
	}

	for _, issue := range issues {
		if hasDesiredLabel(issue, desiredLabel) {
			fmt.Printf("Closing issue #%d with label '%s'\n", *issue.Number, desiredLabel)
			err := closeIssue(ctx, client, *issue.Number)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func hasDesiredLabel(issue *github.Issue, desiredLabel string) bool {
	for _, label := range issue.Labels {
		if *label.Name == desiredLabel {
			return true
		}
	}

	return false
}

func closeIssue(ctx context.Context, client *github.Client, issueNumber int) error {
	state := "closed"
	_, _, err := client.Issues.Edit(ctx, owner, repo, issueNumber, &github.IssueRequest{
		State: &state,
	})

	return err
}
