package main

import "github.com/pkg/errors"

type Game struct {
	ID               string
	Current          string
	Word             string
	IndexByLetter    map[string][]int
	RemainingGuesses int
}

func newGame(word string, remainingGuesses int) (*Game, error) {
	id, err := generateIdentifier()
	if err != nil {
		return nil, errors.Wrap(err, "create new game")
	}

	indexByLetter := make(map[string][]int)
	for index, r := range word {
		str := string(r)
		indexByLetter[str] = append(indexByLetter[str], index)
	}

	var current string
	for i := 0; i < len(word); i++ {
		current = current + "_"
	}

	return &Game{
		ID:               id,
		Word:             word,
		Current:          current,
		IndexByLetter:    indexByLetter,
		RemainingGuesses: remainingGuesses,
	}, nil
}
