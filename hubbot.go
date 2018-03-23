package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.New(&github.Config{Secret: "hogehoge"})
	hook.RegisterEvents(HandlePullRequest, github.PullRequestEvent)
	hook.RegisterEvents(HandlePush, github.PushEvent)
	log.Fatal(ListenAndServe(":80", webhooks.Handler(hook)))
}

func ListenAndServe(addr string, h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		r, err := gateway.NewRequest(ctx, req)
		if err != nil {
			fmt.Printf("%v\n", req)
			return events.APIGatewayProxyResponse{}, err
		}
		w := gateway.NewResponse()
		h.ServeHTTP(w, r)
		resp := w.End()

		if resp.StatusCode == 0 {
			return events.APIGatewayProxyResponse{
				StatusCode:      http.StatusOK,
				Body:            resp.Body,
				Headers:         resp.Headers,
				IsBase64Encoded: resp.IsBase64Encoded,
			}, nil
		}
		return resp, nil
	})

	return nil
}

func HandlePullRequest(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PullRequestPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
}

func HandlePush(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PushPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
}
