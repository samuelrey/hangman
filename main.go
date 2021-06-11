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

	var words []string
	gameHandler := application.NewSimpleGameHandler(words)

	log.Println("Starting server on", serverAddress)

	server := NewServer(serverAddress, gameHandler)
	server.Run()
}
