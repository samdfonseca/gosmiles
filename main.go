package main

import (
	"fmt"
	"os"

	"github.com/samdfonseca/hw-samdfonseca/v2/namer"
	"github.com/samdfonseca/hw-samdfonseca/v2/parser"
)

func main() {
	smiles := os.Args[1:]
	for i := range smiles {
		p := parser.NewParser(smiles[i])
		if err := p.Parse(); err != nil {
			panic(err)
		}
		n := namer.New(p.Root(), false)
		name, err := n.SystematicName()
		if err != nil {
			panic(err)
		}
		fmt.Println(name)
	}
}
