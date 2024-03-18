package service

import (
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

	d, err := factory.NewDeckFactory().Create("standard", Codes)

	if err != nil {
		return nil, err
	}

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

	if sd == nil {
		return nil, nil
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
		return nil, nil
	}

	sd, err := ds.repo.Get(c)

	if err != nil {
		return nil, err
	}

	if sd == nil {
		return nil, nil
	}

	return &dto.OpenDeckDTO{
		RemainingCards: sd.GetRemainingCards(),
		Shuffled:       sd.IsShuffled(),
		UUID:           sd.GetUUID(),
		Cards:          sd.GetCards(),
	}, nil
}
