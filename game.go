package chessago

type Game struct {
	board *Board
}

func NewGame() *Game {
	InitMovesTables()

	return &Game{
		board: NewStartingBoard(),
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
	return g.board.Visualize()
}

// Move executes the given move on the board.
func (g *Game) Move(move *Move) error {
	return g.board.Move(move)
}
