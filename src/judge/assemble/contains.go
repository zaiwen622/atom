package assemble

type Contains struct {
	Judge
}

func (c *Contains) GetResult() bool  {
	left := c.LeftExpress.Do().GetAtom()
	right := c.RightExpress.Do().GetAtom()

	resultLeft :=make(map[float64]bool)
	resultRight :=make(map[float64]bool)
	for _,leftVal:=range left {
		resultLeft[leftVal] =true
	}
	for _,rightVal:=range right {
		resultRight[rightVal] =true
	}

	for _,rightVal:=range right {
		if _,ok :=resultLeft[rightVal];!ok {
			return false
		}
	}

	for _,leftVal:=range left {
		if _,ok :=resultRight[leftVal];!ok {
			return false
		}
	}


	return true
}