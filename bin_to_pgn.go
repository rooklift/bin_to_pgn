package main

import (
	"fmt"
	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
)

func main() {

	board, _ := gcb.BoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	board, _ = board.ForceMove("e2e5")
	board, _ = board.ForceMove("d7d5")
	board, _ = board.ForceMove("e5d6")

	fmt.Printf("%v\n", board)
}
