package model

import (
	"fmt"
	"sort"
)

func SortStandardCards(cards []Card) {
	orderSuit := map[Suit]int{
		"SPADE":   1,
		"DIAMOND": 2,
		"CLUB":    3,
		"HEART":   4,
	}
	orderValue := map[Value]int{
		"ACE":   1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"10":    10,
		"JACK":  11,
		"QUEEN": 12,
		"KING":  13,
	}

	sortFunc := func(i, j int) bool {
		cardI := cards[i]
		cardJ := cards[j]

		if orderSuit[cardI.Suit] != orderSuit[cardJ.Suit] {
			return orderSuit[cardI.Suit] < orderSuit[cardJ.Suit]
		}
		return orderValue[cardI.Value] < orderValue[cardJ.Value]
	}

	// Sort the cards using the custom sorting function
	sort.SliceStable(cards, sortFunc)
}
func GenerateStandardCards(codes []string) ([]Card, error) {
	var cards []Card

	suitMap := SuitMap()
	valueMap := ValueMap()

	if len(codes) == 0 {
		for s, suit := range suitMap {
			for v, value := range valueMap {
				cards = append(cards, Card{Suit: suit, Value: value, Code: fmt.Sprintf("%s%s", v, s)})
			}
		}

		SortStandardCards(cards)

		return cards, nil
	}

	for _, code := range codes {
		if len(code) < 2 {
			return nil, fmt.Errorf("invalid code: %s", code)
		}
		valueCode, suitCode := code[:len(code)-1], code[len(code)-1:]

		suit, suitOk := suitMap[suitCode]
		value, valueOk := valueMap[valueCode]

		if !suitOk || !valueOk {
			return nil, fmt.Errorf("invalid code: %s", code)
		}

		cards = append(cards, Card{Suit: suit, Value: value, Code: code})
	}

	SortStandardCards(cards)

	return cards, nil
}

func SuitMap() map[string]Suit {
	return map[string]Suit{
		"S": "SPADE",
		"D": "DIAMOND",
		"C": "CLUB",
		"H": "HEART",
	}
}

func ValueMap() map[string]Value {
	return map[string]Value{
		"A": "ACE", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "10": "10",
		"J": "JACK", "Q": "QUEEN", "K": "KING",
	}
}
