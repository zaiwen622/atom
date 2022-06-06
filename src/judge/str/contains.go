package str

import (
	"fmt"
	"strings"
)

type Contains struct {
	Judge
}

func (c *Contains) GetResult() bool  {
	left :=c.LeftExpress.Do()
	right :=c.RightExpress.Do()
	c.Result = strings.Contains(left.GetAtom(),right.GetAtom())
	fmt.Println(c.Result,left.GetAtom(),right.GetAtom())
	return c.Result
}
