package numb

type Plus struct {
	Express
}

func (p *Plus) Do()*AtomNumb  {
	var result AtomNumb
	result.SetAtom(p.LeftAtom.GetAtom() + p.RightAtom.GetAtom())
	return &result
}

func (p *Plus) isValid() bool  {
	return true
}
