package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// League stores a collection of players
type League []Player

// Find tries to return a player from a League
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// Encode turns league into JSON
func (l League) Encode() io.Reader {
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(l)
	return &buf
}

// NewLeague creates a League from JSON
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing League, %v", err)
	}

	return league, err
}
