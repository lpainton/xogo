// Package game implements game logic for xogo
package game

import "fmt"

type mark = int

const (
	none mark = iota
	x    mark = iota
	o    mark = iota
)

// GridRef indicates a square on the grid surface
type GridRef = int

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
	Pretty string  `json:"pretty"`
}

// New creates a pristine game board
func New() *Game {
	g := &Game{Next: x}
	g.pretty()
	return g
}

// Mark takes a GridRef and marks it with the next marker
func (g *Game) Mark(square GridRef) error {

	// If the square is already marked we change nothing
	if g.Grid[square] != none {
		return nil
	}

	g.Grid[square] = g.Next
	g.pretty()

	if hasWinner(g.Grid) {
		g.Winner = g.Next
		g.Next = none
		return nil
	}

	g.Next = next(g.Next)
	return nil
}

func next(current mark) mark {
	switch current {
	case x:
		return o
	case o:
		return x
	}
	return none
}

func hasWinner(grid [9]mark) bool {
	if grid[Center] > none {
		if (grid[TopLeft] == grid[Center] && grid[BotRight] == grid[Center]) ||
			(grid[TopRight] == grid[Center] && grid[BotLeft] == grid[Center]) ||
			(grid[TopMid] == grid[Center] && grid[BotMid] == grid[Center]) ||
			(grid[MidLeft] == grid[Center] && grid[MidRight] == grid[Center]) {
			return true
		}
	}
	if grid[TopMid] > none {
		if grid[TopLeft] == grid[TopMid] && grid[TopRight] == grid[TopMid] {
			return true
		}
	}
	if grid[BotMid] > none {
		if grid[BotLeft] == grid[BotMid] && grid[BotRight] == grid[BotMid] {
			return true
		}
	}
	if grid[MidLeft] > none {
		if grid[TopLeft] == grid[MidLeft] && grid[BotLeft] == grid[MidLeft] {
			return true
		}
	}
	if grid[MidRight] > none {
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
		case none:
			chars[i] = ' '
		case x:
			chars[i] = 'X'
		case o:
			chars[i] = 'O'
		}
	}
	g.Pretty = fmt.Sprintf("%c%c%c\n%c%c%c\n%c%c%c", chars[0], chars[1], chars[2], chars[3], chars[4], chars[5], chars[6], chars[7], chars[8])
}

// Empty returns the set of empty references left
func (g *Game) Empty() map[GridRef]bool {
	m := make(map[GridRef]bool)
	for i, v := range g.Grid {
		if v != none {
			m[i] = true
		}
	}
	return m
}
