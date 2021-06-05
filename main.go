package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	startGuesses  = 6
	serverAddress = "localhost:1337"
)

var (
	gameHandler *GameHandler
	words       []string
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gameHandler = newGameHandler()

	var err error
	words, err = loadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on", serverAddress)
	http.HandleFunc("/new", handleNew)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}

type newResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	GuessesRemaining int    `json:"guesses_remaining"`
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(words))
	word := words[i]

	game, err := newGame(word, startGuesses)
	if err != nil {
		log.Println(err)
		return
	}

	gameHandler.register(game)

	resp := newResponse{
		ID:               game.ID,
		Current:          game.Current,
		GuessesRemaining: game.RemainingGuesses,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
