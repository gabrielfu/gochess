package main

type Game struct {
	board Board
}

func NewGame() *Game {
	return &Game{
		board: *NewBoard(),
	}
}

// LegalMoves returns all legal moves for the current player.
func (g *Game) LegalMoves() []*Move {
	candidatePieces := []Piece{}
	var allowedTos Bitboard
	var enemyOccupied Bitboard
	switch g.board.turn {
	case WHITE:
		candidatePieces = WHITE_PIECES
		allowedTos = ^g.board.whiteOccupied
		enemyOccupied = g.board.blackOccupied
	case BLACK:
		candidatePieces = BLACK_PIECES
		allowedTos = ^g.board.blackOccupied
		enemyOccupied = g.board.whiteOccupied
	}

	moves := []*Move{}
	for _, p := range candidatePieces {
		bb := g.board.GetBbForPiece(p)
		// If no more such pieces on the board, skip
		if bb == 0 {
			continue
		}

		// For each "from" square
		for _, from := range bb.Squares() {
			switch p {
			case WHITE_KNIGHT, BLACK_KNIGHT:
				toBb := calcKnightMoves(Square(from)) & allowedTos
				// For each "to" square
				for _, to := range toBb.Squares() {
					moves = append(moves, &Move{
						From:  Square(from),
						To:    Square(to),
						piece: p,
					})
				}
			case WHITE_PAWN:
				attackBb := calcWhitePawnAttacks(Square(from)) & enemyOccupied
				moveBb := (calcWhitePawnMoves(Square(from)) & allowedTos) &^ enemyOccupied
				toBb := attackBb | moveBb
				// For each "to" square
				for _, to := range toBb.Squares() {
					moves = append(moves, &Move{
						From:  Square(from),
						To:    Square(to),
						piece: p,
					})
				}
			}

		}
	}
	return moves
}

// Move executes the given move on the board.
func (g *Game) Move(move *Move) {
	g.board.Move(move)
}
