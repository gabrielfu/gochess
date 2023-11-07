package main

type Game struct {
	board Board
}

func NewGame() *Game {
	return &Game{
		board: *NewBoard(),
	}
}

func (g *Game) LegalMoves() []*Move {
	candidatePieces := []Piece{}
	var allowedTos Bitboard
	switch g.board.turn {
	case WHITE:
		candidatePieces = WHITE_PIECES
		allowedTos = ^g.board.whiteOccupied
	case BLACK:
		candidatePieces = BLACK_PIECES
		allowedTos = ^g.board.blackOccupied
	}

	moves := []*Move{}
	for _, p := range candidatePieces {
		bb := g.board.GetBbForPiece(p)
		// If no more such pieces on the board, skip
		if bb == 0 {
			continue
		}

		// For each "from" square
		for from := 0; from < 64; from++ {
			// If no such piece on this square, skip
			if bb&(1<<uint(from)) == 0 {
				continue
			}

			switch p {
			case WHITE_KNIGHT, BLACK_KNIGHT:
				toBb := calcKnightMoves(Square(from)) & allowedTos
				// should optimize this
				for to := 0; to < 64; to++ {
					if toBb&(1<<uint(to)) == 0 {
						continue
					}

					moves = append(moves, &Move{
						From: Square(from),
						To:   Square(to),
					})
				}
			}

		}
	}
	return moves
}
