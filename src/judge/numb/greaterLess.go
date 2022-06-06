package numb

import (
	"warning/src/calculate"
	"warning/src/judge/andOr"
)

type GreaterLess struct {
	JudgeScope
}

func (e *GreaterLess) GetResult()bool  {
	defer calculate.RecoverCal()
	left := e.LeftExpress.Do()
	right := e.RightExpress.Do()
	base := e.BaseExpress.Do()
	var andor andOr.And
	andor.SetPoint(left.GetAtom() < base.GetAtom()).
		SetPoint(base.GetAtom() < right.GetAtom())

	e.Result  = andor.GetResult()
	return e.Result
}
