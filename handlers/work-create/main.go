package main

import (
	"github.com/Lwchg/gofaas"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(gofaas.NotifyAPIGateway(gofaas.WorkCreate))
}
