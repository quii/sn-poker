package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"
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

// Sort sorts the league by wins
func (l League) Sort() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Wins > l[j].Wins
	})
}

// Encode turns league into JSON
func (l League) Encode() io.Reader {
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(l)
	return &buf
}

// AddWin will update league with new winner
func (l *League) AddWin(name string) {
	player := l.Find(name)

	if player != nil {
		player.Wins++
	} else {
		*l = append(*l, Player{Name: name, Wins: 1})
	}
}

// NewLeague creates a League from JSON
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		return league, fmt.Errorf("problem parsing League, %v", err)
	}

	league.Sort()

	return league, nil
}
