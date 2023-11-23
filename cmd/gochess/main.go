package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	gochess "gochess/internal"

	"github.com/inancgumus/screen"
)

func main() {
	flip := flag.Bool("flip", false, "flip the board at Black's turn")
	flag.Parse()

	screen.Clear()
	reader := bufio.NewReader(os.Stdin)
	var errMsg string

	g := gochess.NewGame()

	for {
		screen.Clear()
		screen.MoveTopLeft()
		if *flip && g.Board().Turn() == gochess.BLACK {
			fmt.Println(g.VisualizeFlipped())
		} else {
			fmt.Println(g.Visualize())
		}
		fmt.Println()
		fmt.Println(g.PGN())

		if g.Ended() {
			fmt.Printf("\033[0;32mGame over! %s won.\033[0;39m\n", g.Winner())
			break
		}

		fmt.Printf("\033[0;31m%s\033[0;39m\n", errMsg)
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
		// reset error message
		errMsg = ""
	}

	fmt.Print("Press Enter key to exit...")
	reader.ReadByte()
}
