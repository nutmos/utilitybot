package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nutmos/utilitybot/config"
	"github.com/nutmos/utilitybot/flightcaller"
	"github.com/nutmos/utilitybot/random"
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
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hi! Please use the following commands:\n/flightcheck <flightNumberIata>\n/random <positiveIntNumber>")
		msg.ParseMode = tgbotapi.ModeMarkdown
		if _, err := Bot.Send(msg); err != nil {
			log.Println(err)
		}
		break
	case "flightcheck":
		flightCheckCommand(message)
		break
	case "random":
		randomCommand(message)
		break
	}
}

func flightCheckCommand(message *tgbotapi.Message) {
	flightNumber := strings.Replace(message.Text, "/flightcheck ", "", 1)
	flightData, err := flightcaller.GetFlightStatus(flightNumber)
	if err != nil {
		sendMessage(message, "Error: Flight Not Found or API Error")
	}
	messageHTML := fmt.Sprintf("Flight: %s\nAirline: %s\nDeparture: %s (%s)\nDeparture Schedule: %s\nDeparture Estimate: %s\nArrival: %s (%s)\nArrival Schedule: %s\nArrival Estimate: %s",
		flightData.Flight.IATA,
		flightData.Airline.Name,
		flightData.Departure.Name,
		flightData.Departure.IATA,
		flightData.Departure.Scheduled.Format(time.ANSIC),
		flightData.Departure.Estimated.Format(time.ANSIC),
		flightData.Arrival.Name,
		flightData.Arrival.IATA,
		flightData.Arrival.Scheduled.Format(time.ANSIC),
		flightData.Arrival.Estimated.Format(time.ANSIC))
	sendMessage(message, messageHTML)
}

func randomCommand(message *tgbotapi.Message) {
	randomNumberRangeString := strings.Replace(message.Text, "/random ", "", 1)
	randomNumberRange, err := strconv.Atoi(randomNumberRangeString)
	if err != nil {
		sendMessage(message, "Error: Please Enter Only Positive Integer")
	}
	result := random.RandomNumber(randomNumberRange)
	sendMessage(message, fmt.Sprintf("%d", result))
}

func sendMessage(receivingMsg *tgbotapi.Message, sendingMessageHTML string) {
	msg := tgbotapi.NewMessage(receivingMsg.Chat.ID, sendingMessageHTML)
	msg.ParseMode = tgbotapi.ModeHTML
	if _, err := Bot.Send(msg); err != nil {
		log.Println(err)
	}
}
