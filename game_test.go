package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct{ suite.Suite }

func (suite *GameTestSuite) TestNewGame() {
	game, err := newGame("paradiddle", 6)

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

func TestGame(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}
