package main

import (
	"fmt"
	// "math/rand"
	"os"

	gcb "github.com/rooklift/bin_to_pgn/gochessboard"
	tree "github.com/rooklift/bin_to_pgn/gochesstree"
	poly "github.com/rooklift/bin_to_pgn/gopolyglot"
)

func main() {

	book, err := poly.LoadFromFile("C:\\Users\\Owner\\Documents\\Misc\\Chess\\Books\\komodo3.bin")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Fprintf(os.Stderr, "Book length: %d\n", len(book.Entries))

	node := tree.NewNode(nil, "")
	board, _ := gcb.BoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	recurse(node, board, book, 0)
	fmt.Fprintf(os.Stderr, "%v nodes\n", node.CountNodes())
	fmt.Printf("%v\n", make_string(node, false))
}

func recurse(node *tree.Node, board *gcb.Board, book *poly.Book, depth int) {

	key := poly.KeyFromBoard(board)
	entries := book.Probe(key)

	if depth > 10 {
		return
	}

	for _, entry := range entries {
		move := entry.MoveString()
		new_node := tree.NewNode(node, move)
		new_board, _ := board.ForceMove(move)
		recurse(new_node, new_board, book, depth + 1)
	}
}

func make_string(node *tree.Node, skip_self bool) string {

	if len(node.Children) == 0 {
		if skip_self {
			return " "
		} else {
			return node.Move + " "
		}
	} else if len(node.Children) == 1 {
		if skip_self {
			return make_string(node.Children[0], false) + " "
		} else {
			return node.Move + " " + make_string(node.Children[0], false) + " "
		}
	} else {
		var s string
		if skip_self {
			s = node.Children[0].Move + " "
		} else {
			s = node.Move + " " + node.Children[0].Move + " "
		}
		for _, child := range node.Children[1:] {
			s += "( " + make_string(child, false) + ") "
		}
		s += make_string(node.Children[0], true)
		return s
	}
}

