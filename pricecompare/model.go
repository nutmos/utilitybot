package pricecompare

import u "github.com/bcicen/go-units"

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
	ConversationTypeProduct
)

type Product struct {
	Number           int
	Name             string
	Price            float32
	Quantity         float32
	Unit             u.Unit
	PricePerBaseUnit float32
}

type QuantityType int

const (
	QuaytityTypeFluid QuantityType = iota
	QuantityTypeWeight
)
