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

	g := chessago.NewGame()
	moves := []*chessago.Move{
		chessago.NewMove(chessago.D2, chessago.D4, chessago.WHITE_PAWN),
		chessago.NewMove(chessago.E7, chessago.E5, chessago.BLACK_PAWN),
		chessago.NewMove(chessago.D4, chessago.E5, chessago.WHITE_PAWN),
		chessago.NewMove(chessago.D7, chessago.D5, chessago.BLACK_PAWN),
		chessago.NewMove(chessago.E5, chessago.E6, chessago.WHITE_PAWN),
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
		g.Move(move)
		i++
	}
}
