package main

import (
	"fmt"
	"strings"
)

type MoveTag uint16

const (
	KingSideCastle MoveTag = iota
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

func MakeMove(s string) (err error) {
	_, err = StrToMove(s)
	return err
}

func StrToMove(s string) (*Move, error) {
	if len(s) != 5 {
		return &Move{}, fmt.Errorf("invalid move length")
	}
	split := strings.Split(s, "-")
	from := strToSquareMap[split[0]]
	to := strToSquareMap[split[1]]
	return &Move{
		from: from,
		to:   to,
	}, nil
}
