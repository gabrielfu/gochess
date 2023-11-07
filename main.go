package main

func main() {
	// board := NewBoard()
	// board.Print()
	// println("=====================================")

	// bb := board.whitePawns
	// println(bb.ToBinaryBoard())
	// println("=====================================")

	sq := D2
	println(sq.ToBinaryBoard())
	println("=====================================")

	moves := calcKnightMoves(sq)
	println(moves.ToBinaryBoard())
	println("=====================================")

	// var squares = []SQUARES{A1, A2, D1, E1, H7, G8}
	// var expected = []PIECES{WHITE_ROOK, WHITE_PAWN, WHITE_KING, WHITE_QUEEN, BLACK_PAWN, BLACK_KNIGHT}
	// for i, square := range squares {
	// 	piece := board.GetPieceAtSquare(uint8(square))
	// 	println(square, ": Expected "+SYMBOLS[expected[i]]+", Actual "+piece)
	// }

	// println(uint8(E1))
	// println(board.GetPieceAtSquare(uint8(E1)))

	// println("=====================================")
	// board.Move(&Move{
	// 	From:  D2,
	// 	To:    D4,
	// 	Piece: WHITE_PAWN,
	// })
	// board.Move(&Move{
	// 	From:  D7,
	// 	To:    D5,
	// 	Piece: BLACK_PAWN,
	// })
	// board.Move(&Move{
	// 	From:  G1,
	// 	To:    F3,
	// 	Piece: WHITE_KNIGHT,
	// })
	// board.Print()
	// println("=====================================")
	// board.Move(&Move{
	// 	From:  D5,
	// 	To:    D4,
	// 	Piece: BLACK_PAWN,
	// })
	// board.Print()
}
