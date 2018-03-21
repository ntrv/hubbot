package main

import (
	"fmt"
	"log"

	"github.com/apex/gateway"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.New(&github.Config{Secret: "hogehoge"})
	hook.RegisterEvents(HandlePullRequest, github.PullRequestEvent)
	hook.RegisterEvents(HandlePush, github.PushEvent)
	log.Fatal(gateway.ListenAndServe(":80", webhooks.Handler(hook)))
}

func HandlePullRequest(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PullRequestPayload)
	fmt.Printf("%+v", pl)
}

func HandlePush(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PushPayload)
	fmt.Printf("%+v", pl)
}
