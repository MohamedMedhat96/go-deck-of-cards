package model

import (
	"testing"
)

func TestNewStandardDeck(t *testing.T) {
	tests := []struct {
		name     string
		codes    []string
		wantLen  int
		wantType string
	}{
		{
			name:     "Full deck without codes",
			codes:    nil,
			wantLen:  52,
			wantType: "standard",
		},
		{
			name:     "Subset of cards",
			codes:    []string{"AS", "KD", "10C"},
			wantLen:  3,
			wantType: "standard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewStandardDeck(tt.codes)

			if gotLen := len(deck.Cards); gotLen != tt.wantLen {
				t.Errorf("NewStandardDeck() gotLen = %v, want %v", gotLen, tt.wantLen)
			}

			if deck.Type != tt.wantType {
				t.Errorf("NewStandardDeck() Type = %v, want %v", deck.Type, tt.wantType)
			}

			if deck.UUID == "" {
				t.Errorf("NewStandardDeck() UUID should not be empty")
			}

			if tt.wantLen != deck.RemainingCards {
				t.Errorf("NewStandardDeck() RemainingCards = %v, want %v", deck.RemainingCards, tt.wantLen)
			}

			if deck.Shuffled != false {
				t.Errorf("NewStandardDeck() Shuffled = %v, want %v", deck.Shuffled, false)
			}
		})
	}
}
