package parser

import (
	"strings"
	"testing"
)

func TestLexerNext(t *testing.T) {
	newAtom := func(s string) *Token {
		return &Token{s, ATOM}
	}
	newBranch := func(s string) *Token {
		return &Token{s, BRANCH}
	}
	tests := []struct {
		smiles   string
		expected []*Token
	}{
		{
			smiles:   "CCCCCCCCC",
			expected: []*Token{newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C"), newAtom("C")},
		},
		{
			smiles:   "CC(C)C",
			expected: []*Token{newAtom("C"), newAtom("C"), newBranch("C"), newAtom("C")},
		},
		{
			smiles:   "CCC(C)CC",
			expected: []*Token{newAtom("C"), newAtom("C"), newAtom("C"), newBranch("C"), newAtom("C"), newAtom("C")},
		},
		{
			smiles:   "CCC(C(CC)(C))(C)C",
			expected: []*Token{newAtom("C"), newAtom("C"), newAtom("C"), newBranch("C(CC)(C)"), newBranch("C"), newAtom("C")},
		},
	}
	for _, tc := range tests {
		t.Run(tc.smiles, func(t *testing.T) {
			l := NewLexer(strings.NewReader(tc.smiles))
			for i, expected := range tc.expected {
				token, err := l.Next()
				if err != nil {
					t.Error(err)
				}
				if token.Type != expected.Type {
					t.Errorf("input: %s, token number: %d, expected: %v, actual: %v\n", tc.smiles, i, expected, token)
				}
				if token.Value != expected.Value {
					t.Errorf("input: %s, token number: %d, expected: %v, actual: %v\n", tc.smiles, i, expected, token)
				}
			}
			if token, err := l.Next(); token != nil || err != nil {
				t.Errorf("input: %s, err: %s, expected nil at end, actual: %v", tc.smiles, err, token)
			}
		})
	}
}
