package service

import deck "go-deck-of-cards/internal/app/model/deck"

type DeckFactory struct{}

var singletonInstance *DeckFactory

func NewDeckFactory() *DeckFactory {
	if singletonInstance == nil {
		singletonInstance = &DeckFactory{}
	}
	return singletonInstance
}

func (df *DeckFactory) Create(Type string, Codes []string) (deck.DeckOperations, error) {
	switch Type {
	case "standard":
		return deck.NewStandardDeck(Codes)
	default:
		return deck.NewStandardDeck(Codes)
	}
}
