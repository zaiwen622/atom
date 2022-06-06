package assemble

import "warning/src/atom"

type Complement struct {
	AssemblesNumb
}

func (i *Complement) Do() *atom.AtomAssembleNumb  {
	var result atom.AtomAssembleNumb

	mp := make(map[float64]bool)
	for _, s := range i.LeftAssemble.Assemble {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range i.RightAssemble.Assemble {
		if _, ok := mp[s]; ok {
			delete(mp, s)
		}
	}

	for key := range mp {
		result.Assemble = append(result.Assemble, key)
	}

	return &result
}
