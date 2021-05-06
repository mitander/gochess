package main

import (
	"fmt"
	"log"
)

func main() {
	// reader = bufio.NewReader(os.Stdin)
	game, err := NewGame(TestPosition)
	if err != nil {
		log.Fatal(err)
	}

	for !game.isDone() {
		fmt.Println(game.DrawBoard())
		// fmt.Print("Enter move: ")
		// input, _ = reader.ReadString('\n')
		// err := game.MakeMove(input)

		fmt.Println(game.Position.board.whiteSqs.Draw())
		game.Quit()
	}

}
