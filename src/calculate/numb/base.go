package numb

import (
	atom1 "warning/src/atom"
)

type AtomNumb = atom1.AtomNumb

const (
	PlusCal ="+"
	DivisionCal ="/"
	MultiCal = "*"
	PowerCal = "^"
	RemainderCal = "%"
	SubCal ="-"
)

type Express struct {
	LeftAtom AtomNumb
	RightAtom AtomNumb
}

func AutoExpress(cal string, numbs... float64) (obj CalculateNumb,ok bool) {
	count :=len(numbs)
	if count ==1 {
		return BeExpress(numbs[0]),true
	}
	if count == 2 {
		return NewExpress(numbs[0] ,numbs[1],cal),true
	}
	return nil,false
}

func NewExpress(numb1 float64,numb2 float64,cal string) CalculateNumb {
	var calculate CalculateNumb


	switch cal {
	case DivisionCal:
		var exp Division
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	case MultiCal:
		var exp Multi
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	case PowerCal:
		var exp Power
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	case RemainderCal:
		var exp Remainder
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	case SubCal:
		var exp Sub
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	case PlusCal:
		var exp Plus
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	default: //用于承接 单个值 转表达式
		var exp Plus
		exp.LeftAtom.SetAtom(numb1)
		exp.RightAtom.SetAtom(numb2)
		calculate = &exp
	}


	return calculate
}

func BeExpress(numb1 float64) CalculateNumb  {
	return NewExpress(numb1,0,"")
}


type CalculateNumb interface {
	Do() *AtomNumb
	isValid() bool
}

