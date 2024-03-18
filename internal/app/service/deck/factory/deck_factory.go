package service

import deck "go-deck-of-cards/internal/app/model/deck"

// DeckFactory struct for creating decks.
type DeckFactory struct{}

var singletonInstance *DeckFactory

func NewDeckFactory() *DeckFactory {
	if singletonInstance == nil {
		singletonInstance = &DeckFactory{}
	}
	return singletonInstance
}

func (df *DeckFactory) Create(Type string, Codes []string) deck.DeckOperations {
	switch Type {
	case "standard":
		return deck.NewStandardDeck(Codes)
	default:
		return deck.NewStandardDeck(Codes)
	}
}
