package main

import (
	"fmt"
	"warning/src/calculate/assemble"
	"warning/src/calculate/numb"
	"warning/src/calculate/str"
	"warning/src/interpreter"
	judgeAssemble "warning/src/judge/assemble"
	judgeNumb "warning/src/judge/numb"
	judgeStr "warning/src/judge/str"
	"warning/src/operation"
)

func main() {
	// 1+2 = 6


	plus,_ :=numb.AutoExpress(numb.PlusCal,1,2)
	result,_ :=numb.AutoExpress(numb.PlusCal,6,plus.Do().GetAtom())

	//fmt.Println(result.Do())
	returns,_ := judgeNumb.AutoJudge(judgeNumb.Elt,plus,result)
	fmt.Println(returns.GetResult())

	// 3 <= 10 < 11
	numb1,_ := numb.AutoExpress(numb.PlusCal,10)
	numb2,_ := numb.AutoExpress("",10)
	numb3,_ := numb.AutoExpress("",11)
	return1,_ := judgeNumb.AutoJudge(judgeNumb.Egtlt,
		numb1,
		numb2,
		numb3)
	fmt.Println(return1.GetResult())


	//str1,_:=str.NewExpress(str.ReplaceCal,"hello world","world","me")
	//fmt.Println(str1)
	str2,_:=str.NewExpress(str.JoinCal,"ll")
	str3,_:=str.NewExpress(str.JoinCal,"hello")
	fmt.Println(str2,str3)
	result2,_:=judgeStr.NewJudge(judgeStr.LikeAll,str3,str2)

	fmt.Println("str",result2.GetResult())

	assemble1 ,_:=assemble.NewExpress(assemble.ComplementCal,[]float64{1,2,3,5,6},[]float64{1,2,3,5})
	//fmt.Println(assemble1)
	assemble2 ,_:=assemble.BeExpress([]float64{7})
	result3,_:=judgeAssemble.NewJudge(judgeAssemble.StrictEq,
		assemble1,
		assemble2,
		)

	fmt.Println(result3.GetResult())



	boo:=operation.NewOperationLink(returns).//true
		AddOper(operation.And).
		AddJudge(return1).//true
		AddOper(operation.Or).
		AddJudge(result2).//false
		AddOper(operation.Or).
		AddJudge(result3).//false
		Exec()


	fmt.Println(boo)





	///
	dataMap :=map[string]any{
		"left":1,
		"right":1,
		"result":2,
		"one":3,
		"two":2,
		"three":2,
	}
	sent :=interpreter.NewSentence("left + right = result || one - two = three",dataMap)
	fmt.Println(sent.Explain().Exec())

}
