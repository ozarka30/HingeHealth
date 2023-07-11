# Hinge Health SRE Code Challenge

Welcome to the Hinge Health SRE Code Challenge, thank you for taking the time to participate.

## Context

In addition to maintainable code we are looking for good coding practices like small commits, sensible comments, clearly written documentation, and ultimately a pull request(s) to publish your work. You can assume the following things:
  * Making a single pull request is fine, but you may use separate pull requests for each part if you prefer
  * Any code you produce would run on a fully-equipped, headless Linux system with all required tooling installed
  * Proper credentials for remote systems are passed as environment variables

There are two parts to the exercise. You should not spend more than 1 hour coding/debugging on each part, but if you do not finish in time please add a `TODO` section in your `README.md` detailing what you would do with more time. Any external setup (a remote system, setting up your IDE, etc.) should not count towards your 1 hour per part.

Examples of major IaaS cloud providers: AWS, Google Cloud, Azure, etc.

Examples of major issue tracking systems: GitHub Issues, Gitlab Issues, JIRA, Pivotal Tracker, Bugzilla, etc. (anything with a publicly documented API to fetch issue details)

## The Challenge, Part 1

Engineers at your company need to be able to frequently and easily provision cloud storage buckets for both production and non-production environments. The team is growing fast and there are not enough infrastructure engineers to handle all requests. You have been tasked with creating a self-service tool engineers can use to provision or destroy their own buckets without direct involvement from the infrastructure team.

Using a toolset and major IaaS cloud provider of your choice, create a tool(s) in the `./part1-cloud-provisioning` sub-directory that your engineers can run either by hand or as part of a CI/CD pipeline. There should be a `README.md` documenting steps engineers need to take to provision or destroy their bucket(s).

## The Challenge, Part 2

You have been asked to create a tool that takes as input a version control commit message and validates that the leading text contains an issue identifier that a) is assigned to a specific person and b) is in a specific state (e.g. `done` or `closed`). Your tool should read standard input to read the commit message. If the message is valid your tool should exit with a `0` exit code, non-zero otherwise. Any error messages should go to the standard error stream.

Using a language/toolset and major issue tracking system of your choice, create a tool in the `./part2-commit-msg-validation` sub-directory called `validate-commit-message` that performs the validation. There should be a `README.md` documenting what issue tracking system is being used, who the expected assignee is (which you can make up or use your own ID), and any other relevant details.

### Good commit message example

    MYJIRA-123: this is an example commit message

If the issue `MYJIRA-123` a) exists, b) is assigned to a specific user and c) is in a `done` state then exit `0`, otherwise `1`.

### Bad commit message example

    This is an example commit message for MYJIRA-123

Invalid by default as the issue ID is not leading the commit message, exit `1`.
