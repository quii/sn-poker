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
const defaultPort = "5000"

func main() {

	binURL := getEnvOrElse("BIN", testBin)
	port := getEnvOrElse("PORT", defaultPort)
	client := &http.Client{}

	bin := &jsonbin.Store{Client: client, BinURL: binURL}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), bin)

	server, err := poker.NewPlayerServer(bin, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	fmt.Println("Listening on http://localhost:" + port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func getEnvOrElse(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("%s environment variable not set, defaulting to %s\n", key, fallback)
		value = fallback
	}

	return value
}
