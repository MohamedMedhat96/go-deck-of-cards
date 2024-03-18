package dto

import card "go-deck-of-cards/internal/app/model/card"

type OpenDeckDTO struct {
	RemainingCards int         `json:"remaining"`
	Shuffled       bool        `json:"shuffled"`
	UUID           string      `json:"deck_id"`
	Cards          []card.Card `json:"cards"`
}
