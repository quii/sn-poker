package jsonbin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/quii/sn-poker"
	"log"
	"net/http"
)

const (
	jsonBinURL  = "https://api.myjson.com/bins"
	emptyBinURL = ""
	emptyJSON   = `[]`
)

// Store acts as a store for a poker league
type Store struct {
	Client *http.Client
	BinURL string
}

// GetPlayerScore returns the score for a player
func (j *Store) GetPlayerScore(name string) int {
	league, err := j.GetLeague()

	if err != nil {
		log.Println(err)
		return 0 //todo: handle this properly
	}

	player := league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin updates json bin with new winner added
func (j *Store) RecordWin(name string) {
	league, err := j.GetLeague()

	//todo: ...
	if err != nil {
		log.Println(err)
		log.Println("did not record win")
		return
	}

	league.AddWin(name)

	req, _ := http.NewRequest(http.MethodPut, string(j.BinURL), league.Encode())
	req.Header.Add("content-type", "application/json")
	res, err := j.Client.Do(req)

	if err != nil {
		log.Println("problem storing league", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Println("did not get 200 when storing league", res)
	}

}

// GetLeague fetches the league from json bin
func (j *Store) GetLeague() (poker.League, error) {
	leagueReq, _ := http.NewRequest(http.MethodGet, string(j.BinURL), nil)
	leagueRes, err := j.Client.Do(leagueReq)

	if err != nil {
		return nil, fmt.Errorf("problem getting league from %s, %v", j.BinURL, err)
	}

	defer leagueRes.Body.Close()
	league, err := poker.NewLeague(leagueRes.Body)

	if err != nil {
		return nil, fmt.Errorf("problem parsing league from %s, %v", j.BinURL, err)
	}

	return league, nil
}

type jsonBinResponse struct {
	URI string `json:"uri"`
}

// CreateNewJSONBin returns the url of a new JSON bin
func CreateNewJSONBin(client *http.Client) (string, error) {
	req, _ := http.NewRequest(http.MethodPost, jsonBinURL, bytes.NewBuffer([]byte(emptyJSON)))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return emptyBinURL, fmt.Errorf("problem creating new JSON bin at %s, %v", jsonBinURL, err)
	}

	defer res.Body.Close()
	var jsonBinResponse jsonBinResponse
	err = json.NewDecoder(res.Body).Decode(&jsonBinResponse)

	if err != nil {
		return emptyBinURL, fmt.Errorf("problem decoding response from json bin %v", err)
	}

	return jsonBinResponse.URI, nil
}
