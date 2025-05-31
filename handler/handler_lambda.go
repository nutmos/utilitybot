package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleLambdaRequest(ctx context.Context, event json.RawMessage) (events.APIGatewayProxyResponse, error) {
	log.Printf("%s\n", string(event))
	// Parse Event
	var e events.LambdaFunctionURLRequest
	if err := json.Unmarshal(event, &e); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Malform Request",
			StatusCode: 400,
		}, err
	}

	// Parse Telegram Message
	log.Printf("%s\n", e.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal([]byte(e.Body), &update); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Malform Messaage",
			StatusCode: 400,
		}, err
	}
	// Send to Telegram Message Handler
	HandleMessage(update.Message)
	return events.APIGatewayProxyResponse{
		Body:       "Message Received",
		StatusCode: 200,
	}, nil
}
