package numb

import (
	"warning/src/calculate/numb"
	judger "warning/src/judge"
)
// JudgeScope LeftExpress (< <= ) BaseExpress (< <= ) RightExpress
type JudgeScope struct {
	LeftExpress numb.CalculateNumb
	BaseExpress numb.CalculateNumb
	RightExpress numb.CalculateNumb
	Result bool
}


func NewJudgeScope(judge string,left,base,right numb.CalculateNumb) (obj judger.Judger,ok bool) {
	var judgerObj judger.Judger
	switch judge {
	case Glt:
		var exp GreaterLess
		exp.LeftExpress = left
		exp.BaseExpress = base
		exp.RightExpress = right
		judgerObj = &exp
	case GtElt:
		var exp GreaterLessEqual
		exp.LeftExpress = left
		exp.BaseExpress = base
		exp.RightExpress = right
		judgerObj = &exp
	case Egtelt:
		var exp GreaterEqualLessEqual
		exp.LeftExpress = left
		exp.BaseExpress = base
		exp.RightExpress = right
		judgerObj = &exp
	case Egtlt:
		var exp GreaterEqualLess
		exp.LeftExpress = left
		exp.BaseExpress = base
		exp.RightExpress = right
		judgerObj = &exp
	default:
		return nil,false

	}


	return judgerObj,true
}






