package pricecompare

import (
	"fmt"

	"github.com/fatih/structs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/nutmos/utilitybot/state"
)

type UserData struct {
	ConversationProgress ConversationProgress
	ProductList          []Product
}

type ConversationProgress struct {
	Product          int
	ConversationType ConversationType
}

type ConversationType int

const (
	ConversationTypeQuantity ConversationType = iota
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

func InitCommand(message *tgbotapi.Message) {
	// start the state
	userID := message.Chat.ID
	emptyUserData := UserData{
		ConversationProgress: ConversationProgress{
			Product:          0,
			ConversationType: ConversationTypeQuantity,
		},
	}
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(emptyUserData))
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
