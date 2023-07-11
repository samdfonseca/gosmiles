package parser

import (
	"bufio"
	"io"
)

const (
	ATOM = iota
	BRANCH
)

type Lexer struct {
	*bufio.Scanner
}

type Token struct {
	Value string
	Type  int
}

func splitAtoms(data []byte, atEOF bool) (int, []byte, error) {
	if len(data) == 0 && atEOF {
		return 0, nil, nil
	}
	if data[0] == 'C' {
		return 1, data[0:1], nil
	}
	if data[0] == '(' {
		open := 1
		for i := 1; i < len(data); i++ {
			if data[i] == '(' {
				open += 1
			} else if data[i] == ')' {
				open -= 1
			}
			if open == 0 {
				return i + 1, data[:i+1], nil
			}
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken
}

func NewLexer(r io.Reader) *Lexer {
	s := bufio.NewScanner(r)
	s.Split(splitAtoms)
	return &Lexer{s}
}

func (l *Lexer) Next() (*Token, error) {
	if !l.Scan() {
		return nil, l.Err()
	}
	err := l.Err()
	if err != nil {
		return nil, err
	}
	token := l.Text()
	if token[0] == '(' && token[len(token)-1] == ')' {
		return &Token{token[1 : len(token)-1], BRANCH}, nil
	}
	return &Token{token, ATOM}, nil
}
