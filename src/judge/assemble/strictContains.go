package assemble

type StrictContains struct {
	Judge
}

func (c *StrictContains) GetResult() bool  {
	left := c.LeftExpress.Do().GetAtom()
	right := c.RightExpress.Do().GetAtom()
	if len(left)!=len(right) {
		return false
	}


	for key,leftVal:=range left {
		rightVal:=right[key]

		if leftVal != rightVal {
			return false
		}
	}

	return true
}