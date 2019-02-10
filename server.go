package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	RecordWin(name string) error
	GetLeague() (League, error)
}

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     Game
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)
	p.game = game
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	router.Handle("/", http.FileServer(http.Dir("./static")))

	p.Handler = router

	return p, nil
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)
	p.game.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()

	if err := p.game.Finish(strings.ToLower(winner)); err != nil {
		fmt.Fprintf(ws, "there was a problem finishing the game - %v", p.game.Finish(strings.ToLower(winner)))
	}
}

func (p *PlayerServer) playGame(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	league, err := p.store.GetLeague()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(league)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, name string) {
	league, err := p.store.GetLeague()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	player := league.Find(name)

	if player == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, player.Wins)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	err := p.store.RecordWin(player)

	if err != nil {
		http.Error(w, fmt.Sprintf("problem recording win %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
