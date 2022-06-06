package interpreter

import (
	"fmt"
	"log"
	"testing"
)

func TestAssemble(t *testing.T) {
	str :=AssembleWord{
		 Word: "[1,2,3]∪[1,2,4] = [1,2,3,4,1]",
	 }
	obj,err:=str.Explain()
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println(obj.GetResult())
}
