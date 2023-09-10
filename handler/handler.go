package handler

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nutmos/utilitybot/config"
	"github.com/nutmos/utilitybot/flightcaller"
)

var (
	Bot *tgbotapi.BotAPI
)

func init() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(config.Config.ApiKey.Telegram)
	if err != nil {
		panic(err)
	}
	Bot.Debug = false
}

func HandleMessage(message *tgbotapi.Message) {
	if !message.IsCommand() {
		return
	}
	switch message.Command() {
	case "flightcheck":
		flightCheckCommand(message)
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hi")
		msg.ParseMode = tgbotapi.ModeHTML
		if _, err := Bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func flightCheckCommand(message *tgbotapi.Message) {
	flightNumber := strings.Replace(message.Text, "/flightcheck ", "", 1)
	flightData, err := flightcaller.GetFlightStatus(flightNumber)
	var messageHTML string
	if err != nil {
		messageHTML = "Error Flight Not Found or API Error"
	} else {
		messageHTML = fmt.Sprintf("Flight: %s\nAirline: %s\nDeparture: %s (%s)\nDeparture Schedule: %s\nArrival: %s (%s)\nArrival Schedule: %s",
			flightData.Flight.IATA,
			flightData.Airline.Name,
			flightData.Departure.Name,
			flightData.Departure.IATA,
			flightData.Departure.Scheduled.Format(time.ANSIC),
			flightData.Arrival.Name,
			flightData.Arrival.IATA,
			flightData.Arrival.Scheduled.Format(time.ANSIC))
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, messageHTML)
	msg.ParseMode = tgbotapi.ModeHTML
	if _, err := Bot.Send(msg); err != nil {
		log.Println(err)
	}
}
