package model

import (
	card "go-deck-of-cards/internal/app/model/card"
	"testing"
)

func TestDeck_Shuffle(t *testing.T) {
	firstCard := &card.Card{Suit: "SPADE", Value: "A", Code: "AS"}
	cards, _ := card.GenerateStandardCards([]string{})
	d := Deck{
		RemainingCards: 3,
		Cards:          cards,
		Shuffled:       false,
	}

	d.Shuffle()

	if !d.IsShuffled() {
		t.Errorf("Expected deck to be shuffled")
	}
	if d.Cards[0].GetCode() == firstCard.GetCode() {
		t.Errorf("Deck was not shuffled")
	}
}

func TestDeck_Deal(t *testing.T) {
	initialDeck := []card.Card{
		{Suit: "SPADE", Value: "A", Code: "AS"},
		{Suit: "DIAMOND", Value: "K", Code: "KD"},
		{Suit: "CLUB", Value: "10", Code: "10C"},
	}

	tests := []struct {
		name          string
		deck          Deck
		numToDeal     int
		wantDealt     int
		wantRemaining int
		wantShuffled  bool
	}{
		{
			name: "Deal less than remaining",
			deck: Deck{
				RemainingCards: len(initialDeck),
				Cards:          initialDeck,
				Shuffled:       false,
			},
			numToDeal:     2,
			wantDealt:     2,
			wantRemaining: 1,
			wantShuffled:  false,
		},
		{
			name: "Deal more than remaining",
			deck: Deck{
				RemainingCards: len(initialDeck),
				Cards:          initialDeck,
				Shuffled:       false,
			},
			numToDeal:     5,
			wantDealt:     3,
			wantRemaining: 0,
			wantShuffled:  false,
		},
		{
			name: "Deal from empty deck",
			deck: Deck{
				RemainingCards: 0,
				Cards:          []card.Card{},
				Shuffled:       false,
			},
			numToDeal:     1,
			wantDealt:     0,
			wantRemaining: 0,
			wantShuffled:  false,
		},
		{
			name: "Deal exact number remaining",
			deck: Deck{
				RemainingCards: len(initialDeck),
				Cards:          initialDeck,
				Shuffled:       true,
			},
			numToDeal:     3,
			wantDealt:     3,
			wantRemaining: 0,
			wantShuffled:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dealtCards := tt.deck.Deal(tt.numToDeal)
			if gotDealt := len(dealtCards); gotDealt != tt.wantDealt {
				t.Errorf("%s: expected %d cards to be dealt, got %d", tt.name, tt.wantDealt, gotDealt)
			}
			if tt.deck.GetRemainingCards() != tt.wantRemaining {
				t.Errorf("%s: expected %d remaining cards, got %d", tt.name, tt.wantRemaining, tt.deck.GetRemainingCards())
			}
			if tt.deck.IsShuffled() != tt.wantShuffled {
				t.Errorf("%s: expected shuffled status to be %v, got %v", tt.name, tt.wantShuffled, tt.deck.IsShuffled())
			}
		})
	}
}
