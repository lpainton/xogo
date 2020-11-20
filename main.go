// Package main implements a server for xogo
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handler("/new", newGame)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newGame(w http.ResponseWriter, r *http.Request) {
}
