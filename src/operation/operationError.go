package operation

import (
	"fmt"
	"log"
)

// OperationError 组合运算异常
type OperationError struct {
	Code uint
	Reason string
}

func (c OperationError) Error() string  {
	return fmt.Sprintf("组合运算异常 自定义 code %d,原因 %s \n",c.Code,c.Reason)
}

func RecoverOper()  {
	var err any = recover()
	if myError,ok := err.(OperationError);ok{
		log.Fatal(myError.Code,myError.Reason)
	}
}