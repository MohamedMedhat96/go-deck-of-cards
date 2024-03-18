package handler

import (
	"fmt"
	"go-deck-of-cards/internal/app/dto"
	"go-deck-of-cards/internal/app/mocks"
	card "go-deck-of-cards/internal/app/model/card"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(mockService *mocks.MockDeckService) *gin.Engine {
	r := gin.Default()

	h := &deckHandler{service: mockService}
	h.RegisterRoutes(r)
	return r
}

func TestCreateNewDeck(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ms := new(mocks.MockDeckService)
		router := setupRouter(ms)

		expectedDeck := &dto.NewDeckDTO{
			RemainingCards: 3,
			Shuffled:       true,
			UUID:           "some-uuid",
		}
		ms.On("CreateNewDeck", mock.Anything, true, []string{"AS", "KD", "10C"}).Return(expectedDeck, nil)

		req, _ := http.NewRequest("POST", "/deck?shuffled=true&cards=AS,KD,10C", nil)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Failure - incorrect code", func(t *testing.T) {
		ms := new(mocks.MockDeckService)
		router := setupRouter(ms)

		ms.On("CreateNewDeck", mock.Anything, mock.Anything, mock.Anything).Return((*dto.NewDeckDTO)(nil), fmt.Errorf("incorrect code(s): 10CS"))

		req, _ := http.NewRequest("POST", "/deck?cards=AS,KD,10CS", nil)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestOpenDeck(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ms := new(mocks.MockDeckService)
		router := setupRouter(ms)

		expectedDeck := &dto.OpenDeckDTO{
			RemainingCards: 3,
			Shuffled:       true,
			UUID:           "some-uuid",
		}
		ms.On("OpenDeck", mock.Anything).Return(expectedDeck, nil)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/deck/%s", "some-uuid"), nil)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Failure - incorrect code", func(t *testing.T) {
		ms := new(mocks.MockDeckService)
		router := setupRouter(ms)

		ms.On("OpenDeck", mock.Anything).Return((*dto.OpenDeckDTO)(nil), nil)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/deck/%s", "some-uuid"), nil)

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDrawCard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		noc := 1
		ms := new(mocks.MockDeckService)
		router := setupRouter(ms)

		expectedDeck := &dto.DrawCardDTO{
			Cards: []card.Card{{Suit: "SPADE", Value: "Ace", Code: "AS"}},
		}

		ms.On("DrawCard", mock.Anything, noc).Return(expectedDeck, nil)

		req, _ := http.NewRequest("POST", fmt.Sprintf("/deck/some-uuid/draw?numberOfCards=%v", noc), nil)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
