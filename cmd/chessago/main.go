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
	var history string
	var message string

	g := chessago.NewGame()

	var i = 0
	for {
		screen.Clear()
		screen.MoveTopLeft()
		fmt.Println(g.Visualize())
		fmt.Println()
		fmt.Println(history)
		fmt.Println(message)
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
		history += fmt.Sprintf("%d. %s ", i/2+1, input)
		// reset message
		message = ""
		i++
	}
}
