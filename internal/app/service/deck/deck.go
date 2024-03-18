package service

import (
	"fmt"
	"go-deck-of-cards/internal/app/dto"
	card "go-deck-of-cards/internal/app/model/card"
	"go-deck-of-cards/internal/app/repository"
	factory "go-deck-of-cards/internal/app/service/deck/factory"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type DeckService interface {
	CreateNewDeck(ctx *gin.Context, shuffled bool, codes []string) (*dto.NewDeckDTO, error)
	DrawCard(ctx *gin.Context, numberOfCards int) (*dto.DrawCardDTO, error)
	OpenDeck(ctx *gin.Context) (*dto.OpenDeckDTO, error)
}

type DeckServiceImpl struct {
	repo repository.DeckRepository
}

func NewDeckService() DeckService {
	return &DeckServiceImpl{repo: &repository.DeckRepositoryImpl{}}
}

func (di *DeckServiceImpl) CreateNewDeck(c *gin.Context, Shuffled bool, Codes []string) (*dto.NewDeckDTO, error) {

	d := factory.NewDeckFactory().Create("standard", Codes)

	if Shuffled {
		d.Shuffle()
	}

	di.repo.Add(c, d)

	return &dto.NewDeckDTO{
		RemainingCards: d.GetRemainingCards(),
		Shuffled:       d.IsShuffled(),
		UUID:           d.GetUUID(),
	}, nil
}

func (ds *DeckServiceImpl) DrawCard(c *gin.Context, NumberOfCards int) (*dto.DrawCardDTO, error) {
	sd, err := ds.repo.Get(c)

	if err != nil {
		return nil, err
	}

	if NumberOfCards > sd.GetRemainingCards() {
		NumberOfCards = sd.GetRemainingCards()
	}

	var cards []card.Card

	if sd.GetRemainingCards() > 0 {
		cards = sd.Deal(NumberOfCards)

		_, err := ds.repo.Update(c, bson.M{"$set": bson.M{"remaining": sd.GetRemainingCards(), "cards": sd.GetCards()}})
		if err != nil {
			return nil, err
		}
	} else {
		cards = []card.Card{}
	}

	return &dto.DrawCardDTO{
		Cards: cards,
	}, nil
}

func (ds *DeckServiceImpl) OpenDeck(c *gin.Context) (*dto.OpenDeckDTO, error) {
	_, exists := c.Get("uuid")

	if !exists {
		return nil, fmt.Errorf("please provide a uuid")
	}

	sd, err := ds.repo.Get(c)

	if err != nil {
		return nil, err
	}

	return &dto.OpenDeckDTO{
		RemainingCards: sd.GetRemainingCards(),
		Shuffled:       sd.IsShuffled(),
		UUID:           sd.GetUUID(),
		Cards:          sd.GetCards(),
	}, nil
}
