package namer

import (
	"fmt"
	"strings"
	"testing"
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
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s-%s", tc.smiles, tc.expected), func(t *testing.T) {
			n := New(tc.smiles)
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
