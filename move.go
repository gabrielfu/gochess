package main

type Move struct {
	From  SQUARES
	To    SQUARES
	Piece PIECES
}

func (m *Move) String() string {
	return SYMBOLS[m.Piece] + " " + SQUARE_NAMES[m.From] + " -> " + SQUARE_NAMES[m.To]
}
