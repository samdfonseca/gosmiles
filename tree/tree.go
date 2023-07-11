package tree

type Node struct {
	Symbol string
	Bonds  []*Node
}

func (n *Node) Insert(next *Node) {
	n.Bonds = append(n.Bonds, next)
}

func (n *Node) LongestChainLength() int {
	if len(n.Bonds) == 0 {
		return 1
	}
	max := 0
	for _, b := range n.Bonds {
		if l := b.LongestChainLength(); l > max {
			max = l
		}
	}
	return max + 1
}

func (n *Node) NextNodeInLongestChain() *Node {
	if len(n.Bonds) == 0 {
		return nil
	}
	if len(n.Bonds) == 1 {
		return n.Bonds[0]
	}
	max := 0
	var next *Node
	for _, b := range n.Bonds {
		if l := b.LongestChainLength(); l > max {
			next = b
			max = l
		}
	}
	return next
}

func (n *Node) Branches() []*Node {
	if len(n.Bonds) == 0 || len(n.Bonds) == 1 {
		return nil
	}
	branches := make([]*Node, 0, 2)
	max := 0
	var next *Node
	for _, b := range n.Bonds {
		if l := b.LongestChainLength(); l > max {
			if next != nil {
				branches = append(branches, b)
			}
			next = b
			max = l
		} else {
			branches = append(branches, b)
		}
	}
	return branches
}
