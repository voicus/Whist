package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Game
// func newGame() *Game
//		creates a new game
//	func (game *Game) addPlayer(player Player)
// 		used to add a player to the current game
// func (game *Game) play()
// 		plays the game according to the rules

type Game struct {
	numberOfCards []int
	deckOfCards   Deck
	players       []Player
	name          string
}

func newGame() *Game {
	game := new(Game)
	game.numberOfCards = []int{8, 8, 8, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 8}
	game.deckOfCards = *NewDeck()
	return game
}

func (game *Game) addPlayer(player Player) {
	game.players = append(game.players, player)
}

func (game *Game) play() {
	fmt.Println("INCEPE JOCULLLLLLLLLLL")
	for _, elem := range game.numberOfCards {

		for j := 0; j < 4; j++ {
			game.players[j].tricks = 0
		}
		round := new(Round)
		game.deckOfCards.index = 0
		game.deckOfCards.ShuffleDeck()

		round.playRound(&game.players, &game.deckOfCards, elem, game.name)
		game.players = append(game.players[1:], game.players[0])
		fmt.Println("s-a terminat runda")
		var scores PlayerScore
		for j := 0; j < 4; j++ {
			scores.Players[j] = game.players[j]
		}

		jsonData, err := json.Marshal(scores)
		os.Stdout.Write(jsonData)
		fmt.Println("trimit scoruruile")
		command := map[string]interface{}{
			"method": "publish",
			"params": map[string]interface{}{
				"channel": game.name,
				"data": map[string]interface{}{
					"data": jsonData,
					"flag": "endRound",
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
		time.Sleep(5)
	}
	var maxScore int
	maxScore = -1
	for j := 0; j < 4; j++ {
		fmt.Println(game.players[j].Score)
		if game.players[j].Score > maxScore {
			maxScore = game.players[j].Score
		}
	}

	for j := 0; j < 4; j++ {
		if game.players[j].Score == maxScore {
			err := incrNumberOfGamesWon(User{Username: game.players[j].Name})
			if err != nil {
				return
			}
		} else {
			err := incrNumberOfGamesLost((User{Username: game.players[j].Name}))
			if err != nil {
				return
			}
		}
	}

	command := map[string]interface{}{
		"method": "publish",
		"params": map[string]interface{}{
			"channel": game.name,
			"data": map[string]interface{}{
				"maxScore": maxScore,
				"flag":     "endgame",
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
	time.Sleep(5)
}
