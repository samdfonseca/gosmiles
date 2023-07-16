package namer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samdfonseca/hw-samdfonseca/v2/parser"
)

func TestNamer(t *testing.T) {
	tests := []struct {
		smiles   string
		expected string
	}{
		{
			smiles:   "CCCCCCCCC",
			expected: "nonane",
		},
		{
			smiles:   "CC(C)C",
			expected: "2-methylpropane",
		},
		{
			smiles:   "CCC(C)CC",
			expected: "3-methylpentane",
		},
		{
			smiles:   "CCC(C)(C)C",
			expected: "2,2-dimethylbutane",
		},
		{
			smiles:   strings.Repeat("C", 23),
			expected: "tricosane",
		},
		{
			smiles:   "CCCC(CCC)" + strings.Repeat("C", 19),
			expected: "4-propyltricosane",
		},
		{
			smiles:   strings.Repeat("C", 20) + "(CCC)CCC",
			expected: "4-propyltricosane",
		},
		{
			smiles:   "CC(C)CCC(CCC)C",
			expected: "2,5-dimethyloctane",
		},
		{
			smiles:   "CC(CC(CC(CC(C)(C)C)C)(C)C)(C)C",
			expected: "2,2,4,4,6,8,8-heptamethylnonane",
		},
		{
			smiles:   "CCC(C)CCC(C)CC(C)C(C)C",
			expected: "2,3,5,8-tetramethyldecane",
		},
		{
			smiles:   "CCCC(C(C)(C)C)C(CC)CCCC",
			expected: "4-(1,1-dimethylethyl)-5-ethylnonane",
		},
		{
			smiles:   "CCC(C)CC(C)C(C(C)(CC)C)C(CC)CCCC",
			expected: "6-(1,1-dimethylpropyl)-7-ethyl-3,5-dimethylundecane",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s-%s", tc.smiles, tc.expected), func(t *testing.T) {
			p := parser.NewParser(tc.smiles)
			if err := p.Parse(); err != nil {
				t.Error(err)
			}
			n := New(p.Root(), false)
			sn, err := n.SystematicName()
			if err != nil {
				t.Error(err)
			}
			if sn != tc.expected {
				t.Errorf("input: %v, expected: %v, actual: %v\n", tc.smiles, tc.expected, sn)
			}
		})
	}
}
