package main

type GameHandler struct {
	gameMap map[string]*Game
}

func newGameHandler() *GameHandler {
	return &GameHandler{make(map[string]*Game)}
}

func (gameHandler *GameHandler) get(id string) (*Game, bool) {
	game, found := gameHandler.gameMap[id]
	return game, found
}

func (gameHandler *GameHandler) register(game *Game) {
	gameHandler.gameMap[game.ID] = game
}

func (gameHandler *GameHandler) delete(id string) {
	delete(gameHandler.gameMap, id)
}
