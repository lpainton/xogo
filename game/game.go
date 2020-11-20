// Package xogo implements game logic for xogo
package xogo

import "encoding/json"

type mark = int

const (
	none mark = iota
	x    mark = iota
	o    mark = iota
)

type gridSquare = int

const (
	// TopLeft is the index of the top left square in the grid
	TopLeft = iota
	// TopMid is the index of the top middle square in the grid
	TopMid = iota
	// TopRight is the index of the top right square in the grid
	TopRight = iota
	// MidLeft is the index of the middle left square in the grid
	MidLeft = iota
	// Center is the index of the center square in the grid
	Center = iota
	// MidRight is the index of the middle right square in the grid
	MidRight = iota
	// BotLeft is the index of the bottom left square in the grid
	BotLeft = iota
	// BotMid is the index of the bottom middle square in the grid
	BotMid = iota
	// BotRight is the index of the bottom right square in the grid
	BotRight = iota
)

// Game represents the state of the game world
type Game struct {
	Grid   [9]mark `json:"gameGrid"`
	Next   mark    `json:"nextMove"`
	Winner mark    `json:"gameWinner"`
}

// New creates a pristine game board
func New() *Game {
	return &Game{Next: x}
}

// FromJSON creates a game from a provided json string
func FromJSON(jsonString []byte) (game *Game, err error) {
	err = json.Unmarshal(jsonString, game)
	if err != nil {
		return nil, err
	}
	return
}

// ToJSON returns a json formatted string representation of a game
func (g *Game) ToJSON() (jsonString []byte, err error) {
	jsonString, err = json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return
}

// Move takes a gridSquare input a generates a new board which reflects the move
func (g Game) Move(square gridSquare) (game Game, err error) {

	// If the square is already marked we change nothing
	if game.Grid[square] != none {
		return g, nil
	}
	return g, nil
}
