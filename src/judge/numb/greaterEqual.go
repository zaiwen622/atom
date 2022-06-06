package numb

import "warning/src/calculate"

type GreaterEqual struct {
	Judge
}

func (e *GreaterEqual) GetResult()bool  {
	defer calculate.RecoverCal()
	left := e.LeftExpress.Do()
	right := e.RightExpress.Do()
	e.Result = left.GetAtom() >= right.GetAtom()
	return e.Result
}
