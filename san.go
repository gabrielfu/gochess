package chessago

import (
	"fmt"
	"regexp"
)

var sanToPieceType = map[string]PieceType{
	"":  PAWN,
	"P": PAWN,
	"N": KNIGHT,
	"B": BISHOP,
	"R": ROOK,
	"Q": QUEEN,
	"K": KING,
}

func parseSAN(san string, b *Board) (*Move, error) {
	if san == "0-0" || san == "O-O" || san == "o-o" {
		return NewCastlingMove(E1, G1, WHITE_KING, WHITE_KING_SIDE), nil
	}
	if san == "0-0-0" || san == "O-O-O" || san == "o-o-o" {
		return NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE), nil
	}

	// cases
	// a4, a8Q, a8=Q, axb4, Nh4, Ng3h4, Ngh4, N3h4, Nxh4, Ngxh4, N3xh4, Ng3xh4

	// color
	// disambiguation (including pinning, e.g., two knights can attack a square but one is pinned, the san is not disambiguated)

	var matches []string
	var pieceType string
	var fromFile string
	var fromRank string
	var toFile string
	var toRank string
	var promotion string

	// parse moves without capture (b4, Nh4)
	if matches = regexp.MustCompile(`^([PRNBKQ])?([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		toFile = matches[2]
		toRank = matches[3]
	} else

	// parse pawn promotion (b8Q, b8=Q)
	if matches = regexp.MustCompile(`^P?([a-h])([18])=?([RNBQ])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[0]
		toFile = matches[1]
		toRank = matches[2]
		promotion = matches[3]
	} else

	// parse pawn moves with capture (axb4)
	if matches = regexp.MustCompile(`^([a-h])?x([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = ""
		toFile = matches[2]
		toRank = matches[3]
		fromFile = matches[1]
	} else

	// parse piece moves with capture (Nxh4)
	if matches = regexp.MustCompile(`^([RNBKQ])?x([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		toFile = matches[2]
		toRank = matches[3]
	} else

	// parse piece moves with & without capture (Ngh4, Ngxh4)
	if matches = regexp.MustCompile(`^([RNBQ])?([a-h])x?([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		fromFile = matches[2]
		toFile = matches[3]
		toRank = matches[4]
	} else

	// parse piece moves with & without capture (N2h4, N2xh4)
	if matches = regexp.MustCompile(`^([RNBQ])?([1-8])x?([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		fromRank = matches[2]
		toFile = matches[3]
		toRank = matches[4]
	} else

	// parse piece moves with & without capture (Ng2h4, Ng2xh4)
	if matches = regexp.MustCompile(`^([RNBQ])?([a-h])([1-8])x?([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		fromFile = matches[2]
		fromRank = matches[3]
		toFile = matches[4]
		toRank = matches[5]
	}

	// fmt.Printf("pieceType=%v toFile=%v toRank=%v\n", sanToPieceType[pieceType], toFile, toRank)

	if matches == nil {
		return nil, fmt.Errorf("Invalid SAN: " + san)
	}

	piece := PieceFromTypeColor(sanToPieceType[pieceType], b.Turn())
	to := SquareFromAlgebraic(toFile + toRank)

	var from Square
	// infer "from" square from the board
	legalMoves := b.LegalMovesForPiece([]Piece{piece})
	legalMoves = FilterMoves(legalMoves, func(move *Move) bool {
		cond := move.To() == to
		// disambiguate
		if fromFile != "" {
			cond = cond && move.From().File() == ParseFile(fromFile)
		}
		if fromRank != "" {
			cond = cond && move.From().Rank() == ParseRank(fromRank)
		}
		return cond
	})
	switch len(legalMoves) {
	case 0:
		return nil, fmt.Errorf("Invalid SAN: " + san)
	case 1:
		from = legalMoves[0].From()
	default:
		return nil, fmt.Errorf("Ambiguous SAN: " + san)
	}

	move := NewMove(from, to, piece)
	if promotion != "" {
		move.SetPromotion(PieceFromTypeColor(sanToPieceType[promotion], b.Turn()))
	}
	return move, nil
}
