package pricecompare

import (
	"fmt"
	"slices"
	"strconv"

	u "github.com/bcicen/go-units"
	"github.com/fatih/structs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/nutmos/utilitybot/state"
)

type UserData struct {
	ConversationProgress ConversationProgress
	ProductCount         int
	BaseUnit             u.Unit
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
	ConversationTypeProductName
	ConversationTypeProductQuantity
	ConversationTypeProductPrice
	ConversationTypeProductUnit
)

type Product struct {
	Number           int
	Name             string
	Price            int
	Quantity         float32
	Unit             u.Unit
	PricePerBaseUnit float32
}

type QuantityType int

const (
	QuaytityTypeFluid QuantityType = iota
	QuantityTypeWeight
)

func StartCommand(message *tgbotapi.Message) []string {
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
	reply := "Great! How many products would you like to compare? Please note that we support up to 5 products at once."
	return []string{reply}
}

func ContinueCommand(message *tgbotapi.Message) []string {
	userID := message.Chat.ID
	userData := getCurrentState(userID)
	switch userData.ConversationProgress.ConversationType {
	case ConversationTypeInit:
		return productCountReply(message, userData)
	case ConversationTypeProductCount:
		return unitReply(message, userData)
	case ConversationTypeBaseUnit, ConversationTypeProductUnit:
		return productNameReply(message, userData)
	case ConversationTypeProductName:
		return productPriceReply(message, userData)
	case ConversationTypeProductPrice:
		return productQuantityReply(message, userData)
	case ConversationTypeProductQuantity:
		return productUnitReply(message, userData)
	}
	return []string{}
}

func productCountReply(message *tgbotapi.Message, userData UserData) []string {
	count, err := strconv.Atoi(message.Text)
	errorReply := "Sorry, please specify only 1 to 5"
	if err != nil {
		return []string{errorReply}
	}
	if !(count >= 1 && count <= 5) {
		return []string{errorReply}
	}
	userData.ConversationProgress.ConversationType = ConversationTypeProductCount
	userData.ProductCount = count
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply := "Great! What is the base unit you would like to use for the comparison?"
	return []string{reply}
}

func unitReply(message *tgbotapi.Message, userData UserData) []string {
	unit, err := u.Find(message.Text)
	if err != nil {
		errReply := "You enter the incorrect unit. Please enter the correct unit."
		return []string{errReply}
	}
	userData.ConversationProgress.ConversationType = ConversationTypeBaseUnit
	userData.BaseUnit = unit
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply1 := "Sure, let's start the unit conversion."
	reply2 := "Product 1: Please enter the product name"
	return []string{reply1, reply2}
}

func productNameReply(message *tgbotapi.Message, userData UserData) []string {
	userData.ConversationProgress.ConversationType = ConversationTypeProductName
	userData.ProductList = append(userData.ProductList, Product{
		Name: message.Text,
	})
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply := "Product 1: Please enter the product price"
	return []string{reply}
}

func productPriceReply(message *tgbotapi.Message, userData UserData) []string {
	userData.ConversationProgress.ConversationType = ConversationTypeProductPrice
	lastProduct := len(userData.ProductList) - 1
	price, err := strconv.Atoi(message.Text)
	if err != nil {
		errReply := "Please enter the correct price"
		return []string{errReply}
	}
	userData.ProductList[lastProduct].Price = price
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply := "Product 1: Please enter the product quantity without a unit"
	return []string{reply}
}

func productQuantityReply(message *tgbotapi.Message, userData UserData) []string {
	userData.ConversationProgress.ConversationType = ConversationTypeProductQuantity
	lastProduct := len(userData.ProductList) - 1
	quantity, err := strconv.ParseFloat(message.Text, 32)
	if err != nil {
		errReply := "Please enter the correct quantity"
		return []string{errReply}
	}
	userData.ProductList[lastProduct].Quantity = float32(quantity)
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply := "Product 1: Please enter the product quantity unit"
	return []string{reply}
}

func productUnitReply(message *tgbotapi.Message, userData UserData) []string {
	unit, err := u.Find(message.Text)
	if err != nil {
		errReply := "You enter the incorrect unit. Please enter the correct unit."
		return []string{errReply}
	}
	userData.ConversationProgress.ConversationType = ConversationTypeProductUnit
	lastProduct := len(userData.ProductList) - 1
	userData.ProductList[lastProduct].Unit = unit
	userID := message.Chat.ID
	state.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	if lastProduct == userData.ProductCount-1 {
		// end - must compile result
		// dismiss cache
		return compileResult(message, userData)
	}
	reply := "Product 2: Please enter the product name."
	return []string{reply}
}

func compileResult(message *tgbotapi.Message, userData UserData) []string {
	userID := message.Chat.ID
	// convert to base unit
	for _, p := range userData.ProductList {
		if p.Unit.Name != userData.BaseUnit.Name {
			// convert unit
			newVal, err := u.ConvertFloat(float64(p.Quantity), p.Unit, userData.BaseUnit)
			if err != nil {
				return []string{"convert unit error!"}
			}
			p.Quantity = float32(newVal.Float())
			p.Unit = userData.BaseUnit
		}
		// compare with price
		p.PricePerBaseUnit = float32(p.Price) / p.Quantity
	}
	slices.SortFunc(userData.ProductList, func(a, b Product) int {
		if a.PricePerBaseUnit < b.PricePerBaseUnit {
			return -1
		} else if a.PricePerBaseUnit > b.PricePerBaseUnit {
			return 1
		}
		return 0
	})
	reply := ""
	for _, p := range userData.ProductList {
		reply += fmt.Sprintf("%.2f\n", p.PricePerBaseUnit)
	}
	state.DelUserState(fmt.Sprintf("%d", userID), "telegram")
	return []string{reply, "Complete Product!"}
}

func getCurrentState(userID int64) UserData {
	// query state
	currentState := state.GetUserState(fmt.Sprintf("%d", userID), "telegram")
	var userData UserData
	mapstructure.Decode(currentState.UserData, &userData)
	return userData
}
