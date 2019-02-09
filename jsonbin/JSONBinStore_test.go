package jsonbin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/quii/sn-poker"
	"log"
	"net/http"
	"testing"
)

const (
	jsonBinURL  = "https://api.myjson.com/bins"
	emptyBinURL = JSONBinURL("")
	emptyJSON   = `[]`
)

type JSONBinURL string

type JSONBinStore struct {
	client *http.Client
	binURL JSONBinURL
}

func (j *JSONBinStore) GetPlayerScore(name string) int {
	league := j.GetLeague()
	player := league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (j *JSONBinStore) RecordWin(name string) {
	//todo: err handling
	league := j.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, poker.Player{Name: name, Wins: 1})
	}

	//todo, refactor above to be a method on league

	req, _ := http.NewRequest(http.MethodPut, string(j.binURL), league.Encode())
	req.Header.Add("content-type", "application/json")
	_, err := j.client.Do(req)

	if err != nil {
		log.Println("problem storing league", err)
	}

}

func (j *JSONBinStore) GetLeague() poker.League {
	leagueReq, _ := http.NewRequest(http.MethodGet, string(j.binURL), nil)
	leagueRes, err := j.client.Do(leagueReq)

	if err != nil {
		log.Println("problem getting league", err)
	}

	defer leagueRes.Body.Close()
	league, err := poker.NewLeague(leagueRes.Body)

	if err != nil {
		log.Println("problem parsing league", err)
	}

	return league
}

type JSONBinResponse struct {
	URI string `json:"uri"`
}

func createNewJSONBin(client http.Client) (JSONBinURL, error) {
	req, _ := http.NewRequest(http.MethodPost, jsonBinURL, bytes.NewBuffer([]byte(emptyJSON)))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return emptyBinURL, fmt.Errorf("problem creating new JSON bin at %s, %v", jsonBinURL, err)
	}

	defer res.Body.Close()
	var jsonBinResponse JSONBinResponse
	err = json.NewDecoder(res.Body).Decode(&jsonBinResponse)

	if err != nil {
		return emptyBinURL, fmt.Errorf("problem decoding response from json bin %v", err)
	}

	fmt.Println("New bin", jsonBinResponse.URI)

	return JSONBinURL(jsonBinResponse.URI), nil
}

func TestCanCreateNewBins(t *testing.T) {
	client := http.Client{}

	bin, err := createNewJSONBin(client)

	if err != nil {
		t.Fatal(err)
	}

	if bin == emptyBinURL {
		t.Error("did not get a bin id")
	}
}

func TestJSONBinStore(t *testing.T) {
	client := &http.Client{}
	binURL := JSONBinURL("https://api.myjson.com/bins/ha5c8")

	bin := JSONBinStore{client: client, binURL: binURL}

	player := "chris"

	currentScore := bin.GetPlayerScore(player)

	bin.RecordWin(player)
	bin.RecordWin(player)

	got := bin.GetPlayerScore(player)
	want := currentScore + 2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
