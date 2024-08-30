package main

type Clock struct {
	cycles       int
	signalCycles int
}

func NewClock() *Clock {
	return &Clock{
		cycles: 1,
	}
}

func (c *Clock) Increment() {
	if c.signalCycles == 40 {
		c.signalCycles = 0
	}

	c.cycles++

	if c.cycles > 20 {
		c.signalCycles++
	}
}

func (c *Clock) GetCycles() int {
	return c.cycles
}

func (c *Clock) SignalStrengthCycle() bool {
	if c.cycles == 20 {
		return true
	} else if c.signalCycles == 40 {
		return true
	}

	return false
}
