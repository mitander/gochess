package main

const (
	startPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

type Game struct {
	moves    []*Move
	position *Position
	outcome  string
	done     bool
}

func NewGame() (*Game, error) {
	pos, err := FEN(startPosition)
	if err != nil {
		return &Game{}, err
	}

	game := &Game{
		moves:    []*Move{},
		position: pos,
	}
	return game, nil
}
