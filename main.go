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

	g := NewGame()
	g.Move(NewMove(D2, D4, WHITE_PAWN))
	g.Move(NewMove(E7, E5, BLACK_PAWN))
	g.Move(NewMove(B1, D2, WHITE_KNIGHT))
	g.Move(NewMove(H8, H3, BLACK_ROOK))
	fmt.Println("Current board:")
	fmt.Println(g.Visualize())
	fmt.Println("=====================")

	sans := []string{"d5", "d8Q", "d8=Q", "dxe5", "Nf3", "Ng1f3", "Ngf3", "N1f3", "Nxf3", "Ngxf3", "N1xf3", "Ng1xf3"}
	for _, san := range sans {
		print(san, "\t")
		move, err := parseSAN(san, g.Board())
		if err != nil {
			println(err.Error())
		} else {
			println(move.String())
		}
	}

}
