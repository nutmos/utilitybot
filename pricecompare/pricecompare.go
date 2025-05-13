package pricecompare

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	u "github.com/bcicen/go-units"
	"github.com/fatih/structs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/nutmos/utilitybot/userstatehandler"
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
	userstatehandler.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(emptyUserData))
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
	case ConversationTypeBaseUnit, ConversationTypeProduct:
		return productReply(message, userData)
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
	userstatehandler.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
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
	userstatehandler.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	reply1 := fmt.Sprintf("You choose %s as a base unit.", unit.Name)
	reply2 := "Let's start the unit conversion."
	reply3 := "Please enter the product with the following format\n\nProduct Name\nPrice (without currency)\nQuantity (without unit)\nUnit"
	reply4 := "Starting with product #1:"
	return []string{reply1, reply2, reply3, reply4}
}

func productReply(message *tgbotapi.Message, userData UserData) []string {
	// parse test
	productInfo := strings.Split(message.Text, "\n")
	userID := message.Chat.ID
	productNum := len(userData.ProductList)
	formatErrMessage := "Wrong product information! Please follow the correct format."
	if len(productInfo) != 4 {
		fmt.Println("Error: arguments not complete")
		return []string{formatErrMessage}
	}
	productName := productInfo[0]
	productPrice, err := strconv.ParseFloat(productInfo[1], 32)
	if err != nil {
		fmt.Println("Error at price")
		return []string{formatErrMessage}
	}
	productQuantity, err := strconv.ParseFloat(productInfo[2], 32)
	if err != nil {
		fmt.Println("Error at quantity")
		return []string{formatErrMessage}
	}
	productUnit, err := u.Find(productInfo[3])
	if err != nil {
		fmt.Println("Error at unit")
		return []string{formatErrMessage}
	}
	// update the text
	userData.ProductList = append(userData.ProductList, Product{
		Number:   productNum,
		Name:     productName,
		Price:    float32(productPrice),
		Quantity: float32(productQuantity),
		Unit:     productUnit,
	})
	userstatehandler.SetUserState(fmt.Sprintf("%d", userID), "telegram", "pricecompare", structs.Map(userData))
	// decide to compile result
	if len(userData.ProductList) == userData.ProductCount {
		return compileResult(message, userData)
	} else {
		nextMessage := fmt.Sprintf("Please enter the product #%d", len(userData.ProductList)+1)
		return []string{nextMessage}
	}
}

func compileResult(message *tgbotapi.Message, userData UserData) []string {
	userID := message.Chat.ID
	// convert to base unit
	for i := range userData.ProductList {
		if userData.ProductList[i].Unit.Name != userData.BaseUnit.Name {
			// convert unit
			newVal, err := u.ConvertFloat(float64(userData.ProductList[i].Quantity), userData.ProductList[i].Unit, userData.BaseUnit)
			if err != nil {
				return []string{"convert unit error!"}
			}
			userData.ProductList[i].Quantity = float32(newVal.Float())
			userData.ProductList[i].Unit = userData.BaseUnit
		}
		// compare with price
		userData.ProductList[i].PricePerBaseUnit = float32(userData.ProductList[i].Price) / userData.ProductList[i].Quantity
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
		reply += fmt.Sprintf("Product %s: %.2f per %s\n", p.Name, p.PricePerBaseUnit, p.Unit.Name)
	}
	userstatehandler.DelUserState(fmt.Sprintf("%d", userID), "telegram")
	return []string{reply, "Thank you for using our tool. Have a nice day!"}
}

func getCurrentState(userID int64) UserData {
	// query state
	currentState := userstatehandler.GetUserState(fmt.Sprintf("%d", userID), "telegram")
	var userData UserData
	mapstructure.Decode(currentState.UserData, &userData)
	return userData
}
