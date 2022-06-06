package interpreter

import (
	"fmt"
	"testing"
)

func TestWord(t *testing.T) {
	str :=NumWord{
		Word: "1 + 1=2",
	}
	obj,err:=str.Explain()
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(obj.GetResult())
}
