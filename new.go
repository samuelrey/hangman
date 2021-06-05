package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/samuelrey/hangman/game"
)

func handleNew(w http.ResponseWriter, r *http.Request) {
	i := rand.Intn(len(words))
	word := words[i]

	game, err := game.NewGame(word, startGuesses)
	if err != nil {
		log.Println(err)
		return
	}

	gameHandler.register(game)

	resp := hangmanResponse{
		ID:               game.ID,
		Current:          game.Current,
		RemainingGuesses: game.RemainingGuesses,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
