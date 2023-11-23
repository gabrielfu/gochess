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

var pieceTypeToSAN = map[PieceType]string{
	PAWN:   "",
	KNIGHT: "N",
	BISHOP: "B",
	ROOK:   "R",
	QUEEN:  "Q",
	KING:   "K",
}

func ParseSAN(san string, b *Board) (*Move, error) {
	// remove check and mate and brilliant symbols
	san = regexp.MustCompile(`([+#]?[!?]{0,2})$`).ReplaceAllString(san, "")

	if san == "0-0" || san == "O-O" || san == "o-o" {
		move := NewCastlingMove(E1, G1, WHITE_KING, WHITE_KING_SIDE)
		if isCastlingValid(move, b) {
			return move, nil
		} else {
			return nil, fmt.Errorf("Illegal SAN: " + san)
		}
	}
	if san == "0-0-0" || san == "O-O-O" || san == "o-o-o" {
		move := NewCastlingMove(E1, C1, WHITE_KING, WHITE_QUEEN_SIDE)
		if isCastlingValid(move, b) {
			return move, nil
		} else {
			return nil, fmt.Errorf("Illegal SAN: " + san)
		}
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

	// parse pawn moves without capture (b4, Nh4)
	if matches = regexp.MustCompile(`^P?([a-h])([2-7])$`).FindStringSubmatch(san); matches != nil {
		pieceType = ""
		toFile = matches[1]
		toRank = matches[2]
	} else

	// parse pawn moves with capture (axb4)
	if matches = regexp.MustCompile(`^P?([a-h])?x([a-h])([2-7])$`).FindStringSubmatch(san); matches != nil {
		pieceType = ""
		fromFile = matches[1]
		toFile = matches[2]
		toRank = matches[3]
	} else

	// parse pawn push promotion (b8Q, b8=Q)
	if matches = regexp.MustCompile(`^P?([a-h])([18])=?([RNBQ])$`).FindStringSubmatch(san); matches != nil {
		pieceType = ""
		toFile = matches[1]
		toRank = matches[2]
		promotion = matches[3]
	} else

	// parse pawn capture promotion (axb8Q, axb8=Q)
	if matches = regexp.MustCompile(`^P?([a-h])?x([a-h])([18])=?([RNBQ])$`).FindStringSubmatch(san); matches != nil {
		pieceType = ""
		fromFile = matches[1]
		toFile = matches[2]
		toRank = matches[3]
		promotion = matches[4]
	} else

	// parse pawn moves without capture (b4, Nh4)
	if matches = regexp.MustCompile(`^P?([a-h])([18])$`).FindStringSubmatch(san); matches != nil {
		return nil, fmt.Errorf("Missing promotion SAN: " + san)
	} else

	// parse pawn moves with capture (axb4)
	if matches = regexp.MustCompile(`^P?([a-h])?x([a-h])([18])$`).FindStringSubmatch(san); matches != nil {
		return nil, fmt.Errorf("Missing promotion SAN: " + san)
	} else

	// parse piece moves without capture (Nh4)
	if matches = regexp.MustCompile(`^([RNBKQ])?([a-h])([1-8])$`).FindStringSubmatch(san); matches != nil {
		pieceType = matches[1]
		toFile = matches[2]
		toRank = matches[3]
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

	if matches == nil {
		return nil, fmt.Errorf("Malformed SAN: " + san)
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
		return nil, fmt.Errorf("Illegal SAN: " + san)
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

// ToSAN converts a move to Standard Algebraic Notation (SAN). b is the board *before* the move.
func (m *Move) ToSAN(b *Board) string {
	switch m.Castle() {
	case WHITE_KING_SIDE, BLACK_KING_SIDE:
		return "O-O"
	case WHITE_QUEEN_SIDE, BLACK_QUEEN_SIDE:
		return "O-O-O"
	default:
	}

	san := pieceTypeToSAN[m.Piece().PieceType()]

	// disambiguate
	disambiguated := false
	legalMoves := b.LegalMovesForPiece([]Piece{m.Piece()})
	legalMoves = FilterMoves(legalMoves, func(move *Move) bool {
		return move.To() == m.To()
	})
	if len(legalMoves) > 1 {
		// disambiguate by file
		disByFile := FilterMoves(legalMoves, func(move *Move) bool {
			return move.From().File() == m.From().File()
		})
		if len(disByFile) == 1 {
			disambiguated = true
			san += m.From().File().String()
		} else {
			// disambiguate by rank
			disByRank := FilterMoves(legalMoves, func(move *Move) bool {
				return move.From().Rank() == m.From().Rank()
			})
			if len(disByRank) == 1 {
				disambiguated = true
				san += m.From().Rank().String()
			} else {
				// disambiguate by file and rank
				disambiguated = true
				san += m.From().String()
			}
		}
	}

	// check for capture
	var enemyOccupied Bitboard
	switch b.Turn() {
	case WHITE:
		enemyOccupied = b.blackOccupied
	case BLACK:
		enemyOccupied = b.whiteOccupied
	}
	isCapture := enemyOccupied.SquareIsSet(m.To()) || (m.Piece().PieceType() == PAWN && b.SquareIsEnpassant(m.To()))
	if isCapture {
		if !disambiguated && m.Piece().PieceType() == PAWN {
			san += m.From().File().String()
		}
		san += "x"
	}

	// add destination square
	san += m.To().String()

	// add promotion
	if m.Promotion() != EMPTY {
		san += "=" + pieceTypeToSAN[m.Promotion().PieceType()]
	}

	// add check and mate
	cpy := b.Copy()
	cpy.Move(m)
	if cpy.IsInCheckmate() {
		san += "#"
	} else if cpy.IsInCheck() {
		san += "+"
	}

	return san
}
