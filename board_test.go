package main

import "testing"

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
	b.Move(&Move{
		From:  D2,
		To:    D4,
		Piece: WHITE_PAWN,
	})
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
	b.Move(&Move{
		From:  D2,
		To:    E3,
		Piece: WHITE_PAWN,
	})
	if b.whitePawns != 1<<E3 {
		t.Errorf("Expected 1<<E3, got %v", b.whitePawns)
	}
	if b.blackRooks != 0 {
		t.Errorf("Expected 0, got %v", b.blackRooks)
	}
}
