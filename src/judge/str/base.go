package str

import (
	"warning/src/calculate/str"
	judger "warning/src/judge"
)

const (
	LikeAll = "%%"
	LikePre ="^%"
	LikeEnd ="%$"
	Eq ="="
)

type Judge struct {
	LeftExpress str.CalculateStr
	RightExpress str.CalculateStr
	Result bool
}


func NewJudge(judge string,left,right str.CalculateStr) (obj judger.Judger,ok bool){
	var judgerObj judger.Judger
	switch judge {
	case LikeAll:
		var exp Contains
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case LikePre:
		var exp LikePrefix
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case LikeEnd:
		var exp LikeLast
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	case Eq:
		var exp Equal
		exp.LeftExpress = left
		exp.RightExpress = right
		judgerObj = &exp
	default:
		return nil,false
	}

	return judgerObj,true
}


