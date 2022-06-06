package numb

import "warning/src/calculate"

type Equal struct {
	Judge
}

func (e *Equal) GetResult()bool  {
	defer calculate.RecoverCal()
	left := e.LeftExpress.Do()
	right := e.RightExpress.Do()
	e.Result = left.GetAtom() == right.GetAtom()
	return e.Result
}
