package gochess

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBitboardBinaryBoard(t *testing.T) {
	b := Bitboard(1 << 22)
	expected := "00000000\n00000000\n00000000\n00000000\n00000000\n01000000\n00000000\n00000000"
	if b.BinaryBoard() != expected {
		t.Errorf("Expected %v, got %v", expected, b.BinaryBoard())
	}
}

func TestBitboardSquares(t *testing.T) {
	var tests = []struct {
		b        Bitboard
		expected []Square
	}{
		{Bitboard(1 << 44), []Square{D6}},
		{Bitboard(1<<22 | 1<<23), []Square{B3, A3}},
	}
	for _, tt := range tests {
		if !cmp.Equal(tt.b.Squares(), tt.expected) {
			t.Errorf("Expected %v, got %v", tt.expected, tt.b.Squares())
		}
	}
}

func TestAddPieceToSquare(t *testing.T) {
	b := NewEmptyBoard()
	// Test adding a piece to an empty square.
	b.AddPieceToSquare(WHITE_PAWN, D2)
	if b.whitePawns != 1<<D2 {
		t.Errorf("Expected 1<<D2, got %v", b.whitePawns)
	}
	// Test adding a piece to a square with a pawn.
	b.AddPieceToSquare(BLACK_ROOK, D2)
	if b.blackRooks != 1<<D2 {
		t.Errorf("Expected 1<<D2, got %v", b.blackRooks)
	}
}

func TestGetPieceAtSquare(t *testing.T) {
	b := NewEmptyBoard()
	// Test getting a piece from an empty square.
	piece := b.GetPieceAtSquare(0)
	if piece != EMPTY {
		t.Errorf("Expected NoPiece, got %v", piece)
	}
	// Test getting a piece from a square with a pawn.
	b.AddPieceToSquare(WHITE_PAWN, D2)
	piece = b.GetPieceAtSquare(D2)
	if piece != WHITE_PAWN {
		t.Errorf("Expected WhitePawn, got %v", piece)
	}
	// Test getting a piece from a square with a black rook.
	b.AddPieceToSquare(BLACK_ROOK, E7)
	piece = b.GetPieceAtSquare(E7)
	if piece != BLACK_ROOK {
		t.Errorf("Expected BlackRook, got %v", piece)
	}
}

func TestMoveWithoutCapture(t *testing.T) {
	b := NewEmptyBoard()
	// Test moving a pawn without capture.
	b.AddPieceToSquare(WHITE_PAWN, D2)
	b.Move(NewMove(D2, D4, WHITE_PAWN))
	if b.whitePawns != 1<<D4 {
		t.Errorf("Expected 1<<D4, got %v", b.whitePawns)
	}
	if b.whitePawns&1<<D2 != 0 {
		t.Errorf("Expected 0, got %v", b.whitePawns&1<<D2)
	}
}

func TestMoveWithCapture(t *testing.T) {
	b := NewEmptyBoard()
	// Test moving a pawn with capture.
	b.AddPieceToSquare(WHITE_PAWN, D2)
	b.AddPieceToSquare(BLACK_ROOK, E3)
	b.Move(NewMove(D2, E3, WHITE_PAWN))
	if b.whitePawns != 1<<E3 {
		t.Errorf("Expected 1<<E3, got %v", b.whitePawns)
	}
	if b.blackRooks != 0 {
		t.Errorf("Expected 0, got %v", b.blackRooks)
	}
}

func TestStartingPosition(t *testing.T) {
	b := NewStartingBoard()
	mapping := map[Square]Piece{
		A1: WHITE_ROOK,
		B1: WHITE_KNIGHT,
		C1: WHITE_BISHOP,
		D1: WHITE_QUEEN,
		E1: WHITE_KING,
		F1: WHITE_BISHOP,
		G1: WHITE_KNIGHT,
		H1: WHITE_ROOK,
		A8: BLACK_ROOK,
		B8: BLACK_KNIGHT,
		C8: BLACK_BISHOP,
		D8: BLACK_QUEEN,
		E8: BLACK_KING,
		F8: BLACK_BISHOP,
		G8: BLACK_KNIGHT,
		H8: BLACK_ROOK,
	}
	for square, piece := range mapping {
		if b.GetPieceAtSquare(square) != piece {
			t.Errorf("Expected %v at Square %v, got %v", piece.Symbol(), square, b.GetPieceAtSquare(square).Symbol())
		}
	}
	for _, square := range []Square{H2, G2, F2, E2, D2, C2, B2, A2} {
		if b.GetPieceAtSquare(square) != WHITE_PAWN {
			t.Errorf("Expected %v at Square %v, got %v", WHITE_PAWN.Symbol(), square, b.GetPieceAtSquare(square).Symbol())
		}
	}
	for _, square := range []Square{H7, G7, F7, E7, D7, C7, B7, A7} {
		if b.GetPieceAtSquare(square) != BLACK_PAWN {
			t.Errorf("Expected %v at Square %v, got %v", BLACK_PAWN.Symbol(), square, b.GetPieceAtSquare(square).Symbol())
		}
	}
	for square := H3; square <= A6; square++ {
		if b.GetPieceAtSquare(square) != EMPTY {
			t.Errorf("Expected %v at Square %v, got %v", EMPTY.Symbol(), square, b.GetPieceAtSquare(square).Symbol())
		}
	}
}

func TestParseFile(t *testing.T) {
	var tests = []struct {
		algebraic string
		expected  File
	}{
		{"a", A},
		{"b", B},
		{"c", C},
		{"d", D},
		{"e", E},
		{"f", F},
		{"g", G},
		{"h", H},
	}
	for _, tt := range tests {
		if file := ParseFile(tt.algebraic); file != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, file)
		}
	}
}

func TestParseRank(t *testing.T) {
	var tests = []struct {
		algebraic string
		expected  Rank
	}{
		{"1", R1},
		{"2", R2},
		{"3", R3},
		{"4", R4},
		{"5", R5},
		{"6", R6},
		{"7", R7},
		{"8", R8},
	}
	for _, tt := range tests {
		if rank := ParseRank(tt.algebraic); rank != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, rank)
		}
	}
}

func TestSquareFromAlgebraic(t *testing.T) {
	var tests = []struct {
		algebraic string
		expected  Square
	}{
		{"a1", A1},
		{"c3", C3},
		{"f6", F6},
		{"h8", H8},
	}
	for _, tt := range tests {
		if square := SquareFromAlgebraic(tt.algebraic); square != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, square)
		}
	}
}
