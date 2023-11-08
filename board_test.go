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
