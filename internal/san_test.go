package gochess

import (
	"errors"
	"testing"
)

func TestParseSAN(t *testing.T) {
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

	illegalError := errors.New("Illegal SAN")
	ambiguousError := errors.New("Ambiguous SAN")
	missingPromotionError := errors.New("Missing promotion SAN")

	testCases := []struct {
		san      string
		expected *Move
		err      error
	}{
		{
			san:      "O-O",
			expected: nil,
			err:      illegalError,
		},
		{
			san:      "O-O-O",
			expected: NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE),
			err:      nil,
		},
		{
			san:      "d5",
			expected: NewMove(D4, D5, WHITE_PAWN),
			err:      nil,
		},
		{
			san:      "b8",
			expected: nil,
			err:      missingPromotionError,
		},
		{
			san:      "axb8",
			expected: nil,
			err:      missingPromotionError,
		},
		{
			san:      "b8Q",
			expected: &Move{B7, B8, WHITE_PAWN, 0, WHITE_QUEEN},
			err:      nil,
		},
		{
			san:      "b8=Q",
			expected: &Move{B7, B8, WHITE_PAWN, 0, WHITE_QUEEN},
			err:      nil,
		},
		{
			san:      "axb8=Q",
			expected: &Move{B7, B8, WHITE_PAWN, 0, WHITE_QUEEN},
			err:      nil,
		},
		{
			san:      "dxe5",
			expected: NewMove(D4, E5, WHITE_PAWN),
			err:      nil,
		},
		{
			san:      "Nh3",
			expected: NewMove(G1, H3, WHITE_KNIGHT),
			err:      nil,
		},
		{
			san:      "Nf3",
			expected: nil,
			err:      ambiguousError,
		},
		{
			san:      "Ng1f3",
			expected: NewMove(G1, F3, WHITE_KNIGHT),
			err:      nil,
		},
		{
			san:      "Na6b8",
			expected: nil,
			err:      illegalError,
		},
		{
			san:      "Ngf3",
			expected: NewMove(G1, F3, WHITE_KNIGHT),
			err:      nil,
		},
		{
			san:      "N1f3",
			expected: NewMove(G1, F3, WHITE_KNIGHT),
			err:      nil,
		},
		{
			san:      "Nxc4",
			expected: NewMove(D2, C4, WHITE_KNIGHT),
			err:      nil,
		},
		{
			san:      "Rxc4",
			expected: nil,
			err:      ambiguousError,
		},
		{
			san:      "R2xc4",
			expected: NewMove(C2, C4, WHITE_ROOK),
			err:      nil,
		},
		{
			san:      "Raxc4",
			expected: NewMove(A4, C4, WHITE_ROOK),
			err:      nil,
		},
		{
			san:      "Ra4xc4",
			expected: NewMove(A4, C4, WHITE_ROOK),
			err:      nil,
		},
	}

	for _, tc := range testCases {
		move, err := ParseSAN(tc.san, b)

		if err != nil && tc.err != nil && (err.Error()[:len(tc.err.Error())] != tc.err.Error()) {
			t.Errorf("Expected error: %v, got: %v", tc.err, err)
		}

		if move != nil && tc.expected != nil && !move.Equals(tc.expected) {
			t.Errorf("Expected move: %v, got: %v", tc.expected, move)
		}
	}
}
