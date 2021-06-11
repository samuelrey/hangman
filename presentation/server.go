package presentation

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
	gameHandler *application.GameHandler
	wordHandler *application.WordHandler
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

func NewServer(
	address string,
	gameHandler application.GameHandler,
	wordHandler application.WordHandler,
) *Server {
	return &Server{
		address:     address,
		gameHandler: &gameHandler,
		wordHandler: &wordHandler,
	}
}

type hangmanResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	RemainingGuesses int    `json:"guesses_remaining"`
}

type guessRequest struct {
	ID    string `json:"id"`
	Guess string `json:"guess"`
}

func (s *Server) handleGuess(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var req guessRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO define response for game not found.
	game, found := (*s.gameHandler).Get(req.ID)
	if !found {
		return
	}

	game.Guess(req.Guess)

	// TODO define response for win/loss cases.
	if game.Loss() || game.Won() {
		(*s.gameHandler).Delete(game.ID)
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
	word := (*s.wordHandler).RandWord()

	game, err := game.NewGame(word, (*s.gameHandler).StartGuesses())
	if err != nil {
		log.Println(err)
		return
	}

	(*s.gameHandler).Register(game)

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
