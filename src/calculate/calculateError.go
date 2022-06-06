package calculate

import (
	"fmt"
	"log"
)

// CalculateError 表达式异常
type CalculateError struct {
	Code uint
	Reason string
}

func (c CalculateError) Error() string  {
	return fmt.Sprintf("计算表达式异常 自定义 code %d,原因 %s \n",c.Code,c.Reason)
}

func RecoverCal()  {
	var err any = recover()
	if myError,ok := err.(CalculateError);ok{
		log.Fatal(myError.Code,myError.Reason)
	}
}