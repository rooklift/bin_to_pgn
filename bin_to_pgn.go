package main

import (
	"fmt"
	// "math/rand"
	"os"

	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
	// tree "github.com/rooklift/bin_to_pgn/gochesstree"
	poly "github.com/rooklift/bin_to_pgn/gopolyglot"
)

func main() {

	book, err := poly.LoadFromFile("C:\\Users\\Owner\\Documents\\Misc\\Chess\\Books\\komodo3.bin")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Fprintf(os.Stderr, "Book length: %d\n", len(book.Entries))

	board, _ := gcb.BoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	print_lines_recurse(board, book, nil)
}

func print_lines_recurse(board *gcb.Board, book *poly.Book, moves []string) {

	key := poly.KeyFromBoard(board)
	entries := book.Probe(key)

	if len(moves) > 10 || len(entries) == 0 {
		fmt.Printf("[Event \"Book from BIN\"]\n\n")
		for _, move := range moves {
			fmt.Printf("%s ", move)
		}
		fmt.Printf("\n\n")
		return
	}

	for _, entry := range entries {
		move := entry.MoveString()
		var new_moves_slice []string
		new_moves_slice = append(new_moves_slice, moves...)
		new_moves_slice = append(new_moves_slice, move)
		new_board, _ := board.ForceMove(move)
		print_lines_recurse(new_board, book, new_moves_slice)
	}
}
