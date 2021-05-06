package main

const (
	StartPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	TestPosition  = "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"
)

type Outcome string

type Status struct {
	Outcome Outcome
	Reason  Reason
	Done    bool
}

const (
	NoOutcome Outcome = "*"
	WhiteWon  Outcome = "1-0"
	BlackWon  Outcome = "0-1"
	Draw      Outcome = "1/2-1/2"
)

type Reason uint8

const (
	Checkmate Reason = iota
	Resignation
	DrawOffer
	Stalemate
)

type Game struct {
	ValidMoves []*Move
	Position   *Position
	Status     Status
}

func NewGame(fen string) (*Game, error) {
	pos, err := FEN(fen)
	if err != nil {
		return &Game{}, err
	}

	return &Game{
		ValidMoves: []*Move{},
		Position:   pos,
	}, nil
}

func (g *Game) DrawBoard() string {
	return g.Position.board.Draw()
}

func (g *Game) isDone() bool {
	return g.Status.Done
}

func (g *Game) Quit() {
	g.Status.Done = true
}
