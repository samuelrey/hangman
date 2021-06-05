package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/guess", handleGuess)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}

type hangmangResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	RemainingGuesses int    `json:"guesses_remaining"`
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

	resp := hangmangResponse{
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

	resp := hangmangResponse{
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
