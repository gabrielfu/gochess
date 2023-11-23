package gochess

type Move struct {
	from      Square
	to        Square
	piece     Piece
	castle    Castle
	promotion Piece
}

func NewMove(from Square, to Square, piece Piece) *Move {
	return &Move{
		from:      from,
		to:        to,
		piece:     piece,
		castle:    0,
		promotion: EMPTY,
	}
}

func NewCastlingMove(from Square, to Square, piece Piece, castle Castle) *Move {
	return &Move{
		from:      from,
		to:        to,
		piece:     piece,
		castle:    castle,
		promotion: EMPTY,
	}
}

func (m *Move) String() string {
	out := m.Piece().Symbol() + " " + m.From().String() + " -> " + m.To().String()
	if m.Promotion() != EMPTY {
		out += " (" + m.Promotion().Symbol() + ")"
	}
	return out
}

func (m *Move) Equals(other *Move) bool {
	if m == nil && other == nil {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return m.from == other.from &&
		m.to == other.to &&
		m.piece == other.piece &&
		m.castle == other.castle
}

func (m *Move) From() Square {
	return m.from
}

func (m *Move) To() Square {
	return m.to
}

func (m *Move) Piece() Piece {
	return m.piece
}

func (m *Move) Castle() Castle {
	return m.castle
}

func (m *Move) SetCastle(castle Castle) {
	m.castle = castle
}

func (m *Move) Promotion() Piece {
	return m.promotion
}

func (m *Move) SetPromotion(p Piece) {
	m.promotion = p
}

func FilterMoves(moves []*Move, filter func(*Move) bool) []*Move {
	out := make([]*Move, 0, len(moves))
	for _, move := range moves {
		if filter(move) {
			out = append(out, move)
		}
	}
	return out
}
