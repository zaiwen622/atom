package operation

import (
	judger "warning/src/judge"
)

const (
	And ="&&"
	Or ="||"
)

type Operation struct {
	Oper string
	Judger judger.Judger
}


type OperationLink struct {
	Count int
	link []*Operation
}

func (o *Operation) IsEmpty()bool  {
	var temp *Operation
	return temp == o
}

func NewOperationLink(judgerObj judger.Judger) *OperationLink  {
	newLink :=  &Operation{
		Judger:judgerObj,
	}
	link :=&OperationLink{
		link: make([]*Operation,0),
	}
	link.link = append(link.link,newLink)
	link.Count ++
	return link
}

func (o *OperationLink) AddOper(oper string) *OperationLink {
	//要单数
	if o.Count % 2 !=1 {
		panic(any(OperationError{
			Code: 201,
			Reason: "and oper err",
		}))
	}

	newLink :=  &Operation{
		Oper:oper,
	}

	o.link = append(o.link,newLink)
	o.Count ++


	return o
}

func (o *OperationLink) AddJudge(judge judger.Judger) *OperationLink {
	//要双数
	if o.Count % 2 !=0 {
		panic(any(OperationError{
			Code: 201,
			Reason: "and judge err",
		}))
	}

	newLink :=  &Operation{
		Judger:judge,
	}

	o.link = append(o.link,newLink)
	o.Count ++

	return o
}


type Executor interface {
	Exec()bool
}


func (o *OperationLink) Exec() bool {

	if o.Count % 2 != 1 {
		panic(any(OperationError{
			Code: 201,
			Reason: "and exec count err",
		}))
	}
	if o.Count == 1 {
		return o.link[0].Judger.GetResult()
	}

	result :=o.link[0].Judger.GetResult()
	for i:=1;i<o.Count;i=i+2 {
		step :=i+1
		result = doExec(result,o.link[i].Oper,o.link[step].Judger)
	}


	return result
}

func doExec(leftResult bool,op string ,right judger.Judger) bool {
	rightResult := right.GetResult()
	switch op {
	case And:
		return leftResult && rightResult
	case Or:
		return leftResult || rightResult
	default:
		panic(any(OperationError{
			Code: 201,
			Reason:"Operation unknown",
		}))
	}
}



