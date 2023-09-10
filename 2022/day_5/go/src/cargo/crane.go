package cargo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Crane struct {
	instructions []CraneInstruction
	Model        int
	Debug        bool
	Visualize    bool
}

type CraneInstruction struct {
	count int
	src   int
	dest  int
}

func (c *Crane) LoadInstructions(instructionData []string) {
	for _, v := range instructionData {
		tmp := strings.Split(v, " ")
		c.instructions = append(c.instructions, CraneInstruction{
			count: c.processInstructionData(tmp[1]),
			src:   c.processInstructionData(tmp[3]),
			dest:  c.processInstructionData(tmp[5]),
		})
	}
}

func (c *Crane) ExecuteInstructions(cargo *Cargo) {
	for _, v := range c.instructions {
		bufferZone := []Crate{}
		for i := 0; i < v.count; i += 1 {

			if c.Visualize {
				time.Sleep(1 * time.Second)
			}

			if c.Debug {
				fmt.Printf("Instruction %d: move %d from %d to %d\n", i, v.count, v.src, v.dest)
				cargo.ShowCargo()
			}

			if c.Model != 9001 {
				crate, err := cargo.Pop(v.src)
				if err != nil {
					log.Fatalf("intruction not valid: `move %d from %d to %d`.\n error: %s",
						v.count, v.src, v.dest, err)
				}
				cargo.Push(v.dest, crate)
			} else {
				crate, err := cargo.Pop(v.src)
				bufferZone = append(bufferZone, crate)

				if err != nil {
					log.Fatalf("intruction not valid: `move %d from %d to %d`.\n error: %s",
						v.count, v.src, v.dest, err)
				}

				if i == v.count-1 {
					for j := len(bufferZone) - 1; j >= 0; j -= 1 {
						cargo.Push(v.dest, bufferZone[j])
					}
				}
			}
		}
	}
}

func (c *Crane) processInstructionData(data string) int {
	value, err := strconv.Atoi(data)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (c *Crane) ShowInstructions() {
	fmt.Println("======= Crane Instructions =======")
	for _, v := range c.instructions {
		fmt.Printf("move %d from %d to %d\n", v.count, v.src, v.dest)
	}
	fmt.Println("==================================")
	fmt.Println()
}
