package namer

type Namer struct {
	Smiles string
}

func New(smiles string) *Namer {
	return &Namer{smiles}
}

func (n *Namer) SystematicName() string {
	return ""
}
