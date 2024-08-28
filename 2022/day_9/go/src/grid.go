package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]int
	Rope   *Rope
	Render bool
}

func NewGrid(width, height int, rope *Rope) *Grid {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
	}
	return &Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
		Rope:   rope,
		Render: false,
	}
}

func (g *Grid) CountVisited() int {
	count := 0
	for _, row := range g.Cells {
		for _, cell := range row {
			if cell > 0 {
				count += 1
			}
		}
	}
	return count
}

func (g *Grid) Update() {

	var args []string
	if g.Render {
		args = []string{"-c", "sleep 1 && clear"}
		cmd := exec.Command("bash", args...)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	g.Cells[g.Rope.Tail.Y][g.Rope.Tail.X] += 1

	if g.Render {
		fmt.Println(g)
	}
}

func (g *Grid) String() string {
	var str string
	for i, row := range g.Cells {
		for j, cell := range row {
			if i == g.Rope.Head.Y && j == g.Rope.Head.X {
				str += "H"
			} else if id := g.bodyInCell(i, j); id != -1 {
				str += strconv.Itoa(id)
			} else if i == g.Rope.Tail.Y && j == g.Rope.Tail.X {
				str += "T"
			} else if cell == 0 {
				str += "."
			} else {
				str += "#"
			}
		}
		str += "\n"
	}
	return str
}

func (g *Grid) bodyInCell(i, j int) int {
	for k := 1; k < len(g.Rope.Knots)-1; k++ {
		if g.Rope.Knots[k].X == j && g.Rope.Knots[k].Y == i {
			return k
		}
	}

	return -1
}
