package main

type MoveTag uint16

const (
	KingSideCastle MoveTag = 1 << iota
	QueenSideCastle
	Capture
	EnPassant
	Check
	inCheck
)

type Move struct {
	from  Square
	to    Square
	promo PieceType
	tags  MoveTag
}

func (m *Move) From() Square {
	return m.from
}

func (m *Move) To() Square {
	return m.to
}

func (m *Move) Promo() PieceType {
	return m.promo
}

func (m *Move) HasTag(tag MoveTag) bool {
	return (tag & m.tags) > 0
}

func (m *Move) addTag(tag MoveTag) {
	m.tags = m.tags | tag
}

func (m *Move) String() string {
	return m.from.String() + m.to.String() + m.promo.String()
}
