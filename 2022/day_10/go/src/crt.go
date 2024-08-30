package main

import (
	"fmt"
	"os"
	"os/exec"
)

type CRT struct {
	clock      *Clock
	registers  map[string]*Register
	sprite     *Sprite
	pixels     [6][40]rune
	pixelIndex int
	pixelRow   int
}

func NewCRT(clock *Clock, registers map[string]*Register, sprite *Sprite) *CRT {
	pixs := [6][40]rune{}
	for i := 0; i < len(pixs); i++ {
		for j := 0; j < len(pixs[i]); j++ {
			pixs[i][j] = '.'
		}
	}
	return &CRT{
		clock:     clock,
		registers: registers,
		sprite:    sprite,
		pixels:    pixs,
	}
}

func (c *CRT) Update() {
	c.pixelIndex = (c.clock.GetCycles() - 1) % 40
	c.pixelRow = (c.clock.GetCycles() - 1) / 40

	for i := 0; i < len(c.sprite.pixels); i++ {
		spriteIndex := c.sprite.pixels[i]
		if spriteIndex > 0 {
			spriteIndex--
		}
		if spriteIndex == c.pixelIndex {
			c.pixels[c.pixelRow][c.pixelIndex] = '#'
			break
		}
	}

	if c.pixelIndex > 0 && c.pixels[c.pixelRow][c.pixelIndex-1] == ':' {
		c.pixels[c.pixelRow][c.pixelIndex-1] = '.'
	} else if c.pixels[c.pixelRow][c.pixelIndex] != '#' {
		c.pixels[c.pixelRow][c.pixelIndex] = ':'
	}

	args := []string{"-c", "sleep .1 && clear"}
	cmd := exec.Command("bash", args...)
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println(c)
}

func (c *CRT) MoveSprite() {
	c.sprite.Move(c.registers["x"].GetValue())
}

func (c *CRT) String() string {
	var output string
	for i := 0; i < len(c.pixels); i++ {
		for j := 0; j < len(c.pixels[i]); j++ {
			output += string(c.pixels[i][j])
		}
		output += "\n"
	}

	return output
}
