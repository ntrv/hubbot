package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ntrv/hubbot/chatwork"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.New(&github.Config{Secret: "hogehoge"})
	hook.RegisterEvents(chatwork.HandlePullRequest, github.PullRequestEvent)
	hook.RegisterEvents(chatwork.HandlePush, github.PushEvent)
	lambda.Start(handleHubbot(webhooks.Handler(hook)))
}
