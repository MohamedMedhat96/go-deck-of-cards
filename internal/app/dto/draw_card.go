package dto

import card "go-deck-of-cards/internal/app/model/card"

type DrawCardDTO struct {
	Cards []card.Card `json:"cards"`
}
