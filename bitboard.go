package main

import (
	"math/bits"
	"strconv"
)

type bitboard uint64

func NewBitboard(m map[Square]bool) (bitboard, error) {
	s := ""
	for sq := 0; sq < 64; sq++ {
		if m[Square(sq)] {
			s += "1"
		} else {
			s += "0"
		}
	}
	bb, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return bitboard(bb), nil
}

func (b bitboard) Occupied(sq Square) bool {
	return (bits.RotateLeft64(uint64(b), int(sq)+1) & 1) == 1
}

func (b bitboard) Draw() string {
	s := "  A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += Rank(r).String()
		for f := 0; f < TotalSquareRows; f++ {
			sq := getSquare(File(f), Rank(r))
			if b.Occupied(sq) {
				s += "1"
			} else {
				s += "-"
			}
			s += " "
		}
		s += "\n"
	}
	return s
}
