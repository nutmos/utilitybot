package handler

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Event struct {
	Message  *tgbotapi.Message `json:"message"`
	UpdateId int               `json:"update_id"`
}

func HandleLambdaRequest(ctx context.Context, event json.RawMessage) error {
	var e Event
	if err := json.Unmarshal(event, &e); err != nil {
		return err
	}
	HandleMessage(e.Message)
	return nil
}
