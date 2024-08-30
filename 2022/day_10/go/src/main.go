package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2022 - Day 10")

	var test bool
	var part int

	flag.BoolVar(&test, "test", false, "Run the test input")
	flag.IntVar(&part, "part", 1, "Run part 1 or 2")
	flag.Parse()

	// test = true
	part = 2
	var input string
	if test {
		input = "input/test.txt"
		fmt.Println("Running test input")
	} else {
		input = "input/prod.txt"
		fmt.Println("Running puzzle input")
	}

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	if part == 1 {
		fmt.Println("Part 1")
		emu := NewEmulator(data)
		emu.Run(nil, emu.SignalStrength)
		fmt.Println("Total signal strength:", emu.GetTotalSignalStrength())

	} else {
		fmt.Println("Part 2")

		emu := NewEmulator(data)
		sprite := NewSprite()
		crt := NewCRT(emu.clock, emu.registers, sprite)
		emu.Run(crt.Update, crt.MoveSprite)
	}
}
