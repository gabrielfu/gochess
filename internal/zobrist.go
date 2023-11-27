package gochess

import "math/rand"

var pieceHashes [12][64]Bitboard
var enPassantHashes [64]Bitboard
var castlingHashes = make(map[Castle]Bitboard)
var whitesTurnHash Bitboard

func random() Bitboard {
	return Bitboard(rand.Uint64())
}

func init() {
	for i := 0; i < 64; i++ {
		enPassantHashes[i] = random()
		for j := 0; j < 12; j++ {
			pieceHashes[j][i] = random()
		}
	}
	castlingHashes[WHITE_KING_SIDE] = random()
	castlingHashes[WHITE_QUEEN_SIDE] = random()
	castlingHashes[BLACK_KING_SIDE] = random()
	castlingHashes[BLACK_QUEEN_SIDE] = random()
	whitesTurnHash = random()
}

func ZobristHash(b *Board) uint64 {
	hash := Bitboard(0)

	// castling rights
	if b.castlingRights.Has(WHITE_KING_SIDE) {
		hash ^= castlingHashes[WHITE_KING_SIDE]
	}
	if b.castlingRights.Has(WHITE_QUEEN_SIDE) {
		hash ^= castlingHashes[WHITE_QUEEN_SIDE]
	}
	if b.castlingRights.Has(BLACK_KING_SIDE) {
		hash ^= castlingHashes[BLACK_KING_SIDE]
	}
	if b.castlingRights.Has(BLACK_QUEEN_SIDE) {
		hash ^= castlingHashes[BLACK_QUEEN_SIDE]
	}

	// pieces
	for _, piece := range ALL_PIECES {
		bb := b.GetBbForPiece(piece)
		for _, sq := range bb.Squares() {
			hash ^= pieceHashes[piece][sq]
		}
	}

	// en passant
	for _, sq := range b.EnPassantSquare().Squares() {
		hash ^= enPassantHashes[int(sq)]
	}

	if b.Turn() == WHITE {
		hash ^= whitesTurnHash
	}
	return uint64(hash)
}
