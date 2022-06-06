package interpreter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"warning/src/calculate/str"
	judger "warning/src/judge"
	judgeStr "warning/src/judge/str"
)

const (
	StrReg =`\'.*\'`
)

var strRegStr string
var strJudgeRegStr string

func init()  {
	strJudgeRegStr = `(`+judgeStr.LikeAll+`|\`+judgeStr.LikePre+`|`+
		judgeStr.LikeEnd+`|`+judgeStr.Eq+`){1}`
	strRegStr = StrReg +`\s*`+ strJudgeRegStr +`\s*`+ StrReg
	fmt.Println(strRegStr)
}

type StrWord struct {
	Word string
}

func (w *StrWord) Explain() (judge judger.Judger,err error) {
	reg := regexp.MustCompile(strRegStr)
	data :=reg.FindAllString(w.Word,-1)

	if len(data)!=1 {
		return nil,errors.New("cannot explain ")
	}

	reg = regexp.MustCompile(`(?U)`+StrReg)
	strs :=reg.FindAllString(w.Word,-1)
	fmt.Println(strs)
	if len(strs) != 2{
		return nil,errors.New("cannot str count err ")
	}

	var leftStr,rightStr string

	leftStr = strings.Trim(strs[0],"'")
	rightStr = strings.Trim(strs[1],"'")

	reg = regexp.MustCompile(strJudgeRegStr)
	judgeArr :=reg.FindAllString(w.Word,-1)

	if len(judgeArr)!=1 {
		return nil,errors.New("cannot explain judge ")
	}

	judgeWay :=strings.Trim(judgeArr[0],"")


	str1,ok:=str.NewExpress(str.JoinCal,leftStr)
	if !ok {
		return nil,errors.New("cannot leftStr judge ")
	}
	str2,ok:=str.NewExpress(str.JoinCal,rightStr)
	if !ok {
		return nil,errors.New("cannot rightStr judge ")
	}

	judge,ok =judgeStr.NewJudge(judgeWay,str1,str2)
	if !ok {
		return nil,errors.New("cannot str judge ")
	}
	return judge,nil
}
