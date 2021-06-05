package main

type hangmanResponse struct {
	ID               string `json:"id"`
	Current          string `json:"current"`
	RemainingGuesses int    `json:"guesses_remaining"`
}
