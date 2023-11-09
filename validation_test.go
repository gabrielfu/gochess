package chessago

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitMovesTables()
	code := m.Run()
	os.Exit(code)
}

func TestIsSquareAttacked(t *testing.T) {
	b := NewStartingBoard()
	if isSquareAttacked(E4, b) {
		t.Errorf("E4 should be not attacked")
	}
	if !isSquareAttacked(E6, b) {
		t.Errorf("E6 should be attacked by black pawn")
	}
}
