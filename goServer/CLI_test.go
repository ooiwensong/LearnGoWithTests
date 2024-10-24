package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/ooiwensong/LearnGoWithTests/goServer"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

// var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyBlindAlert = &poker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("3\nChris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("8\nCleo wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 8)
		assertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when user does not enter '{name} wins' format and does not start", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("1\nBaby cries\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinInputErrMsg)
	})
}

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	got := game.StartedWith

	if got != want {
		t.Errorf("game started with %d players, want %d", got, want)
	}
}

func assertGameFinishedWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()
	got := game.FinishedWith

	if got != want {
		t.Errorf("game finished with %q as winner, want %q", got, want)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}
