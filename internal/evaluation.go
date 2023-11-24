package gochess

import (
	"math"
)

var BASE_VALUES = map[Piece]int{
	WHITE_PAWN:   100,
	WHITE_KNIGHT: 320,
	WHITE_BISHOP: 330,
	WHITE_ROOK:   500,
	WHITE_QUEEN:  900,
	WHITE_KING:   5000,
	BLACK_PAWN:   -100,
	BLACK_KNIGHT: -320,
	BLACK_BISHOP: -330,
	BLACK_ROOK:   -500,
	BLACK_QUEEN:  -900,
	BLACK_KING:   -5000,
}

const (
	MAX_EVAL = 100000
	MIN_EVAL = -100000
)

var PAWN_W_ADJ = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 25, 25, 10, 5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}
var KNIGHT_W_ADJ = []int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}
var BISHOP_W_ADJ = []int{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}
var ROOK_W_ADJ = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, 10, 10, 10, 10, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 0, 0, 0,
}
var QUEEN_W_ADJ = []int{
	-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 5, 5, 5, 0, -5,
	0, 0, 5, 5, 5, 5, 0, -5,
	-10, 5, 5, 5, 5, 5, 0, -10,
	-10, 0, 5, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20,
}
var KING_W_ADJ = []int{
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20,
}

var PAWN_B_ADJ []int
var KNIGHT_B_ADJ []int
var BISHOP_B_ADJ []int
var ROOK_B_ADJ []int
var QUEEN_B_ADJ []int
var KING_B_ADJ []int

func flip(arr []int) []int {
	flipped := make([]int, len(arr))
	for i, val := range arr {
		flipped[len(arr)-1-i] = val
	}
	return flipped
}

func negative(arr []int) []int {
	negative := make([]int, len(arr))
	for i, val := range arr {
		negative[i] = -val
	}
	return negative
}

func init() {
	PAWN_B_ADJ = negative(flip(PAWN_W_ADJ))
	KNIGHT_B_ADJ = negative(flip(KNIGHT_W_ADJ))
	BISHOP_B_ADJ = negative(flip(BISHOP_W_ADJ))
	ROOK_B_ADJ = negative(flip(ROOK_W_ADJ))
	QUEEN_B_ADJ = negative(flip(QUEEN_W_ADJ))
	KING_B_ADJ = negative(flip(KING_W_ADJ))
}

func lookupAdj(piece Piece, sq Square) int {
	switch piece {
	case WHITE_PAWN:
		return PAWN_W_ADJ[sq]
	case WHITE_KNIGHT:
		return KNIGHT_W_ADJ[sq]
	case WHITE_BISHOP:
		return BISHOP_W_ADJ[sq]
	case WHITE_ROOK:
		return ROOK_W_ADJ[sq]
	case WHITE_QUEEN:
		return QUEEN_W_ADJ[sq]
	case WHITE_KING:
		return KING_W_ADJ[sq]
	case BLACK_PAWN:
		return PAWN_B_ADJ[sq]
	case BLACK_KNIGHT:
		return KNIGHT_B_ADJ[sq]
	case BLACK_BISHOP:
		return BISHOP_B_ADJ[sq]
	case BLACK_ROOK:
		return ROOK_B_ADJ[sq]
	case BLACK_QUEEN:
		return QUEEN_B_ADJ[sq]
	case BLACK_KING:
		return KING_B_ADJ[sq]
	default:
		return 0
	}
}

func Evaluate(b *Board) int {
	eval := 0
	for _, piece := range ALL_PIECES {
		bb := b.GetBbForPiece(piece)
		for _, sq := range bb.Squares() {
			eval += BASE_VALUES[piece] + lookupAdj(piece, sq)
		}
	}
	return eval
}

func EvaluationBar(eval int, width int) string {
	var rescaled float64
	if eval == MAX_EVAL {
		rescaled = 1
	} else if eval == MIN_EVAL {
		rescaled = 0
	} else {
		pad := 1 / float64(width)
		if eval > 640 {
			rescaled = 1
		} else if eval < -640 {
			rescaled = 0
		} else {
			absEval := math.Abs(float64(eval))
			rescaled = math.Log10(absEval/100+1)/2.2 + 0.5
			if eval < 0 {
				rescaled = 1 - rescaled
			}
		}
		// transform rescale from [0, 1] to [pad, 1-pad]
		rescaled = rescaled*(1-pad*2) + pad
	}
	var rounded int
	if eval >= 0 {
		rounded = int(math.Floor(rescaled * float64(width)))
	} else {
		rounded = int(math.Ceil(rescaled * float64(width)))
	}
	out := ""
	for i := 0; i < width; i++ {
		if i < rounded {
			out += "\u2588"
		} else {
			out += "\u2591"
		}
	}
	return out
}
