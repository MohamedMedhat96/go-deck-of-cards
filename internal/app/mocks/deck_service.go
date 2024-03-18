package mocks

import (
	"go-deck-of-cards/internal/app/dto"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockDeckService struct {
	mock.Mock
}

func (m *MockDeckService) CreateNewDeck(c *gin.Context, shuffled bool, codes []string) (*dto.NewDeckDTO, error) {
	args := m.Called(c, shuffled, codes)
	return args.Get(0).(*dto.NewDeckDTO), args.Error(1)
}

func (m *MockDeckService) DrawCard(c *gin.Context, numberOfCards int) (*dto.DrawCardDTO, error) {
	args := m.Called(c, numberOfCards)
	return args.Get(0).(*dto.DrawCardDTO), args.Error(1)
}

func (m *MockDeckService) OpenDeck(c *gin.Context) (*dto.OpenDeckDTO, error) {
	args := m.Called(c)
	return args.Get(0).(*dto.OpenDeckDTO), args.Error(1)
}
