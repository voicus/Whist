package main

type PlayerCards struct {
	Player string `json:"username"`
	Cards  []Card `json:"cards"`
}

type GameDTO struct {
	Trump   Card           `json:"trump"`
	Players [4]PlayerCards `json:"players"`
}

type PlayerScore struct {
	Players [4]Player `json:"players"`
}
