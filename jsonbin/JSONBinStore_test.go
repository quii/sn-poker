package jsonbin

import (
	"net/http"
	"testing"
)

func TestJSONBinStore(t *testing.T) {
	client := &http.Client{}
	binURL := "https://api.myjson.com/bins/ha5c8"
	bin := Store{Client: client, BinURL: binURL}

	playerName := "bob"

	league, err := bin.GetLeague()

	if err != nil {
		t.Fatal(err)
	}

	player := league.Find(playerName)

	currentScore := 0
	if player != nil {
		currentScore = player.Wins
	}

	bin.RecordWin(playerName)
	bin.RecordWin(playerName)

	league, err = bin.GetLeague()

	if err != nil {
		t.Fatal(err)
	}

	got := league.Find(playerName).Wins
	want := currentScore + 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
