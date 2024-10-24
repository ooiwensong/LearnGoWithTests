package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
	BadWinInputErrMsg    = "Bad value received for winner of game, please try again with a '{name} wins' format"
)

// our CLI is only responsible for input/ output while delegating
// game specifics to Game
type CLI struct {
	// playerStore PlayerStore
	// alerter     BlindAlerter
	in   *bufio.Scanner // to read user input
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		// playerStore: store,
		// alerter:     alerter,
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	// Atoi parses a string into an integer
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	/**

	// when PlayPoker is called, this function will be called which in turn
	// calls whichever particular implementation of ScheduleAlertAt
	// cli.scheduleBlindAlerts(numberOfPlayers)
	// userInput := cli.readLine()
	// cli.playerStore.RecordWin(extractWinner(userInput))

	**/

	winnerInput := cli.readLine()

	err = checkCorrectWinnerInput(winnerInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinInputErrMsg)
		return
	}

	winner := extractWinner(winnerInput)

	cli.game.Start(numberOfPlayers)
	cli.game.Finish(winner)
}

// func (cli *CLI) scheduleBlindAlerts(numerOfPlayers int) {
// 	blindIncrement := time.Duration(5+numerOfPlayers) * time.Minute

// 	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
// 	blindTime := 0 * time.Second
// 	for _, blind := range blinds {
// 		cli.alerter.ScheduleAlertAt(blindTime, blind)
// 		blindTime += blindIncrement
// 	}
// }

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func checkCorrectWinnerInput(userInput string) error {
	input := strings.SplitN(userInput, " ", 2)
	if !strings.Contains(input[1], "wins") {
		return fmt.Errorf("bad input error")
	}
	return nil
}

func (cli *CLI) readLine() string {
	// reads up to a newline
	cli.in.Scan()
	return cli.in.Text()
}
