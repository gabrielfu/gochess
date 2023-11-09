package chessago

type Move struct {
	From  Square
	To    Square
	Piece Piece

	castle      Castle
	validated   bool
	isCapture   bool
	isPromotion bool
}

func NewMove(from Square, to Square, piece Piece) *Move {
	return &Move{
		From:  from,
		To:    to,
		Piece: piece,
	}
}

func (m *Move) String() string {
	return m.Piece.Symbol() + " " + m.From.String() + " -> " + m.To.String()
}

func (m *Move) Castle() Castle {
	return m.castle
}

func (m *Move) SetCastle(castle Castle) {
	m.castle = castle
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

	piece := board.GetPieceAtSquare(m.From)
	if piece == WHITE_PAWN {
		if m.To == m.From+8 {
			m.validated = true
		} else if m.To == m.From+16 && m.From >= A2 && m.From <= H2 {
			m.validated = true
		} else if m.To == m.From+7 && m.From%8 != 0 && board.GetPieceAtSquare(m.To) != EMPTY {
			m.validated = true
			m.isCapture = true
		} else if m.To == m.From+9 && m.From%8 != 7 && board.GetPieceAtSquare(m.To) != EMPTY {
			m.validated = true
			m.isCapture = true
		}
	}
	return true
}
