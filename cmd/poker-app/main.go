package main

import (
	"fmt"
	"github.com/quii/sn-poker"
	"github.com/quii/sn-poker/jsonbin"
	"log"
	"net/http"
	"os"
)

const testBin = "https://api.myjson.com/bins/ha5c8"

func main() {

	binURL := os.Getenv("BIN")
	if binURL == "" {
		fmt.Println("Warning: You have not set a BIN environment variable")
		fmt.Printf("Defaulting to test bin %s\n", testBin)
		binURL = testBin
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	client := &http.Client{}

	bin := &jsonbin.Store{Client: client, BinURL: binURL}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), bin)

	server, err := poker.NewPlayerServer(bin, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
