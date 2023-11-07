package main

type Game struct {
	board Board
}

func NewGame() *Game {
	return &Game{
		board: *NewBoard(),
	}
}

// func (g *Game) LegalMoves() []*Move {
// 	candidatePieces := []Piece{}
// 	var allowedTos bitboard
// 	switch g.board.turn {
// 	case WHITE:
// 		candidatePieces = WHITE_PIECES
// 		allowedTos = ^g.board.whiteOccupied
// 	case BLACK:
// 		candidatePieces = BLACK_PIECES
// 		allowedTos = ^g.board.blackOccupied
// 	}

// 	moves := []*Move{}
// 	for _, p := range candidatePieces {
// 		bb := g.board.GetBbForPiece(p)
// 		if bb == 0 {
// 			continue
// 		}

// 		for to := range allowedTos.Iterate() {
// 		}

// 	}
// }
