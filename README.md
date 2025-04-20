# github-bot

Repository that contains a Go Application for a github code review bot that listens to a GitHub Webhook configured to be triggered based on Pull Requests, then fetches the files alongwith its contents that were changed as part of the PR and sends the conent to a LLM in AWS Bedrock for review. The LLM Response or review comments are then added to the GitHub PR as a comment by the Go Application.
