package main

import (
	"sort"
	"strings"
	"testing"
)

var decks = []*Deck{
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
	NewDeck(),
}

type giveCardsTestCase struct {
	request []int
}

var giveCardsTestCases = []giveCardsTestCase{
	{request: []int{8, 8, 8, 8}},
	{request: []int{7, 7, 7, 7}},
	{request: []int{6, 6, 6, 6}},
	{request: []int{5, 5, 5, 5}},
	{request: []int{4, 4, 4, 4}},
	{request: []int{2, 2, 2, 2}},
	{request: []int{1, 1, 1, 1}},
	{request: []int{3, 4, 2, 5}},
	{request: []int{9, 8, 7, 6}},
	{request: []int{2, 4, 1, 3}},
	{request: []int{12, 8, 3, 9}},
}

func TestDeck_ShuffleDeck(t *testing.T) {
	for _, d := range decks {
		d.ShuffleDeck()
	}
	for i, d1 := range decks {
		for j, d2 := range decks {
			if i != j && d1.Equals(*d2) {
				t.Error("Very small chance of getting the same decks twice!")
			}
		}
	}

}

type Cards []Card

func (p Cards) Len() int           { return len(p) }
func (p Cards) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p Cards) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func testEq(a Cards, b Cards) bool {
	sort.Sort(a)
	sort.Sort(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if (a[i].Value != b[i].Value) || (a[i].Suite != b[i].Suite) {
			return false
		}
	}
	return true
}

func getString(a []Card) string {
	strArr := make([]string, len(a))
	for i, v := range a {
		strArr[i] = v.String()
	}

	return strings.Join(strArr, ", ")
}

func TestDeck_GiveCards(t *testing.T) {
	for _, c := range giveCardsTestCases {
		var deck = NewDeck()
		deck.ShuffleDeck()
		expected1 := make([]Card, c.request[0])
		expected2 := make([]Card, c.request[1])
		expected3 := make([]Card, c.request[2])
		expected4 := make([]Card, c.request[3])

		var index1 = 0
		var index2 = index1 + c.request[0]
		var index3 = index2 + c.request[1]
		var index4 = index3 + c.request[2]
		copy(expected1, deck.cards[index1:index2])
		copy(expected2, deck.cards[index2:index3])
		copy(expected3, deck.cards[index3:index4])
		copy(expected4, deck.cards[index4:index4+c.request[3]])
		var requestedCards1 = deck.GiveCards(c.request[0])
		var requestedCards2 = deck.GiveCards(c.request[1])
		var requestedCards3 = deck.GiveCards(c.request[2])
		var requestedCards4 = deck.GiveCards(c.request[3])
		if !testEq(expected1, requestedCards1) {
			t.Error("Expected: " + getString(expected1) + " Got: " + getString(requestedCards1))
		}
		if !testEq(expected2, requestedCards2) {
			t.Error("Expected: " + getString(expected2) + " Got: " + getString(requestedCards2))
		}
		if !testEq(expected3, requestedCards3) {
			t.Error("Expected: " + getString(expected3) + " Got: " + getString(requestedCards3))
		}
		if !testEq(expected4, requestedCards4) {
			t.Error("Expected: " + getString(expected4) + " Got: " + getString(requestedCards4))
		}
	}
}
