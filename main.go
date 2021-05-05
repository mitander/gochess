package main

import (
	"fmt"
	"log"
)

func main() {

	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	for !game.isDone() {
		fmt.Println(game.DrawBoard())
		fmt.Println(game.Position.board.blackBishop.Draw())
		game.Quit()
	}

}
