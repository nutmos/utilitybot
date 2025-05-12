package main

import (
	"bufio"
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nutmos/utilitybot/handler"
)

func main() {
	bot := handler.Bot
	botData, _ := bot.GetMe()
	log.Printf("Bot %s", botData.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	updates := bot.GetUpdatesChan(u)
	go receiveUpdates(ctx, updates)
	log.Println("Start listening for updates. Press enter to stop")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handler.HandleMessage(update.Message)
		break

		// Handle button clicks
		/*
			case update.CallbackQuery != nil:
				handleButton(update.CallbackQuery)
				break
		*/
	}
}
