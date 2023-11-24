package gochess

import (
	"testing"
)

func TestIsSquareAttacked(t *testing.T) {
	b := NewStartingBoard()
	if isSquareAttacked(E4, b) {
		t.Errorf("E4 should be not attacked")
	}
	if !isSquareAttacked(E6, b) {
		t.Errorf("E6 should be attacked by black pawn")
	}
}

func TestIsCastlingValid(t *testing.T) {
	b := NewEmptyBoard()
	b.AddPieceToSquare(WHITE_ROOK, A1)
	b.AddPieceToSquare(WHITE_ROOK, H1)
	b.AddPieceToSquare(WHITE_KING, E1)

	queenSideCastleMove := NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE)
	kingSideCastleMove := NewCastlingMove(E1, G1, WHITE_KING, WHITE_KING_SIDE)

	if !isCastlingValid(queenSideCastleMove, b) {
		t.Errorf("Castling whtie queen side should be valid")
	}

	if !isCastlingValid(kingSideCastleMove, b) {
		t.Errorf("Castling whtie king side should be valid")
	}

	var cpy *Board
	cpy = NewEmptyBoard()
	b.AddPieceToSquare(WHITE_KING, E1)
	if isCastlingValid(NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE), cpy) {
		t.Errorf("Castling should be invalid without rook")
	}

	cpy = NewEmptyBoard()
	cpy.AddPieceToSquare(WHITE_ROOK, A1)
	cpy.AddPieceToSquare(WHITE_KING, D1)
	if isCastlingValid(NewCastlingMove(D1, C1, WHITE_KING, WHITE_QUEEN_SIDE), cpy) {
		t.Errorf("Castling should be invalid from incorrect position")
	}

	cpy = NewEmptyBoard()
	cpy.AddPieceToSquare(WHITE_ROOK, B1)
	cpy.AddPieceToSquare(WHITE_KING, E1)
	if isCastlingValid(NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE), cpy) {
		t.Errorf("Castling should be invalid from incorrect position")
	}

	cpy = NewEmptyBoard()
	cpy.AddPieceToSquare(WHITE_ROOK, A1)
	cpy.AddPieceToSquare(WHITE_QUEEN, E1)
	if isCastlingValid(NewCastlingMove(E1, C1, WHITE_QUEEN, WHITE_QUEEN_SIDE), cpy) {
		t.Errorf("Castling should be invalid for non-king piece")
	}

	cpy = b.Copy()
	cpy.Move(NewMove(E1, D1, WHITE_KING))
	cpy.Move(NewMove(D1, E1, WHITE_KING))
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid after king has been moved")
	}

	cpy = b.Copy()
	cpy.Move(NewMove(H1, G1, WHITE_ROOK))
	cpy.Move(NewMove(G1, H1, WHITE_ROOK))
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid after rook has been moved")
	}

	cpy = b.Copy()
	cpy.AddPieceToSquare(WHITE_KNIGHT, G1)
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid with knight in the way")
	}

	cpy = b.Copy()
	cpy.AddPieceToSquare(BLACK_QUEEN, E8)
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid when in check")
	}

	cpy = b.Copy()
	cpy.AddPieceToSquare(BLACK_QUEEN, F8)
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid through check")
	}

	cpy = b.Copy()
	cpy.AddPieceToSquare(BLACK_QUEEN, G8)
	if isCastlingValid(kingSideCastleMove, cpy) {
		t.Errorf("Castling should be invalid when destination is in check")
	}
}
