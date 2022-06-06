package interpreter

import (
	"fmt"
	"strconv"
	"strings"
	judger "warning/src/judge"
	"warning/src/operation"
)

type Sentence struct {
	Data string
	DataMap map[string]any
}


func (s *Sentence) GetData() string  {
	return s.Data
}

func (s *Sentence) SetData(data string)  {
	s.Data = data
}

func NewSentence(str string,dataMap map[string]any) *Sentence {

	obj :=&Sentence{
		Data: str,
		DataMap: dataMap,
	}
	obj.replaceData()
	return obj
}

func (s *Sentence) Explain() operation.Executor {
	defer operation.RecoverOper()

	result := ExploredData(s.Data)
	//make judge
	var link *operation.OperationLink
	var sent ,oper string
	for idx,singleStr :=range result {

		sent,oper = explodeSingle(singleStr)
		if idx ==0 {
			judge :=makeJudge(sent)
			link =operation.NewOperationLink(judge)
			if oper !="" {
				link = link.AddOper(oper)
			}
			continue
		}

		if sent !="" {
			link = link.AddJudge(makeJudge(sent))
		}
		if oper != "" {
			link = link.AddOper(oper)
		}

	}


	return link
}

func ExploredData(str string) (result []string) {

	and :=strings.SplitAfter(str,operation.And)
	for _,val :=range and {
		result = append(result,strings.SplitAfter(val,operation.Or)...)
	}
	return
}

func (s *Sentence) replaceData()*Sentence  {
	for str,val :=range s.DataMap {
		var valStr string
		switch val.(type) {
		case string:
			valStr = val.(string)
		case int:
			valStr = strconv.Itoa(val.(int))

		case []float64,[]int,[]float32:
			valStr = fmt.Sprintf("%s",val)
		}
		s.Data = strings.ReplaceAll(s.Data,str,valStr)
	}

	return s
}

func makeJudge(singleStr string) (judge judger.Judger)  {
	str:=StrWord{
		Word: singleStr,

	}
	judge,err :=str.Explain()
	if err ==nil{
		return
	}

	strNum :=NumWord{
		Word: singleStr,
	}
	judge,err =strNum.Explain()
	if err ==nil {
		return
	}

	assembleStr :=AssembleWord{
		Word: singleStr,
	}
	judge,err =assembleStr.Explain()
	if err ==nil {
		return
	}


	return nil

}

func explodeSingle(singleStr string) (sent ,oper string){
	andBool:=strings.HasSuffix(singleStr,operation.And)
	if andBool {
		idx :=strings.LastIndex(singleStr, operation.And)
		sent = singleStr[:idx]
		oper = singleStr[idx:]
		return
	}

	orBool:=strings.HasSuffix(singleStr,operation.Or)
	if orBool {
		idx :=strings.LastIndex(singleStr, operation.Or)
		sent = singleStr[:idx]
		oper = singleStr[idx:]
		return
	}

	return singleStr,""
}

