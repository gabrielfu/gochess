package chessago

import (
	"fmt"
	"regexp"
)

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
	var piece string
	var toFile string
	var toRank string

	// parse moves without capture
	r := regexp.MustCompile(`^([PRNBKQ])?([a-h])([1-8])$`)
	if matches = r.FindStringSubmatch(san); matches != nil {
		piece = matches[1]
		toFile = matches[2]
		toRank = matches[3]
	}

	println(piece, toFile, toRank)

	// fromFile := san[0]
	return nil, fmt.Errorf("Invalid SAN: " + san)
}
