package deck

import (
	card "go-deck-of-cards/internal/app/model/card"

	"github.com/google/uuid"
)

type StandardDeck struct {
	Deck `bson:",inline"`
}

func NewStandardDeck(codes []string) *StandardDeck {

	cards, _ := card.GenerateStandardCards(codes)

	return &StandardDeck{Deck{RemainingCards: len(cards), Shuffled: false, UUID: uuid.New().String(), Cards: cards, Type: "standard"}}

}
