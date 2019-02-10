package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

// StubPlayerStore implements PlayerStore for testing purposes
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

// RecordWin will record a win to WinCalls
func (s *StubPlayerStore) RecordWin(name string) error {
	s.WinCalls = append(s.WinCalls, name)
	return nil
}

// GetLeague returns League
func (s *StubPlayerStore) GetLeague() (League, error) {
	return s.League, nil
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got '%s' want '%s'", store.WinCalls[0], winner)
	}
}

// ScheduledAlert holds information about when an alert is scheduled
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// SpyBlindAlerter allows you to spy on ScheduleAlertAt calls
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

// ScheduleAlertAt records alerts that have been scheduled
func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}

// GameSpy will spy on calls to games
type GameSpy struct {
	t               *testing.T
	startCalled     bool
	startCalledWith int
	BlindAlert      []byte

	finishedCalled   bool
	finishCalledWith string
}

var (
	// DummyGame should be used when you dont care how a game is interacted with
	DummyGame = NewGameSpy(nil)
)

// NewGameSpy will create a new game to make assertions on how its used
func NewGameSpy(t *testing.T) *GameSpy {
	return &GameSpy{t: t}
}

// Start records how a game was started
func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.startCalled = true
	g.startCalledWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

// Finish records who won the game
func (g *GameSpy) Finish(winner string) error {
	g.finishedCalled = true
	g.finishCalledWith = winner
	return nil
}

// AssertFinishCalledWith tests who won the game
func (g *GameSpy) AssertFinishCalledWith(winner string) {
	g.t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return g.finishCalledWith == winner
	})

	if !passed {
		g.t.Errorf("expected finish called with '%s' but got '%s'", winner, g.finishCalledWith)
	}
}

// AssertStartedWith tests who started the game
func (g *GameSpy) AssertStartedWith(numberOfPlayersWanted int) {
	g.t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return g.startCalledWith == numberOfPlayersWanted
	})

	if !passed {
		g.t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, g.startCalledWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
