package parser

import (
	"bufio"
	"io"
)

type Parser struct {
	*bufio.Scanner
}

func splitAtoms(data []byte, atEOF bool) (int, []byte, error) {
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

func New(r io.Reader) *Parser {
	s := bufio.NewScanner(r)
	s.Split(splitAtoms)
	return &Parser{s}
}
