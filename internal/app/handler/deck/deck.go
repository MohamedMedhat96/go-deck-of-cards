package handler

import (
	deck "go-deck-of-cards/internal/app/service/deck"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterRoutes(router *gin.Engine)
}

type deckHandler struct {
	service deck.DeckService
}

func NewDeckHandler() Handler {
	s := deck.NewDeckService()
	return &deckHandler{service: s}
}

func (h *deckHandler) CreateNewDeck(c *gin.Context) {
	shuffled := c.DefaultQuery("shuffled", "false") == "true"

	var codes []string
	if c.Query("cards") != "" {
		codes = strings.Split(c.Query("cards"), ",")
	} else {
		codes = []string{}
	}

	nd, err := h.service.CreateNewDeck(c, shuffled, codes)

	if err != nil {
		c.JSON(http.StatusBadRequest, &map[string]string{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nd)
}

type Val struct {
	Name string
}

func (h *deckHandler) OpenDeck(c *gin.Context) {
	nd, err := h.service.OpenDeck(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &map[string]string{"error": err.Error()})
		return
	}

	if nd == nil {
		c.JSON(http.StatusNotFound, &map[string]string{"error": "The deck you are trying to open does not exist"})
		return
	}

	c.JSON(http.StatusOK, nd)

}
func (h *deckHandler) DrawCard(c *gin.Context) {
	ns, exists := c.GetQuery("numberOfCards")
	noc := 1

	if exists {
		var err error
		noc, err = strconv.Atoi(ns)
		if err != nil {
			c.JSON(http.StatusBadRequest, &map[string]string{"error": "The value you have entered for numberOfCards is not a number, please try again with a number"})
			return
		}
	}

	nd, err := h.service.DrawCard(c, noc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &map[string]string{"error": err.Error()})
		return
	}

	if nd == nil {
		c.JSON(http.StatusNotFound, &map[string]string{"error": "The deck you are trying to open does not exist"})
		return
	}
	c.JSON(http.StatusOK, nd)
}

func (h *deckHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/decks", h.CreateNewDeck)
	router.GET("/decks/:uuid", h.OpenDeck)
	router.POST("/decks/:uuid/draw", h.DrawCard)
}
