package main

import (
	"os"
	s "strings"

	jira "github.com/andygrunwald/go-jira"
)

// Check Message text start with an issue (MYJIRA-123: )
func validCommitText(message string) bool {
	jiraIssueID := s.Split(message, ":")[0]
	return s.Contains(jiraIssueID, os.Getenv("JIRA_PREFIX"))
}

// Check Issue state
// TODO: Verify if they really only want 0 and -1, what about error messages?
func validIssueStatus(message string) bool {
	// Get Issue ID
	jiraIssueID := s.Split(message, ":")[0]
	// Set creds for HTTP client
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USER"),
		Password: os.Getenv("JIRA_TOKEN"),
	}
	jiraClient, _ := jira.NewClient(tp.Client(), os.Getenv("JIRA_URL"))
	// Get Issue
	issue, _, err := jiraClient.Issue.Get(jiraIssueID, nil)
	if err != nil {
		// log.Fatalf("Error: %v", err)
		return false
	}
	// Check Assignee and Status
	if issue.Fields.Assignee.DisplayName == os.Getenv("JIRA_ASSIGNEE") && issue.Fields.Status.StatusCategory.Name == "Done" {
		return true
	}

	return false
}

func main() {
	commitMessage := os.Args[1]

	if validCommitText(commitMessage) && validIssueStatus(commitMessage) {
		println(0)
	} else {
		println(-1)
	}

}
