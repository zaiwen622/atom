package assemble

import (
	"warning/src/calculate/assemble"
	judger "warning/src/judge"
)

const (
	Eq ="=" // 左集合被包含在右集合
	StrictEq = "==" // 左元素 与右元素 个数相同
)

type Judge struct {
	LeftExpress assemble.CalculateAssembleNumb
	RightExpress assemble.CalculateAssembleNumb
	Result bool
}


func NewJudge(judge string,left,right assemble.CalculateAssembleNumb) (obj judger.Judger,ok bool){
	var judgerObj judger.Judger
	switch judge {
	case Eq:
		var exp Contains
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case StrictEq:
		var exp StrictContains
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp

	default:
		return nil,false
	}

	return judgerObj,true
}