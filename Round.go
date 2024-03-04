package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"time"
)

// Round
// playRound
// plays the round according to the rules

type Round struct {
	sumOfBids int
	trump     Card
	first     Card
}

func (round *Round) playRound(players *[]Player, deck *Deck, numberOfCards int, gameID string) {
	fmt.Println("incepe runda pentru jocul ", gameID)
	round.sumOfBids = 0
	round.trump.Value = -1
	round.first.Value = -1
	indexMap := make(map[string]int)
	for i := 0; i < 4; i++ {
		indexMap[(*players)[i].Name] = i
		(*players)[i].GiveCards(deck.GiveCards(numberOfCards))
	}

	if numberOfCards < 8 {
		round.trump = deck.GiveCards(1)[0]
	}

	/// trimit trump + cartile fiecauria
	var gameDTO GameDTO
	gameDTO.Trump = round.trump
	for i := 0; i < 4; i++ {
		gameDTO.Players[i].Player = (*players)[i].Name
		gameDTO.Players[i].Cards = (*players)[i].cards
	}
	jsonData, err := json.Marshal(gameDTO)
	os.Stdout.Write(jsonData)

	command := map[string]interface{}{
		"method": "publish",
		"params": map[string]interface{}{
			"channel": gameID,
			"data": map[string]interface{}{
				"data": jsonData,
				"flag": "carti_joc",
			},
		},
	}

	dataA, err := json.Marshal(command)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8000/api", bytes.NewBuffer(dataA))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "apikey a3d9c270-52df-45f8-9a66-a1bb8e9e04ce")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	sum := 0

	fmt.Println("cer biduri")
	for i := 0; i < 3; i++ {
		fmt.Println("cer bid playerului ", (*players)[i].Name)
		gameMapMu.Lock()
		ch := gameMap[(*players)[i].Name+"bid"]
		gameMapMu.Unlock()
		(*players)[i].makeBid(false, 0, numberOfCards, gameID, ch)

		(*players)[i].tricks = 0
		sum += (*players)[i].getBid()
	}
	gameMapMu.Lock()
	ch := gameMap[(*players)[3].Name+"bid"]
	gameMapMu.Unlock()
	fmt.Println("cer bid playerului ", (*players)[3].Name)
	(*players)[3].makeBid(true, sum, numberOfCards, gameID, ch)
	(*players)[3].tricks = 0
	fmt.Println("incepe jucatul cartilor")
	for i := 0; i < numberOfCards; i++ {
		time.Sleep(1 * time.Second)
		var winningCard Card
		var winningPlayer *Player
		winningPlayer = nil
		isFirst := 1
		for i := 0; i < 4; i++ {
			var played Card
			fmt.Println("Cer carte jucatorului", (*players)[i].Name)
			if isFirst == 1 {
				gameMapMu.Lock()
				ch := gameMap[(*players)[i].Name+"card"]
				gameMapMu.Unlock()
				if round.trump.Value == -1 {
					played = (*players)[i].playCard(nil, nil, gameID, ch)
				} else {
					played = (*players)[i].playCard(nil, &round.trump, gameID, ch)
				}
				round.first = played
				winningCard = played
				winningPlayer = &(*players)[i]
				isFirst = 0
			} else {
				gameMapMu.Lock()
				ch := gameMap[(*players)[i].Name+"card"]
				gameMapMu.Unlock()
				played = (*players)[i].playCard(&round.first, &round.trump, gameID, ch)
			}
			var trumpCard *Card
			if round.trump.Value == -1 {
				trumpCard = nil
			} else {
				trumpCard = &round.trump
			}
			if isFirst == 0 && played.Compare(winningCard, trumpCard, round.first) {
				winningCard = played
				winningPlayer = &(*players)[i]
			}
		}

		command := map[string]interface{}{
			"method": "publish",
			"params": map[string]interface{}{
				"channel": gameID,
				"data": map[string]interface{}{
					"flag": "endHand",
				},
			},
		}

		dataA, err := json.Marshal(command)
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", "http://localhost:8000/api", bytes.NewBuffer(dataA))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "apikey a3d9c270-52df-45f8-9a66-a1bb8e9e04ce")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		winningPlayer.tricks++
		command1 := map[string]interface{}{
			"method": "publish",
			"params": map[string]interface{}{
				"channel": gameID,
				"data": map[string]interface{}{
					"flag":             "tricks",
					(*players)[0].Name: (*players)[0].tricks,
					(*players)[1].Name: (*players)[1].tricks,
					(*players)[2].Name: (*players)[2].tricks,
					(*players)[3].Name: (*players)[3].tricks,
				},
			},
		}

		dataA1, err1 := json.Marshal(command1)
		if err != nil {
			panic(err)
		}
		req1, err1 := http.NewRequest("POST", "http://localhost:8000/api", bytes.NewBuffer(dataA1))
		if err != nil {
			panic(err)
		}
		req1.Header.Set("Content-Type", "application/json")
		req1.Header.Set("Authorization", "apikey a3d9c270-52df-45f8-9a66-a1bb8e9e04ce")
		client1 := http.Client{}
		resp1, err1 := client1.Do(req1)
		if err1 != nil {
			panic(err)
		}
		defer resp1.Body.Close()
		var j int
		j = -1
		for i := 0; i < 4; i++ {
			if (*players)[i].Name == winningPlayer.Name {
				j = i
			}
		}
		*players = append((*players)[j:], (*players)[:j]...)
	}

	for i := 0; i < 4; i++ {
		if (*players)[i].tricks == (*players)[i].bid {
			(*players)[i].Score += 10 + (*players)[i].bid
		} else {
			(*players)[i].Score -= int(math.Abs(float64((*players)[i].tricks - (*players)[i].bid)))
		}
	}

	sort.Slice((*players), func(i, j int) bool {
		return indexMap[(*players)[i].Name] < indexMap[(*players)[j].Name]
	})
}
