package service

import (
	"go-deck-of-cards/internal/app/mocks"
	card "go-deck-of-cards/internal/app/model/card"
	deck "go-deck-of-cards/internal/app/model/deck"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDrawCard(t *testing.T) {
	mockRepo := new(mocks.MockDeckRepo)
	service := DeckServiceImpl{repo: mockRepo}

	tests := []struct {
		name          string
		numberOfCards int
		setupMock     func()
		wantErr       bool
		wantCardsLen  int
	}{
		{
			name:          "Draw from non-empty deck",
			numberOfCards: 2,
			setupMock: func() {
				mockRepo.On("Get", mock.Anything).Return(DeckWithCards(5), nil)
				mockRepo.On("Update", mock.Anything, mock.Anything).Return(true, nil)
			},
			wantErr:      false,
			wantCardsLen: 2,
		},
		{
			name:          "Request more cards than available",
			numberOfCards: 10,
			setupMock: func() {
				mockRepo.On("Get", mock.Anything).Return(DeckWithCards(3), nil)
				mockRepo.On("Update", mock.Anything, mock.Anything).Return(true, nil)
			},
			wantErr:      false,
			wantCardsLen: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &gin.Context{}
			tc.setupMock()

			result, err := service.DrawCard(ctx, tc.numberOfCards)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result.Cards, tc.wantCardsLen)
			}
		})
	}
}

func TestOpenDeck(t *testing.T) {
	mockRepo := new(mocks.MockDeckRepo)
	service := DeckServiceImpl{repo: mockRepo}

	tests := []struct {
		name    string
		setup   func(ctx *gin.Context)
		wantErr bool
	}{
		{
			name: "Successfully open a deck with valid UUID",
			setup: func(ctx *gin.Context) {
				doc := DeckWithCards(52)
				ctx.Set("uuid", doc.GetUUID())
				mockRepo.On("Get", ctx).Return(doc, nil)
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &gin.Context{}
			tc.setup(ctx)

			d, err := service.OpenDeck(ctx)
			if tc.wantErr {
				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.Equal(t, d.UUID, ctx.GetString("uuid"))
			}
		})
	}
}

func TestCreateNewStandardDeck(t *testing.T) {
	mockRepo := new(mocks.MockDeckRepo)
	service := DeckServiceImpl{repo: mockRepo}

	mockRepo.On("Add", mock.Anything, mock.AnythingOfType("*model.StandardDeck")).Return(true, nil)

	tests := []struct {
		name     string
		shuffled bool
		codes    []string
		wantErr  bool
	}{
		{
			name:     "Create with shuffle and specific codes",
			shuffled: true,
			codes:    []string{"AS", "KD", "QH"},
			wantErr:  false,
		},
		{
			name:     "Create new unshuffled deck",
			shuffled: false,
			codes:    []string{},
			wantErr:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.CreateNewDeck(&gin.Context{}, tc.shuffled, tc.codes)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				mockRepo.AssertCalled(t, "Add", mock.Anything, mock.AnythingOfType("*model.StandardDeck"))

			}
		})
	}
}

func DeckWithCards(n int) deck.DeckOperations {
	cards := make([]card.Card, n)
	for i := 0; i < n; i++ {
		cards[i] = card.Card{
			Suit:  "SPADE",
			Value: "A",
			Code:  "AS",
		}
	}
	d := &deck.Deck{Cards: cards, Shuffled: false, Type: "standard", UUID: uuid.NewString(), RemainingCards: n}
	return d
}
