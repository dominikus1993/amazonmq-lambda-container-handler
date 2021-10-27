package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event = events.ActiveMQEvent

func handler(event Event) error {
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("Received event: %s", eventJson)
	return nil
}

func main() {
	lambda.Start(handler)
}
