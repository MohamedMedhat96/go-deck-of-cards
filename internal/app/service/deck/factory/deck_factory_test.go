package service

import (
	deck "go-deck-of-cards/internal/app/model/deck"
	"testing"
)

func TestNewDeckFactory_Singleton(t *testing.T) {
	firstInstance := NewDeckFactory()
	secondInstance := NewDeckFactory()

	if firstInstance != secondInstance {
		t.Error("NewDeckFactory should return the same instance")
	}
}

func TestDeckFactory_CreateStandardDeck(t *testing.T) {
	factory := NewDeckFactory()
	deckType := "standard"
	codes := []string{"AS", "KD", "10C"}

	createdDeck, _ := factory.Create(deckType, codes)

	if standardDeck, ok := createdDeck.(*deck.StandardDeck); ok {
		if len(standardDeck.GetCards()) != len(codes) {
			t.Errorf("Expected deck length %d, got %d", len(codes), len(standardDeck.GetCards()))
		}

	} else {
		t.Errorf("Expected *deck.StandardDeck, got %T", createdDeck)
	}
}
