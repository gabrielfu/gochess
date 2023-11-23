package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"chessago"

	"github.com/inancgumus/screen"
)

func main() {
	flip := flag.Bool("flip", false, "flip the board at Black's turn")
	flag.Parse()

	screen.Clear()
	reader := bufio.NewReader(os.Stdin)
	var errMsg string

	g := chessago.NewGame()

	var i = 0
	for {
		screen.Clear()
		screen.MoveTopLeft()
		if *flip && g.Board().Turn() == chessago.BLACK {
			fmt.Println(g.VisualizeFlipped())
		} else {
			fmt.Println(g.Visualize())
		}
		fmt.Println()
		fmt.Println(g.PGN())

		if g.Board().IsInCheckmate() {
			winner := 1 - g.Turn()
			fmt.Printf("\033[0;32mGame over! %s won.\033[0;39m\n", winner)
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

		move, err := chessago.ParseSAN(input, g.Board())
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
		i++
	}

	fmt.Print("Press Enter key to exit...")
	reader.ReadByte()
}
