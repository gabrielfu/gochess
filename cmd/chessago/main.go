package main

import (
	"bufio"
	"fmt"
	"os"

	"chessago"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	reader := bufio.NewReader(os.Stdin)
	var message string

	g := chessago.NewGame()

	var i = 0
	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println(g.Visualize())
		fmt.Println()
		fmt.Println(g.PGN())
		fmt.Printf("\033[0;31m%s\033[0;39m\n", message)
		fmt.Print("Your move: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			message = "Error reading input: " + err.Error()
			continue
		}
		input = input[:len(input)-1]

		move, err := chessago.ParseSAN(input, g.Board())
		if err != nil {
			message = err.Error()
			continue
		}

		if err := g.Move(move); err != nil {
			message = err.Error()
			continue
		}
		// reset message
		message = ""
		i++
	}
}
