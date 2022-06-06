package assemble

import "warning/src/atom"

type Inter struct {
	AssemblesNumb
}

func (i *Inter) Do() *atom.AtomAssembleNumb  {
	var result atom.AtomAssembleNumb

	mp := make(map[float64]bool)

	for _, s := range i.LeftAssemble.Assemble {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range i.RightAssemble.Assemble {
		if _, ok := mp[s]; ok {
			result.Assemble = append(result.Assemble, s)
		}
	}

	return &result
}