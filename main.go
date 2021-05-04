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

	fmt.Println(game.position.board.Draw())
	fmt.Println(game.position)
}
