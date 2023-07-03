package tree

type Node struct {
	Bonds []Node
}

func (n *Node) Insert(next Node) {
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
