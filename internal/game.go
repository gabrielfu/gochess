package gochess

import "fmt"

type Game struct {
	board   *Board
	history []*Move
	pgn     string
	ended   bool
	winner  Color
}

func NewGame() *Game {
	return &Game{
		board:   NewStartingBoard(),
		history: []*Move{},
		pgn:     "",
		ended:   false,
		winner:  NO_COLOR,
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

	if g.Board().IsInCheckmate() {
		g.ended = true
		winner := 1 - g.Turn()
		g.winner = winner
		g.pgn += fmt.Sprintf("%d-%d", 1-winner, winner)
	}
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

func (g *Game) Ended() bool {
	return g.ended
}

func (g *Game) Winner() Color {
	return g.winner
}
