package main

import "fmt"

// Card
// holds the value and suite of a card
// func (card *Card) compare(winningCard Card, trump *Card, first Card) bool
// 		used to compare the current card with the winning card -> return true if the current card is better
//															   -> return false, otherwise

// func (card *Card) equals(otherCard Card) bool
//		used to compare the suites of two cards

type suites int

const (
	HEARTS   = 1
	CLUBS    = 2
	DIAMONDS = 3
	SPADES   = 4
)

type Card struct {
	Suite suites `json:"suite"`
	Value int    `json:"value"`
}

func NewCard(suite suites, value int) *Card {
	return &Card{
		Suite: suite,
		Value: value,
	}
}

func (card *Card) Compare(winningCard Card, trump *Card, first Card) bool {
	if trump != nil && card.Suite != trump.Suite && card.Suite != first.Suite {
		return false
	}
	if card.Suite == winningCard.Suite {
		return card.Value > winningCard.Value
	}
	if trump != nil && card.Suite == first.Suite && winningCard.Suite == trump.Suite {
		return false
	}
	if trump != nil && winningCard.Suite == first.Suite && card.Suite == trump.Suite {
		return true
	}
	return false
}

func (card *Card) Equals(otherCard Card) bool {
	return card.Suite == otherCard.Suite
}

func (card *Card) String() string {
	return fmt.Sprintf("Suite: %d, Value: %d", card.Suite, card.Value)
}

func (card *Card) showCard() { // nu avem nevoie de functia asta
	fmt.Printf("%d %d\n", card.Suite, card.Value)
}
