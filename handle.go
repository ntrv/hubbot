package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
)

type lambdaHandleFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func handleHubbot(h http.Handler) lambdaHandleFunc {
	if h == nil {
		h = http.DefaultServeMux
	}
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	}
}
