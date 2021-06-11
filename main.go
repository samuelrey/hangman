package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/samuelrey/hangman/application"
)

const (
	startGuesses  = 6
	serverAddress = "localhost:1337"
	wordsFile     = "words.txt"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	wordHandler := application.NewSimpleWordHandler(wordsFile)
	gameHandler := application.NewSimpleGameHandler()

	log.Println("Starting server on", serverAddress)

	server := NewServer(serverAddress, gameHandler, wordHandler)
	server.Run()
}
