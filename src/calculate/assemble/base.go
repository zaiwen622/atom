package assemble

import "warning/src/atom"

const (
	InterCal ="∩"
	UnionCal ="∪"
	ComplementCal = "CuP"
)

type AssemblesNumb struct {
	LeftAssemble atom.AtomAssembleNumb
	RightAssemble atom.AtomAssembleNumb
}


func NewExpress(cal string ,left []float64,right []float64) (obj CalculateAssembleNumb, ok bool) {
	var result  CalculateAssembleNumb
	switch cal {
	case InterCal:
		 exp :=new(Inter)
		 exp.LeftAssemble = atom.NewAtomAssembleNumb(left...)
		 exp.RightAssemble = atom.NewAtomAssembleNumb(right...)
		 result = exp
	case UnionCal:
		exp :=new(Union)
		exp.LeftAssemble = atom.NewAtomAssembleNumb(left...)
		exp.RightAssemble = atom.NewAtomAssembleNumb(right...)
		result = exp
	case ComplementCal:
		exp :=new(Complement)
		exp.LeftAssemble = atom.NewAtomAssembleNumb(left...)
		exp.RightAssemble = atom.NewAtomAssembleNumb(right...)
		result = exp
	default:
		return BeExpress(left)

	}

	return result,true
}

func BeExpress(numbs []float64) (obj CalculateAssembleNumb, ok bool) {
	var result  CalculateAssembleNumb
	exp :=new(None)
	newAtom :=atom.NewAtomAssembleNumb(numbs...)
	exp.Assemble= newAtom.GetAtom()
	result = exp
	return result,true
}


type CalculateAssembleNumb interface {
	Do() *atom.AtomAssembleNumb
}
