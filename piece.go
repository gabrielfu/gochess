package main

type Color uint8

const (
	WHITE Color = iota
	BLACK
)

type Piece uint8

const (
	WHITE_PAWN Piece = iota
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
	EMPTY
)

var WHITE_PIECES = []Piece{
	WHITE_PAWN, WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN, WHITE_KING,
}

var BLACK_PIECES = []Piece{
	BLACK_PAWN, BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN, BLACK_KING,
}

var ALL_PIECES = append(WHITE_PIECES, BLACK_PIECES...)

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

func (p Piece) Color() Color {
	if p < 6 {
		return WHITE
	} else {
		return BLACK
	}
}

func (p Piece) Symbol() string {
	return SYMBOLS[p]
}
