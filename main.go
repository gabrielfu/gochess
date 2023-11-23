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

	b := NewEmptyBoard()
	b.AddPieceToSquare(WHITE_KING, E1)
	b.AddPieceToSquare(BLACK_KING, E8)
	b.AddPieceToSquare(WHITE_PAWN, D4)
	b.AddPieceToSquare(BLACK_PAWN, E5)
	b.AddPieceToSquare(WHITE_KNIGHT, D2)
	b.AddPieceToSquare(WHITE_KNIGHT, G1)
	b.AddPieceToSquare(WHITE_PAWN, B7)
	b.AddPieceToSquare(BLACK_ROOK, C4)
	b.AddPieceToSquare(WHITE_ROOK, C2)
	b.AddPieceToSquare(WHITE_ROOK, A4)
	b.AddPieceToSquare(WHITE_ROOK, A1)

	g := NewGame()
	g.SetBoard(b)
	fmt.Println("Current board:")
	fmt.Println(g.Visualize())
	fmt.Println("=====================")

	sans := []string{"O-O", "O-O-O", "d5", "b8Q", "b8=Q", "dxe5", "Nh3", "Nf3", "Ng1f3", "Na6b8", "Ngf3", "N1f3", "Nxc4", "Rxc4", "R1xc4", "Raxc4", "Ra4xc4"}
	for _, san := range sans {
		print(san, "\t")
		move, err := ParseSAN(san, g.Board())
		if err != nil {
			println(err.Error())
		} else {
			println(move.String())
		}
	}

}
