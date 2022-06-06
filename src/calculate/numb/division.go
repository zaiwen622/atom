package numb

import (
	"warning/src/calculate"
)

type Division struct {
	Express
}

func (p *Division) Do() *AtomNumb  {
	var result AtomNumb
	if !p.isValid() {
		errorCal := calculate.CalculateError{
			Code:   201,
			Reason: "分母不能为零",
		}
		panic(any(errorCal))
	}
	result.SetAtom(p.LeftAtom.GetAtom() / p.RightAtom.GetAtom())
	return &result
}

func (p *Division) isValid() bool  {
	if p.RightAtom.GetAtom() == 0 {
		return false
	}
	return true
}

