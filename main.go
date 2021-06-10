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

var (
	gameHandler *SimpleGameHandler
	words       []string
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gameHandler = newSimpleGameHandler()

	var err error
	words, err = loadWords("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on", serverAddress)
	server := NewServer(serverAddress, gameHandler)
	server.Run()
}
