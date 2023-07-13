package main

import "os"

// Check Message text start with an issue (MYJIRA-123: )
func validCommitText(message string) bool {
	return false
}

// Check Issue state
func validIssueStatus(issueId string) bool {
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
