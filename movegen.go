package main

var KingMovesTable [64]Bitboard
var KnightMovesTable [64]Bitboard
var WhitePawnAttacksTable [64]Bitboard
var BlackPawnAttacksTable [64]Bitboard

/*
 * Non sliding pieces
 */

// InitMovesTables initializes the moves tables for non-sliding pieces (Kings, Knights, Pawns).
func InitMovesTables() {
	initKingMovesTable()
	initKnightMovesTable()
	initWhitePawnAttacksTable()
	initBlackPawnAttacksTable()
}

func initKingMovesTable() {
	for i := 0; i < 64; i++ {
		KingMovesTable[i] = calcKingMoves(Square(i))
	}
}

func initKnightMovesTable() {
	for i := 0; i < 64; i++ {
		KnightMovesTable[i] = calcKnightMoves(Square(i))
	}
}

func initWhitePawnAttacksTable() {
	for i := 0; i < 56; i++ {
		WhitePawnAttacksTable[i] = calcWhitePawnAttacks(Square(i))
	}
}

func initBlackPawnAttacksTable() {
	for i := 8; i < 64; i++ {
		BlackPawnAttacksTable[i] = calcBlackPawnAttacks(Square(i))
	}
}

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

func calcWhitePawnAttacks(from Square) Bitboard {
	// TODO: Need to add en passant
	moves := Bitboard(0)
	file := from.File()
	if file < A {
		moves |= 1 << (from + 9) // LEFT
	}
	if file > H {
		moves |= 1 << (from + 7) // RIGHT
	}
	return moves
}

func calcBlackPawnAttacks(from Square) Bitboard {
	// TODO: Need to add en passant
	moves := Bitboard(0)
	file := from.File()
	if file < A {
		moves |= 1 << (from - 7) // LEFT
	}
	if file > H {
		moves |= 1 << (from - 9) // RIGHT
	}
	return moves
}

func calcWhitePawnMoves(from Square) Bitboard {
	moves := Bitboard(0)
	if from <= A7 {
		moves |= 1 << (from + 8) // UP
	}
	if from.Rank() == R2 {
		moves |= 1 << (from + 16) // UP 2
	}
	return moves
}

func calcBlackPawnMoves(from Square) Bitboard {
	moves := Bitboard(0)
	if from >= A2 {
		moves |= 1 << (from - 8) // DOWN
	}
	if from.Rank() == R7 {
		moves |= 1 << (from - 16) // DOWN 2
	}
	return moves
}

func GetKnightMoves(from Square) Bitboard {
	return KnightMovesTable[from]
}

func GetKingMoves(from Square) Bitboard {
	return KingMovesTable[from]
}

func GetWhitePawnAttacks(from Square) Bitboard {
	return WhitePawnAttacksTable[from]
}

func GetBlackPawnAttacks(from Square) Bitboard {
	return BlackPawnAttacksTable[from]
}

func GetWhitePawnMoves(from Square) Bitboard {
	return calcWhitePawnMoves(from)
}

func GetBlackPawnMoves(from Square) Bitboard {
	return calcBlackPawnMoves(from)
}

/*
 * Sliding pieces
 */

func calcRookMoves(from Square, occupany Bitboard) Bitboard {
	moves := Bitboard(0)
	file := from.File()
	rank := from.Rank()
	// LEFT
	for i := file + 1; i <= A; i++ {
		newSq := uint8(rank)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// RIGHT
	for i := file - 1; i >= H; i-- {
		newSq := uint8(rank)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// UP
	for i := rank + 1; i <= R8; i++ {
		newSq := uint8(i)*8 + uint8(file)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// DOWN
	for i := rank - 1; i >= R1; i-- {
		newSq := uint8(i)*8 + uint8(file)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	return moves
}

func calcBishopMoves(from Square, occupany Bitboard) Bitboard {
	moves := Bitboard(0)
	file := from.File()
	rank := from.Rank()
	// UP LEFT
	for i, j := file+1, rank+1; i <= A && j <= R8; i, j = i+1, j+1 {
		newSq := uint8(j)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// UP RIGHT
	for i, j := file-1, rank+1; i >= H && j <= R8; i, j = i-1, j+1 {
		newSq := uint8(j)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// DOWN LEFT
	for i, j := file+1, rank-1; i <= A && j >= R1; i, j = i+1, j-1 {
		newSq := uint8(j)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	// DOWN RIGHT
	for i, j := file-1, rank-1; i >= H && j >= R1; i, j = i-1, j-1 {
		newSq := uint8(j)*8 + uint8(i)
		moves |= 1 << newSq
		if occupany&(1<<newSq) != 0 {
			break
		}
	}
	return moves
}

func calcQueenMoves(from Square, occupany Bitboard) Bitboard {
	return calcRookMoves(from, occupany) | calcBishopMoves(from, occupany)
}

func GetRookMoves(from Square, occupany Bitboard) Bitboard {
	return calcRookMoves(from, occupany)
}

func GetBishopMoves(from Square, occupany Bitboard) Bitboard {
	return calcBishopMoves(from, occupany)
}

func GetQueenMoves(from Square, occupany Bitboard) Bitboard {
	return calcQueenMoves(from, occupany)
}
