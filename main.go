package main

import (
	"fmt"
	gochess "gochess/internal"
)

func main() {
	g := gochess.NewGame()
	g.Move(gochess.MustParseSAN("g4", g.Board()))
	g.Move(gochess.MustParseSAN("g5", g.Board()))
	g.Move(gochess.MustParseSAN("h4", g.Board()))
	g.Move(gochess.MustParseSAN("h5", g.Board()))
	g.Move(gochess.MustParseSAN("g5", g.Board()))
	g.Move(gochess.MustParseSAN("g4", g.Board()))
	g.Move(gochess.MustParseSAN("Rh8", g.Board()))

	fmt.Println("WHITE_KING_SIDE:", g.Board().CastlingRights().Has(gochess.WHITE_KING_SIDE))
	fmt.Println("WHITE_QUEEN_SIDE:", g.Board().CastlingRights().Has(gochess.WHITE_QUEEN_SIDE))
	fmt.Println("BLACK_KING_SIDE:", g.Board().CastlingRights().Has(gochess.BLACK_KING_SIDE))
	fmt.Println("BLACK_QUEEN_SIDE:", g.Board().CastlingRights().Has(gochess.BLACK_QUEEN_SIDE))

	// score := gochess.Evaluate(g.Board())
	// fmt.Println(score)
	fmt.Println(g.Visualize())
}
