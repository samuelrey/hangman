package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/samuelrey/hangman/application"
	"github.com/samuelrey/hangman/game"
)

type Server struct {
	address     string
	gameHandler application.GameHandler
}

func (s *Server) handle() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/new", s.handleNew)
	mux.HandleFunc("/guess", s.handleGuess)
	return mux
}

func (s *Server) Run() {
	http.ListenAndServe(s.address, s.handle())
}

func NewServer(address string, gameHandler application.GameHandler) *Server {
	return &Server{
		address:     address,
		gameHandler: gameHandler,
	}
}

type hangmanResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	RemainingGuesses int    `json:"guesses_remaining"`
}

type guessRequest struct {
	ID     string `json:"id"`
	Letter string `json:"guess"`
}

func (s *Server) handleGuess(w http.ResponseWriter, r *http.Request) {
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

	game, found := s.gameHandler.Get(guess.ID)
	if !found {
		return
	}

	game.Guess(guess.Letter)

	if game.Loss() || game.Won() {
		s.gameHandler.Delete(game.ID)
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

func (s *Server) handleNew(w http.ResponseWriter, r *http.Request) {
	word := s.gameHandler.RandWord()

	game, err := game.NewGame(word, startGuesses)
	if err != nil {
		log.Println(err)
		return
	}

	s.gameHandler.Register(game)

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
