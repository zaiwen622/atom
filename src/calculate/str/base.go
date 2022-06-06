package str

import (
	"warning/src/atom"
	"warning/src/calculate"
)

type AtomStr = atom.AtomStr
const (
	ReplaceCal = "<>"
	JoinCal ="+"
)

type Strs struct {
	AllStr []AtomStr
}

func (s *Strs) ToStringArr() []string {
	var result []string
	for _,v :=range s.AllStr {
		if !v.IsEmpty(){
			result = append(result,v.GetAtom())
		}

	}
	return result
}

func (s *Strs) Append(str string) *Strs {
	var atomStr AtomStr
	atomStr.SetAtom(str)
	if atomStr.IsEmpty() {
		panic(any(calculate.CalculateError{
			Code:   201,
			Reason: "str can not empty",
		}))
	}
	s.AllStr = append(s.AllStr,atomStr)
	return s
}

func (s *Strs) Remove(idx int) *Strs  {
	a  := s.AllStr
	if idx >= len(a) || idx < 0 {
		return s
	}
	v :=a[idx]
	if v.IsEmpty() {
		return s
	}
	a = a[:idx+copy(a[idx:], a[idx+1:])] // 删除中间1个元素

	s.AllStr = a

	return s
}

func NewExpress(cal string,str...string) (obj CalculateStr,ok bool) {
	var calculateStr CalculateStr


	switch cal {
	case ReplaceCal:
		if len(str) !=3 {
			return nil,false
		}

		var exp Replace
		exp.New(str[0],str[1],str[2])
		calculateStr = &exp
	case JoinCal:
		var exp Join
		exp.New(str...)
		calculateStr = &exp

	default: //用于承接 单个值 转表达式
		var exp Join
		exp.New(str...)
		calculateStr = &exp
	}


	return calculateStr,true
}

type CalculateStr interface {
	Do() *AtomStr
}