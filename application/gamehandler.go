package application

import (
	"github.com/samuelrey/hangman/game"
)

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

func NewSimpleGameHandler() *SimpleGameHandler {
	return &SimpleGameHandler{
		gameMap: make(map[string]*game.Game),
	}
}

func (gh *SimpleGameHandler) Get(id string) (*game.Game, bool) {
	game, found := gh.gameMap[id]
	return game, found
}

func (gh *SimpleGameHandler) Register(g *game.Game) {
	gh.gameMap[g.ID] = g
}

func (gh *SimpleGameHandler) Delete(id string) {
	delete(gh.gameMap, id)
}
