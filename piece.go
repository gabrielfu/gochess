package chessago

type Color uint8

const (
	WHITE Color = iota
	BLACK
	NO_COLOR
)

func (c Color) String() string {
	switch c {
	case WHITE:
		return "White"
	case BLACK:
		return "Black"
	default:
		return "No Color"
	}
}

type PieceType uint8

const (
	PAWN PieceType = iota
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
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

func (p Piece) PieceType() PieceType {
	return PieceType(uint8(p) % 6)
}

func (p Piece) Symbol() string {
	return SYMBOLS[p]
}

func PieceFromTypeColor(pieceType PieceType, color Color) Piece {
	return Piece(uint8(pieceType) + uint8(color)*6)
}
