package model

type Suit string
type Value string

type Card struct {
	Suit  Suit   `json:"suit" bson:"suit"`
	Value Value  `json:"value" bson:"value"`
	Code  string `json:"code" bson:"code"`
}
type CardOperations interface {
	GetValue() Value
	GetSuit() Suit
	GetCode() string
}

func (c *Card) GetSuit() Suit {
	return c.Suit
}

func (c *Card) GetValue() Value {
	return c.Value
}

func (c *Card) GetCode() string {
	return c.Code
}
