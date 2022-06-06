package str

type Equal struct {
	Judge
}

func (c *Equal) GetResult() bool  {
	left :=c.LeftExpress.Do()
	right :=c.RightExpress.Do()
	c.Result = left.GetAtom() == right.GetAtom()
	return c.Result
}
