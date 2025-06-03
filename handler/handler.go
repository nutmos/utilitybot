package handler

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nutmos/utilitybot/confighandler"
	"github.com/nutmos/utilitybot/flightcaller"
	"github.com/nutmos/utilitybot/myflights"
	"github.com/nutmos/utilitybot/pricecompare"
	"github.com/nutmos/utilitybot/random"
	state "github.com/nutmos/utilitybot/userstatehandler"
)

var (
	Bot *tgbotapi.BotAPI
)

func init() {
	var err error
	log.Printf("confighandler.Config.ApiKey.Telegram: %s", confighandler.Config.ApiKey.Telegram)
	Bot, err = tgbotapi.NewBotAPI(confighandler.Config.ApiKey.Telegram)
	if err != nil {
		panic(err)
	}
	Bot.Debug = false
	state.InitUserStateHandler(state.UserStateHandlerTypeInMemory)
}

func HandleMessage(message *tgbotapi.Message) {
	if !message.IsCommand() {
		// check the current state
		userID := message.Chat.ID
		currentState := state.GetUserState(fmt.Sprintf("%d", userID), "telegram")
		if currentState == nil {
			SendMessage(message, "Error: Please enter a command before proceeding.")
			return
		}
		currentCommand := currentState.CurrentCommand
		switch *currentCommand {
		case "pricecompare":
			contPriceCompare(message)
			break
		}
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
	case "showmyflights":
		showMyFlightsCommand(message)
		break
	case "pricecompare":
		priceCompare(message)
	}
}

func flightCheckCommand(message *tgbotapi.Message) {
	flightNumber := strings.Replace(message.Text, "/flightcheck ", "", 1)
	flightResp, err := flightcaller.GetFlight(&flightcaller.FlightRequest{
		FlightIATA: &flightNumber,
	})
	flightDataArr := flightResp.Data
	if err != nil {
		SendMessage(message, "Error: Flight Not Found or API Error")
	}
	messageHTML := ""
	for i, flightData := range flightDataArr {
		if i > 1 {
			messageHTML += "\n\n"
		}
		messageHTML += fmt.Sprintf("Flight: %s\nAirline: %s\nDeparture: %s (%s)\nDeparture Schedule: %s\nDeparture Estimate: %s\nArrival: %s (%s)\nArrival Schedule: %s\nArrival Estimate: %s",
			flightData.Flight.IATA,
			flightData.Airline.Name,
			flightData.Departure.Airport,
			flightData.Departure.IATA,
			flightData.Departure.Scheduled,
			flightData.Departure.Estimated,
			flightData.Arrival.Airport,
			flightData.Arrival.IATA,
			flightData.Arrival.Scheduled,
			flightData.Arrival.Estimated,
		)
	}
	SendMessage(message, messageHTML)
}

func randomCommand(message *tgbotapi.Message) {
	randomNumberRangeString := strings.Replace(message.Text, "/random ", "", 1)
	log.Println("randomNumberRangeString: %s\n", randomNumberRangeString)
	randomNumberRange, err := strconv.Atoi(randomNumberRangeString)
	if err != nil {
		log.Println("%v", err)
		SendMessage(message, "Error: Please Enter Only Positive Integer")
		return
	}
	result := random.RandomNumber(randomNumberRange)
	SendMessage(message, fmt.Sprintf("%d", result))
}

func showMyFlightsCommand(message *tgbotapi.Message) {
	mf, err := myflights.MyFlightQuery(&myflights.MyFlightQueryRequest{
		UserId: 1234,
	})
	if err != nil {
		log.Println(err)
		SendMessage(message, "eror")
		return
	}
	msgForSend := fmt.Sprintf("Flight Found: %d\n\n",
		len(mf.Data),
	)
	for _, d := range mf.Data {
		msgForSend += formatFlightMsg(d) + "\n\n"
	}
	SendMessage(message, msgForSend)
}

func SendMessage(receivingMsg *tgbotapi.Message, sendingMessageHTML string) {
	msg := tgbotapi.NewMessage(receivingMsg.Chat.ID, sendingMessageHTML)
	msg.ParseMode = tgbotapi.ModeHTML
	if _, err := Bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func formatFlightMsg(flightData *flightcaller.FlightResponseData) string {
	return fmt.Sprintf("Flight: %s\nAirline: %s\nDeparture: %s (%s)\nDeparture Schedule: %s\nDeparture Estimate: %s\nArrival: %s (%s)\nArrival Schedule: %s\nArrival Estimate: %s",
		flightData.Flight.IATA,
		flightData.Airline.Name,
		flightData.Departure.Airport,
		flightData.Departure.IATA,
		flightData.Departure.Scheduled,
		flightData.Departure.Estimated,
		flightData.Arrival.Airport,
		flightData.Arrival.IATA,
		flightData.Arrival.Scheduled,
		flightData.Arrival.Estimated,
	)
}

func priceCompare(message *tgbotapi.Message) {
	reply := pricecompare.StartCommand(message)
	for _, r := range reply {
		SendMessage(message, r)
	}
}

func contPriceCompare(message *tgbotapi.Message) {
	reply := pricecompare.ContinueCommand(message)
	for _, r := range reply {
		SendMessage(message, r)
	}
}
