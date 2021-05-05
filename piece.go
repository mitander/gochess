package main

type Color int8

const (
	NoColor Color = iota
	White
	Black
)

func (c Color) Other() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	return NoColor
}

func (c Color) Name() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return "No Color"
}

type PieceType int8

const (
	NoPieceType PieceType = iota
	King
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func (p PieceType) String() string {
	switch p {
	case King:
		return "k"
	case Queen:
		return "q"
	case Rook:
		return "r"
	case Bishop:
		return "b"
	case Knight:
		return "n"
	case Pawn:
		return "p"
	}
	return ""
}

type Piece int8

const (
	NoPiece Piece = iota
	WhiteKing
	WhiteQueen
	WhiteRook
	WhiteBishop
	WhiteKnight
	WhitePawn
	BlackKing
	BlackQueen
	BlackRook
	BlackBishop
	BlackKnight
	BlackPawn
)

var (
	PiecesMap = []Piece{
		WhiteKing, WhiteQueen, WhiteRook, WhiteBishop, WhiteKnight, WhitePawn,
		BlackKing, BlackQueen, BlackRook, BlackBishop, BlackKnight, BlackPawn,
	}
)

func getPiece(t PieceType, c Color) Piece {
	for _, p := range PiecesMap {
		if p.Color() == c && p.Type() == t {
			return p
		}
	}
	return NoPiece
}

func (p Piece) Type() PieceType {
	switch p {
	case WhiteKing, BlackKing:
		return King
	case WhiteQueen, BlackQueen:
		return Queen
	case WhiteRook, BlackRook:
		return Rook
	case WhiteBishop, BlackBishop:
		return Bishop
	case WhiteKnight, BlackKnight:
		return Knight
	case WhitePawn, BlackPawn:
		return Pawn
	}
	return NoPieceType
}

func (p Piece) Color() Color {
	switch p {
	case WhiteKing, WhiteQueen, WhiteRook, WhiteBishop, WhiteKnight, WhitePawn:
		return White
	case BlackKing, BlackQueen, BlackRook, BlackBishop, BlackKnight, BlackPawn:
		return Black
	}
	return NoColor
}

func (p Piece) Unicode() string {
	uc := []string{" ", "♚", "♛", "♜", "♝", "♞", "♟", "♔", "♕", "♖", "♗", "♘", "♙"}
	return uc[int(p)]
}
