package main

import (
	"errors"
	"strings"
)

type Piece byte

func (p Piece) value() int {
	return map[Piece]int{'P': 100, 'N': 280, 'B': 320, 'R': 479, 'Q': 929, 'K': 30000}[p]
}

func (p Piece) ours(t Turn) bool {
	if t == WhiteTurn {
		return p.value() != 0
	} else {
		return p.Flip().value() != 0
	}
}

type Board [120]Piece

func (p Piece) Flip() Piece {
	return map[Piece]Piece{'P': 'p', 'N': 'n', 'B': 'b', 'R': 'r', 'Q': 'q', 'K': 'k', 'p': 'P', 'n': 'N', 'b': 'B', 'r': 'R', 'q': 'Q', 'k': 'K', ' ': ' ', '.': '.'}[p]
}

func FEN(fen string) (b Board, t Turn, err error) {
	parts := strings.Split(fen, " ")
	rows := strings.Split(parts[0], "/")
	if len(rows) != 8 {
		return b, t, errors.New("FEN should have 8 rows")
	}
	for i := 0; i < len(b); i++ {
		b[i] = ' '
	}
	for i := 0; i < 8; i++ {
		index := i*10 + 21
		for _, c := range rows[i] {
			q := Piece(c)
			if q >= '1' && q <= '8' {
				for j := Piece(0); q-j >= '1'; j++ {
					b[index] = '.'
					index += 1
				}
			} else if q.value() == 0 && q.Flip().value() == 0 {
				return b, t, errors.New("invalid piece value: " + string(c))
			} else {
				b[index] = q
				index += 1
			}
		}
		if index%10 != 9 {
			return b, t, errors.New("invalid row length")
		}
	}
	if len(parts) > 1 && parts[1] == "b" {
		return b, BlackTurn, nil
	} else {
		return b, WhiteTurn, nil
	}
}

func (b Board) String() (s string) {
	s = "\n"
	for row := 2; row < 10; row++ {
		for col := 1; col < 9; col++ {
			s += string(b[row*10+col])
		}
		s += "\n"
	}
	return s
}
