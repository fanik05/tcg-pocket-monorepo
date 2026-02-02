package main

import "fmt"

type Player struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Points int `json:"points"`
	Opponents []string `json:"opponents"` // ids of players faced
}

type Match struct {
	ID string `json:"id"`
	Round int `json:"round"` // TODO: change to string for swiss rounds?
	Player1ID string `json:"player1_id"`
	Player2ID string `json:"player2_id"`
	WinnerID string `json:"winner_id"`
}

func main() {
	fmt.Println("TCG Pocket Tournament Engine: Booting up...")
}