package numb

type Remainder struct {
	Express
	Way int //way  1 向上取整  2 向下取整  其他 四舍五入
}

func (p *Remainder) Do() *AtomNumb  {
	var result AtomNumb
	result.SetAtom(float64(p.LeftAtom.GetAtomToInt(p.Way) % p.RightAtom.GetAtomToInt(p.Way)))
	return &result
}

func (p *Remainder) isValid() bool  {


	return true
}
