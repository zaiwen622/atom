package numb

type Multi struct {
	Express
}

func (p *Multi) Do() *AtomNumb  {
	var result AtomNumb
	result.SetAtom(p.LeftAtom.GetAtom() * p.RightAtom.GetAtom())
	return &result
}

func (p *Multi) isValid() bool  {
	return true
}
