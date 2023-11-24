package main

import (
	"fmt"
	gochess "gochess/internal"
)

func main() {
	// b := gochess.NewEmptyBoard()
	// b.AddPieceToSquare(gochess.WHITE_KING, gochess.E1)
	// b.AddPieceToSquare(gochess.BLACK_KING, gochess.E8)
	// b.AddPieceToSquare(gochess.WHITE_PAWN, gochess.D4)
	// b.AddPieceToSquare(gochess.BLACK_PAWN, gochess.E5)
	// b.AddPieceToSquare(gochess.WHITE_KNIGHT, gochess.D2)
	// b.AddPieceToSquare(gochess.WHITE_KNIGHT, gochess.G1)
	// b.AddPieceToSquare(gochess.WHITE_PAWN, gochess.B7)
	// b.AddPieceToSquare(gochess.BLACK_ROOK, gochess.C4)
	// b.AddPieceToSquare(gochess.WHITE_ROOK, gochess.C2)
	// b.AddPieceToSquare(gochess.WHITE_ROOK, gochess.A4)
	// b.AddPieceToSquare(gochess.WHITE_ROOK, gochess.A1)
	b := gochess.NewStartingBoard()
	score := gochess.Evaluate(b)
	fmt.Println(score)
	fmt.Println(b.Visualize())
}
