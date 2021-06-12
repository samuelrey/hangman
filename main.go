package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/samuelrey/hangman/application"
	"github.com/samuelrey/hangman/presentation"
)

const (
	startGuesses  = 6
	serverAddress = "0.0.0.0:14420"
	wordsFile     = "words.txt"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	wordHandler := application.NewSimpleWordHandler(wordsFile)
	gameHandler := application.NewSimpleGameHandler(startGuesses)

	log.Println("Starting server on", serverAddress)

	server := presentation.NewServer(serverAddress, gameHandler, wordHandler)
	server.Run()
}
