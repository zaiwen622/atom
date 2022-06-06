package assemble

import "warning/src/atom"

type None struct {
	atom.AtomAssembleNumb
}

func (n *None) Do() *atom.AtomAssembleNumb {
	return &atom.AtomAssembleNumb{
		Assemble: n.Assemble,
	}
}
