package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ntrv/hubbot/chatwork"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.New(&github.Config{Secret: "hogehoge"})
	cw := chatwork.New(&chatwork.Config{ApiKey: "hogehoge"})
	hook.RegisterEvents(cw.HandlePullRequest, github.PullRequestEvent)
	hook.RegisterEvents(cw.HandlePush, github.PushEvent)
	lambda.Start(handleHubbot(webhooks.Handler(hook)))
}
