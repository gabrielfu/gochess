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
		{From: chessago.D2, To: chessago.D4, Piece: chessago.WHITE_PAWN},
		{From: chessago.E7, To: chessago.E5, Piece: chessago.BLACK_PAWN},
		{From: chessago.D4, To: chessago.E5, Piece: chessago.WHITE_PAWN},
		{From: chessago.D7, To: chessago.D5, Piece: chessago.BLACK_PAWN},
		{From: chessago.E5, To: chessago.E6, Piece: chessago.WHITE_PAWN},
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
