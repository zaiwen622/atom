package str

import (
	"strings"
)

type Replace struct {
	Strs
}

func (r *Replace) New(s , old , new string) *Replace {
	r.Append(s).Append(old).Append(new)
	return r
}

func (r *Replace) Do() *AtomStr {
	var result AtomStr
	if r.AllStr[0].IsEmpty() || r.AllStr[1].IsEmpty() || r.AllStr[2].IsEmpty() {
		return &result
	}
	result.SetAtom(strings.ReplaceAll(r.AllStr[0].GetAtom(),r.AllStr[1].GetAtom(),r.AllStr[2].GetAtom()))

	return &result
}