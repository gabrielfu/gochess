package chessago

import "fmt"

type Game struct {
	board   *Board
	history []*Move
	pgn     string
}

func NewGame() *Game {
	InitMovesTables()

	return &Game{
		board:   NewStartingBoard(),
		history: []*Move{},
		pgn:     "",
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

// Move executes the given move on the board.
func (g *Game) Move(move *Move) error {
	san := move.ToSAN(g.Board())
	turnNotation := ""
	if g.Turn() == WHITE {
		turnNotation = fmt.Sprintf("%d. ", len(g.history)/2+1)
	}

	err := g.Board().Move(move)
	if err != nil {
		return err
	}
	g.history = append(g.history, move)
	g.pgn += turnNotation + san + " "
	return err
}

func (g *Game) Turn() Color {
	return g.Board().Turn()
}

func (g *Game) History() []*Move {
	return g.history
}

func (g *Game) PGN() string {
	return g.pgn
}
