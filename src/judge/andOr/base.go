package andOr

type AndOr struct {
	Point []bool
}

func (o *AndOr) SetPoint(b bool) *AndOr {
	o.Point=append(o.Point, b)
	return o
}

type And struct {
	AndOr
}

type Or struct {
	AndOr
}

func (o *And) GetResult() bool {
	result := o.Point[0]
	for _,bo :=range o.Point {
		result = result && bo
		if result == false { // 一假则假
			return false
		}
	}
	return result
}

func (o *Or) GetResult() bool {
	result := o.Point[0]
	for _,bo :=range o.Point {
		result = result || bo
		if result == true { // 一真则真
			return true
		}
	}
	return result
}