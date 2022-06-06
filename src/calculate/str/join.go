package str

import "strings"

type Join struct {
	Strs
}

func (j *Join) New(str... string) *Join {
	for _,s :=range str {
		j.Append(s)
	}
	return j
}

func (j *Join) Do() *AtomStr {
	var result AtomStr
	result.SetAtom(strings.Join(j.ToStringArr(),""))

	return &result
}