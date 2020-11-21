// Package game implements game logic for xogo
package game

import "fmt"

type mark = int

const (
	// None represents no marking on a square
	None mark = iota
	// X represents an 'X' marking on a square
	X mark = iota
	// O represents an 'O' marking on a square
	O mark = iota
)

// GridRef indicates a square on the grid surface
type GridRef = int

const (
	// TopLeft is the index of the top left square in the grid
	TopLeft GridRef = iota
	// TopMid is the index of the top middle square in the grid
	TopMid GridRef = iota
	// TopRight is the index of the top right square in the grid
	TopRight GridRef = iota
	// MidLeft is the index of the middle left square in the grid
	MidLeft GridRef = iota
	// Center is the index of the center square in the grid
	Center GridRef = iota
	// MidRight is the index of the middle right square in the grid
	MidRight GridRef = iota
	// BotLeft is the index of the bottom left square in the grid
	BotLeft GridRef = iota
	// BotMid is the index of the bottom middle square in the grid
	BotMid GridRef = iota
	// BotRight is the index of the bottom right square in the grid
	BotRight GridRef = iota
)

// Game represents the state of the game world
type Game struct {
	Grid   [9]mark `json:"gameGrid"`
	Next   mark    `json:"nextUp"`
	Turn   mark    `json:"whoseTurn"`
	Winner mark    `json:"winner"`
	Pretty string  `json:"pretty"`
}

// New creates a pristine game board
func New() *Game {
	g := &Game{Turn: X, Next: O}
	g.pretty()
	return g
}

// Mark takes a GridRef and marks it with the next marker
func (g *Game) Mark(square GridRef) error {

	// If the square is already marked we change nothing
	if g.Grid[square] != None {
		return nil
	}

	g.Grid[square] = g.Turn
	g.pretty()

	if won(g.Grid) {
		g.Winner = g.Turn
		g.Turn = None
		g.Next = None
		return nil
	}

	t := g.Turn
	g.Turn = g.Next
	g.Next = t

	return nil
}

// Won returns true if the board has a winner
func won(grid [9]mark) bool {
	if grid[Center] > None {
		if (grid[TopLeft] == grid[Center] && grid[BotRight] == grid[Center]) ||
			(grid[TopRight] == grid[Center] && grid[BotLeft] == grid[Center]) ||
			(grid[TopMid] == grid[Center] && grid[BotMid] == grid[Center]) ||
			(grid[MidLeft] == grid[Center] && grid[MidRight] == grid[Center]) {
			return true
		}
	}
	if grid[TopMid] > None {
		if grid[TopLeft] == grid[TopMid] && grid[TopRight] == grid[TopMid] {
			return true
		}
	}
	if grid[BotMid] > None {
		if grid[BotLeft] == grid[BotMid] && grid[BotRight] == grid[BotMid] {
			return true
		}
	}
	if grid[MidLeft] > None {
		if grid[TopLeft] == grid[MidLeft] && grid[BotLeft] == grid[MidLeft] {
			return true
		}
	}
	if grid[MidRight] > None {
		if grid[TopRight] == grid[MidRight] && grid[BotRight] == grid[MidRight] {
			return true
		}
	}
	return false
}

func (g *Game) pretty() {
	var chars [9]rune
	for i, v := range g.Grid {
		switch v {
		case None:
			chars[i] = ' '
		case X:
			chars[i] = 'X'
		case O:
			chars[i] = 'O'
		}
	}
	g.Pretty = fmt.Sprintf("%c%c%c,%c%c%c,%c%c%c", chars[0], chars[1], chars[2], chars[3], chars[4], chars[5], chars[6], chars[7], chars[8])
}

// Valid returns the set of valid marks remaining
func (g *Game) Valid() map[GridRef]bool {
	m := make(map[GridRef]bool)
	for i, v := range g.Grid {
		if v == None {
			m[i] = true
		}
	}
	return m
}
