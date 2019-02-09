package main

import (
	"github.com/quii/sn-poker"
	"github.com/quii/sn-poker/jsonbin"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	binURL := "https://api.myjson.com/bins/ha5c8"
	bin := &jsonbin.Store{Client: client, BinURL: binURL}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), bin)

	server, err := poker.NewPlayerServer(bin, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
