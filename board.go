package main

import "fmt"

type Bitboard uint64

func (bb Bitboard) BinaryBoard() string {
	bin := fmt.Sprintf("%064b", bb)
	out := ""
	for i := 0; i < 8; i++ {
		out += bin[i*8:i*8+8] + "\n"
	}
	return out[:len(out)-1]
}

// Squares returns a slice of squares that are set in the bitboard.
func (bb Bitboard) Squares() []Square {
	squares := []Square{}
	for i := 0; i < 64; i++ {
		if bb&(1<<uint(i)) != 0 {
			squares = append(squares, Square(i))
		}
	}
	return squares
}

// A8 B8 C8 D8 E8 F8 G8 H8 ;
// A7 B7 C7 D7 E7 F7 G7 H7 ;
// A6 B6 C6 D6 E6 F6 G6 H6 ;
// A5 B5 C5 D5 E5 F5 G5 H5 ;
// A4 B4 C4 D4 E4 F4 G4 H4 ;
// A3 B3 C3 D3 E3 F3 G3 H3 ;
// A2 B2 C2 D2 E2 F2 G2 H2 ;
// A1 B1 C1 D1 E1 F1 G1 H1 ;
type Square int

// Count from LSB.
// H1 = 0, G1 = 1, ..., B8 = 62, A8 = 63.
const (
	H1 Square = iota
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

func (sq Square) BinaryBoard() string {
	return Bitboard(1 << sq).BinaryBoard()
}

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

func (sq Square) String() string {
	return SQUARE_NAMES[sq]
}

type File int

const (
	H File = iota
	G
	F
	E
	D
	C
	B
	A
)

var FILE_NAMES = []string{
	"h", "g", "f", "e", "d", "c", "b", "a",
}

func (f File) String() string {
	return FILE_NAMES[f]
}

type Rank int

const (
	R1 Rank = iota
	R2
	R3
	R4
	R5
	R6
	R7
	R8
)

var RANK_NAMES = []string{
	"R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8",
}

func (r Rank) String() string {
	return RANK_NAMES[r]
}

func (sq Square) File() File {
	return File(sq % 8)
}

func (sq Square) Rank() Rank {
	return Rank(sq / 8)
}

// Bitboard representation of a chess board.
type Board struct {
	whitePawns   Bitboard
	whiteKnights Bitboard
	whiteBishops Bitboard
	whiteRooks   Bitboard
	whiteQueens  Bitboard
	whiteKing    Bitboard
	blackPawns   Bitboard
	blackKnights Bitboard
	blackBishops Bitboard
	blackRooks   Bitboard
	blackQueens  Bitboard
	blackKing    Bitboard

	turn           Color
	castlingRights uint8
	enPassant      uint8

	whiteOccupied Bitboard
	blackOccupied Bitboard
	allOccupied   Bitboard
}

// NewEmptyBoard returns a new board with no pieces.
func NewEmptyBoard() *Board {
	return &Board{
		whitePawns:     0,
		whiteKnights:   0,
		whiteBishops:   0,
		whiteRooks:     0,
		whiteQueens:    0,
		whiteKing:      0,
		blackPawns:     0,
		blackKnights:   0,
		blackBishops:   0,
		blackRooks:     0,
		blackQueens:    0,
		blackKing:      0,
		turn:           WHITE,
		castlingRights: 0,
		enPassant:      0,
		whiteOccupied:  0,
		blackOccupied:  0,
		allOccupied:    0,
	}
}

// NewStartingBoard returns a new board with the starting position.
func NewStartingBoard() *Board {
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
		turn:           WHITE,
		castlingRights: 0,
		enPassant:      0,
		whiteOccupied:  0x000000000000ffff,
		blackOccupied:  0xffff000000000000,
		allOccupied:    0xffff00000000ffff,
	}
}

func (b *Board) AddPieceToSquare(p Piece, sq Square) {
	bb := b.GetBbForPiece(p)
	if bb == nil {
		return
	}
	*bb |= (1 << sq)
	b.UpdateOccupied()
}

func (b *Board) UpdateOccupied() {
	b.whiteOccupied = b.whitePawns | b.whiteKnights | b.whiteBishops | b.whiteRooks | b.whiteQueens | b.whiteKing
	b.blackOccupied = b.blackPawns | b.blackKnights | b.blackBishops | b.blackRooks | b.blackQueens | b.blackKing
	b.allOccupied = b.whiteOccupied | b.blackOccupied
}

func (b *Board) Visualize() string {
	out := ""
	for i := 63; i >= 0; i-- {
		piece := b.GetPieceAtSquare(Square(i))
		out += piece.Symbol()
		if i%8 == 0 {
			out += "\n"
		}
	}
	return out
}

// GetPieceAtSquare returns the piece at the given square (0-63).
func (b *Board) GetPieceAtSquare(square Square) Piece {
	var mask Bitboard = 1 << square
	if b.whitePawns&mask != 0 {
		return WHITE_PAWN
	} else if b.whiteKnights&mask != 0 {
		return WHITE_KNIGHT
	} else if b.whiteBishops&mask != 0 {
		return WHITE_BISHOP
	} else if b.whiteRooks&mask != 0 {
		return WHITE_ROOK
	} else if b.whiteQueens&mask != 0 {
		return WHITE_QUEEN
	} else if b.whiteKing&mask != 0 {
		return WHITE_KING
	} else if b.blackPawns&mask != 0 {
		return BLACK_PAWN
	} else if b.blackKnights&mask != 0 {
		return BLACK_KNIGHT
	} else if b.blackBishops&mask != 0 {
		return BLACK_BISHOP
	} else if b.blackRooks&mask != 0 {
		return BLACK_ROOK
	} else if b.blackQueens&mask != 0 {
		return BLACK_QUEEN
	} else if b.blackKing&mask != 0 {
		return BLACK_KING
	} else {
		return EMPTY
	}
}

func (b *Board) GetBbForPiece(p Piece) *Bitboard {
	switch p {
	case WHITE_PAWN:
		return &b.whitePawns
	case WHITE_KNIGHT:
		return &b.whiteKnights
	case WHITE_BISHOP:
		return &b.whiteBishops
	case WHITE_ROOK:
		return &b.whiteRooks
	case WHITE_QUEEN:
		return &b.whiteQueens
	case WHITE_KING:
		return &b.whiteKing
	case BLACK_PAWN:
		return &b.blackPawns
	case BLACK_KNIGHT:
		return &b.blackKnights
	case BLACK_BISHOP:
		return &b.blackBishops
	case BLACK_ROOK:
		return &b.blackRooks
	case BLACK_QUEEN:
		return &b.blackQueens
	case BLACK_KING:
		return &b.blackKing
	default:
		return nil
	}
}

func (b *Board) Move(move *Move) error {
	piece := b.GetPieceAtSquare(move.From)
	if piece != move.Piece {
		return fmt.Errorf("Piece mismatch: " + piece.Symbol() + " != " + move.Piece.Symbol())
	}

	target := b.GetPieceAtSquare(move.To)

	bb := b.GetBbForPiece(piece)
	if bb == nil || *bb == 0 {
		return fmt.Errorf("No piece at square: " + move.From.String())
	}
	if target != EMPTY {
		targetBb := b.GetBbForPiece(target)
		if targetBb == nil || *targetBb == 0 {
			return fmt.Errorf("Invalid target piece: " + target.Symbol())
		}
		// capture target
		*targetBb &^= (1 << move.To)
	}
	// clear from "From" and set to "To"
	*bb &^= (1 << move.From)
	*bb |= (1 << move.To)

	// Update occupied bitboards
	b.UpdateOccupied()
	// Next player's turn
	b.turn = 1 - b.turn
	return nil
}
