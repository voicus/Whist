package main

import (
	"testing"
)

type compareTestCase struct {
	card        *Card
	winningCard Card
	trump       *Card
	first       *Card
	expected    bool
}

type equalsTestCase struct {
	card      Card
	otherCard Card
	expected  bool
}

var compareTestCases = []compareTestCase{
	{NewCard(SPADES, J), *NewCard(CLUBS, J), nil, NewCard(CLUBS, 7), false},
	{NewCard(HEARTS, K), *NewCard(HEARTS, 9), nil, NewCard(HEARTS, 7), true},
	{NewCard(DIAMONDS, 7), *NewCard(DIAMONDS, J), nil, NewCard(DIAMONDS, 9), false},
	{NewCard(CLUBS, A), *NewCard(DIAMONDS, J), NewCard(DIAMONDS, 7), NewCard(DIAMONDS, 9), false},
	{NewCard(SPADES, K), *NewCard(HEARTS, 10), NewCard(SPADES, 7), NewCard(HEARTS, J), true},
	{NewCard(SPADES, K), *NewCard(HEARTS, 10), NewCard(CLUBS, K), NewCard(HEARTS, 11), false},
	{NewCard(CLUBS, K), *NewCard(SPADES, J), NewCard(CLUBS, 10), NewCard(SPADES, 7), true},
	{NewCard(DIAMONDS, 4), *NewCard(HEARTS, Q), NewCard(HEARTS, K), NewCard(SPADES, A), false},
	{NewCard(DIAMONDS, K), *NewCard(HEARTS, 8), NewCard(HEARTS, 9), NewCard(DIAMONDS, 8), false},
	{NewCard(HEARTS, J), *NewCard(HEARTS, 7), NewCard(HEARTS, 10), NewCard(CLUBS, Q), true},
	{NewCard(DIAMONDS, Q), *NewCard(SPADES, K), NewCard(SPADES, Q), NewCard(CLUBS, 9), false},
	{NewCard(CLUBS, A), *NewCard(CLUBS, 8), NewCard(CLUBS, J), NewCard(SPADES, 8), true},
	{NewCard(DIAMONDS, A), *NewCard(DIAMONDS, 9), NewCard(HEARTS, A), NewCard(DIAMONDS, 10), true},
	{NewCard(CLUBS, 7), *NewCard(DIAMONDS, 7), NewCard(SPADES, 9), NewCard(DIAMONDS, J), false},
	{NewCard(SPADES, 10), *NewCard(HEARTS, 10), NewCard(DIAMONDS, 10), NewCard(HEARTS, A), false},
	{NewCard(SPADES, 8), *NewCard(HEARTS, Q), NewCard(DIAMONDS, K), NewCard(CLUBS, J), false},
	{NewCard(DIAMONDS, 9), *NewCard(DIAMONDS, J), NewCard(HEARTS, 9), NewCard(DIAMONDS, A), false},
	{NewCard(HEARTS, J), *NewCard(HEARTS, 8), NewCard(DIAMONDS, 8), NewCard(HEARTS, 7), true},
	{NewCard(DIAMONDS, Q), *NewCard(SPADES, A), NewCard(CLUBS, 7), NewCard(SPADES, 7), false},
	{NewCard(DIAMONDS, 7), *NewCard(CLUBS, 8), NewCard(SPADES, 9), NewCard(SPADES, K), false},
	{NewCard(CLUBS, K), *NewCard(SPADES, J), NewCard(CLUBS, 9), NewCard(SPADES, Q), true},
	{NewCard(CLUBS, A), *NewCard(CLUBS, Q), NewCard(HEARTS, K), NewCard(CLUBS, 10), true},
	{NewCard(HEARTS, 8), *NewCard(SPADES, K), NewCard(HEARTS, 7), NewCard(SPADES, 7), true},
}

var equalTestCases = []equalsTestCase{
	{*NewCard(SPADES, 7), *NewCard(CLUBS, 7), false},
	{*NewCard(SPADES, 8), *NewCard(CLUBS, 10), false},
	{*NewCard(SPADES, 9), *NewCard(SPADES, K), true},
	{*NewCard(SPADES, 9), *NewCard(SPADES, 11), true},
	{*NewCard(CLUBS, A), *NewCard(CLUBS, J), true},
	{*NewCard(CLUBS, 10), *NewCard(HEARTS, 8), false},
	{*NewCard(CLUBS, 11), *NewCard(HEARTS, 7), false},
	{*NewCard(CLUBS, A), *NewCard(CLUBS, 9), true},
	{*NewCard(HEARTS, 8), *NewCard(HEARTS, Q), true},
	{*NewCard(HEARTS, 9), *NewCard(SPADES, A), false},
	{*NewCard(HEARTS, K), *NewCard(CLUBS, 8), false},
	{*NewCard(HEARTS, J), *NewCard(DIAMONDS, 10), false},
	{*NewCard(DIAMONDS, 7), *NewCard(DIAMONDS, J), true},
	{*NewCard(DIAMONDS, 7), *NewCard(HEARTS, J), false},
	{*NewCard(DIAMONDS, 11), *NewCard(SPADES, 8), false},
	{*NewCard(DIAMONDS, 9), *NewCard(CLUBS, 10), false},
	{*NewCard(SPADES, 10), *NewCard(DIAMONDS, 11), false},
	{*NewCard(DIAMONDS, 8), *NewCard(DIAMONDS, Q), true},
	{*NewCard(HEARTS, Q), *NewCard(HEARTS, 7), true},
	{*NewCard(CLUBS, K), *NewCard(CLUBS, 11), true},
}

func testCompare(t *testing.T, card *Card, winningCard Card, trump *Card, first Card, expected bool) {
	if card.Compare(winningCard, trump, first) != expected {
		if expected {
			trumpStr := " (nil) "
			if trump != nil {
				trumpStr = trump.String()
			}
			t.Error("Given WinningCard: " + winningCard.String() + " , TrumpCard: " + trumpStr + " , FirstCard: " + first.String() + ". Card: " + card.String() + " is supposed to win.")
		} else {
			trumpStr := " (nil) "
			if trump != nil {
				trumpStr = trump.String()
			}
			t.Error("Given WinningCard: " + winningCard.String() + " , TrumpCard: " + trumpStr + " , FirstCard: " + first.String() + ". Card: " + card.String() + " is not supposed to win.")
		}
	}
}

func testEqual(t *testing.T, card Card, otherCard Card, expected bool) {
	if card.Equals(otherCard) != expected {
		if expected {
			t.Error("Card " + card.String() + " is supposed to be equal to " + otherCard.String())
		} else {
			t.Error("Card " + card.String() + " is not supposed to be equal to " + otherCard.String())
		}
	}
}

func TestCard_Compare(t *testing.T) {
	for _, c := range compareTestCases {
		testCompare(t, c.card, c.winningCard, c.trump, *c.first, c.expected)
	}
}

func TestCard_Equals(t *testing.T) {
	for _, c := range equalTestCases {
		testEqual(t, c.card, c.otherCard, c.expected)
	}
}
