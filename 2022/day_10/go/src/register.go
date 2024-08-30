package main

type Register struct {
	Name  string
	value int
}

func NewRegister(name string) *Register {
	return &Register{
		Name:  name,
		value: 1,
	}
}

func (r *Register) Increment() {
	r.value++
}

func (r *Register) Decrement() {
	r.value--
}

func (r *Register) GetValue() int {
	return r.value
}

func (r *Register) SetValue(value int) {
	r.value = value
}

func (r *Register) Add(value int) {
	r.value += value
}
