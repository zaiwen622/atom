package numb

import (
	"math"
)

type Power struct {
	Express
}

func (p *Power) Do() *AtomNumb  {
	var result AtomNumb
	result.SetAtom(math.Pow(p.LeftAtom.GetAtom(), p.RightAtom.GetAtom()))
	return &result
}

func (p *Power) isValid() bool  {
	return true
}
