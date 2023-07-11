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
	p := n.Bonds[0]
	for i := 1; i < len(n.Bonds); i++ {
		if p.LongestChainLength() == n.Bonds[i].LongestChainLength() {
			if len(p.Branches()) < len(n.Bonds[i].Branches()) {
				p = n.Bonds[i]
			}
		} else if p.LongestChainLength() < n.Bonds[i].LongestChainLength() {
			p = n.Bonds[i]
		}
	}
	return p
}

func (n *Node) Branches() []*Node {
	if len(n.Bonds) == 0 || len(n.Bonds) == 1 {
		return nil
	}
	branches := make([]*Node, 0, 2)
	p := n.Bonds[0]
	for i := 1; i < len(n.Bonds); i++ {
		if p.LongestChainLength() == n.Bonds[i].LongestChainLength() {
			if len(p.Branches()) < len(n.Bonds[i].Branches()) {
				branches = append(branches, p)
				p = n.Bonds[i]
			}
		} else if p.LongestChainLength() < n.Bonds[i].LongestChainLength() {
			branches = append(branches, p)
			p = n.Bonds[i]
		} else {
			branches = append(branches, n.Bonds[i])
		}
	}
	return branches
}
