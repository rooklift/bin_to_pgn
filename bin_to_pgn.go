package main

import (
	"fmt"
	// "math/rand"

	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
	tree "github.com/rooklift/bin_to_pgn/gochesstree"
	poly "github.com/rooklift/bin_to_pgn/gopolyglot"
)

func main() {

	book, err := poly.LoadFromFile("C:\\Users\\Owner\\Documents\\Misc\\Chess\\Books\\komodo3.bin")

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Book length: %d\n", len(book.Entries))

	node := tree.NewNode(nil, "")
	board, _ := gcb.BoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	recurse(node, board, book, 0)
	fmt.Printf("%v nodes\n", node.CountNodes())

	for {
		fmt.Printf("%v ", node.Move)
		if len(node.Children) == 0 {
			fmt.Printf("\n")
			break
		}
		node = node.Children[0]
	}
}

func recurse(node *tree.Node, board *gcb.Board, book *poly.Book, depth int) {

	key := poly.KeyFromBoard(board)
	entries := book.Probe(key)

	if depth > 50 {
		return
	}

	for _, entry := range entries {
		move := entry.MoveString()
		new_node := tree.NewNode(node, move)
		new_board, _ := board.ForceMove(move)
		recurse(new_node, new_board, book, depth + 1)
	}
}