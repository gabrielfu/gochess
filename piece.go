package chessago

type Color uint8

const (
	WHITE Color = 0b0
	BLACK Color = 0b1
)

type PieceType uint8

const (
	PAWN   PieceType = 0b0010
	KNIGHT PieceType = 0b0100
	BISHOP PieceType = 0b0110
	ROOK   PieceType = 0b1000
	QUEEN  PieceType = 0b1010
	KING   PieceType = 0b1100
)

type Piece uint8

const (
	WHITE_PAWN   Piece = Piece(uint8(WHITE) | uint8(PAWN))
	WHITE_KNIGHT Piece = Piece(uint8(WHITE) | uint8(KNIGHT))
	WHITE_BISHOP Piece = Piece(uint8(WHITE) | uint8(BISHOP))
	WHITE_ROOK   Piece = Piece(uint8(WHITE) | uint8(ROOK))
	WHITE_QUEEN  Piece = Piece(uint8(WHITE) | uint8(QUEEN))
	WHITE_KING   Piece = Piece(uint8(WHITE) | uint8(KING))
	BLACK_PAWN   Piece = Piece(uint8(BLACK) | uint8(PAWN))
	BLACK_KNIGHT Piece = Piece(uint8(BLACK) | uint8(KNIGHT))
	BLACK_BISHOP Piece = Piece(uint8(BLACK) | uint8(BISHOP))
	BLACK_ROOK   Piece = Piece(uint8(BLACK) | uint8(ROOK))
	BLACK_QUEEN  Piece = Piece(uint8(BLACK) | uint8(QUEEN))
	BLACK_KING   Piece = Piece(uint8(BLACK) | uint8(KING))
	EMPTY        Piece = 0
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
	return SYMBOLS[p>>1-1]
}
