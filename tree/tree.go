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

func (n *Node) BuildParentChain(chain []*Node) []*Node {
	chain = append(chain, n)
	if n.NextNodeInLongestChain() == nil {
		return chain
	}
	return n.NextNodeInLongestChain().BuildParentChain(chain)
}

func (n *Node) Branches() []*Node {
	if len(n.Bonds) == 0 || len(n.Bonds) == 1 {
		return nil
	}
	branches := make([]*Node, 0, 2)
	p := n.Bonds[0]
	for i := 1; i < len(n.Bonds); i++ {
		if p.LongestChainLength() == n.Bonds[i].LongestChainLength() && len(p.Branches()) < len(n.Bonds[i].Branches()) {
			branches = append(branches, p)
			p = n.Bonds[i]
		} else if p.LongestChainLength() < n.Bonds[i].LongestChainLength() {
			branches = append(branches, p)
			p = n.Bonds[i]
		} else {
			branches = append(branches, n.Bonds[i])
		}
	}
	return branches
}

// BranchLocations maps branch length to a list of locations where a branch of that length occurs.
// To build the branch names, we need to know both the branch locations as well as how many times similar
// (same length) branchs occur. For example, CC(C)C(C)C has two 1-length branches, "di"+"methyl", at locations 2 and 3,
// "2,3", to become 2,3-dimethylbutane. A data structure that maps branch length to a list of locations contains all the
// information needed to produce branch naming. If a branch is closer to the end of the chain than any branch is to the
// beginning of the chain, the location numbering is reversed to ensure the lowest possible numbers for the position of
// the branches.
func (n *Node) BranchLocations() map[int][]int {
	i := 2
	substituents := make(map[int][]int)
	pcl := n.LongestChainLength()
	min := pcl
	max := 0
	for node := n.NextNodeInLongestChain(); node.NextNodeInLongestChain() != nil; node = node.NextNodeInLongestChain() {
		if len(node.Bonds) >= 2 {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
			for _, branch := range node.Branches() {
				blcl := branch.LongestChainLength()
				locations, ok := substituents[blcl]
				if !ok {
					substituents[blcl] = []int{i}
				} else {
					substituents[blcl] = append(locations, i)
				}
			}
		}
		i++
	}
	if (pcl - max) < (min - 1) {
		revSubstituents := make(map[int][]int, len(substituents))
		for k, v := range substituents {
			for i := range v {
				v[i] = pcl - v[i] + 1
			}
			revSubstituents[k] = v
		}
		return revSubstituents
	}
	return substituents
}
