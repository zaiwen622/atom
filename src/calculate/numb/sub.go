package numb

type Sub struct {
	Express
}

func (p *Sub) Do() *AtomNumb  {
	var result AtomNumb
	result.SetAtom(p.LeftAtom.GetAtom() - p.RightAtom.GetAtom())
	return &result
}

func (p *Sub) isValid() bool  {
	return true
}
