package game

import "github.com/pkg/errors"

type Game struct {
	ID               string
	Current          string
	Word             string
	IndexByLetter    map[string][]int
	RemainingGuesses int
}

func NewGame(word string, remainingGuesses int) (*Game, error) {
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

// Guess checks whether the given letter is part of the secret word. If it is,
// then we replace the underscores with the given letter to display to the user.
// Otherwise we reduce the number of remaining guesses.
//
// Note that if the user guesses a letter correctly then guesses that same
// letter again, we consider the second guess incorrect.
func (game *Game) Guess(letter string) {
	if indexes, ok := game.IndexByLetter[letter]; ok {
		// replace underscores with the correctly guessed letter
		for _, i := range indexes {
			game.Current = game.Current[:i] + letter + game.Current[i+1:]
		}

		delete(game.IndexByLetter, letter)
	} else {
		game.RemainingGuesses--
	}
}

// Loss will return true if the user has no more guesses and there are still
// letters that have not been guessed.
func (game *Game) Loss() bool {
	return game.RemainingGuesses < 1 && len(game.IndexByLetter) > 0
}

// Won will return true if the user has guessed each of the letters within the
// limited number of guesses provided.
func (game *Game) Won() bool {
	return game.RemainingGuesses >= 0 && len(game.IndexByLetter) == 0
}
