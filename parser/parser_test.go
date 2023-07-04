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
			p := New(strings.NewReader(tc.smiles))
			for i, token := range tc.expected {
				p.Scan()
				if actual := p.Text(); actual != token {
					t.Errorf("input: %s, token number: %d, expected: %s, actual: %s", tc.smiles, i, token, actual)
				}
			}
		})
	}
}
