package poker

import (
	"io"
	"time"
)

var (
	blinds = []int{100, 200, 300, 400, 500, 600, 800, 1000, 1500, 2000, 4000, 8000}
)

// TexasHoldem manages a game of poker
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewTexasHoldem returns a new game
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

// Start will schedule blind alerts dependant on the number of players
func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	var blindTime time.Duration

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends the game, recording the winner
func (p *TexasHoldem) Finish(winner string) error {
	return p.store.RecordWin(winner)
}
