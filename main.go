// Package main implements a server for xogo
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/xogo/game"
)

const host = "http://localhost:8080/"

var (
	newURL      = fmt.Sprintf("%s%s", host, "new")
	topLeftURL  = fmt.Sprintf("%s%s%d", host, "mark/", game.TopLeft)
	topMidURL   = fmt.Sprintf("%s%s%d", host, "mark/", game.TopMid)
	topRightURL = fmt.Sprintf("%s%s%d", host, "mark/", game.TopRight)
	midLeftURL  = fmt.Sprintf("%s%s%d", host, "mark/", game.MidLeft)
	centerURL   = fmt.Sprintf("%s%s%d", host, "mark/", game.Center)
	midRightURL = fmt.Sprintf("%s%s%d", host, "mark/", game.MidRight)
	botLeftURL  = fmt.Sprintf("%s%s%d", host, "mark/", game.BotLeft)
	botMidURL   = fmt.Sprintf("%s%s%d", host, "mark/", game.BotMid)
	botRightURL = fmt.Sprintf("%s%s%d", host, "mark/", game.BotRight)
)

// State holds the passed state representation
type State struct {
	GameBoard    *game.Game `json:"gameBoard"`
	ValidActions *Actions   `json:"validActions"`
}

// Actions enumerate a set of legal actions for the game
type Actions struct {
	NewGame  string `json:"newGame"`
	TopLeft  string `json:"markTopLeft,omitempty"`
	TopMid   string `json:"markTopMid,omitempty"`
	TopRight string `json:"markTopRight,omitempty"`
	MidLeft  string `json:"markMidLeft,omitempty"`
	Center   string `json:"markCenter,omitempty"`
	MidRight string `json:"markMidRight,omitempty"`
	BotLeft  string `json:"markBotLeft,omitempty"`
	BotMid   string `json:"markBotMid,omitempty"`
	BotRight string `json:"markBotRight,omitempty"`
}

func main() {
	http.HandleFunc("/new", newGame)
	handleByRef(game.TopLeft)
	handleByRef(game.TopMid)
	handleByRef(game.TopRight)
	handleByRef(game.MidLeft)
	handleByRef(game.Center)
	handleByRef(game.MidRight)
	handleByRef(game.BotLeft)
	handleByRef(game.BotMid)
	handleByRef(game.BotRight)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newState() *State {
	g := game.New()
	a := &Actions{
		NewGame:  newURL,
		TopLeft:  topLeftURL,
		TopMid:   topMidURL,
		TopRight: topRightURL,
		MidLeft:  midLeftURL,
		Center:   centerURL,
		MidRight: midRightURL,
		BotLeft:  botLeftURL,
		BotMid:   botMidURL,
		BotRight: botRightURL,
	}
	return &State{
		GameBoard:    g,
		ValidActions: a,
	}
}

func newGame(w http.ResponseWriter, r *http.Request) {
	s := newState()
	str, err := json.Marshal(s)
	if err != nil {
		// TODO(lee.painton): Return a legit status code
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}
	fmt.Fprint(w, string(str))
	return
}

func handleByRef(ref game.GridRef) {
	str := fmt.Sprintf("/mark/%d", ref)
	f := func(w http.ResponseWriter, r *http.Request) {
		mark(w, r, ref)
	}
	http.HandleFunc(str, f)
	return
}

func mark(w http.ResponseWriter, r *http.Request, ref game.GridRef) {
	if r.Body == nil {
		http.Error(w, "mark handler: no request body found", http.StatusBadRequest)
		return
	}

	var s State
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.GameBoard.Mark(ref)
	s.updateActions()

	str, err := json.Marshal(&s)
	if err != nil {
		// TODO(lee.painton): Return a legit status code
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	fmt.Fprint(w, string(str))
	return
}

func (s *State) updateActions() {
	newActions := &Actions{
		NewGame: fmt.Sprintf("%s%s", host, "new"),
	}
	if s.GameBoard.Winner != 0 {
		s.ValidActions = newActions
		return
	}

	isMarked := s.GameBoard.Empty()
	if !isMarked[game.TopLeft] {
		newActions.TopLeft = topLeftURL
	}
	if !isMarked[game.TopMid] {
		newActions.TopMid = topMidURL
	}
	if !isMarked[game.TopRight] {
		newActions.TopRight = topRightURL
	}
	if !isMarked[game.MidLeft] {
		newActions.MidLeft = midLeftURL
	}
	if !isMarked[game.Center] {
		newActions.Center = centerURL
	}
	if !isMarked[game.MidRight] {
		newActions.MidRight = midRightURL
	}
	if !isMarked[game.BotLeft] {
		newActions.BotLeft = botLeftURL
	}
	if !isMarked[game.BotMid] {
		newActions.BotMid = botMidURL
	}
	if !isMarked[game.BotRight] {
		newActions.BotRight = botRightURL
	}

	s.ValidActions = newActions
	return
}
