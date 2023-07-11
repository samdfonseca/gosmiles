package namer

import (
	"strings"

	"github.com/samdfonseca/hw-samdfonseca/v2/parser"
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
	Smiles string
}

func New(smiles string) *Namer {
	return &Namer{smiles}
}

func (n *Namer) SystematicName() (string, error) {
	p := parser.NewParser(n.Smiles)
	if err := p.Parse(); err != nil {
		return "", err
	}
	root := p.Root()
	lcl := root.LongestChainLength()
	parentNamePrefix := NumericalTerm(lcl)
	if parentNamePrefix[len(parentNamePrefix)-1] == 'a' {
		parentNamePrefix = parentNamePrefix[:len(parentNamePrefix)-1]
	}
	parentName := parentNamePrefix + "ane"
	return parentName, nil
}

func NumericalTerm(nCarbon int) string {
	if nCarbon <= 20 {
		return PREFIXES[nCarbon]
	}
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
			if val[len(val)-1] == 'a' {
				numerals[i] = val[:len(val)-1]
			}
			break
		}
		if val[len(val)-1] == numerals[i+1][0] {
			numerals[i] = val[:len(val)-1]
		}
	}
	return strings.Join(numerals, "")
}
