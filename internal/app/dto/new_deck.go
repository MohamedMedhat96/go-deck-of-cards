package dto

type NewDeckDTO struct {
	RemainingCards int    `json:"remaining"`
	Shuffled       bool   `json:"shuffled"`
	UUID           string `json:"deck_id"`
}
