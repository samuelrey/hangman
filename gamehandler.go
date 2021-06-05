package main

import "github.com/samuelrey/hangman/game"

// GameHandler defines the set of behavior we need to implement to manage games.
// This interface could be implemented for a database or an external storage
// service.
type GameHandler interface {
	Get(id string) (*game.Game, bool)
	Register(g *game.Game)
	Delete(id string)
}

// SimpleGameHandler implements the GameHandler behavior using a map.
type SimpleGameHandler struct {
	gameMap map[string]*game.Game
}

func newSimpleGameHandler() *SimpleGameHandler {
	return &SimpleGameHandler{make(map[string]*game.Game)}
}

func (gameHandler *SimpleGameHandler) Get(id string) (*game.Game, bool) {
	game, found := gameHandler.gameMap[id]
	return game, found
}

func (gameHandler *SimpleGameHandler) Register(g *game.Game) {
	gameHandler.gameMap[g.ID] = g
}

func (gameHandler *SimpleGameHandler) Delete(id string) {
	delete(gameHandler.gameMap, id)
}
