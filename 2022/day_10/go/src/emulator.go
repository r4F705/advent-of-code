package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: Could also be named as CPU
type Emulator struct {
	clock               *Clock
	curInstruction      int
	instructions        []Instruction
	registers           map[string]*Register
	opcodes             []byte
	totalSignalStrength int
}

func NewEmulator(data []byte) *Emulator {
	return &Emulator{
		clock:   NewClock(),
		opcodes: data,
		registers: map[string]*Register{
			"x": NewRegister("x"),
		},
		instructions:   parseOpCodes(data),
		curInstruction: 0,
	}
}

func parseOpCodes(opcodes []byte) []Instruction {
	var instructions []Instruction
	for _, opcode := range strings.Split(string(opcodes), "\n") {
		opdata := strings.Split(opcode, " ")

		if len(opdata) == 0 {
			panic("Invalid opcode")
		}

		switch opdata[0] {
		case "noop":
			instructions = append(instructions, &Nop{Op: opdata[0]})
		case "addx":
			value, err := strconv.Atoi(opdata[1])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, &Add{Op: opdata[0], Value: value})
		}
	}

	return instructions
}

func (e *Emulator) GetClock() *Clock {
	return e.clock
}

func (e *Emulator) GetTotalSignalStrength() int {
	return e.totalSignalStrength
}

func (e *Emulator) SignalStrength() {
	if e.clock.SignalStrengthCycle() {
		signalStrength := e.registers["x"].GetValue() * e.clock.GetCycles()
		fmt.Printf("Signal strength: %s(%d) at %d : %d \n", e.registers["x"].Name, e.registers["x"].GetValue(), e.clock.GetCycles(), signalStrength)
		e.totalSignalStrength += signalStrength
	}
}

func (e *Emulator) Run(beforeCycle, afterCycle func()) {
	for {
		if e.curInstruction >= len(e.instructions) {
			break
		}

		beforeCycle()

		instruction := e.instructions[e.curInstruction]
		if instruction.Execute(e.registers["x"]) {
			e.curInstruction++
			fmt.Printf("Cycle: %d, Instruction: %s, Register: %s(%d)\n", e.clock.GetCycles(), instruction, e.registers["x"].Name, e.registers["x"].GetValue())
		} else {
			fmt.Printf("Cycle: %d, Register: %s(%d)\n", e.clock.GetCycles(), e.registers["x"].Name, e.registers["x"].GetValue())
		}

		e.clock.Increment()

		afterCycle()
	}
}
