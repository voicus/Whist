package main

import (
	"testing"
)

type hasSuiteTestCase struct {
	playerCards []Card
	card        Card
	expected    bool
}

type validCardsTestCase struct {
	playerCards []Card
	first       *Card
	trump       *Card
	expected    []Card
}

var hasSuiteTestCases = []hasSuiteTestCase{
	{[]Card{*NewCard(SPADES, 8), *NewCard(DIAMONDS, 11), *NewCard(DIAMONDS, A), *NewCard(HEARTS, K)}, *NewCard(SPADES, 7), true},
	{[]Card{*NewCard(HEARTS, 8), *NewCard(CLUBS, 11), *NewCard(CLUBS, A)}, *NewCard(SPADES, 7), false},
	{[]Card{*NewCard(SPADES, 8)}, *NewCard(SPADES, 7), true},
	{[]Card{*NewCard(SPADES, 8)}, *NewCard(DIAMONDS, 7), false},
	{[]Card{*NewCard(CLUBS, 8), *NewCard(DIAMONDS, A), *NewCard(DIAMONDS, K)}, *NewCard(DIAMONDS, 7), true},
	{[]Card{*NewCard(HEARTS, K)}, *NewCard(HEARTS, 7), true},
	{[]Card{*NewCard(CLUBS, K)}, *NewCard(HEARTS, 7), false},
	{[]Card{*NewCard(DIAMONDS, 11), *NewCard(CLUBS, A), *NewCard(HEARTS, K)}, *NewCard(SPADES, 7), false},
	{[]Card{*NewCard(DIAMONDS, A), *NewCard(HEARTS, K)}, *NewCard(DIAMONDS, 7), true},
	{[]Card{*NewCard(HEARTS, K)}, *NewCard(SPADES, 7), false},
	{[]Card{*NewCard(SPADES, 8), *NewCard(CLUBS, 11), *NewCard(DIAMONDS, A), *NewCard(HEARTS, K)}, *NewCard(SPADES, 7), true},
}

var validCardsTestCases = []validCardsTestCase{
	{[]Card{*NewCard(DIAMONDS, K), *NewCard(HEARTS, A), *NewCard(HEARTS, K)}, NewCard(CLUBS, Q), NewCard(CLUBS, 10), []Card{*NewCard(DIAMONDS, K), *NewCard(HEARTS, A), *NewCard(HEARTS, K)}},
	{[]Card{*NewCard(CLUBS, 7), *NewCard(SPADES, A), *NewCard(HEARTS, 9), *NewCard(CLUBS, 8)}, NewCard(SPADES, 8), NewCard(DIAMONDS, 8), []Card{*NewCard(SPADES, A)}},
	{[]Card{*NewCard(HEARTS, 10)}, NewCard(DIAMONDS, A), NewCard(SPADES, K), []Card{*NewCard(HEARTS, 10)}},
	{[]Card{*NewCard(DIAMONDS, J), *NewCard(HEARTS, 7), *NewCard(DIAMONDS, 7), *NewCard(HEARTS, Q), *NewCard(HEARTS, 8), *NewCard(CLUBS, 9)}, NewCard(SPADES, J), NewCard(DIAMONDS, 10), []Card{*NewCard(DIAMONDS, 7), *NewCard(DIAMONDS, J)}},
	{[]Card{*NewCard(CLUBS, Q), *NewCard(HEARTS, A), *NewCard(CLUBS, 10), *NewCard(CLUBS, A), *NewCard(DIAMONDS, Q), *NewCard(DIAMONDS, K), *NewCard(CLUBS, 8), *NewCard(SPADES, J)}, NewCard(CLUBS, K), NewCard(HEARTS, 8), []Card{*NewCard(CLUBS, Q), *NewCard(CLUBS, 10), *NewCard(CLUBS, A), *NewCard(CLUBS, 8)}},
	{[]Card{*NewCard(DIAMONDS, A), *NewCard(DIAMONDS, 9), *NewCard(SPADES, 9), *NewCard(SPADES, 7), *NewCard(SPADES, Q), *NewCard(DIAMONDS, 10), *NewCard(DIAMONDS, 7), *NewCard(HEARTS, J)}, NewCard(HEARTS, 9), NewCard(CLUBS, 9), []Card{*NewCard(HEARTS, J)}},
	{[]Card{*NewCard(SPADES, 10)}, NewCard(CLUBS, J), NewCard(DIAMONDS, 8), []Card{*NewCard(SPADES, 10)}},
	{[]Card{*NewCard(SPADES, 8), *NewCard(HEARTS, 7), *NewCard(CLUBS, 7)}, NewCard(HEARTS, Q), NewCard(HEARTS, K), []Card{*NewCard(HEARTS, 7)}},
	{[]Card{*NewCard(HEARTS, 10), *NewCard(DIAMONDS, J)}, NewCard(SPADES, Q), NewCard(SPADES, K), []Card{*NewCard(HEARTS, 10), *NewCard(DIAMONDS, J)}},
	{[]Card{*NewCard(SPADES, 8), *NewCard(HEARTS, K), *NewCard(CLUBS, 9), *NewCard(HEARTS, Q), *NewCard(HEARTS, 10), *NewCard(SPADES, K), *NewCard(SPADES, A)}, NewCard(DIAMONDS, J), NewCard(HEARTS, A), []Card{*NewCard(HEARTS, Q), *NewCard(HEARTS, 10), *NewCard(HEARTS, K)}},
	{[]Card{*NewCard(HEARTS, 10), *NewCard(CLUBS, 10), *NewCard(HEARTS, 7), *NewCard(CLUBS, K), *NewCard(CLUBS, A), *NewCard(HEARTS, J), *NewCard(CLUBS, 7)}, NewCard(CLUBS, Q), NewCard(HEARTS, A), []Card{*NewCard(CLUBS, 10), *NewCard(CLUBS, K), *NewCard(CLUBS, A), *NewCard(CLUBS, 7)}},
	{[]Card{*NewCard(SPADES, 10)}, NewCard(DIAMONDS, J), NewCard(HEARTS, Q), []Card{*NewCard(SPADES, 10)}},
	{[]Card{*NewCard(SPADES, J)}, NewCard(SPADES, 9), NewCard(SPADES, 7), []Card{*NewCard(SPADES, J)}},
	{[]Card{*NewCard(DIAMONDS, 7), *NewCard(CLUBS, 8), *NewCard(DIAMONDS, Q), *NewCard(SPADES, Q), *NewCard(SPADES, 8), *NewCard(HEARTS, K)}, NewCard(HEARTS, 8), NewCard(SPADES, A), []Card{*NewCard(HEARTS, K)}},
	{[]Card{*NewCard(DIAMONDS, 9), *NewCard(DIAMONDS, 10), *NewCard(DIAMONDS, 8), *NewCard(CLUBS, J), *NewCard(SPADES, A)}, NewCard(SPADES, K), NewCard(DIAMONDS, K), []Card{*NewCard(SPADES, A)}},
	{[]Card{*NewCard(HEARTS, K), *NewCard(DIAMONDS, 7), *NewCard(SPADES, J), *NewCard(SPADES, 9), *NewCard(DIAMONDS, A), *NewCard(SPADES, A), *NewCard(HEARTS, 7)}, NewCard(CLUBS, 9), nil, []Card{*NewCard(HEARTS, K), *NewCard(DIAMONDS, 7), *NewCard(SPADES, J), *NewCard(SPADES, 9), *NewCard(DIAMONDS, A), *NewCard(SPADES, A), *NewCard(HEARTS, 7)}},
	{[]Card{*NewCard(HEARTS, 9), *NewCard(DIAMONDS, J), *NewCard(SPADES, 10), *NewCard(DIAMONDS, 8)}, NewCard(DIAMONDS, 10), nil, []Card{*NewCard(DIAMONDS, J), *NewCard(DIAMONDS, 8)}},
	{[]Card{*NewCard(DIAMONDS, A), *NewCard(HEARTS, 7), *NewCard(SPADES, 9)}, NewCard(SPADES, A), nil, []Card{*NewCard(SPADES, 9)}},
	{[]Card{*NewCard(CLUBS, 7), *NewCard(DIAMONDS, J), *NewCard(HEARTS, J), *NewCard(DIAMONDS, Q), *NewCard(DIAMONDS, A), *NewCard(SPADES, 9), *NewCard(HEARTS, Q), *NewCard(HEARTS, 10)}, NewCard(SPADES, Q), nil, []Card{*NewCard(SPADES, 9)}},
	{[]Card{*NewCard(CLUBS, Q), *NewCard(HEARTS, 8), *NewCard(SPADES, J), *NewCard(DIAMONDS, K)}, nil, NewCard(HEARTS, K), []Card{*NewCard(HEARTS, 8)}},
	{[]Card{*NewCard(HEARTS, Q), *NewCard(HEARTS, 10), *NewCard(SPADES, 9)}, nil, NewCard(DIAMONDS, Q), []Card{*NewCard(HEARTS, Q), *NewCard(HEARTS, 10), *NewCard(SPADES, 9)}},
	{[]Card{*NewCard(SPADES, J), *NewCard(CLUBS, A), *NewCard(HEARTS, J), *NewCard(HEARTS, Q)}, nil, nil, []Card{*NewCard(SPADES, J), *NewCard(CLUBS, A), *NewCard(HEARTS, J), *NewCard(HEARTS, Q)}},
	{[]Card{*NewCard(HEARTS, 10), *NewCard(SPADES, 9), *NewCard(HEARTS, 7)}, nil, nil, []Card{*NewCard(HEARTS, 10), *NewCard(SPADES, 9), *NewCard(HEARTS, 7)}},
}

func TestPlayer_GetValidCards(t *testing.T) {
	player := Player{}
	for _, c := range validCardsTestCases {
		player.GiveCards(c.playerCards)
		if !testEq(player.GetValidCards(c.first, c.trump), c.expected) {
			t.Error("Got: " + getString(player.GetValidCards(c.first, c.trump)) + " expected: " + getString(c.expected))
		}
	}
}

func TestPlayer_HasSuite(t *testing.T) {
	player := Player{}
	for _, c := range hasSuiteTestCases {
		player.GiveCards(c.playerCards)
		if player.HasSuite(c.card) != c.expected {
			if c.expected {
				t.Error("Given cards: " + getString(c.playerCards) + " is supposed to have suite same as " + c.card.String())
			} else {
				t.Error("Given cards: " + getString(c.playerCards) + " is not supposed to have suite same as " + c.card.String())

			}
		}
	}
}
