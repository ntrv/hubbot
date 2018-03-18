package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func HelloHandler(ctx context.Context, req events.APIGatewayProxyRequest) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}

func main() {
	lambda.Start(HelloHandler)
}
