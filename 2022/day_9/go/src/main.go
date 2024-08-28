package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2022 - Day 9")
	var mode bool
	var render bool
	var part int

	flag.BoolVar(&mode, "test", false, "Runs the program with test input")
	flag.BoolVar(&render, "render", false, "Creates a visual representation of the grid")
	flag.IntVar(&part, "part", 1, "Runs the program with the specified part")
	flag.Parse()

	var dataFilepath string
	var rope *Rope
	var grid *Grid

	if part == 1 {
		if mode {
			// Test mode
			dataFilepath = "input/test.txt"
			rope = NewRope(Vector2{0, 4}, 0)
			grid = NewGrid(6, 5, rope)

		} else {
			dataFilepath = "input/prod.txt"
			rope = NewRope(Vector2{500, 500}, 0)
			grid = NewGrid(1000, 1000, rope)
		}

		// Read data from file
		data, err := os.ReadFile(dataFilepath)
		if err != nil {
			panic(err)
		}

		grid.Render = render

		for _, line := range strings.Split(string(data), "\n") {
			moveOrder := ParseMove(line)
			moveOrder.Execute(grid)
		}

		count := grid.CountVisited()
		fmt.Printf("Total visited cells: %d\n", count)
	} else {
		if mode {
			// Test mode
			dataFilepath = "input/test_2.txt"
			rope = NewRope(Vector2{11, 15}, 8)
			grid = NewGrid(27, 21, rope)

		} else {
			dataFilepath = "input/prod.txt"
			rope = NewRope(Vector2{500, 500}, 8)
			grid = NewGrid(1000, 1000, rope)
		}

		// Read data from file
		data, err := os.ReadFile(dataFilepath)
		if err != nil {
			panic(err)
		}

		grid.Render = render

		for _, line := range strings.Split(string(data), "\n") {
			moveOrder := ParseMove(line)
			moveOrder.Execute(grid)
		}

		count := grid.CountVisited()
		fmt.Printf("Total visited cells: %d\n", count)
	}
}
