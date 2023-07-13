# Validate Commit Message

This tool is used to verify an issue is in a commit message, and that the issue has been assigned and closed

## Requirements

Set the following ENV Variables
Jira Credentials
export JIRA_USER=""
export JIRA_TOKEN=""
export JIRA_URL=""

Validation Elements
export JIRA_PREFIX="MYJIRA-"
export JIRA_ASSIGNEE="Adam Henderson"

## Running the Tool

Install the project
`go install validate-commit-message`

Run the project
`/go/bin/validate-commit-message "MYJIRA-1:This is some text"`

SDK used for JIRA Integration
https://pkg.go.dev/github.com/andygrunwald/go-jira#readme-features