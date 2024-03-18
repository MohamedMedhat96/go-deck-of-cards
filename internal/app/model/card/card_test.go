package model

import (
	"testing"
)

func TestCard_GetSuit(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want Suit
	}{
		{
			name: "Test Hearts Suit",
			card: Card{Suit: "Hearts", Value: "Ace", Code: "AH"},
			want: "Hearts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.GetSuit(); got != tt.want {
				t.Errorf("Card.GetSuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetValue(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want Value
	}{
		{
			name: "Test Ace Value",
			card: Card{Suit: "Hearts", Value: "Ace", Code: "AH"},
			want: "Ace",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.GetValue(); got != tt.want {
				t.Errorf("Card.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetCode(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want string
	}{
		{
			name: "Test Code",
			card: Card{Suit: "Hearts", Value: "Ace", Code: "AH"},
			want: "AH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.GetCode(); got != tt.want {
				t.Errorf("Card.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
