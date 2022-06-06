package interpreter

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"warning/src/calculate/numb"
	judger "warning/src/judge"
	judgeNum "warning/src/judge/numb"
)

const (
	NumbReg =`([+-]*\d+)(\.\d+)?`
)

var numRegStr string
var numCalRegStr string
var numJudgeRegStr string

func init()  {
	numCalRegStr =`[\`+numb.PlusCal+`|\`+numb.DivisionCal+`|\`+numb.MultiCal+
		`|\`+numb.PowerCal+`|\`+numb.RemainderCal+`|`+numb.SubCal+`]`
	numJudgeRegStr =`[`+judgeNum.Eq+`|`+judgeNum.Neq+`|`+judgeNum.Egt+`|`+judgeNum.Gt+`|`+judgeNum.Lt+
		`|`+judgeNum.Elt+
		`]`
	numRegStr = NumbReg +`\s*`+numCalRegStr+`\s*` + NumbReg+
		`\s*`+numJudgeRegStr+`\s*`+ NumbReg

	fmt.Println(numRegStr)
}


type NumWord struct {
	Word string
}

func (w *NumWord) Explain() (judge judger.Judger,err error) {
	reg := regexp.MustCompile(numRegStr)
	data :=reg.FindAllString(w.Word,-1)
	if len(data)!=1 {
		return nil,errors.New("cannot explain ")
	}
	var left,right, result float64

	reg = regexp.MustCompile(NumbReg)
	numbs :=reg.FindAllString(w.Word,-1)
	left,err = strconv.ParseFloat(numbs[0], 64)
	if err !=nil {
		return nil,err
	}

	right,err = strconv.ParseFloat(numbs[1], 64)
	if err !=nil {
		return nil,err
	}

	result,err = strconv.ParseFloat(numbs[2], 64)
	if err !=nil {
		return nil,err
	}


	reg = regexp.MustCompile(numJudgeRegStr)
	judgeArr :=reg.FindAllString(w.Word,-1)

	if len(judgeArr)!=1 {
		return nil,errors.New("cannot explain judge ")
	}

	judgeWay :=strings.Trim(judgeArr[0],"")

	reg = regexp.MustCompile(numCalRegStr)
	calArr :=reg.FindAllString(w.Word,-1)

	if len(calArr)!=1 {
		return nil,errors.New("cannot explain cal ")
	}

	calWay :=strings.Trim(calArr[0],"")


	leftResult,_ :=numb.AutoExpress(calWay,left,right)
	rightResult,_ :=numb.AutoExpress(numb.PlusCal,result)

	judge,ok := judgeNum.AutoJudge(judgeWay,leftResult,rightResult)
	if !ok{
		return nil,errors.New("cannot explain judge ")
	}

	return judge,nil
}

