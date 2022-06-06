package interpreter

import (
	"fmt"
	"testing"
)

func TestExploredData(t *testing.T) {
	data:=ExploredData(`'oPrice' - 'sPrice' >=0 && 'state'  ∩ [1,3,4] || 'id' = "444"`)
	fmt.Printf("%#v \n",data)
}
