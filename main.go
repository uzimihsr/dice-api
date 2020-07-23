package main

import (
	"net/http"

	"github.com/uzimihsr/dice-api/dice"
	"github.com/uzimihsr/dice-api/handler"
)

func main() {
	d := dice.NewDice(6)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.DiceHandler(d))
	mux.HandleFunc("/cheat", handler.CheatDiceHandler(d))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
