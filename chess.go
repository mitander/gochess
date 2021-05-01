package main

import (
	"errors"
	"strings"
)

type Piece byte
type Board [120]Piece
type Square int

type Position struct {
	board Board   // current board
	wc    [2]bool // white castling possibilities
	bc    [2]bool // black castling possibilities
	check bool    // check active
	ep    Square  // en-passant square where pawn can be captured
}

// parses fen string
func FEN(f string) (b Board, err error) {
	parts := strings.Split(f, " ")
	rows := strings.Split(parts[0], "/")

	if len(rows) != 8 {
		return b, errors.New("Invalid FEN length")
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
					index++
				}
			} else {
				b[index] = q
				index++
			}
		}
		if index%10 != 9 {
			return b, errors.New("Invalid row length")
		}
	}
	return b, nil
}

func newBoard() Position {
	board, _ := FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBKQBNR")
	return Position{
		board: board,
	}
}

// returns human readable board representation
func (a Board) String() (s string) {
	s = "\n"
	for row := 2; row < 10; row++ {
		for col := 1; col < 9; col++ {
			s = s + string(a[row*10+col])
		}
		s = s + "\n"
	}
	return s
}
