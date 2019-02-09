package jsonbin

import (
	"net/http"
	"testing"
)

// uncomment me to create a new bin
//func TestCanCreateNewBins(t *testing.T) {
//	client := http.Client{}
//
//	bin, err := createNewJSONBin(client)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if bin == emptyBinURL {
//		t.Error("did not get a bin id")
//	}
//}

func TestJSONBinStore(t *testing.T) {
	client := &http.Client{}
	binURL := "https://api.myjson.com/bins/ha5c8"
	bin := Store{Client: client, BinURL: binURL}

	player := "bob"

	currentScore := bin.GetPlayerScore(player)

	bin.RecordWin(player)
	bin.RecordWin(player)

	got := bin.GetPlayerScore(player)
	want := currentScore + 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
