package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	startGuesses  = 6
	serverAddress = "localhost:1337"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error
	words, err := loadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	gameHandler := newSimpleGameHandler(words)

	log.Println("Starting server on", serverAddress)
	server := NewServer(serverAddress, gameHandler)
	server.Run()
}
