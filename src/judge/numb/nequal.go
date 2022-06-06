package numb

import "warning/src/calculate"

type Nequal struct {
	Judge
}

func (e *Nequal) GetResult()bool  {
	defer calculate.RecoverCal()
	left := e.LeftExpress.Do()
	right := e.RightExpress.Do()
	e.Result = left.GetAtom() != right.GetAtom()
	return e.Result
}
