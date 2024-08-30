package main

import "strconv"

type Instruction interface {
	Execute(reg *Register) bool
}

type Nop struct {
	Op string
}

func (n *Nop) Execute(reg *Register) bool {
	return true
}

func (n *Nop) String() string {
	return n.Op
}

type Add struct {
	Op     string
	Value  int
	cycles int
}

func (a *Add) Execute(reg *Register) bool {
	a.cycles++

	if a.cycles == 2 {
		reg.Add(a.Value)
		return true
	}

	return false
}

func (a *Add) String() string {
	return a.Op + " " + strconv.Itoa(a.Value)
}
