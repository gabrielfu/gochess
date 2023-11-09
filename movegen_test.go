package chessago

import "testing"

func TestCalcKingMoves(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 49216},
		{D8, 2898066360212914176},
		{H3, 50463488},
		{F4, 60298231808},
		{E5, 30872694685696},
	}
	for _, tt := range tests {
		moves := calcKingMoves(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}

func TestCalcKnightMoves(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 4202496},
		{D8, 19184278881435648},
		{H3, 8657044482},
		{F4, 11068131838464},
		{E5, 5666883501293568},
	}
	for _, tt := range tests {
		moves := calcKnightMoves(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}

func TestCalcWhitePawnMoves(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 32768},
		{D8, 0},
		{H3, 16777216},
		{F4, 17179869184},
		{E5, 8796093022208},
	}
	for _, tt := range tests {
		moves := calcWhitePawnMoves(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}

func TestCalcWhitePawnAttacks(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 16384},
		{D8, 0},
		{H3, 33554432},
		{F4, 42949672960},
		{E5, 21990232555520},
	}
	for _, tt := range tests {
		moves := calcWhitePawnAttacks(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}

func TestCalcBlackPawnMoves(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 0},
		{D8, 4503599627370496},
		{H3, 256},
		{F4, 262144},
		{E5, 134217728},
	}
	for _, tt := range tests {
		moves := calcBlackPawnMoves(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}

func TestCalcBlackPawnAttacks(t *testing.T) {
	var tests = []struct {
		sq       Square
		expected Bitboard
	}{
		{A1, 0},
		{D8, 11258999068426240},
		{H3, 512},
		{F4, 655360},
		{E5, 335544320},
	}
	for _, tt := range tests {
		moves := calcBlackPawnAttacks(tt.sq)
		if moves != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, moves)
		}
	}
}
