package chessago

import "fmt"

func Main() {
	// captures
	// sliding pieces
	// TODO: castling
	// TODO: en passant
	// TODO: promotion

	// TODO: evaluation
	// TODO: search

	// TODO: move validation
	// TODO: check detection
	// TODO: checkmate detection
	// TODO: stalemate detection
	// TODO: draw conditions

	// TODO: FEN parsing
	// TODO: PGN parsing

	// board := NewStartingBoard()
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

	// b := NewStartingBoard()
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

	b := NewEmptyBoard()
	b.AddPieceToSquare(WHITE_ROOK, A1)
	b.AddPieceToSquare(WHITE_KING, E1)
	b.AddPieceToSquare(WHITE_ROOK, H1)
	g.SetBoard(b)

	legalMoves := g.LegalMoves()
	for _, move := range legalMoves {
		fmt.Println(move)
		cpy := b.Copy()
		cpy.Move(move)
		println(cpy.Visualize())
		println("=====================================")
	}
	println(len(legalMoves), "moves")
	return

	println(g.Visualize())
	println(g.board.castlingRights)

	if err = g.Move(&Move{
		from:   E1,
		to:     C1,
		piece:  WHITE_KING,
		castle: WHITE_QUEEN_SIDE,
	}); err != nil {
		panic(err)
	}
	println(g.Visualize())
	println(g.board.castlingRights)
	return

	if err = g.Move(&Move{
		from:  D2,
		to:    D4,
		piece: WHITE_PAWN,
	}); err != nil {
		panic(err)
	}
	if err = g.Move(&Move{
		from:  E7,
		to:    E5,
		piece: BLACK_PAWN,
	}); err != nil {
		panic(err)
	}
	if err = g.Move(&Move{
		from:  D4,
		to:    E5,
		piece: WHITE_PAWN,
	}); err != nil {
		panic(err)
	}
	println(g.board.Visualize())
	println(g.board.blackPawns.BinaryBoard())
	// legalMoves := g.LegalMoves()
	// for _, move := range legalMoves {
	// 	fmt.Println(move)
	// 	// board := NewStartingBoard()
	// 	// board.Move(move)
	// 	// board.Print()
	// 	println("=====================================")
	// }
	// println(len(legalMoves))

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
