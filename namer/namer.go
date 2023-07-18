package namer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/samdfonseca/hw-samdfonseca/v2/tree"
)

var (
	PREFIXES = []string{
		"",
		"meth",
		"eth",
		"prop",
		"but",
		"pent",
		"hex",
		"hept",
		"oct",
		"non",
		"dec",
		"undec",
		"dodec",
		"tridec",
		"tetradec",
		"pentadec",
		"hexadec",
		"heptadec",
		"octadec",
		"nonadec",
		"icos",
	}
	NUMERICAL_TERMS = map[int]string{
		1:    "hen",
		2:    "do",
		3:    "tri",
		4:    "tetra",
		5:    "penta",
		6:    "hexa",
		7:    "hepta",
		8:    "octa",
		9:    "nona",
		10:   "deca",
		20:   "icosa",
		30:   "triaconta",
		40:   "tetraconta",
		50:   "pentaconta",
		60:   "hexaconta",
		70:   "heptaconta",
		80:   "octaconta",
		90:   "nonaconta",
		100:  "hecta",
		200:  "dicta",
		300:  "tricta",
		400:  "tetracta",
		500:  "pentacta",
		600:  "hexacta",
		700:  "heptacta",
		800:  "octacta",
		900:  "nonacta",
		1000: "kilia",
		2000: "dilia",
		3000: "trilia",
		4000: "tetralia",
		5000: "pentalia",
		6000: "hexalia",
		7000: "heptalia",
		8000: "octalia",
		9000: "nonalia",
	}
)

type Namer struct {
	Node   *tree.Node
	Branch bool
}

func New(node *tree.Node, branch bool) *Namer {
	return &Namer{node, branch}
}

func (n *Namer) SystematicName() (string, error) {
	root := n.Node
	chain := root.BuildParentChain([]*tree.Node{})
	lchain := len(chain)
	parentNamePrefix := ""
	if lchain <= 20 {
		parentNamePrefix = PREFIXES[lchain]
	} else {
		parentNamePrefix = NumericalTerm(lchain)
	}

	if parentNamePrefix[len(parentNamePrefix)-1] == 'a' {
		parentNamePrefix = parentNamePrefix[:len(parentNamePrefix)-1]
	}

	parentName := ""
	if n.Branch {
		parentName = parentNamePrefix + "yl"
	} else {
		parentName = parentNamePrefix + "ane"
	}

	// complexSubstituentNodes := []*tree.Node{}
	// for i := range chain {
	// 	substituents := chain[i].Branches()
	// 	if len(substituents) == 0 {
	// 		continue
	// 	}
	// 	for j := range substituents {
	// 		if len(substituents[j].Branches()) > 0 {
	// 			nn := New(substituents[j], true)
	// 			complexSubstituent, err := nn.SystematicName()
	// 			if err != nil {
	// 				return "", err
	// 			}
	// 			complexSubstituent = complexSubstituent
	// 		}
	// 	}
	// }

	substituents := root.BranchLocations()

	temp := parentName

	branchPrefixes := []branchPrefix{}
	for nCarbon, locations := range substituents {
		branchPrefixes = append(branchPrefixes, BranchPrefix(nCarbon, locations))
	}
	sort.Slice(branchPrefixes, func(i, j int) bool {
		r := strings.Compare(branchPrefixes[i].base, branchPrefixes[j].base)
		return r == -1
	})

	branchPrefixStrings := []string{}
	for i := range branchPrefixes {
		branchPrefixStrings = append(branchPrefixStrings, branchPrefixes[i].String())
	}
	temp = strings.Join(branchPrefixStrings, "-") + temp

	return temp, nil
}

type branchPrefix struct {
	locations  []int
	base       string
	multiplier string
	// prefix     string
}

func (bp branchPrefix) String() string {
	sort.Ints(bp.locations)
	locations := []string{}
	for _, l := range bp.locations {
		locations = append(locations, fmt.Sprintf("%d", l))
	}
	locationsString := strings.Join(locations, ",")
	return fmt.Sprintf("%s-%s%syl", locationsString, bp.multiplier, bp.base)
}

func BranchPrefix(nCarbon int, locations []int) branchPrefix {
	prefix := ""
	if len(locations) == 2 {
		prefix = "di"
	} else if len(locations) > 2 {
		prefix = NumericalTerm(len(locations))
	}

	base := ""
	if nCarbon <= 20 {
		base = PREFIXES[nCarbon]
	} else {
		base = NumericalTerm(nCarbon)
	}
	return branchPrefix{locations, base, prefix}
}

func NumericalTerm(nCarbon int) string {
	numerals := make([]string, 0, 4)
	m := 1
	for i := 0; i < 4; i++ {
		n := (nCarbon - (nCarbon % m)) % (m * 10)
		if n != 0 {
			numerals = append(numerals, NUMERICAL_TERMS[n])
		}
		m = m * 10
	}
	for i, val := range numerals {
		if i == len(numerals)-1 {
			break
		}
		if val[len(val)-1] == numerals[i+1][0] {
			numerals[i] = val[:len(val)-1]
		}
	}
	return strings.Join(numerals, "")
}
