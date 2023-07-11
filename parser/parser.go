package parser

import (
	"strings"

	"github.com/samdfonseca/hw-samdfonseca/v2/tree"
)

type Nodeable interface {
	Node() (*tree.Node, error)
}

type Parser struct {
	l    *Lexer
	root *tree.Node
}

func NewParser(s string) *Parser {
	l := NewLexer(strings.NewReader(s))
	return &Parser{l, nil}
}

func (p *Parser) Parse() error {
	var current *tree.Node
	var parent *tree.Node
	for token, err := p.l.Next(); token != nil && err == nil; token, err = p.l.Next() {
		if token.Type == ATOM {
			current = &tree.Node{
				Symbol: token.Value,
				Bonds:  make([]*tree.Node, 0, 4),
			}
			if p.root == nil {
				p.root = current
				parent = current
				continue
			}
			parent.Insert(current)
			parent = current
		} else if token.Type == BRANCH {
			sub := NewParser(token.Value)
			if err := sub.Parse(); err != nil {
				return err
			}
			current = sub.Root()
			parent.Insert(current)
		}
	}
	return nil
}

func (p *Parser) Root() *tree.Node {
	return p.root
}
