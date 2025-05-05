package pricecompare

import (
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/nutmos/utilitybot/state"
)

type UserData struct {
	ConversationProgress ConversationProgress
	ProductCount         int
	BaseUnit             string
	ProductList          []Product
}

type ConversationProgress struct {
	Product          int
	ConversationType ConversationType
}

type ConversationType int

const (
	ConversationTypeInit ConversationType = iota
	ConversationTypeProductCount
	ConversationTypeBaseUnit
	ConversationTypeQuantity
	ConversationTypePrice
)

type Product struct {
	Number       int
	Price        int
	Quantity     int
	QuantityType QuantityType
}

type QuantityType int

const (
	QuaytityTypeFluid QuantityType = iota
	QuantityTypeWeight
)

func StartCommand(message *tgbotapi.Message) string {
	// start the state
	userID := message.Chat.ID
	emptyUserData := UserData{
		ConversationProgress: ConversationProgress{
			Product:          0,
			ConversationType: ConversationTypeInit,
		},
	}
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(emptyUserData))
	// ask the user for how many products
	return "Great! How many products would you like to compare? Please note that we support up to 5 products at once."
}

func ContinueCommand(message *tgbotapi.Message) string {
	userID := message.Chat.ID
	userData := getCurrentState(userID)
	switch userData.ConversationProgress.ConversationType {
	case ConversationTypeInit:
		return productCountReply(message, userData)
	case ConversationTypeBaseUnit:
		return unitReply(message, userData)
	}
	return ""
}

func productCountReply(message *tgbotapi.Message, userData UserData) string {
	count, err := strconv.Atoi(message.Text)
	if err != nil {
		return "Sorry, please specify only 1 to 5"
	}
	if !(count >= 1 && count <= 5) {
		return "Sorry, please specify only 1 to 5"
	}
	userData.ConversationProgress.ConversationType = ConversationTypeProductCount
	userData.ProductCount = count
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	return "Great! What is the base unit you would like to use for the comparison?"
}

func unitReply(message *tgbotapi.Message) string {

}

func ProgressCommand(message *tgbotapi.Message) {
	// query state
	userID := message.Chat.ID
	currentState := state.GetUserState(fmt.Sprintf("%d", userID), "telegram")
	var userData UserData
	mapstructure.Decode(currentState.UserData, &userData)
	if userData.ConversationProgress.ConversationType == ConversationTypeQuantity {

	} else if userData.ConversationProgress.ConversationType == ConversationTypePrice {

	}
}

func getCurrentState(userID int64) UserData {
	// query state
	currentState := state.GetUserState(fmt.Sprintf("%d", userID), "telegram")
	var userData UserData
	mapstructure.Decode(currentState.UserData, &userData)
	return userData
}
