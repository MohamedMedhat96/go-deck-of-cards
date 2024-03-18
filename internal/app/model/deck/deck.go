package model

import (
	"fmt"
	card "go-deck-of-cards/internal/app/model/card"
	"math/rand"
	"time"
)

type Deck struct {
	RemainingCards int         `json:"remaining" bson:"remaining"`
	Cards          []card.Card `json:"cards" bson:"cards"`
	Shuffled       bool        `json:"shuffled" bson:"shuffled"`
	UUID           string      `json:"deck_id" bson:"uuid"`
	Type           string      `json:"type" bson:"type"`
}

type DeckOperations interface {
	Shuffle()
	GetRemainingCards() int
	IsShuffled() bool
	GetUUID() string
	GetCards() []card.Card
	Deal(int) []card.Card
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})

	d.Shuffled = true
}

func (d *Deck) GetRemainingCards() int {
	return d.RemainingCards
}

func (d *Deck) GetUUID() string {
	return d.UUID
}

func (d *Deck) IsShuffled() bool {
	return d.Shuffled
}

func (d *Deck) GetCards() []card.Card {
	return d.Cards
}

func (d *Deck) Deal(NumberOfCards int) []card.Card {
	if d.RemainingCards == 0 {
		return []card.Card{}
	}

	if NumberOfCards > d.GetRemainingCards() {
		NumberOfCards = d.GetRemainingCards()
	}
	fmt.Println(NumberOfCards)
	fmt.Println(d.RemainingCards)
	c := d.Cards[:NumberOfCards]
	d.Cards = d.Cards[NumberOfCards:]
	d.RemainingCards = d.RemainingCards - NumberOfCards

	return c
}
