package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	game.guess(guess.Letter)
	fmt.Printf("%+v", game)

	if game.loss() || game.won() {
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
