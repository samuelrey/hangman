package main

import (
	"log"
	"math/rand"
	"net/http"
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
	http.HandleFunc("/new", handleNew)
	http.HandleFunc("/guess", handleGuess)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
