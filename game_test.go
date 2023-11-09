package chessago

import "testing"

func BenchmarkLegalMovesOnSmallBoard(b *testing.B) {
	board := NewEmptyBoard()
	board.AddPieceToSquare(WHITE_ROOK, A1)
	board.AddPieceToSquare(WHITE_KING, E1)
	board.AddPieceToSquare(WHITE_ROOK, H1)
	g := NewGame()
	g.SetBoard(board)

	for i := 0; i < b.N; i++ {
		g.LegalMoves()
	}
}

func BenchmarkLegalMovesOnStartingBoard(b *testing.B) {
	g := NewGame()

	for i := 0; i < b.N; i++ {
		g.LegalMoves()
	}
}
