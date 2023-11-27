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

const TranpositionSize = 100_000

type TranspositionTable struct {
	entries [TranpositionSize]*TranspositionTableEntry
}

type TranspositionTableEntry struct {
	hash  uint64
	depth int
	eval  int
	move  *Move
}

func (tt *TranspositionTable) Get(hash uint64) *TranspositionTableEntry {
	entry := tt.entries[hash%TranpositionSize]
	if entry != nil && entry.hash == hash {
		return entry
	}
	return nil
}

func (tt *TranspositionTable) Put(hash uint64, depth int, eval int, move *Move) {
	tt.entries[hash%TranpositionSize] = &TranspositionTableEntry{
		hash:  hash,
		depth: depth,
		eval:  eval,
		move:  move,
	}
}

var transpositionTable = &TranspositionTable{}

func Search(b *Board, depth int) *SearchResult {
	eval, move := alphabeta(b, depth, MIN_EVAL-1, MAX_EVAL+1)
	return &SearchResult{
		eval: eval,
		move: move,
	}
}

func alphabeta(b *Board, depth int, alpha int, beta int) (int, *Move) {
	// tt lookup
	hash := ZobristHash(b)
	ttEntry := transpositionTable.Get(hash)
	if ttEntry != nil && ttEntry.depth >= depth {
		return ttEntry.eval, ttEntry.move
	}

	// Check for game over or drawn positions
	if b.IsInCheckmate() {
		if b.Turn() == WHITE {
			return Evaluate(b, BlackWon), nil
		} else {
			return Evaluate(b, WhiteWon), nil
		}
	}
	if b.IsInStalemate() || b.IsInsufficientMaterial() {
		return Evaluate(b, Draw), nil
	}

	// game in progress
	if depth == 0 {
		return Evaluate(b, InProgress), nil
	}

	var bestMove *Move
	var best int
	if b.Turn() == WHITE {
		best = MIN_EVAL - 1
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
	} else {
		best = MAX_EVAL + 1
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
	}

	transpositionTable.Put(hash, depth, best, bestMove)
	return best, bestMove
}
