package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/samuelrey/hangman/game"
)

type hangmanResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	RemainingGuesses int    `json:"guesses_remaining"`
}

type guessRequest struct {
	ID     string `json:"id"`
	Letter string `json:"guess"`
}

func handleGuess(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var guess guessRequest
	err = json.Unmarshal(body, &guess)
	if err != nil {
		log.Println(err)
		return
	}

	game, found := gameHandler.get(guess.ID)
	if !found {
		return
	}

	game.Guess(guess.Letter)
	fmt.Printf("%+v", game)

	if game.Loss() || game.Won() {
		gameHandler.delete(game.ID)
	}

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
