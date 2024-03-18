package model

import (
	"reflect"
	"testing"
)

func TestSortStandardCards(t *testing.T) {
	cards := []Card{
		{Suit: "HEART", Value: "2", Code: "2H"},
		{Suit: "SPADE", Value: "KING", Code: "KS"},
		{Suit: "DIAMOND", Value: "ACE", Code: "AD"},
		{Suit: "CLUB", Value: "10", Code: "10C"},
		{Suit: "SPADE", Value: "1", Code: "1S"},
	}
	expected := []Card{
		{Suit: "SPADE", Value: "1", Code: "1S"},
		{Suit: "SPADE", Value: "KING", Code: "KS"},
		{Suit: "DIAMOND", Value: "ACE", Code: "AD"},
		{Suit: "CLUB", Value: "10", Code: "10C"},
		{Suit: "HEART", Value: "2", Code: "2H"},
	}

	SortStandardCards(cards)
	if !reflect.DeepEqual(cards, expected) {
		t.Errorf("SortStandardCards() = %v, want %v", cards, expected)
	}
}

func TestGenerateStandardCards(t *testing.T) {
	tests := []struct {
		name    string
		codes   []string
		want    []Card
		wantErr bool
	}{
		{
			name:  "Valid single code",
			codes: []string{"KS"},
			want:  []Card{{Suit: "SPADE", Value: "KING", Code: "KS"}},
		},
		{
			name:  "Valid multiple codes",
			codes: []string{"KS", "AD"},
			want: []Card{
				{Suit: "SPADE", Value: "KING", Code: "KS"},
				{Suit: "DIAMOND", Value: "ACE", Code: "AD"},
			},
		},
		{
			name:    "Invalid code length",
			codes:   []string{"K"},
			wantErr: true,
		},
		{
			name:    "Invalid suit code",
			codes:   []string{"KT"},
			wantErr: true,
		},
		{
			name:    "Invalid value code",
			codes:   []string{"XS"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateStandardCards(tt.codes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateStandardCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStandardCards() got = %v, want %v", got, tt.want)
			}
		})
	}
}
