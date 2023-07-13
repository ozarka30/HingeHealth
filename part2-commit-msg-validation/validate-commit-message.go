package main

import (
	"os"
	s "strings"
)

// Check Message text start with an issue (MYJIRA-123: )
func validCommitText(message string) bool {
	jiraIssueID := s.Split(message, ":")[0]
	return s.Contains(jiraIssueID, "MYJIRA-")
}

// Check Issue state
func validIssueStatus(message string) bool {
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
