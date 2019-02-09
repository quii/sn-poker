package poker_test

import (
	"github.com/quii/sn-poker/jsonbin"
	"github.com/thanhpk/randstr"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	client := &http.Client{}
	binURL := "https://api.myjson.com/bins/ha5c8"
	bin := &jsonbin.Store{Client: client, BinURL: binURL}

	server := mustMakePlayerServer(t, bin, dummyGame)
	player := randstr.Hex(16)

	t.Log("new player", player)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})
}
