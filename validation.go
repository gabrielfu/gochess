package chessago

// isSquareAttacked returns true if the square is being attacked by the opponent
func isSquareAttacked(sq Square, b *Board) bool {
	queens := b.whiteQueens
	rooks := b.whiteRooks
	bishops := b.whiteBishops
	knights := b.whiteKnights
	pawns := b.whitePawns
	king := b.whiteKing
	if b.turn == WHITE {
		queens = b.blackQueens
		rooks = b.blackRooks
		bishops = b.blackBishops
		knights = b.blackKnights
		pawns = b.blackPawns
		king = b.blackKing
	}

	for _, from := range queens.Squares() {
		if GetQueenMoves(from, b.allOccupied).SquareIsSet(sq) {
			return true
		}
	}

	for _, from := range rooks.Squares() {
		if GetRookMoves(from, b.allOccupied).SquareIsSet(sq) {
			return true
		}
	}

	for _, from := range bishops.Squares() {
		if GetBishopMoves(from, b.allOccupied).SquareIsSet(sq) {
			return true
		}
	}

	for _, from := range knights.Squares() {
		if GetKnightMoves(from).SquareIsSet(sq) {
			return true
		}
	}

	for _, from := range pawns.Squares() {
		if b.turn == WHITE {
			if GetBlackPawnAttacks(from).SquareIsSet(sq) {
				return true
			}
		} else {
			if GetWhitePawnAttacks(from).SquareIsSet(sq) {
				return true
			}
		}
	}

	return GetKingMoves(king.Squares()[0]).SquareIsSet(sq)
}
