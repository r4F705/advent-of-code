package cargo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cargo struct {
	StacksCount int
	Stacks      [][]Crate
	initialized bool
}

func (c *Cargo) Init(data []string) {
	stacksLine := data[len(data)-1]
	stackCountStr := stacksLine[len(stacksLine)-4:]
	stackCountStr = strings.Trim(stackCountStr, " ")
	stackCount, _ := strconv.Atoi(stackCountStr)
	c.StacksCount = stackCount
	c.Stacks = make([][]Crate, c.StacksCount)

	for i := len(data) - 2; i >= 0; i -= 1 {
		seek := 0
		content := ""
		for j := 0; j < c.StacksCount; j += 1 {
			if seek+4 < len(data[i]) {
				content = data[i][seek : seek+4]
				seek += 4
			} else {
				content = data[i][seek : seek+3]
			}

			if strings.Contains(content, "[") {
				c.Push(j+1, Crate{Content: content})
			}
		}
	}

	c.initialized = true
}

func (c *Cargo) Push(stack int, crate Crate) {
	if stack > 0 {
		stack -= 1
	}
	c.Stacks[stack] = append(c.Stacks[stack], crate)
}

func (c *Cargo) Pop(stack int) (Crate, error) {
	if stack > 0 {
		stack -= 1
	}

	if len(c.Stacks[stack]) > 0 {
		crate := c.Stacks[stack][len(c.Stacks[stack])-1]
		c.Stacks[stack] = c.Stacks[stack][:len(c.Stacks[stack])-1]
		return crate, nil
	}

	return Crate{}, errors.New("cannot pop empty stack")
}

func (c *Cargo) ShowTopLayer() {
	if c.initialized {
		topLayerCrates := ""
		for i := 0; i < c.StacksCount; i += 1 {
			crate := c.Stacks[i][len(c.Stacks[i])-1]
			topLayerCrates += crate.Content
		}
		topLayerCrates = strings.ReplaceAll(topLayerCrates, "[", "")
		topLayerCrates = strings.ReplaceAll(topLayerCrates, "]", "")
		topLayerCrates = strings.ReplaceAll(topLayerCrates, " ", "")
		fmt.Printf("Top layer crates: %s\n\n", topLayerCrates)
	}
}

func (c *Cargo) ShowCargo() {
	fmt.Println("======= Cargo contents =======")
	for i := 0; i < c.StacksCount; i += 1 {
		for j := 0; j < len(c.Stacks[i]); j += 1 {
			fmt.Printf("%s ", c.Stacks[i][j].Content)
		}
		fmt.Println()
	}
	fmt.Println("==============================")
	fmt.Println()
}
