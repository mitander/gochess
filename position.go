package main

type Result []string
type CastleStatus string

// Position is current game position
type Position struct {
	board        *Board       // current board
	castleStatus CastleStatus // white castling possibilities [white,black]
	check        bool         // check is active
	ep           Square       // en-passant square where pawn can be captured
	turn         Color        // current turn
}
