package atom

import (
	"math"
)


// AtomNumb 计算的最小单元
type AtomNumb struct {

	Numb float64
}

func (n *AtomNumb) GetAtom() float64  {
	return n.Numb
}
// GetAtomToInt 转整数
// way  1 向上取整  2 向下取整  其他 四舍五入
func (n *AtomNumb) GetAtomToInt(way int) int  {
	var result float64
	switch way {
	case 1 :
		result = math.Round(n.GetAtom())
	case 2:
		result = math.Floor(n.GetAtom())
	default:
		result = math.Floor(n.GetAtom() + 0.5)
	}

	return int(result)
}

func (n *AtomNumb) SetAtom(numb float64) {
	n.Numb = numb
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------

// AtomStr 字符串
type AtomStr struct {
	Str string
}


func (n *AtomStr) GetAtom() string  {
	return n.Str
}

func (n *AtomStr) SetAtom(str string) {
	n.Str = str
}

func (n AtomStr) IsEmpty()bool  {
	var str AtomStr
	return str == n
}


//-------------------------------------------------------------------------------------------------------------------------------------------------------

// AtomAssembleNumb 数值集合
type AtomAssembleNumb struct {
	Assemble []float64
}

func NewAtomAssembleNumb(numb... float64) AtomAssembleNumb {
	return AtomAssembleNumb{
		Assemble: numb,
	}
}

func (a *AtomAssembleNumb) append(numb... float64) *AtomAssembleNumb {
	a.Assemble = append(a.Assemble,numb...)
	return a
}

func (a *AtomAssembleNumb) GetAtom() []float64 {
	return a.Assemble
}


func (a *AtomAssembleNumb) Remove(idx int) *AtomAssembleNumb  {

	if idx >= len(a.Assemble) || idx < 0 {
		return a
	}
	v :=a.Assemble[idx]
	if v == 0 {
		return a
	}
	a.Assemble = a.Assemble[:idx+copy(a.Assemble[idx:], a.Assemble[idx+1:])] // 删除中间1个元素

	return a
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------


// AtomAssembleStr 字符串集合
type AtomAssembleStr struct {

	Assemble []string
}

func (a *AtomAssembleStr) new(numb... string) *AtomAssembleStr {
	return &AtomAssembleStr{Assemble: numb}
}
func (a *AtomAssembleStr) append(numb... string) *AtomAssembleStr {
	a.Assemble = append(a.Assemble,numb...)
	return a
}


func (a *AtomAssembleStr) Remove(idx int) *AtomAssembleStr  {

	if idx >= len(a.Assemble) || idx < 0 {
		return a
	}
	v :=a.Assemble[idx]
	if v =="" {
		return a
	}
	a.Assemble = a.Assemble[:idx+copy(a.Assemble[idx:], a.Assemble[idx+1:])] // 删除中间1个元素

	return a
}
