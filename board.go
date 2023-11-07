package main

import "fmt"

// Bitboard representation of a chess board.
type Board struct {
	whitePawns   uint64
	whiteKnights uint64
	whiteBishops uint64
	whiteRooks   uint64
	whiteQueens  uint64
	whiteKing    uint64
	blackPawns   uint64
	blackKnights uint64
	blackBishops uint64
	blackRooks   uint64
	blackQueens  uint64
	blackKing    uint64

	whiteTurn      bool
	castlingRights uint8
	enPassant      uint8

	whiteOccupied uint64
	blackOccupied uint64
	allOccupied   uint64
}

// A8 B8 C8 D8 E8 F8 G8 H8 ;
// A7 B7 C7 D7 E7 F7 G7 H7 ;
// A6 B6 C6 D6 E6 F6 G6 H6 ;
// A5 B5 C5 D5 E5 F5 G5 H5 ;
// A4 B4 C4 D4 E4 F4 G4 H4 ;
// A3 B3 C3 D3 E3 F3 G3 H3 ;
// A2 B2 C2 D2 E2 F2 G2 H2 ;
// A1 B1 C1 D1 E1 F1 G1 H1 ;
type SQUARES int

// Count from LSB.
// H1 = 0, G1 = 1, ..., B8 = 62, A8 = 63.
const (
	H1 SQUARES = iota
	G1
	F1
	E1
	D1
	C1
	B1
	A1
	H2
	G2
	F2
	E2
	D2
	C2
	B2
	A2
	H3
	G3
	F3
	E3
	D3
	C3
	B3
	A3
	H4
	G4
	F4
	E4
	D4
	C4
	B4
	A4
	H5
	G5
	F5
	E5
	D5
	C5
	B5
	A5
	H6
	G6
	F6
	E6
	D6
	C6
	B6
	A6
	H7
	G7
	F7
	E7
	D7
	C7
	B7
	A7
	H8
	G8
	F8
	E8
	D8
	C8
	B8
	A8
)

var SQUARE_NAMES = []string{
	"h1", "g1", "f1", "e1", "d1", "c1", "b1", "a1",
	"h2", "g2", "f2", "e2", "d2", "c2", "b2", "a2",
	"h3", "g3", "f3", "e3", "d3", "c3", "b3", "a3",
	"h4", "g4", "f4", "e4", "d4", "c4", "b4", "a4",
	"h5", "g5", "f5", "e5", "d5", "c5", "b5", "a5",
	"h6", "g6", "f6", "e6", "d6", "c6", "b6", "a6",
	"h7", "g7", "f7", "e7", "d7", "c7", "b7", "a7",
	"h8", "g8", "f8", "e8", "d8", "c8", "b8", "a8",
}

type PIECES uint8

const (
	WHITE_PAWN PIECES = iota
	WHITE_KNIGHT
	WHITE_BISHOP
	WHITE_ROOK
	WHITE_QUEEN
	WHITE_KING
	BLACK_PAWN
	BLACK_KNIGHT
	BLACK_BISHOP
	BLACK_ROOK
	BLACK_QUEEN
	BLACK_KING
)

var SYMBOLS = []string{
	"♟", "♞", "♝", "♜", "♛", "♚",
	"♙", "♘", "♗", "♖", "♕", "♔",
	".",
}

// var SYMBOLS = []string{
// 	"P", "N", "B", "R", "Q", "K",
// 	"p", "n", "b", "r", "q", "k",
// 	".",
// }

func NewBoard() *Board {
	return &Board{
		whitePawns:     0x000000000000ff00,
		whiteKnights:   0x0000000000000042,
		whiteBishops:   0x0000000000000024,
		whiteRooks:     0x0000000000000081,
		whiteQueens:    0x0000000000000008,
		whiteKing:      0x0000000000000010,
		blackPawns:     0x00ff000000000000,
		blackKnights:   0x4200000000000000,
		blackBishops:   0x2400000000000000,
		blackRooks:     0x8100000000000000,
		blackQueens:    0x0800000000000000,
		blackKing:      0x1000000000000000,
		whiteTurn:      true,
		castlingRights: 0,
		enPassant:      0,
		whiteOccupied:  0x000000000000ffff,
		blackOccupied:  0xffff000000000000,
		allOccupied:    0xffff00000000ffff,
	}
}

func (b *Board) UpdateOccupied() {
	b.whiteOccupied = b.whitePawns | b.whiteKnights | b.whiteBishops | b.whiteRooks | b.whiteQueens | b.whiteKing
	b.blackOccupied = b.blackPawns | b.blackKnights | b.blackBishops | b.blackRooks | b.blackQueens | b.blackKing
	b.allOccupied = b.whiteOccupied | b.blackOccupied
}

func (b *Board) Print() {
	for i := 63; i >= 0; i-- {
		print(b.GetPieceAtSquare(uint8(i)))
		if i%8 == 0 {
			println()
		}
	}
}

// GetPieceAtSquare returns the piece at the given square (0-63).
func (b *Board) GetPieceAtSquare(square uint8) string {
	var mask uint64 = 1 << square
	if b.whitePawns&mask != 0 {
		return SYMBOLS[WHITE_PAWN]
	} else if b.whiteKnights&mask != 0 {
		return SYMBOLS[WHITE_KNIGHT]
	} else if b.whiteBishops&mask != 0 {
		return SYMBOLS[WHITE_BISHOP]
	} else if b.whiteRooks&mask != 0 {
		return SYMBOLS[WHITE_ROOK]
	} else if b.whiteQueens&mask != 0 {
		return SYMBOLS[WHITE_QUEEN]
	} else if b.whiteKing&mask != 0 {
		return SYMBOLS[WHITE_KING]
	} else if b.blackPawns&mask != 0 {
		return SYMBOLS[BLACK_PAWN]
	} else if b.blackKnights&mask != 0 {
		return SYMBOLS[BLACK_KNIGHT]
	} else if b.blackBishops&mask != 0 {
		return SYMBOLS[BLACK_BISHOP]
	} else if b.blackRooks&mask != 0 {
		return SYMBOLS[BLACK_ROOK]
	} else if b.blackQueens&mask != 0 {
		return SYMBOLS[BLACK_QUEEN]
	} else if b.blackKing&mask != 0 {
		return SYMBOLS[BLACK_KING]
	} else {
		return SYMBOLS[len(SYMBOLS)-1]
	}
}

func (b *Board) Move(move *Move) {
	fmt.Println(move)
	piece := b.GetPieceAtSquare(uint8(move.From))
	switch piece {
	case SYMBOLS[WHITE_PAWN]:
		b.whitePawns ^= 1 << move.From
		b.whitePawns |= 1 << move.To
	case SYMBOLS[WHITE_KNIGHT]:
		b.whiteKnights ^= 1 << move.From
		b.whiteKnights |= 1 << move.To
	case SYMBOLS[BLACK_PAWN]:
		b.blackPawns ^= 1 << move.From
		b.blackPawns |= 1 << move.To
	default:
		println("Unknown piece: " + piece)
	}
}
