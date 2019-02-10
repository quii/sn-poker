package poker

import "testing"
import "github.com/google/go-cmp/cmp"

func TestLeague_AddWin(t *testing.T) {
	t.Run("adding wins for players", func(t *testing.T) {
		p1 := "Cleo"
		p2 := "Tiest"

		league := League{}

		league.AddWin(p1)
		league.AddWin(p2)
		league.AddWin(p1)

		assertPlayerScore(t, league, p1, 2)
		assertPlayerScore(t, league, p2, 1)
	})
}

func TestLeague_Encode_Decode(t *testing.T) {
	p1 := "Cleo"
	p2 := "Tiest"

	league := League{}

	league.AddWin(p1)
	league.AddWin(p1)
	league.AddWin(p1)
	league.AddWin(p2)
	league.AddWin(p1)
	league.AddWin(p2)

	encodedAndDecodeLeague, err := DecodeLeague(league.Encode())

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(encodedAndDecodeLeague, league) {
		diff := cmp.Diff(encodedAndDecodeLeague, league)
		t.Errorf("leagues were not equal %v", diff)
	}

}

func assertPlayerScore(t *testing.T, league League, name string, want int) {
	t.Helper()

	player := league.Find(name)

	if player == nil {
		t.Fatalf("could not find %s in league", name)
	}

	if player.Wins != want {
		t.Errorf("got %d, want %d", player.Wins, want)
	}
}
