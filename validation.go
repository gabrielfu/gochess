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

	if king != 0 {
		return GetKingMoves(king.Squares()[0]).SquareIsSet(sq)
	}

	return false
}

func isCastlingPathClear(from Square, to Square, b *Board) bool {
	if from.Rank() != to.Rank() {
		return false
	}
	if isSquareAttacked(from, b) {
		return false
	}
	if from < to {
		for i := from + 1; i <= to; i++ {
			if isSquareAttacked(i, b) || b.allOccupied.SquareIsSet(i) {
				return false
			}
		}
	} else {
		for i := from - 1; i >= to; i-- {
			if isSquareAttacked(i, b) || b.allOccupied.SquareIsSet(i) {
				return false
			}
		}
	}
	return true
}

func isCastlingValid(move *Move, b *Board) bool {
	if move.Piece() == WHITE_KING {
		if b.turn != WHITE {
			return false
		}
		if move.From() != E1 {
			return false
		}
		if move.To() == C1 {
			if move.Castle() != WHITE_QUEEN_SIDE {
				return false
			}
			if !b.castlingRights.Has(WHITE_QUEEN_SIDE) {
				return false
			}
		} else if move.To() == G1 {
			if move.Castle() != WHITE_KING_SIDE {
				return false
			}
			if !b.castlingRights.Has(WHITE_KING_SIDE) {
				return false
			}
		} else {
			return false
		}
	} else if move.Piece() == BLACK_KING {
		if b.turn != BLACK {
			return false
		}
		if move.From() != E8 {
			return false
		}
		if move.To() == C8 {
			if move.Castle() != BLACK_QUEEN_SIDE {
				return false
			}
			if !b.castlingRights.Has(BLACK_QUEEN_SIDE) {
				return false
			}
		} else if move.To() == G8 {
			if move.Castle() != BLACK_KING_SIDE {
				return false
			}
			if !b.castlingRights.Has(BLACK_KING_SIDE) {
				return false
			}
		} else {
			return false
		}
	}

	if !isCastlingPathClear(move.From(), move.To(), b) {
		return false
	}
	return true
}

// isMovable returns true if the from, to and piece are valid
func isMovable(move *Move, b *Board) bool {
	// validate not friendly occupied
	allowedTos := ^b.whiteOccupied
	if b.turn == BLACK {
		allowedTos = ^b.blackOccupied
	}
	if !allowedTos.SquareIsSet(move.To()) {
		return false
	}

	// validate piece is currently at "from"
	p := b.GetPieceAtSquare(move.From())
	if p != move.Piece() {
		return false
	}

	// validate piece can move to "to"
	switch p {
	case WHITE_KNIGHT, BLACK_KNIGHT:
		if !GetKnightMoves(move.From()).SquareIsSet(move.To()) {
			return false
		}
	case WHITE_PAWN:
		if !GetWhitePawnAttacks(move.From()).SquareIsSet(move.To()) &&
			!GetWhitePawnMoves(move.From()).SquareIsSet(move.To()) {
			return false
		}
	case BLACK_PAWN:
		if !GetBlackPawnAttacks(move.From()).SquareIsSet(move.To()) &&
			!GetBlackPawnMoves(move.From()).SquareIsSet(move.To()) {
			return false
		}
	case WHITE_BISHOP, BLACK_BISHOP:
		if !GetBishopMoves(move.From(), b.allOccupied).SquareIsSet(move.To()) {
			return false
		}
	case WHITE_ROOK, BLACK_ROOK:
		if !GetRookMoves(move.From(), b.allOccupied).SquareIsSet(move.To()) {
			return false
		}
	case WHITE_QUEEN, BLACK_QUEEN:
		if !GetQueenMoves(move.From(), b.allOccupied).SquareIsSet(move.To()) {
			return false
		}
	case WHITE_KING, BLACK_KING:
		if !GetKingMoves(move.From()).SquareIsSet(move.To()) {
			return false
		}
	}
	return true
}

func isMoveValid(move *Move, b *Board) bool {
	if move.Castle() != 0 {
		return isCastlingValid(move, b)
	}
	// TODO
	return true
}
