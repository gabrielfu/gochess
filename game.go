package chessago

type Game struct {
	board *Board
}

func NewGame() *Game {
	InitMovesTables()

	return &Game{
		board: NewStartingBoard(),
	}
}

func (g *Game) SetBoard(board *Board) {
	g.board = board
}

// LegalMoves returns all legal moves for the current player.
func (g *Game) LegalMoves() []*Move {
	var candidatePieces []Piece
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
		if bb == nil || *bb == 0 {
			continue
		}

		// For each "from" square
		for _, from := range bb.Squares() {
			var toBb Bitboard = Bitboard(0)
			switch p {
			case WHITE_KNIGHT, BLACK_KNIGHT:
				toBb = GetKnightMoves(Square(from)) & allowedTos
			case WHITE_PAWN:
				attackBb := GetWhitePawnAttacks(Square(from)) & enemyOccupied
				moveBb := (GetWhitePawnMoves(Square(from)) & allowedTos) &^ enemyOccupied
				toBb = attackBb | moveBb
			case BLACK_PAWN:
				attackBb := GetBlackPawnAttacks(Square(from)) & enemyOccupied
				moveBb := (GetBlackPawnMoves(Square(from)) & allowedTos) &^ enemyOccupied
				toBb = attackBb | moveBb
			case WHITE_BISHOP, BLACK_BISHOP:
				toBb = GetBishopMoves(Square(from), g.board.allOccupied) & allowedTos
			case WHITE_ROOK, BLACK_ROOK:
				toBb = GetRookMoves(Square(from), g.board.allOccupied) & allowedTos
			case WHITE_QUEEN, BLACK_QUEEN:
				toBb = GetQueenMoves(Square(from), g.board.allOccupied) & allowedTos
			case WHITE_KING, BLACK_KING:
				toBb = GetKingMoves(Square(from)) & allowedTos
			}

			// For each "to" square
			for _, to := range toBb.Squares() {
				moves = append(moves, &Move{
					From:  Square(from),
					To:    Square(to),
					Piece: p,
				})
			}

			// castling
			if p == WHITE_KING && from == E1 {
				if g.board.castlingRights.Has(WHITE_QUEEN_SIDE) {
					// TODO: check castling path condition
					moves = append(moves, &Move{
						From:   E1,
						To:     C1,
						Piece:  WHITE_KING,
						castle: WHITE_QUEEN_SIDE,
					})
				}
				if g.board.castlingRights.Has(WHITE_KING_SIDE) {
					moves = append(moves, &Move{
						From:   E1,
						To:     G1,
						Piece:  WHITE_KING,
						castle: WHITE_KING_SIDE,
					})
				}
			} else if p == BLACK_KING && from == E8 {
				if g.board.castlingRights.Has(BLACK_QUEEN_SIDE) {
					moves = append(moves, &Move{
						From:   E8,
						To:     C8,
						Piece:  BLACK_KING,
						castle: BLACK_QUEEN_SIDE,
					})
				}
				if g.board.castlingRights.Has(BLACK_KING_SIDE) {
					moves = append(moves, &Move{
						From:   E8,
						To:     G8,
						Piece:  BLACK_KING,
						castle: BLACK_KING_SIDE,
					})
				}
			}
		}
	}
	return moves
}

func (g *Game) Visualize() string {
	return g.board.Visualize()
}

// Move executes the given move on the board.
func (g *Game) Move(move *Move) error {
	return g.board.Move(move)
}
