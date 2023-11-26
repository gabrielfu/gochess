package gochess

import (
	"errors"
	"fmt"
)

type Status string

const (
	InProgress Status = "*"
	WhiteWon   Status = "1-0"
	BlackWon   Status = "0-1"
	Draw       Status = "1/2-1/2"
)

func (s Status) String() string {
	return string(s)
}

type Game struct {
	board *Board

	moves         []*Move
	positions     []*Board
	halfMoveClock int
	moveCount     int
	status        Status
	pgns          []string
}

func NewGame() *Game {
	b := NewStartingBoard()
	return &Game{
		board:         b,
		moves:         []*Move{},
		positions:     []*Board{b.Copy()},
		halfMoveClock: 0,
		moveCount:     0,
		status:        InProgress,
		pgns:          []string{},
	}
}

func (g *Game) Board() *Board {
	return g.board
}

func (g *Game) SetBoard(board *Board) {
	g.board = board
}

// LegalMoves returns all legal moves for the current player.
func (g *Game) LegalMoves() []*Move {
	return g.Board().LegalMoves()
}

func (g *Game) Visualize() string {
	return g.Board().Visualize()
}

func (g *Game) VisualizeFlipped() string {
	return g.Board().VisualizeFlipped()
}

func (g *Game) MoveSAN(san string) error {
	move, err := ParseSAN(san, g.Board())
	if err != nil {
		return err
	}
	return g.Move(move)
}

// Move executes the given move on the board.
func (g *Game) Move(move *Move) error {
	g.moves = append(g.moves, move)
	if g.Turn() == BLACK {
		g.moveCount += 1
	}
	if move.Piece().PieceType() == PAWN || move.Captured() != EMPTY {
		g.halfMoveClock = 0
	} else {
		g.halfMoveClock += 1
	}

	san := move.ToSAN(g.Board())
	turnNotation := ""
	if g.Turn() == WHITE {
		turnNotation = fmt.Sprintf("%d. ", g.MoveCount()+1)
	}
	pgn := g.PGN()
	pgn += turnNotation + san + " "

	err := g.Board().Move(move)
	if err != nil {
		return err
	}
	g.positions = append(g.positions, g.Board().Copy())

	if g.Board().IsInCheckmate() {
		winner := 1 - g.Turn()
		if winner == WHITE {
			g.status = WhiteWon
		} else {
			g.status = BlackWon
		}
		pgn += g.Status().String()
	}
	g.pgns = append(g.pgns, pgn)
	return err
}

func (g *Game) Undo() error {
	if len(g.moves) == 0 {
		return errors.New("no moves to undo")
	}

	g.moves = g.moves[:len(g.moves)-1]
	g.positions = g.positions[:len(g.positions)-1]
	g.pgns = g.pgns[:len(g.pgns)-1]
	g.halfMoveClock = max(g.halfMoveClock-1, 0)
	if g.Turn() == BLACK {
		g.moveCount -= 1
	}
	g.status = InProgress

	board := g.positions[len(g.positions)-1]
	g.SetBoard(board.Copy())
	return nil
}

func (g *Game) Turn() Color {
	return g.Board().Turn()
}

func (g *Game) MoveCount() int {
	return g.moveCount
}

func (g *Game) HalfMoveClock() int {
	return g.halfMoveClock
}

func (g *Game) Moves() []*Move {
	return g.moves
}

func (g *Game) Positions() []*Board {
	return g.positions
}

func (g *Game) Status() Status {
	return g.status
}

func (g *Game) PGN() string {
	if len(g.pgns) == 0 {
		return ""
	}
	return g.pgns[len(g.pgns)-1]
}
