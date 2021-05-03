package main

// square is an index of the board
type Square int

// movement from one square to another
type Move struct {
	from Square
	to   Square
}

// move directions, to simplify move calculation
const N, E, S, W = -10, 1, 10, -1

// corner squares
const A1, H1, A8, H8 Square = 91, 98, 21, 28

// string represenation of a square - e.g: h4
func (s Square) String() string { return string([]byte{" abcdefgh "[s%10], "  87654321  "[s/10]}) }

// string represenation of a move - e.g: h4h6
func (m Move) String() string { return m.from.String() + m.to.String() }
