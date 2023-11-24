package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	gochess "gochess/internal"

	"github.com/inancgumus/screen"
)

func main() {
	flip := flag.Bool("flip", false, "flip the board at Black's turn")
	eval := flag.Bool("eval", false, "show the evaluation bar")
	whiteEngine := flag.Bool("w", false, "white will be played by an engine")
	blackEngine := flag.Bool("b", false, "black will be played by an engine")
	flag.Parse()

	screen.Clear()
	reader := bufio.NewReader(os.Stdin)
	var errMsg string

	g := gochess.NewGame()

	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println(gochess.Banner())
		fmt.Println()
		if *flip && g.Board().Turn() == gochess.BLACK {
			fmt.Println(g.VisualizeFlipped())
		} else {
			fmt.Println(g.Visualize())
		}

		if *eval {
			e := gochess.Evaluate(g.Board())
			bar := gochess.EvaluationBar(e, 18)
			fmt.Println()
			fmt.Printf("%s %.2f\n", bar, float32(e)/100)
		}

		fmt.Println()
		fmt.Println(g.PGN())

		if g.Ended() {
			fmt.Printf("\033[0;32mGame over! %s won.\033[0;39m\n", g.Winner())
			break
		}

		fmt.Printf("\033[0;31m%s\033[0;39m\n", errMsg)

		playerTurn := (g.Turn() == gochess.WHITE && !*whiteEngine) || (g.Turn() == gochess.BLACK && !*blackEngine)
		if playerTurn {
			fmt.Print("Your move: ")

			input, err := reader.ReadString('\n')
			if err != nil {
				errMsg = "Error reading input: " + err.Error()
				continue
			}
			input = input[:len(input)-1]

			move, err := gochess.ParseSAN(input, g.Board())
			if err != nil {
				errMsg = err.Error()
				continue
			}

			if err := g.Move(move); err != nil {
				errMsg = err.Error()
				continue
			}
		} else {
			fmt.Print("Engine is thinking...")
			result := gochess.Search(g.Board(), 4)
			if result.Move() == nil {
				errMsg = "Internal Error: Engine could not find a move!"
				break
			}
			fmt.Println(result.Move().ToSAN(g.Board()))
			g.Move(result.Move())
			time.Sleep(50 * time.Millisecond)
		}

		// reset error message
		errMsg = ""
	}

	fmt.Print("Press Enter key to exit...")
	reader.ReadByte()
}
