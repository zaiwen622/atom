package str

import "strings"

type LikeLast struct {
	Judge
}

func (c *LikeLast) GetResult() bool  {
	left :=c.LeftExpress.Do()
	right :=c.RightExpress.Do()
	c.Result = strings.LastIndex(left.GetAtom(),right.GetAtom()) ==
		(len(left.GetAtom())-len(right.GetAtom()))
	return c.Result
}
