package gochess

import "fmt"

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
	pgn           string
}

func NewGame() *Game {
	b := NewStartingBoard()
	return &Game{
		board:         b,
		moves:         []*Move{},
		positions:     []*Board{b},
		halfMoveClock: 0,
		moveCount:     0,
		status:        InProgress,
		pgn:           "",
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
	g.pgn += turnNotation + san + " "

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
		g.pgn += g.Status().String()
	}
	return err
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
	return g.pgn
}
