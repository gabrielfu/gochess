package main

import "fmt"

func main() {
	// TODO: captures
	// TODO: sliding pieces
	// TODO: castling
	// TODO: en passant
	// TODO: promotion

	// TODO: move validation
	// TODO: check detection
	// TODO: checkmate detection
	// TODO: stalemate detection
	// TODO: draw conditions

	// TODO: FEN parsing
	// TODO: PGN parsing

	// board := NewBoard()
	// board.Print()
	// println("=====================================")

	// // bb := board.whitePawns
	// // println(bb.BinaryBoard())
	// // println("=====================================")

	// moves := calcWhitePawnAttacks(D8)
	// println(moves.BinaryBoard())
	// println("=====================================")

	// allowedTos := ^board.whiteOccupied
	// dest := moves & allowedTos
	// println(dest.BinaryBoard())
	// println("=====================================")

	// b := NewBoard()
	// b.Move(&Move{
	// 	From:  E1,
	// 	To:    F5,
	// 	Piece: WHITE_QUEEN,
	// })
	// b.Move(&Move{
	// 	From:  C2,
	// 	To:    C4,
	// 	Piece: WHITE_PAWN,
	// })
	// b.Move(&Move{
	// 	From:  G2,
	// 	To:    G4,
	// 	Piece: WHITE_PAWN,
	// })
	// println(b.Visualize())
	// moves := GetQueenMoves(F5, b.allOccupied)
	// println(moves.BinaryBoard())
	// return

	g := NewGame()
	var err error
	if err = g.Move(&Move{
		From:  D2,
		To:    D4,
		Piece: WHITE_PAWN,
	}); err != nil {
		panic(err)
	}
	if err = g.Move(&Move{
		From:  E7,
		To:    E5,
		Piece: BLACK_PAWN,
	}); err != nil {
		panic(err)
	}
	if err = g.Move(&Move{
		From:  D4,
		To:    E5,
		Piece: WHITE_PAWN,
	}); err != nil {
		panic(err)
	}
	println(g.board.Visualize())
	legalMoves := g.LegalMoves()
	for _, move := range legalMoves {
		fmt.Println(move)
		// board := NewBoard()
		// board.Move(move)
		// board.Print()
		println("=====================================")
	}
	println(len(legalMoves))

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
