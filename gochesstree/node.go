package gochesstree

type Node struct {
	move		string			// The move that got us here. "" for root.
	parent		*Node
	children	[]*Node
}

func NewNode(parent *Node, move string) {
	ret := new(Node)
	ret.move = move
	ret.parent = parent
	parent.children = append(parent.children, ret)
	return ret
}
