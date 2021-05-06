package main

type CastleStatus string

type Position struct {
	board        *Board       // current board
	castleStatus CastleStatus // white castling possibilities [white,black]
	check        bool         // check is active
	enPassant    Square       // en-passant square where pawn can be captured
	turn         Color        // current turn
}

func (p *Position) getBoard() *Board {
	return p.board
}

func (p *Position) setBoard(b *Board) {
	p.board = b
}

func (p *Position) getCastleStatus() CastleStatus {
	return p.castleStatus
}

func (p *Position) setCastleStatus(cs CastleStatus) {
	p.castleStatus = cs
}

func (p *Position) getEnPassant() Square {
	return p.enPassant
}

func (p *Position) setEnPassant(ep Square) {
	p.enPassant = ep
}

func (p *Position) getTurn() Color {
	return p.turn
}

func (p *Position) setTurn(c Color) {
	p.turn = c
}
