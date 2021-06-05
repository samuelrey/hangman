# Fleet Hangman challenge

This challenge requires you to implement a small API for the guessing game Hangman. If you are unfamiliar with the game, please have a look at [Hangman (game) on Wikipedia](https://en.wikipedia.org/wiki/Hangman_(game)).

### Background

Please spend a maximum of **2 hours** on this coding challenge. We are interested in understanding the strategies you use and tradeoffs you make developing during this short time period. Keep in mind that we will read your code and some comments can be helpful for understanding.

Though the problem domain does not directly represent the work that engineers at Fleet perform on a day-to-day basis, as a small and growing company we constantly evaluate tradeoffs and manage risk in our development process.

After your submission, a Fleet engineer will have a conversation with you about the strategies you chose, tradeoffs you made, and technical details of your implementation. We will use this as a jumping off point to understanding more about your engineering background and development process.

In the interest of _respecting your time_, _fairness to all candidates_, and _setting expectations appropriately upon your joining the team_, please stick to the time limit. We ask that you spend no more than **2 hours** working on your submission.

## Challenge


### For frontend engineers

Please implement a React single page web application that provides a user interface for the Hangman game using the API specified below.  Please make the game responsive and usable down to 320px width.  Don't worry about optimizing for touch events; imagine your audience will be pointing and clicking with a mouse cursor.

Your game UI should talk to `http://localhost:1337` at appropriate times given the API documentation below.  Since the API isn't actually implemented yet, you can ignore error responses and mock your own data.  But please be sure your code is actually sending requests.

Please get as far as you can in the time limit, and focus on building a minimally viable solution that is a working game, while making it as attractive and usable as you have time for.

### For backend engineers

Please implement the API as specified below.

You may use the provided `main.go` file as the start of your implementation. This file loads the word list and starts an HTTP listener. The `generateIdentifier()` function in `identifier.go` may be used for generating game IDs.

The server can be started with `go run .` and is configured to listen on [http://localhost:1337](http://localhost:1337).


## API

Below, the API endpoints are described:

### New game

Start a new game.

#### Request

```
POST /new
```

Request body is ignored.

#### Response

The response contains the initial state for the newly started game.

* `id` - Identifier for this game. Use this to make guesses in the game.
* `current` - The current board state. Always consists of only `_` characters at start of game.
* `guesses_remaining` - How many guesses remain before the player loses.

#### Example

``` sh
curl -X POST https://fleet-hangman.herokuapp.com/new
{"id":"f8302916-69f1-462b-b640-e503faa94397","current":"________","guesses_remaining":6}
```

### Make guess

Guess a letter in an ongoing game. Any game that is completed (when all letters are guessed, or no guesses remain) can be cleared from the data store.

#### Request

```
POST /guess
{"id":"<game_id>","guess":"<[A-Z]>"}
```

* `id` - Identifier for this game. Get a new game identifier from the `/new` endpoint.
* `guess` - The character to guess. Must be a single ASCII character (A-Z). 

#### Response

The response contains the updated game state.

* `id` - Identifier for this game. This will be unchanged from the request.
* `current` - The current board state.
* `guesses_remaining` - How many guesses remain before the player loses.

#### Example

``` sh
 curl -X POST https://fleet-hangman.herokuapp.com/guess -d '{"id":"f8302916-69f1-462b-b640-e503faa94397","guess":"A"}'
{"id":"f8302916-69f1-462b-b640-e503faa94397","current":"______A_","guesses_remaining":6}
```
