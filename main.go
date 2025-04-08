package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubPRPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	URL    string `json:"url"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

func webhookHandleFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a webhook request")
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request Body", http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	var payload GitHubPRPayload

	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(writer, "Failed to parse JSON encoded data", http.StatusBadRequest)
		return
	}

	if payload.Action == "opened" || payload.Action == "reopened" {
		fmt.Printf("PR %s is ready for review", payload.PullRequest.Title)
	}

}

func main() {
	http.HandleFunc("/webhook", webhookHandleFunc)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening at Port 8080 for GitHub PR events")
}
