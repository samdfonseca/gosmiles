package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestSplitAtoms(t *testing.T) {
	tests := []struct {
		data            []byte
		atEOF           bool
		expectedAdvance int
		expectedToken   []byte
		expectedErr     error
	}{
		{
			data:            []byte("CC(C)C"),
			atEOF:           false,
			expectedAdvance: 1,
			expectedToken:   []byte("C"),
			expectedErr:     nil,
		},
		{
			data:            []byte("CC(C)C"),
			atEOF:           true,
			expectedAdvance: 1,
			expectedToken:   []byte("C"),
			expectedErr:     nil,
		},
		{
			data:            []byte("(C)C"),
			atEOF:           false,
			expectedAdvance: 3,
			expectedToken:   []byte("(C)"),
			expectedErr:     nil,
		},
		{
			data:            []byte("(C)C"),
			atEOF:           true,
			expectedAdvance: 3,
			expectedToken:   []byte("(C)"),
			expectedErr:     nil,
		},
		{
			data:            []byte("(C(C)"),
			atEOF:           false,
			expectedAdvance: 0,
			expectedToken:   nil,
			expectedErr:     nil,
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s-%v", tc.data, tc.atEOF), func(t *testing.T) {
			advance, token, err := splitAtoms(tc.data, tc.atEOF)
			if advance != tc.expectedAdvance || !bytes.Equal(token, tc.expectedToken) || !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected advance: %d, actual advance: %d\nexpected token: %s, actual token: %s\nexpected err: %v, actual err: %v\n", tc.expectedAdvance, advance, tc.expectedToken, token, tc.expectedErr, err)
			}
		})
	}
}

func TestParserScan(t *testing.T) {
	tests := []struct {
		smiles   string
		expected []string
	}{
		{
			smiles:   "CCCCCCCCC",
			expected: []string{"C", "C", "C", "C", "C", "C", "C", "C", "C"},
		},
		{
			smiles:   "CC(C)C",
			expected: []string{"C", "C", "(C)", "C"},
		},
		{
			smiles:   "CCC(C)CC",
			expected: []string{"C", "C", "C", "(C)", "C", "C"},
		},
		{
			smiles:   "CCC(C(CC)(C))(C)C",
			expected: []string{"C", "C", "C", "(C(CC)(C))", "(C)", "C"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.smiles, func(t *testing.T) {
			l := NewLexer(strings.NewReader(tc.smiles))
			for i, token := range tc.expected {
				l.Scan()
				if actual := l.Text(); actual != token {
					t.Errorf("input: %s, token number: %d, expected: %s, actual: %s", tc.smiles, i, token, actual)
				}
			}
		})
	}
}

func TestParserParse(t *testing.T) {
	p := NewParser("CCC(C(CC))(C)C")
	if err := p.Parse(); err != nil {
		t.Error(err)
	}
	n := p.Root() // [C]CC(C(CC))(C)C
	if len(n.Bonds) != 1 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 1, len(n.Bonds))
	}
	n = n.Bonds[0] // C[C]C(C(CC))(C)C
	if len(n.Bonds) != 1 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 1, len(n.Bonds))
	}
	n = n.Bonds[0] // CC[C](C(CC))(C)C
	if len(n.Bonds) != 3 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 3, len(n.Bonds))
	}
	n1 := n.Bonds[0] // CCC([C](CC))(C)C
	if len(n1.Bonds) != 1 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 1, len(n1.Bonds))
	}
	n1 = n1.Bonds[0] // CCC(C([C]C))(C)C
	if len(n1.Bonds) != 1 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 1, len(n1.Bonds))
	}
	n1 = n1.Bonds[0] // CCC(C(C[C]))(C)C
	if len(n1.Bonds) != 0 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 0, len(n1.Bonds))
	}
	n1 = n.Bonds[1] // CCC(C(CC))([C])C
	if len(n1.Bonds) != 0 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 0, len(n1.Bonds))
	}
	n = n.Bonds[2] // CCC(C(CC))(C)[C]
	if len(n.Bonds) != 0 {
		t.Errorf("number of bonds, expected: %d, actual: %d", 0, len(n.Bonds))
	}
}
