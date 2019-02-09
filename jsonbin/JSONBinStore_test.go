package jsonbin

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type JSONBinResponse struct {
	URI string   `json:"uri"`
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
