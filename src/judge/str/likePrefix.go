package str

import "strings"

type LikePrefix struct {
	Judge
}

func (c *LikePrefix) GetResult() bool  {
	left :=c.LeftExpress.Do()
	right :=c.RightExpress.Do()
	c.Result = strings.Index(left.GetAtom(),right.GetAtom()) == 0
	return c.Result
}
