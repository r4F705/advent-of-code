package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Directions [4]string

type Forest struct {
	trees [][]int
	dirs  Directions
}

func NewForest(data *string) *Forest {
	forest := Forest{
		trees: [][]int{},
		dirs: Directions{
			"right",
			"down",
			"left",
			"up",
		},
	}

	forest.parseData(data)

	return &forest
}

func (f *Forest) parseData(data *string) {
	for _, line := range strings.Split((*data), "\n") {
		treeRow := []int{}
		for _, char := range line {
			treeHeight, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			treeRow = append(treeRow, treeHeight)
		}

		f.addRow(treeRow)
	}

	fmt.Println("Forest:")
	for _, row := range f.trees {
		fmt.Println(row)
	}
}

func (f *Forest) addRow(row []int) {
	f.trees = append(f.trees, row)
}

func (f *Forest) compareTree(treeA, treeB int) bool {
	return treeA > treeB
}

func (f *Forest) GetTree(x, y int) int {
	return f.trees[x][y]
}

func (f *Forest) IsTreeVisible(row, col int) bool {

	// If tree is on the edge of the forest, it is visible
	if row == 0 || row == len(f.trees)-1 || col == 0 || col == len(f.trees[row])-1 {
		return true
	}

	visibleFromDirection := []bool{false, false, false, false}
	curTreeHeight := f.GetTree(row, col)

	for _, dir := range f.dirs {
		switch dir {

		case "right":
			visible := true
			for i := len(f.trees[row]) - 1; i > col; i-- {
				if !f.compareTree(curTreeHeight, f.GetTree(row, i)) {
					visible = false
					break
				}
			}
			visibleFromDirection[0] = visible
			// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "down":
			visible := true
			for i := len(f.trees) - 1; i > row; i-- {
				if !f.compareTree(curTreeHeight, f.GetTree(i, col)) {
					visible = false
					break
				}
			}
			visibleFromDirection[1] = visible
			// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "left":
			visible := true
			for i := 0; i < col; i++ {
				if !f.compareTree(curTreeHeight, f.GetTree(row, i)) {
					visible = false
					break
				}
			}
			visibleFromDirection[2] = visible
			// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "up":
			visible := true
			for i := 0; i < row; i++ {
				if !f.compareTree(curTreeHeight, f.GetTree(i, col)) {
					visible = false
					break
				}
			}
			visibleFromDirection[3] = visible
			// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)
		}
	}

	// If tree is visible from any direction, return true
	for _, visible := range visibleFromDirection {
		if visible {
			return true
		}
	}

	return false
}

func (f *Forest) TreeScenicScore(row, col int) int {
	score := 1
	curTreeHeight := f.GetTree(row, col)

	for _, dir := range f.dirs {
		switch dir {

		case "right":
			visibleTrees := 0
			for i := min(col+1, len(f.trees[row])); i < len(f.trees[row]); i++ {
				visibleTrees++
				if !f.compareTree(curTreeHeight, f.GetTree(row, i)) {
					break
				}
			}
			score *= visibleTrees
		// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "down":
			visibleTrees := 0
			for i := min(row+1, len(f.trees)); i < len(f.trees); i++ {
				visibleTrees++
				if !f.compareTree(curTreeHeight, f.GetTree(i, col)) {
					break
				}
			}
			score *= visibleTrees
		// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "left":
			visibleTrees := 0
			for i := max(col-1, 0); i >= 0; i-- {
				visibleTrees++
				if !f.compareTree(curTreeHeight, f.GetTree(row, i)) {
					break
				}
			}
			score *= visibleTrees
		// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)

		case "up":
			visibleTrees := 0
			for i := max(row-1, 0); i >= 0; i-- {
				visibleTrees++
				if !f.compareTree(curTreeHeight, f.GetTree(i, col)) {
					break
				}
			}
			score *= visibleTrees
			// fmt.Printf("Tree[%d] (%d, %d) visibility from %s: %v\n", curTreeHeight, row, col, dir, visible)
		}
	}

	return score
}

func (f *Forest) CountVisibleTrees() int {
	count := 0

	for row := range f.trees {
		for col := range f.trees[row] {
			if f.IsTreeVisible(row, col) {
				// fmt.Printf("Tree at (%d, %d) is visible\n", row, col)
				count++
			}
		}
	}

	return count
}

func (f *Forest) MaxScenicScore() (int, []int) {
	maxScore := 0
	coords := []int{0, 0}

	for row := range f.trees {
		for col := range f.trees[row] {
			score := f.TreeScenicScore(row, col)
			if score > maxScore {
				maxScore = score
				coords[0] = row
				coords[1] = col
			}
		}
	}

	return maxScore, coords
}
