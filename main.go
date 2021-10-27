package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event = events.ActiveMQEvent

func handler(ctx context.Context, event Event) error {
	log.Println("Received event ", event)
	return nil
}

func main() {
	lambda.Start(handler)
}
