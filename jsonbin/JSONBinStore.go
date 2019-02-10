package jsonbin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/quii/sn-poker"
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

// RecordWin updates json bin with new winner added
func (j *Store) RecordWin(name string) error {
	league, err := j.GetLeague()

	if err != nil {
		return err
	}

	league.AddWin(name)

	req, _ := http.NewRequest(http.MethodPut, string(j.BinURL), league.Encode())
	req.Header.Add("content-type", "application/json")
	res, err := j.Client.Do(req)

	if err != nil {
		return fmt.Errorf("problem storing new json at %s %v", j.BinURL, err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("did not get 200 when storing new json at %s, got %v", j.BinURL, res.Status)
	}

	return nil
}

// GetLeague fetches the league from json bin
func (j *Store) GetLeague() (poker.League, error) {
	leagueReq, _ := http.NewRequest(http.MethodGet, string(j.BinURL), nil)
	leagueRes, err := j.Client.Do(leagueReq)

	if err != nil {
		return nil, fmt.Errorf("problem getting league from %s, %v", j.BinURL, err)
	}

	defer leagueRes.Body.Close()
	league, err := poker.Decode(leagueRes.Body)

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
