package gochesstree

type Node struct {
	Move		string			// The move that got us here. "" for root.
	Parent		*Node
	Children	[]*Node
}

func NewNode(parent *Node, move string) *Node {
	ret := new(Node)
	ret.Move = move
	ret.Parent = parent
	if parent != nil {
		parent.Children = append(parent.Children, ret)
	}
	return ret
}

func (self *Node) CountNodes() int {
	count := 1
	for _, child := range self.Children {
		count += child.CountNodes()
	}
	return count
}
