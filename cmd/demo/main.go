package main

import (
	"bufio"
	"fmt"
	"os"

	gochess "gochess/internal"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	reader := bufio.NewReader(os.Stdin)

	g := gochess.NewGame()
	moves := []*gochess.Move{
		gochess.NewMove(gochess.D2, gochess.D4, gochess.WHITE_PAWN),
		gochess.NewMove(gochess.E7, gochess.E5, gochess.BLACK_PAWN),
		gochess.NewMove(gochess.D4, gochess.E5, gochess.WHITE_PAWN),
		gochess.NewMove(gochess.D7, gochess.D5, gochess.BLACK_PAWN),
		gochess.NewMove(gochess.E5, gochess.E6, gochess.WHITE_PAWN),
	}

	var i = 0
	for i < len(moves) {
		screen.MoveTopLeft()
		fmt.Println(g.Visualize())
		reader.ReadString('\n')

		move := moves[i]
		if err := g.Move(move); err != nil {
			panic(err)
		}
		i++
	}
}
