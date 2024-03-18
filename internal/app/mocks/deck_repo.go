package mocks

import (
	deck "go-deck-of-cards/internal/app/model/deck"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockDeckRepo struct {
	mock.Mock
}

func (m *MockDeckRepo) Add(c *gin.Context, d deck.DeckOperations) (bool, error) {
	args := m.Called(c, d)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockDeckRepo) Get(c *gin.Context) (deck.DeckOperations, error) {
	args := m.Called(c)
	return args.Get(0).(deck.DeckOperations), args.Error(1)
}

func (m *MockDeckRepo) Update(c *gin.Context, update bson.M) (bool, error) {
	args := m.Called(c, update)
	return args.Bool(0), args.Error(1)
}
