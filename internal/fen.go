package gochess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var pieceFEN = []string{
	"P", "N", "B", "R", "Q", "K",
	"p", "n", "b", "r", "q", "k",
	".",
}

func pieceToFEN(p Piece) string {
	return pieceFEN[p]
}

func pieceFromFEN(fen string) Piece {
	for i, piece := range pieceFEN {
		if piece == fen {
			return Piece(i)
		}
	}
	return EMPTY
}

func colorToFEN(c Color) string {
	if c == WHITE {
		return "w"
	} else {
		return "b"
	}
}

func castlingToFEN(b *Board) string {
	fen := ""
	castlingRights := b.CastlingRights()
	if castlingRights.Has(WHITE_KING_SIDE) {
		fen += "K"
	}
	if castlingRights.Has(WHITE_QUEEN_SIDE) {
		fen += "Q"
	}
	if castlingRights.Has(BLACK_KING_SIDE) {
		fen += "k"
	}
	if castlingRights.Has(BLACK_QUEEN_SIDE) {
		fen += "q"
	}
	if fen == "" {
		fen = "-"
	}
	return fen
}

func enPassantToFEN(b *Board) string {
	fen := "-"
	enPassantSquare := b.EnPassantSquare()
	if !enPassantSquare.IsEmpty() {
		fen = enPassantSquare.Squares()[0].String()
	}
	return fen
}

func toFEN(b *Board, halfMoveClock int, moveCount int) string {
	fen := ""
	for rank := 7; rank >= 0; rank-- {
		emptySquares := 0
		for file := 7; file >= 0; file-- {
			square := Square(file + rank*8)
			piece := b.GetPieceAtSquare(square)
			if piece == EMPTY {
				emptySquares++
			} else {
				if emptySquares > 0 {
					fen += fmt.Sprint(emptySquares)
					emptySquares = 0
				}
				fen += pieceToFEN(piece)
			}
		}
		if emptySquares > 0 {
			fen += fmt.Sprint(emptySquares)
		}
		if rank > 0 {
			fen += "/"
		}
	}
	fen += " "
	fen += colorToFEN(b.Turn())
	fen += " "
	fen += castlingToFEN(b)
	fen += " "
	fen += enPassantToFEN(b)
	fen += " "
	fen += fmt.Sprint(halfMoveClock)
	fen += " "
	fen += fmt.Sprint(moveCount)
	return fen
}

func ToFEN(g *Game) string {
	return toFEN(g.Board(), g.HalfMoveClock(), g.MoveCount())
}

func parsePieces(b *Board, fen string) {
	rank := 7
	file := 7
	for _, c := range fen {
		if c == '/' {
			rank--
			file = 7
		} else if c >= '1' && c <= '8' {
			file -= int(c - '0')
		} else {
			piece := pieceFromFEN(string(c))
			b.AddPieceToSquare(piece, Square(file+rank*8))
			file--
		}
	}
}

func parseTurn(b *Board, fen string) {
	if fen == "w" {
		b.SetTurn(WHITE)
	} else {
		b.SetTurn(BLACK)
	}
}

func parseCastlingRights(b *Board, fen string) {
	if fen == "-" {
		return
	}
	for _, c := range fen {
		switch c {
		case 'K':
			b.SetCastlingRights(b.CastlingRights().Add(WHITE_KING_SIDE))
		case 'Q':
			b.SetCastlingRights(b.CastlingRights().Add(WHITE_QUEEN_SIDE))
		case 'k':
			b.SetCastlingRights(b.CastlingRights().Add(BLACK_KING_SIDE))
		case 'q':
			b.SetCastlingRights(b.CastlingRights().Add(BLACK_QUEEN_SIDE))
		}
	}
}

func parseEnPassantSquare(b *Board, fen string) {
	if fen == "-" {
		return
	}
	square := SquareFromAlgebraic(fen)
	b.SetEnPassantSquare(square)
}

func ParseFEN(fen string) (*Game, error) {
	fenParts := strings.Split(fen, " ")
	if len(fenParts) != 6 {
		return nil, errors.New("invalid FEN")
	}

	b := NewEmptyBoard()
	parsePieces(b, fenParts[0])
	parseTurn(b, fenParts[1])
	parseCastlingRights(b, fenParts[2])
	parseEnPassantSquare(b, fenParts[3])
	halfMoveClock, err := strconv.Atoi(fenParts[4])
	if err != nil {
		return nil, err
	}
	moveCount, err := strconv.Atoi(fenParts[5])
	if err != nil {
		return nil, err
	}

	g := NewGame()
	g.SetBoard(b)
	g.SetHalfMoveClock(halfMoveClock)
	g.SetMoveCount(moveCount)
	return g, nil
}
