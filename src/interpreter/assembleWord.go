package interpreter

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"warning/src/calculate/assemble"
	judger "warning/src/judge"
	judgeAssemble "warning/src/judge/assemble"
)

const (
	AssembleReg =`\[(\s*`+NumbReg+`\s*,?)*\]`
)

var assRegStr string
var assCalRegStr string
var assJudgeRegStr string

func init()  {
	assCalRegStr =`(`+assemble.InterCal+`|`+assemble.UnionCal+`|`+assemble.ComplementCal+`){1}`
	assJudgeRegStr =`(`+judgeAssemble.Eq+`|`+judgeAssemble.StrictEq+`){1}`

	assRegStr = AssembleReg +`\s*`+ assCalRegStr +`\s*`+
		AssembleReg +`\s*`+ assJudgeRegStr +`\s*`+ AssembleReg

	fmt.Println(assRegStr)

}

type AssembleWord struct {
	Word string
}

func (a *AssembleWord) Explain() (judge judger.Judger,err error) {
	reg := regexp.MustCompile(assRegStr)
	data :=reg.FindAllString(a.Word,-1)
	if len(data)!=1 {
		return nil,errors.New("cannot explain ")
	}

	var left,right, result []float64
	reg = regexp.MustCompile(AssembleReg)
	ass :=reg.FindAllString(a.Word,-1)
	left,right,result,err = a.toFloat(ass)
	if err !=nil {
		return nil,err
	}

	reg = regexp.MustCompile(assJudgeRegStr)
	judgeArr :=reg.FindAllString(a.Word,-1)

	if len(judgeArr)!=1 {
		return nil,errors.New("cannot explain judge ")
	}

	judgeWay :=strings.Trim(judgeArr[0],"")

	reg = regexp.MustCompile(assCalRegStr)
	calArr :=reg.FindAllString(a.Word,-1)

	if len(calArr)!=1 {
		return nil,errors.New("cannot explain cal ")
	}

	calWay :=strings.Trim(calArr[0],"")


	assemble1 ,ok :=assemble.NewExpress(calWay,left,right)
	if !ok {
		return nil,errors.New("cannot explain assemble1 ")
	}
	//fmt.Println(assemble1)
	assemble2 ,ok :=assemble.BeExpress(result)
	if !ok {
		return nil,errors.New("cannot explain assemble2 ")
	}
	judge,ok =judgeAssemble.NewJudge(judgeWay,
		assemble1,
		assemble2,
	)
	if !ok {
		return nil,errors.New("cannot explain judge ")
	}

	return judge,nil
}

func (a *AssembleWord) toFloat(ass []string) (left,right, result []float64, err error) {
	if len(ass) !=3 {
		return nil,nil,nil,errors.New("assemble nums err")
	}

	left,err =toFloatArr(ass[0])
	if err !=nil {
		return nil,nil,nil,errors.New("assemble nums left err")
	}
	right,err =toFloatArr(ass[1])
	if err !=nil {
		return nil,nil,nil,errors.New("assemble nums right err")
	}
	result,err =toFloatArr(ass[2])
	if err !=nil {
		return nil,nil,nil,errors.New("assemble nums result err")
	}
	return

}

func toFloatArr(str string)(resul []float64,err error)  {
	reg := regexp.MustCompile(NumbReg)
	numbs :=reg.FindAllString(str,-1)
	var result []float64
	for _,numb :=range numbs {
		flo,err :=strconv.ParseFloat(numb,64)
		if err !=nil {
			return nil,err
		}
		result = append(result,flo)
	}
	return result,nil
}