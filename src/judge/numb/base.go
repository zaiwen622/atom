package numb

import (
	"warning/src/calculate/numb"
	judger "warning/src/judge"
)

const (
	Eq  = "="
	Neq = "!="
	Egt = ">="
	Gt = ">"
	Lt = "<"
	Elt = "<="
	Glt = "<<" // a<x<b
	Egtlt = "<=<" //a<=x<b
	GtElt = "<<=" //a<x<=b
	Egtelt = "<=<=" //a<=x<=b
)


type Judge struct {
	LeftExpress numb.CalculateNumb
	RightExpress numb.CalculateNumb
	Result bool
}

func AutoJudge(judge string,numbs... numb.CalculateNumb) (ojb judger.Judger,ok bool) {
	count :=len(numbs)
	if count == 2 {

		return NewJudge(judge,numbs[0],numbs[1])
	}
	if count == 3 {
		return NewJudgeScope(judge,numbs[0],numbs[1],numbs[2])
	}


	return nil,false
}


func NewJudge(judge string,left,right numb.CalculateNumb) (obj judger.Judger,ok bool){
	var judgerObj judger.Judger
	switch judge {
	case Eq:
		var exp Equal
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Lt:
		var exp Less
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Gt:
		var exp Greater
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Egt:
		var exp GreaterEqual
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Neq:
		var exp Nequal
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Elt:
		var exp LessEqual
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	default:
		return nil,false
	}

	return judgerObj,true
}





