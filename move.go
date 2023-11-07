package main

type Move struct {
	From Square
	To   Square

	piece       Piece
	validated   bool
	isCapture   bool
	isPromotion bool
}

func NewMove(from Square, to Square) *Move {
	return &Move{
		From: from,
		To:   to,
	}
}

func (m *Move) String() string {
	return SYMBOLS[m.piece] + " " + SQUARE_NAMES[m.From] + " -> " + SQUARE_NAMES[m.To]
}

func (m *Move) Piece() Piece {
	return m.piece
}

func (m *Move) Validated() bool {
	return m.validated
}

func (m *Move) IsCapture() bool {
	return m.isCapture
}

func (m *Move) IsPromotion() bool {
	return m.isPromotion
}

func (m *Move) Validate(board *Board) bool {
	m.validated = false
	m.isCapture = false
	m.isPromotion = false

	piece := board.GetPieceAtSquare(uint8(m.From))
	if piece == WHITE_PAWN {
		if m.To == m.From+8 {
			m.validated = true
		} else if m.To == m.From+16 && m.From >= A2 && m.From <= H2 {
			m.validated = true
		} else if m.To == m.From+7 && m.From%8 != 0 && board.GetPieceAtSquare(uint8(m.To)) != EMPTY {
			m.validated = true
			m.isCapture = true
		} else if m.To == m.From+9 && m.From%8 != 7 && board.GetPieceAtSquare(uint8(m.To)) != EMPTY {
			m.validated = true
			m.isCapture = true
		}
	}
	return true
}
