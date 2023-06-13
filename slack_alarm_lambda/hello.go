package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, body any) (string, error) {
	return fmt.Sprintf("Hello %s!", body), nil
}

func main() {
	lambda.Start(HandleRequest)
}
