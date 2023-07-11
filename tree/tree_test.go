package tree

import (
	"fmt"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	n := &Node{Symbol: "C", Bonds: make([]*Node, 0, 4)}
	n1 := &Node{Symbol: "C1", Bonds: make([]*Node, 0, 4)}
	n.Insert(n1)
	if len(n.Bonds) != 1 {
		t.Errorf("len(n.Bonds), expected: %d, actual: %d\n", 1, len(n.Bonds))
	}
	if n.Bonds[0].Symbol != "C1" {
		t.Errorf("n.Bonds[0].Symbol, expected: %s, actual: %s\n", "C1", n.Bonds[0].Symbol)
	}
}

func TestTreeLongestChainLength(t *testing.T) {
	root := &Node{Symbol: "C0", Bonds: make([]*Node, 0, 4)}
	n := root
	for i := 1; i < 6; i++ {
		n.Insert(&Node{Symbol: fmt.Sprintf("C%d", i), Bonds: make([]*Node, 0, 4)})
		n = n.Bonds[0]
	}
	branch := &Node{Symbol: "B1-1", Bonds: make([]*Node, 0, 4)}
	branch.Insert(&Node{Symbol: "B1-2", Bonds: make([]*Node, 0, 4)})
	root.Bonds[0].Bonds[0].Insert(branch)
	// CCC(CC)CCC
	lcl := root.LongestChainLength()
	if lcl != 6 {
		t.Errorf("root.LongestChainLength(), expected: %d, actual: %d\n", 6, lcl)
	}
}

func TestTreeBranches(t *testing.T) {
	root := &Node{Symbol: "C0", Bonds: make([]*Node, 0, 4)}
	n := root
	for i := 1; i < 6; i++ {
		n.Insert(&Node{Symbol: fmt.Sprintf("C%d", i), Bonds: make([]*Node, 0, 4)})
		n = n.Bonds[0]
	}
	branch := &Node{Symbol: "B1-1", Bonds: make([]*Node, 0, 4)}
	branch.Insert(&Node{Symbol: "B1-2", Bonds: make([]*Node, 0, 4)})
	root.Bonds[0].Bonds[0].Insert(branch)
	branch = &Node{Symbol: "B2-1", Bonds: make([]*Node, 0, 4)}
	root.Bonds[0].Bonds[0].Insert(branch)
	// CCC(CC)(C)CCC
	n = root.Bonds[0].Bonds[0]
	nBranches := len(n.Branches())
	if nBranches != 2 {
		t.Errorf("len(n.Branches()), expected: %d, actual: %d\n", 2, nBranches)
	}
	s := n.Branches()[0].Symbol
	if s != "B1-1" {
		t.Errorf("n.Branches()[0].Symbol, expected: %s, actual: %s\n", "B1-1", s)
	}
	s = n.Branches()[1].Symbol
	if s != "B2-1" {
		t.Errorf("n.Branches()[1].Symbol, expected: %s, actual: %s\n", "B2-1", s)
	}
}

func TestTreeNextNodeInLongestChain(t *testing.T) {
	root := &Node{Symbol: "C0", Bonds: make([]*Node, 0, 4)}
	n := root
	for i := 1; i < 6; i++ {
		n.Insert(&Node{Symbol: fmt.Sprintf("C%d", i), Bonds: make([]*Node, 0, 4)})
		n = n.Bonds[0]
	}
	branch := &Node{Symbol: "B1-1", Bonds: make([]*Node, 0, 4)}
	branch.Insert(&Node{Symbol: "B1-2", Bonds: make([]*Node, 0, 4)})
	root.Bonds[0].Bonds[0].Insert(branch)
	branch = &Node{Symbol: "B2-1", Bonds: make([]*Node, 0, 4)}
	root.Bonds[0].Bonds[0].Insert(branch)

	n = root
	for i := 1; i < 6; i++ {
		n = n.NextNodeInLongestChain()
		if n.Symbol != fmt.Sprintf("C%d", i) {
			t.Errorf("incorrect symbol, expected: %s, actual: %s\n", fmt.Sprintf("C%d", i), n.Symbol)
		}
	}
}
