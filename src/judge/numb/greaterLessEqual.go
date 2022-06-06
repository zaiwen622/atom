package numb

import (
	"fmt"
	"warning/src/calculate"
	"warning/src/judge/andOr"
)

type GreaterLessEqual struct {
	JudgeScope
}

func (e *GreaterLessEqual) GetResult()bool  {
	defer calculate.RecoverCal()
	left := e.LeftExpress.Do()
	right := e.RightExpress.Do()
	base := e.BaseExpress.Do()
	var andor andOr.And
	andor.SetPoint(left.GetAtom() < base.GetAtom()).
		SetPoint(base.GetAtom() <= right.GetAtom())

	fmt.Println(andor)

	e.Result  = andor.GetResult()
	return e.Result
}
