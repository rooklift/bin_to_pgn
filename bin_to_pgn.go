package main

import (
	"fmt"
	// gcb "github.com/rooklift/bin_to_pgn/gochessboard"
	poly "github.com/rooklift/bin_to_pgn/gopolyglot"
)

func main() {

	book, err := poly.LoadFromFile("C:\\Users\\Owner\\Documents\\Misc\\Chess\\Books\\komodo3.bin")

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("%v\n", book[0].Key)

	// Start at startpos
	// Find entries in polyglot book
	// Recursively add next positions to tree
	// Continue
}
