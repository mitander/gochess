package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	turnsMap = map[string]Color{
		"w": White,
		"b": Black,
	}
	pieceMap = map[string]Piece{
		"K": WhiteKing,
		"Q": WhiteQueen,
		"R": WhiteRook,
		"B": WhiteBishop,
		"N": WhiteKnight,
		"P": WhitePawn,
		"k": BlackKing,
		"q": BlackQueen,
		"r": BlackRook,
		"b": BlackBishop,
		"n": BlackKnight,
		"p": BlackPawn,
	}
)

func FEN(fen string) (*Position, error) {
	fen = strings.TrimSpace(fen)
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("fen: invalid fen length [%v]", len(fen))
	}
	b, err := fenBoard(parts[0])
	if err != nil {
		return nil, err
	}
	turn, ok := turnsMap[parts[1]]
	if !ok {
		return nil, fmt.Errorf("fen: error parsing turns")
	}
	castle, err := fenCastleStatus(parts[2])
	if err != nil {
		return nil, err
	}
	sq, err := fenEnPassant(parts[3])
	if err != nil {
		return nil, err
	}
	return &Position{
		board:        b,
		turn:         turn,
		castleStatus: castle,
		enPassant:    sq,
	}, nil
}

func fenBoard(board string) (*Board, error) {
	ranks := strings.Split(board, "/")
	if len(ranks) != 8 {
		return nil, fmt.Errorf("fen - invalid board: %s", board)
	}
	m := map[Square]Piece{}
	for i, str := range ranks {
		rank := Rank(7 - i)
		fileMap, err := fenRank(str)
		if err != nil {
			return nil, err
		}
		for file, piece := range fileMap {
			m[getSquare(file, rank)] = piece
		}
	}
	return NewBoard(m)
}

func (b *Board) setBBForSquareUtil(m *Move) {
	// Set bitboards for white, black and empty squares
	b.whiteSqs = b.whiteKing | b.whiteQueen | b.whiteRook | b.whiteBishop | b.whiteKnight | b.whitePawn
	b.blackSqs = b.blackKing | b.blackQueen | b.blackRook | b.blackBishop | b.blackKnight | b.blackPawn
	b.emptySqs = ^(b.whiteSqs | b.blackSqs)

	// Initialization of king square bitboards on nil Move
	if m == nil {
		b.whiteKingSq = NoSquare
		b.blackKingSq = NoSquare

		for i := 0; i < TotalSquares; i++ {
			sq := Square(i)
			// Set white king square
			if b.whiteKing.Occupied(sq) {
				b.whiteKingSq = sq
				// Set black king square
			} else if b.blackKing.Occupied(sq) {
				b.blackKingSq = sq
			}
		}
		// Update king square bitboards on king Move
	} else if m.from == b.whiteKingSq {
		b.whiteKingSq = m.to
	} else if m.from == b.blackKingSq {
		b.blackKingSq = m.to
	}
}

func fenCastleStatus(castle string) (CastleStatus, error) {
	for _, r := range castle {
		c := fmt.Sprintf("%c", r)
		switch c {
		case "K", "Q", "k", "q", "-":
		default:
			return "-", fmt.Errorf("fen: invalid castle rights [%s]", castle)
		}
	}
	return CastleStatus(castle), nil
}

func fenEnPassant(ep string) (Square, error) {
	if ep == "-" {
		return NoSquare, nil
	}
	sq := strToSquareMap[ep]
	if sq == NoSquare || !(sq.GetRank() == Rank3 || sq.GetRank() == Rank6) {
		return NoSquare, fmt.Errorf("fen: invalid En Passant square [%s]", ep)
	}
	return sq, nil
}

func fenRank(rank string) (map[File]Piece, error) {
	count := 0
	m := map[File]Piece{}
	err := fmt.Errorf("fen - invalid rank [%s]", rank)
	for _, r := range rank {
		c := fmt.Sprintf("%c", r)
		piece := pieceMap[c]
		if piece == NoPiece {
			skip, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			count += skip
			continue
		}
		m[File(count)] = piece
		count++
	}
	if count != 8 {
		return nil, err
	}
	return m, nil
}
