package gochess

import (
	"fmt"
)

type Bitboard uint64

func (bb Bitboard) BinaryBoard() string {
	bin := fmt.Sprintf("%064b", bb)
	out := ""
	for i := 0; i < 8; i++ {
		out += bin[i*8:i*8+8] + "\n"
	}
	return out[:len(out)-1]
}

func (bb Bitboard) IsEmpty() bool {
	return bb == 0
}

func (bb Bitboard) IsSingular() bool {
	return bb&(bb-1) == 0
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

func (bb Bitboard) SquareIsSet(sq Square) bool {
	return bb&(1<<uint(sq)) != 0
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

var LIGHT_SQUARES = Bitboard(0x55aa55aa55aa55aa)
var DARK_SQUARES = Bitboard(0xaa55aa55aa55aa55)

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

func ParseFile(file string) File {
	return File(7 - (file[0] - 'a'))
}

func ParseRank(rank string) Rank {
	return Rank(rank[0] - '1')
}

func SquareFromAlgebraic(algebraic string) Square {
	return SquareFromFileRank(ParseFile(algebraic[0:1]), ParseRank(algebraic[1:2]))
}

func SquareFromFileRank(file File, rank Rank) Square {
	return Square(uint8(rank)*8 + uint8(file))
}

type Castle uint8

const (
	WHITE_KING_SIDE Castle = 1 << iota
	WHITE_QUEEN_SIDE
	BLACK_KING_SIDE
	BLACK_QUEEN_SIDE
)

type CastlingRights uint8

func (c CastlingRights) Has(right Castle) bool {
	return uint8(c)&uint8(right) != 0
}

func (c CastlingRights) Remove(right Castle) CastlingRights {
	return CastlingRights(uint8(c) &^ uint8(right))
}

func (c CastlingRights) Add(right Castle) CastlingRights {
	return CastlingRights(uint8(c) | uint8(right))
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
	castlingRights CastlingRights

	enPassantSquare Bitboard
	whiteOccupied   Bitboard
	blackOccupied   Bitboard
	allOccupied     Bitboard
}

// NewEmptyBoard returns a new board with no pieces.
func NewEmptyBoard() *Board {
	return &Board{
		whitePawns:      0,
		whiteKnights:    0,
		whiteBishops:    0,
		whiteRooks:      0,
		whiteQueens:     0,
		whiteKing:       0,
		blackPawns:      0,
		blackKnights:    0,
		blackBishops:    0,
		blackRooks:      0,
		blackQueens:     0,
		blackKing:       0,
		turn:            WHITE,
		castlingRights:  CastlingRights(WHITE_KING_SIDE | WHITE_QUEEN_SIDE | BLACK_KING_SIDE | BLACK_QUEEN_SIDE),
		enPassantSquare: 0,
		whiteOccupied:   0,
		blackOccupied:   0,
		allOccupied:     0,
	}
}

// NewStartingBoard returns a new board with the starting position.
func NewStartingBoard() *Board {
	return &Board{
		whitePawns:      0x000000000000ff00,
		whiteKnights:    0x0000000000000042,
		whiteBishops:    0x0000000000000024,
		whiteRooks:      0x0000000000000081,
		whiteQueens:     0x0000000000000010,
		whiteKing:       0x0000000000000008,
		blackPawns:      0x00ff000000000000,
		blackKnights:    0x4200000000000000,
		blackBishops:    0x2400000000000000,
		blackRooks:      0x8100000000000000,
		blackQueens:     0x1000000000000000,
		blackKing:       0x0800000000000000,
		turn:            WHITE,
		castlingRights:  CastlingRights(WHITE_KING_SIDE | WHITE_QUEEN_SIDE | BLACK_KING_SIDE | BLACK_QUEEN_SIDE),
		enPassantSquare: 0,
		whiteOccupied:   0x000000000000ffff,
		blackOccupied:   0xffff000000000000,
		allOccupied:     0xffff00000000ffff,
	}
}

func (b *Board) Copy() *Board {
	return &Board{
		whitePawns:      b.whitePawns,
		whiteKnights:    b.whiteKnights,
		whiteBishops:    b.whiteBishops,
		whiteRooks:      b.whiteRooks,
		whiteQueens:     b.whiteQueens,
		whiteKing:       b.whiteKing,
		blackPawns:      b.blackPawns,
		blackKnights:    b.blackKnights,
		blackBishops:    b.blackBishops,
		blackRooks:      b.blackRooks,
		blackQueens:     b.blackQueens,
		blackKing:       b.blackKing,
		turn:            b.turn,
		castlingRights:  b.castlingRights,
		enPassantSquare: b.enPassantSquare,
		whiteOccupied:   b.whiteOccupied,
		blackOccupied:   b.blackOccupied,
		allOccupied:     b.allOccupied,
	}
}

func (b *Board) Turn() Color {
	return b.turn
}

func (b *Board) SetTurn(turn Color) {
	b.turn = turn
}

// SwitchTurn switches the turn to the next player.
func (b *Board) SwitchTurn() {
	b.turn = 1 - b.turn
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
		if i%8 == 7 {
			out += fmt.Sprint(i/8+1) + " "
		}
		piece := b.GetPieceAtSquare(Square(i))
		out += piece.Symbol() + " "
		if i%8 == 0 {
			out += "\n"
		}
	}
	out += "  a b c d e f g h"
	return out
}

func (b *Board) VisualizeFlipped() string {
	out := ""
	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			out += fmt.Sprint(i/8+1) + " "
		}
		piece := b.GetPieceAtSquare(Square(i))
		out += piece.Symbol() + " "
		if i%8 == 7 {
			out += "\n"
		}
	}
	out += "  h g f e d c b a"
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

func (b *Board) makeCastleMove(move *Move) error {
	var rookMove *Move
	switch move.Castle() {
	case WHITE_KING_SIDE:
		rookMove = NewMove(H1, F1, WHITE_ROOK)
	case WHITE_QUEEN_SIDE:
		rookMove = NewMove(A1, D1, WHITE_ROOK)
	case BLACK_KING_SIDE:
		rookMove = NewMove(H8, F8, BLACK_ROOK)
	case BLACK_QUEEN_SIDE:
		rookMove = NewMove(A8, D8, BLACK_ROOK)
	}
	if err := b.makeMove(move); err != nil {
		return err
	}
	if err := b.makeMove(rookMove); err != nil {
		return err
	}
	b.UnsetEnPassantSquare()
	return nil
}

func (b *Board) SquareIsEnpassant(sq Square) bool {
	return b.enPassantSquare.SquareIsSet(sq)
}

func (b *Board) makeMove(move *Move) error {
	// move this to validation
	piece := b.GetPieceAtSquare(move.From())
	if piece != move.Piece() {
		return fmt.Errorf("Piece mismatch: " + piece.Symbol() + " != " + move.Piece().Symbol())
	}

	bb := b.GetBbForPiece(piece)
	if bb == nil || *bb == 0 {
		return fmt.Errorf("No piece at square: " + move.From().String())
	}

	// capture target
	if b.SquareIsEnpassant(move.To()) {
		// en passant captures
		if piece == WHITE_PAWN {
			move.SetCaptured(BLACK_PAWN)
			targetBb := b.GetBbForPiece(BLACK_PAWN)
			*targetBb &^= (1 << SquareFromFileRank(move.To().File(), R5))
		} else if piece == BLACK_PAWN {
			move.SetCaptured(WHITE_PAWN)
			targetBb := b.GetBbForPiece(WHITE_PAWN)
			*targetBb &^= (1 << SquareFromFileRank(move.To().File(), R4))
		}
	} else {
		// other captures
		target := b.GetPieceAtSquare(move.To())
		if target != EMPTY {
			move.SetCaptured(target)
			targetBb := b.GetBbForPiece(target)
			if targetBb == nil || *targetBb == 0 {
				return fmt.Errorf("Invalid target piece: " + target.Symbol())
			}
			*targetBb &^= (1 << move.To())

			// remove castling right if rook is captured
			if target == WHITE_ROOK && move.To() == A1 {
				b.castlingRights = b.castlingRights.Remove(WHITE_QUEEN_SIDE)
			} else if target == WHITE_ROOK && move.To() == H1 {
				b.castlingRights = b.castlingRights.Remove(WHITE_KING_SIDE)
			} else if target == BLACK_ROOK && move.To() == A8 {
				b.castlingRights = b.castlingRights.Remove(BLACK_QUEEN_SIDE)
			} else if target == BLACK_ROOK && move.To() == H8 {
				b.castlingRights = b.castlingRights.Remove(BLACK_KING_SIDE)
			}
		}
	}

	// move the piece
	// clear from "From" and set to "To"
	*bb &^= (1 << move.From())
	*bb |= (1 << move.To())

	// remove castling right if king or rook is moved
	if piece == WHITE_KING {
		b.castlingRights = b.castlingRights.Remove(WHITE_KING_SIDE).Remove(WHITE_QUEEN_SIDE)
	} else if piece == BLACK_KING {
		b.castlingRights = b.castlingRights.Remove(BLACK_KING_SIDE).Remove(BLACK_QUEEN_SIDE)
	} else if piece == WHITE_ROOK {
		if move.From() == A1 {
			b.castlingRights = b.castlingRights.Remove(WHITE_QUEEN_SIDE)
		} else if move.From() == H1 {
			b.castlingRights = b.castlingRights.Remove(WHITE_KING_SIDE)
		}
	} else if piece == BLACK_ROOK {
		if move.From() == A8 {
			b.castlingRights = b.castlingRights.Remove(BLACK_QUEEN_SIDE)
		} else if move.From() == H8 {
			b.castlingRights = b.castlingRights.Remove(BLACK_KING_SIDE)
		}
	}

	// promotion
	if move.Promotion() != EMPTY {
		*bb &^= (1 << move.To())
		promoBb := b.GetBbForPiece(move.Promotion())
		if promoBb == nil {
			return fmt.Errorf("Invalid promotion piece: " + move.Promotion().Symbol())
		}
		*promoBb |= (1 << move.To())
	}

	// mark en passant squares
	if piece == WHITE_PAWN && move.From().Rank() == R2 && move.To().Rank() == R4 {
		b.SetEnPassantSquare(SquareFromFileRank(move.From().File(), R3))
	} else if piece == BLACK_PAWN && move.From().Rank() == R7 && move.To().Rank() == R5 {
		b.SetEnPassantSquare(SquareFromFileRank(move.From().File(), R6))
	} else {
		b.UnsetEnPassantSquare()
	}

	return nil
}

func (b *Board) MoveSAN(san string) error {
	move, err := ParseSAN(san, b)
	if err != nil {
		return err
	}
	return b.Move(move)
}

// Move moves a piece on the board without legality validation.
func (b *Board) Move(move *Move) error {
	if move.Castle() != 0 {
		if err := b.makeCastleMove(move); err != nil {
			return err
		}
	} else {
		if err := b.makeMove(move); err != nil {
			return err
		}
	}

	// Update occupied bitboards
	b.UpdateOccupied()
	// Next player's turn
	b.turn = 1 - b.turn
	return nil
}

func (b *Board) IsInCheck() bool {
	if b.Turn() == WHITE {
		return isSquareAttacked(b.whiteKing.Squares()[0], b)
	}
	return isSquareAttacked(b.blackKing.Squares()[0], b)
}

func (b *Board) IsInCheckmate() bool {
	return b.IsInCheck() && (len(b.LegalMoves()) == 0)
}

func (b *Board) IsInStalemate() bool {
	return !b.IsInCheck() && (len(b.LegalMoves()) == 0)
}

func (b *Board) IsInsufficientMaterial() bool {
	if b.whitePawns != 0 || b.blackPawns != 0 || b.whiteRooks != 0 || b.blackRooks != 0 || b.whiteQueens != 0 || b.blackQueens != 0 {
		return false
	}

	bishops := b.whiteBishops | b.blackBishops
	knights := b.whiteKnights | b.blackKnights

	numBishops := len(bishops.Squares())
	numKnights := len(knights.Squares())

	if numBishops == 0 && numKnights == 0 {
		// K v. K
		return true
	} else if numBishops == 1 && numKnights == 0 {
		// KB v. K
		return true
	} else if numBishops == 0 && numKnights == 1 {
		// KN v. K
		return true
	} else if numBishops == 0 && numKnights == 2 && b.whiteKnights.IsSingular() && b.blackKnights.IsSingular() {
		// KNN v. K
		// theoretically not drawn with a blender, but we'll consider it drawn
		return true
	} else if numKnights == 0 && numBishops == 2 && b.whiteBishops.IsSingular() && b.blackBishops.IsSingular() {
		// KB v. KB
		// if bishops are opposite colours, theoretically not drawn with a blunder, but we'll consider it drawn
		return true
	}
	return false
}

// LegalMoves returns all legal moves for the current player and specified candidate pieces.
func (b *Board) LegalMovesForPiece(candidatePieces []Piece) []*Move {
	var allowedTos Bitboard
	var enemyOccupied Bitboard
	switch b.Turn() {
	case WHITE:
		allowedTos = ^b.whiteOccupied
		enemyOccupied = b.blackOccupied
	case BLACK:
		allowedTos = ^b.blackOccupied
		enemyOccupied = b.whiteOccupied
	}

	moves := []*Move{}
	for _, p := range candidatePieces {
		bb := b.GetBbForPiece(p)
		// If no more such pieces on the board, skip
		if bb == nil || *bb == 0 {
			continue
		}

		// For each "from" square
		for _, from := range bb.Squares() {
			var toBb Bitboard = Bitboard(0)
			switch p {
			case WHITE_KNIGHT, BLACK_KNIGHT:
				toBb = GetKnightMoves(Square(from)) & allowedTos
			case WHITE_PAWN:
				attackBb := GetWhitePawnAttacks(Square(from)) & (enemyOccupied | b.enPassantSquare)
				moveBb := (GetWhitePawnMoves(Square(from), b.allOccupied) & allowedTos) &^ enemyOccupied
				toBb = attackBb | moveBb
			case BLACK_PAWN:
				attackBb := GetBlackPawnAttacks(Square(from)) & (enemyOccupied | b.enPassantSquare)
				moveBb := (GetBlackPawnMoves(Square(from), b.allOccupied) & allowedTos) &^ enemyOccupied
				toBb = attackBb | moveBb
			case WHITE_BISHOP, BLACK_BISHOP:
				toBb = GetBishopMoves(Square(from), b.allOccupied) & allowedTos
			case WHITE_ROOK, BLACK_ROOK:
				toBb = GetRookMoves(Square(from), b.allOccupied) & allowedTos
			case WHITE_QUEEN, BLACK_QUEEN:
				toBb = GetQueenMoves(Square(from), b.allOccupied) & allowedTos
			case WHITE_KING, BLACK_KING:
				toBb = GetKingMoves(Square(from)) & allowedTos
			}

			// For each "to" square
			for _, to := range toBb.Squares() {
				// promotion
				if p == WHITE_PAWN && to.Rank() == R8 {
					for _, promo := range []Piece{WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN} {
						move := NewMove(Square(from), Square(to), p)
						move.SetPromotion(promo)
						moves = append(moves, move)
					}
				} else if p == BLACK_PAWN && to.Rank() == R1 {
					for _, promo := range []Piece{BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN} {
						move := NewMove(Square(from), Square(to), p)
						move.SetPromotion(promo)
						moves = append(moves, move)
					}
				} else {
					moves = append(moves, NewMove(Square(from), Square(to), p))
				}
			}

			// castling
			if p == WHITE_KING && from == E1 {
				if b.castlingRights.Has(WHITE_QUEEN_SIDE) {
					castlingMove := NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE)
					if isCastlingValid(castlingMove, b) {
						moves = append(moves, castlingMove)
					}
				}
				if b.castlingRights.Has(WHITE_KING_SIDE) {
					castlingMove := NewCastlingMove(E1, G1, WHITE_KING, WHITE_KING_SIDE)
					if isCastlingValid(castlingMove, b) {
						moves = append(moves, castlingMove)
					}
				}
			} else if p == BLACK_KING && from == E8 {
				if b.castlingRights.Has(BLACK_QUEEN_SIDE) {
					castlingMove := NewCastlingMove(E8, C8, BLACK_KING, BLACK_QUEEN_SIDE)
					if isCastlingValid(castlingMove, b) {
						moves = append(moves, castlingMove)
					}
				}
				if b.castlingRights.Has(BLACK_KING_SIDE) {
					castlingMove := NewCastlingMove(E8, G8, BLACK_KING, BLACK_KING_SIDE)
					if isCastlingValid(castlingMove, b) {
						moves = append(moves, castlingMove)
					}
				}
			}
		}
	}

	// After making a move, cannot be in check
	moves = FilterMoves(moves, func(move *Move) bool {
		cpy := b.Copy()
		err := cpy.Move(move)
		if err != nil {
			return false
		}
		// go back to original turn
		cpy.turn = 1 - cpy.turn
		return !cpy.IsInCheck()
	})
	return moves
}

var legalMovesCache = make(map[uint64][]*Move)

func (b *Board) legalMoves() []*Move {
	var candidatePieces []Piece
	switch b.Turn() {
	case WHITE:
		candidatePieces = WHITE_PIECES
	case BLACK:
		candidatePieces = BLACK_PIECES
	}
	return b.LegalMovesForPiece(candidatePieces)
}

// LegalMoves returns all legal moves for the current player.
func (b *Board) LegalMoves() []*Move {
	if moves, ok := legalMovesCache[ZobristHash(b)]; ok {
		return moves
	}
	moves := b.legalMoves()
	legalMovesCache[ZobristHash(b)] = moves
	return moves
}

func (b *Board) CastlingRights() CastlingRights {
	return b.castlingRights
}

func (b *Board) SetCastlingRights(cr CastlingRights) {
	b.castlingRights = cr
}

func (b *Board) EnPassantSquare() Bitboard {
	return b.enPassantSquare
}

func (b *Board) SetEnPassantSquare(sq Square) {
	b.enPassantSquare = Bitboard(1 << sq)
}

func (b *Board) UnsetEnPassantSquare() {
	b.enPassantSquare = Bitboard(0)
}
