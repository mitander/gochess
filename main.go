package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cli() {
	pos, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(pos.board)
		fmt.Println(pos.Moves())
		valid := false
		for !valid {
			fmt.Printf("Enter move (%v):", pos.turn)
			input, _ := r.ReadString('\n')
			input = strings.TrimSpace(input)
			for _, m := range pos.Moves() {
				if input == m.String() {
					pos = pos.Move(m)
					valid = true
					break
				}
			}
		}
	}
}

func main() {
	cli()
}
