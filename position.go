package main

type Turn string

const (
	WhiteTurn Turn = "White"
	BlackTurn Turn = "Black"
	startFEN       = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

type Position struct {
	board Board   // current board
	turn  Turn    // Current turn ("White"/"Black")
	wc    [2]bool // white castling - [kingside, queenside]
	bc    [2]bool // black castling - [kingside, queenside]
	ep    Square  // square of avaliable en-passant
	kp    Square  // king passent during castling, where kind can be captured
	check bool    // boolean if there is a check
}

func abs(n int) int { return int((int64(n) ^ int64(n)>>63) - int64(n)>>63) }

// starts a new game
func NewGame() (Position, error) {
	board, turn, err := FEN(startFEN)
	if err != nil {
		return Position{}, err
	}
	return Position{
		board: board,
		turn:  turn,
	}, nil
}

var wdir = map[Piece][]Square{
	'P': {N, N + N, N + W, N + E},
	'N': {N + N + E, E + N + E, E + S + E, S + S + E, S + S + W, W + S + W, W + N + W, N + N + W},
	'B': {N + E, S + E, S + W, N + W},
	'R': {N, E, S, W},
	'Q': {N, E, S, W, N + E, S + E, S + W, N + W},
	'K': {N, E, S, W, N + E, S + E, S + W, N + W},
}

var bdir = map[Piece][]Square{
	'p': {S, S + S, S + W, S + E},
	'n': {N + N + E, E + N + E, E + S + E, S + S + E, S + S + W, W + S + W, W + N + W, N + N + W},
	'b': {N + E, S + E, S + W, N + W},
	'r': {N, E, S, W},
	'q': {N, E, S, W, N + E, S + E, S + W, N + W},
	'k': {N, E, S, W, N + E, S + E, S + W, N + W},
}

func getDirections(t Turn) map[Piece][]Square {
	if t == WhiteTurn {
		return wdir
	} else {
		return bdir
	}
}

func valid(s Piece, t Turn) bool {
	return s != ' ' || (s == '.' && !s.ours(t))
}

func (pos Position) Moves() (moves []Move) {
	directions := getDirections(pos.turn)
	for index, p := range pos.board {
		i := Square(index)
		if !p.ours(pos.turn) {
			continue
		}
		for _, dir := range directions[p] {
			for j := i + dir; ; j = j + dir {
				q := pos.board[j]
				if !valid(q, pos.turn) {
					break
				}
				if p == 'P' {
					if (dir == N || dir == N+N) && q != '.' {
						break
					}
					if dir == N+N && (i < A1+N || pos.board[i+N] != '.') {
						break
					}
					if (dir == N+W || dir == N+E) && q == '.' && (j != pos.ep && j != pos.kp && j != pos.kp-1 && j != pos.kp+1) {
						break
					}
				}
				if p == 'p' {
					if (dir == S || dir == S+S) && q != '.' {
						break
					}
					if dir == S+S && (i < A8+S || pos.board[i+S] != '.') {
						break
					}
					if (dir == S+W || dir == S+E) && q == '.' && (j != pos.ep && j != pos.kp && j != pos.kp-1 && j != pos.kp+1) {
						break
					}
				}
				moves = append(moves, Move{from: i, to: j})
				// Crawling pieces should stop after a single move
				if p == 'P' || p == 'N' || p == 'K' || (q != ' ' && q != '.' && !q.ours(pos.turn)) {
					break
				}
				if p == 'p' || p == 'n' || p == 'k' || (q != ' ' && q != '.' && !q.ours(pos.turn)) {
					break
				}
				// Castling rules
				if i == A1 && pos.board[j+E] == 'K' && pos.wc[0] {
					moves = append(moves, Move{from: j + E, to: j + W})
				}
				if i == A8 && pos.board[j+E] == 'k' && pos.wc[0] {
					moves = append(moves, Move{from: j + E, to: j + W})
				}
				if i == H1 && pos.board[j+W] == 'K' && pos.wc[1] {
					moves = append(moves, Move{from: j + W, to: j + E})
				}
				if i == H8 && pos.board[j+W] == 'k' && pos.wc[1] {
					moves = append(moves, Move{from: j + W, to: j + E})
				}
			}
		}
	}
	return moves
}

func (pos Position) Move(m Move) (np Position) {
	from, to, piece := m.from, m.to, pos.board[m.from]
	np = pos
	np.ep = 0
	np.kp = 0
	np.board[m.to] = pos.board[m.from]
	np.board[m.from] = '.'

	switch from {
	case A1:
		np.wc[0] = false
	case H1:
		np.wc[1] = false
	case A8:
		np.bc[1] = false
	case H8:
		np.bc[0] = false
	}

	switch piece {
	case 'K', 'k':
		np.wc[0], np.wc[1] = false, false // if king moved, disable castling
		if abs(int(to-from)) == 2 {
			if to < from {
				np.board[H1] = '.'
			} else {
				np.board[A1] = '.'
			}
			np.board[(from+to)/2] = 'R'
		}

	case 'P', 'p':
		// Pawn promotion
		if A8 <= to && to <= H8 { // if pawn moves to 8th row: update to Queen
			np.board[to] = 'Q'
		}
		// En-passant
		if to-from == 2*N {
			np.ep = from + N // set avaliable en-passant Square
		}
		// En-passant capture
		if to == pos.ep {
			np.board[to+S] = '.'
		}
	}

	if pos.turn == "White" {
		np.turn = BlackTurn
	} else {
		np.turn = WhiteTurn
	}

	nm := np.Moves()

	// check if move is a check - if so set check to checked king Square
	for _, m := range nm {
		switch np.board[m.to] {
		case 'K', 'k':
			np.check = true
		default:
			np.check = false
		}
	}
	return np
}
