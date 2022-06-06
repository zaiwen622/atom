package interpreter

import (
	"fmt"
	"log"
	"testing"
)

func TestStrWord(t *testing.T) {
	str:=StrWord{
		Word: "'hello' %% 'll'",

	}
	obj,err:=str.Explain()
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(obj.GetResult())
}