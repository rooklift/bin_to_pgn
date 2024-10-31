package main

import (
	"fmt"
	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
	poly "github.com/rooklift/bin_to_pgn/gopolyglot"
)

func main() {

/*
	book, err := poly.LoadFromFile("C:\\Users\\Owner\\Documents\\Misc\\Chess\\Books\\komodo3.bin")

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
*/

	board, _ := gcb.BoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	fmt.Printf("%v\n", poly.KeyFromBoard(board))



	// Start at startpos
	// Find entries in polyglot book
	// Recursively add next positions to tree
	// Continue
}
