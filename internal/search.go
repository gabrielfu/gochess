package gochess

type SearchResult struct {
	eval int
	move *Move
}

func (sr *SearchResult) Eval() int {
	return sr.eval
}

func (sr *SearchResult) Move() *Move {
	return sr.move
}

func Search(b *Board, depth int) *SearchResult {
	eval, move := alphabeta(b, depth, MIN_EVAL-1, MAX_EVAL+1)
	return &SearchResult{
		eval: eval,
		move: move,
	}
}

func alphabeta(b *Board, depth int, alpha int, beta int) (int, *Move) {
	if depth == 0 {
		return Evaluate(b), nil
	}

	var bestMove *Move
	if b.Turn() == WHITE {
		best := MIN_EVAL
		for _, move := range b.LegalMoves() {
			// TODO: should implement undo move so that we don't need a copy
			cpy := b.Copy()
			cpy.Move(move)
			eval, _ := alphabeta(cpy, depth-1, alpha, beta)
			if eval > best {
				best = eval
				bestMove = move
			}
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return best, bestMove
	} else {
		best := MAX_EVAL
		for _, move := range b.LegalMoves() {
			cpy := b.Copy()
			cpy.Move(move)
			eval, _ := alphabeta(cpy, depth-1, alpha, beta)
			if eval < best {
				best = eval
				bestMove = move
			}
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return best, bestMove
	}
}
