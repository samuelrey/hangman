package main

import "github.com/samuelrey/hangman/game"

type GameHandler struct {
	gameMap map[string]*game.Game
}

func newGameHandler() *GameHandler {
	return &GameHandler{make(map[string]*game.Game)}
}

func (gameHandler *GameHandler) get(id string) (*game.Game, bool) {
	game, found := gameHandler.gameMap[id]
	return game, found
}

func (gameHandler *GameHandler) register(game *game.Game) {
	gameHandler.gameMap[game.ID] = game
}

func (gameHandler *GameHandler) delete(id string) {
	delete(gameHandler.gameMap, id)
}
