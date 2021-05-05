package main

type Board struct {
	whiteKing   bitboard
	whiteQueen  bitboard
	whiteRook   bitboard
	whiteBishop bitboard
	whiteKnight bitboard
	whitePawn   bitboard
	blackKing   bitboard
	blackQueen  bitboard
	blackRook   bitboard
	blackBishop bitboard
	blackKnight bitboard
	blackPawn   bitboard
	whiteSqs    bitboard
	blackSqs    bitboard
	emptySqs    bitboard
	whiteKingSq Square
	blackKingSq Square
}

func NewBoard(m map[Square]Piece) (*Board, error) {
	b := &Board{}
	for _, p1 := range pieceMap {
		bm := map[Square]bool{}
		for sq, p2 := range m {
			if p1 == p2 {
				bm[sq] = true
			}
		}
		bb, err := NewBitboard(bm)
		if err != nil {
			return &Board{}, err
		}
		b.setPieceBB(p1, bb)
	}
	b.setBBForSquareUtil(nil)
	return b, nil
}

func (b *Board) Piece(sq Square) Piece {
	for _, p := range pieceMap {
		bb := b.getPieceBB(p)
		if bb.Occupied(sq) {
			return p
		}
	}
	return NoPiece
}

func (b *Board) Draw() string {
	s := "\n  A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += Rank(r).String()
		for f := 0; f < TotalSquareRows; f++ {
			p := b.Piece(getSquare(File(f), Rank(r)))
			if p == NoPiece {
				s += "-"
			} else {
				s += p.Unicode()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

func (b *Board) getPieceBB(p Piece) bitboard {
	switch p {
	case WhiteKing:
		return b.whiteKing
	case WhiteQueen:
		return b.whiteQueen
	case WhiteRook:
		return b.whiteRook
	case WhiteBishop:
		return b.whiteBishop
	case WhiteKnight:
		return b.whiteKnight
	case WhitePawn:
		return b.whitePawn
	case BlackKing:
		return b.blackKing
	case BlackQueen:
		return b.blackQueen
	case BlackRook:
		return b.blackRook
	case BlackBishop:
		return b.blackBishop
	case BlackKnight:
		return b.blackKnight
	case BlackPawn:
		return b.blackPawn
	}
	return bitboard(0)
}

func (b *Board) setPieceBB(p Piece, bb bitboard) {
	switch p {
	case WhiteKing:
		b.whiteKing = bb
	case WhiteQueen:
		b.whiteQueen = bb
	case WhiteRook:
		b.whiteRook = bb
	case WhiteBishop:
		b.whiteBishop = bb
	case WhiteKnight:
		b.whiteKnight = bb
	case WhitePawn:
		b.whitePawn = bb
	case BlackKing:
		b.blackKing = bb
	case BlackQueen:
		b.blackQueen = bb
	case BlackRook:
		b.blackRook = bb
	case BlackBishop:
		b.blackBishop = bb
	case BlackKnight:
		b.blackKnight = bb
	case BlackPawn:
		b.blackPawn = bb
	}
}
