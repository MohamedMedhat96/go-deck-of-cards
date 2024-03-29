package repository

import (
	"fmt"
	deck "go-deck-of-cards/internal/app/model/deck"
	"go-deck-of-cards/internal/pkg/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DeckCollectionName = "deck"

type DeckRepository interface {
	Add(ctx *gin.Context, deck deck.DeckOperations) (bool, error)

	Get(ctx *gin.Context) (deck.DeckOperations, error)

	Update(ctx *gin.Context, filters bson.M) (bool, error)
}

type DeckRepositoryImpl struct{}

func (dr *DeckRepositoryImpl) Add(ctx *gin.Context, deck deck.DeckOperations) (bool, error) {
	done, err := db.GetCollection(DeckCollectionName).InsertOne(ctx, deck)
	return done != nil, err
}

func (dr *DeckRepositoryImpl) Get(ctx *gin.Context) (deck.DeckOperations, error) {
	var d deck.Deck

	uuid := ctx.GetString("uuid")

	collection := db.GetCollection(DeckCollectionName)
	err := collection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&d)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &d, nil
}

func (dr *DeckRepositoryImpl) Update(ctx *gin.Context, filters bson.M) (bool, error) {
	uuid := ctx.GetString("uuid")

	collection := db.GetCollection(DeckCollectionName)
	result, err := collection.UpdateOne(ctx, bson.M{"uuid": uuid}, filters)

	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, fmt.Errorf("no documents were found for the filter")
	}

	return true, nil
}
