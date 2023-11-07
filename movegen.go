package main

// func calcKnightMoves(from Square) {
// 	moves := Bitboard(0)

// }

func calcKingMoves(from Square) Bitboard {
	moves := Bitboard(0)
	file := from.File()
	if from <= A7 {
		moves |= 1 << (from + 8) // UP
	}
	if from >= H2 {
		moves |= 1 << (from - 8) // DOWN
	}
	if file < A {
		moves |= 1 << (from + 1) // LEFT
	}
	if file > H {
		moves |= 1 << (from - 1) // RIGHT
	}
	if file < A && from <= A7 {
		moves |= 1 << (from + 9) // UP LEFT
	}
	if file > H && from <= A7 {
		moves |= 1 << (from + 7) // UP RIGHT
	}
	if file < A && from >= H2 {
		moves |= 1 << (from - 7) // DOWN LEFT
	}
	if file > H && from >= H2 {
		moves |= 1 << (from - 9) // DOWN RIGHT
	}
	return moves
}

func calcKnightMoves(from Square) Bitboard {
	moves := Bitboard(0)
	file := from.File()
	if file != A && from <= A6 {
		moves |= 1 << (from + 17) // UP LEFT
	}
	if file != H && from <= A6 {
		moves |= 1 << (from + 15) // UP RIGHT
	}
	if file != A && from >= A2 {
		moves |= 1 << (from - 15) // DOWN LEFT
	}
	if file != H && from > H3 {
		moves |= 1 << (from - 17) // DOWN RIGHT
	}
	if file < B && from <= A7 {
		moves |= 1 << (from + 10) // LEFT UP
	}
	if file < B && from >= H2 {
		moves |= 1 << (from - 6) // LEFT DOWN
	}
	if file > G && from >= H2 {
		moves |= 1 << (from - 10) // RIGHT DOWN
	}
	if file > G && from <= A7 {
		moves |= 1 << (from + 6) // RIGHT UP
	}
	return moves
}
