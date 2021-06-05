package game

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct{ suite.Suite }

func (suite *GameTestSuite) TestNewGame() {
	game, err := NewGame("paradiddle", 6)

	suite.Require().NoError(err)
	suite.Require().Equal("paradiddle", game.Word)
	suite.Require().Equal("__________", game.Current)
	suite.Require().Equal(6, game.RemainingGuesses)

	indexByLetter := map[string][]int{
		"p": {0},
		"a": {1, 3},
		"r": {2},
		"d": {4, 6, 7},
		"i": {5},
		"l": {8},
		"e": {9},
	}
	suite.Require().Equal(indexByLetter, game.IndexByLetter)
}

func (suite *GameTestSuite) TestGuessRight() {
	game, err := NewGame("paradiddle", 6)
	suite.Require().NoError(err)

	game.Guess("d")

	suite.Require().Equal("____d_dd__", game.Current)
	suite.Require().Equal(6, game.RemainingGuesses)

	indexByLetter := map[string][]int{
		"p": {0},
		"a": {1, 3},
		"r": {2},
		"i": {5},
		"l": {8},
		"e": {9},
	}
	suite.Require().Equal(indexByLetter, game.IndexByLetter)
}

func (suite *GameTestSuite) TestGuessWrong() {
	game, err := NewGame("paradiddle", 6)
	suite.Require().NoError(err)

	game.Guess("z")

	suite.Require().Equal("__________", game.Current)
	suite.Require().Equal(5, game.RemainingGuesses)

	indexByLetter := map[string][]int{
		"p": {0},
		"a": {1, 3},
		"r": {2},
		"d": {4, 6, 7},
		"i": {5},
		"l": {8},
		"e": {9},
	}
	suite.Require().Equal(indexByLetter, game.IndexByLetter)
}

// Test that we reduce the number of remaining guesses if a correct letter is
// given more than once.
func (suite *GameTestSuite) TestGuessLetterTwice() {
	game, err := NewGame("paradiddle", 6)
	suite.Require().NoError(err)

	game.Guess("d")
	game.Guess("d")

	suite.Require().Equal("____d_dd__", game.Current)
	suite.Require().Equal(5, game.RemainingGuesses)

	indexByLetter := map[string][]int{
		"p": {0},
		"a": {1, 3},
		"r": {2},
		"i": {5},
		"l": {8},
		"e": {9},
	}
	suite.Require().Equal(indexByLetter, game.IndexByLetter)
}

func TestGame(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}
