package assemble

import "warning/src/atom"

type Union struct {
	AssemblesNumb
}

func (i *Union) Do() *atom.AtomAssembleNumb  {
	var result atom.AtomAssembleNumb

	result.Assemble = append(result.Assemble, i.LeftAssemble.Assemble...)
	result.Assemble = append(result.Assemble, i.RightAssemble.Assemble...)

	return &result
}