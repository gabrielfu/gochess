package main

import (
	"fmt"
	gochess "gochess/internal"
)

func main() {
	g := gochess.NewGame()
	g.MoveSAN("g4")
	g.MoveSAN("g5")
	g.MoveSAN("h4")
	g.MoveSAN("h5")
	g.MoveSAN("g5")
	g.MoveSAN("g4")
	g.MoveSAN("Rh8")

	fmt.Println("WHITE_KING_SIDE:", g.Board().CastlingRights().Has(gochess.WHITE_KING_SIDE))
	fmt.Println("WHITE_QUEEN_SIDE:", g.Board().CastlingRights().Has(gochess.WHITE_QUEEN_SIDE))
	fmt.Println("BLACK_KING_SIDE:", g.Board().CastlingRights().Has(gochess.BLACK_KING_SIDE))
	fmt.Println("BLACK_QUEEN_SIDE:", g.Board().CastlingRights().Has(gochess.BLACK_QUEEN_SIDE))

	// score := gochess.Evaluate(g.Board())
	// fmt.Println(score)
	fmt.Println(g.Visualize())
}
