package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Board [][]Node

type Pos struct {
	X int
	Y int
}

type Node struct {
	Elevation int8
	Distance  uint
	Visited   bool
	IsStart   bool
	IsEnd     bool
	Neighbors []*Node
}

func createBoard(filename string) (Board, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	if runtime.GOOS == "windows" {
		lines = strings.Split(strings.TrimSpace(string(data)), "\r\n")
	} else {
		lines = strings.Split(strings.TrimSpace(string(data)), "\n")
	}
	board := make(Board, len(lines))

	for i, line := range lines {
		nodes := []Node{}
		for _, v := range line {

			isStart := false
			isEnd := false

			if v == 'S' {
				isStart = true
			} else if v == 'E' {
				isEnd = true
			}

			var distance uint
			if isStart {
				distance = 0
			} else {
				distance = ^uint(0)
			}

			nodes = append(nodes, Node{
				Elevation: int8(v),
				Distance:  distance,
				Visited:   false,
				IsStart:   isStart,
				IsEnd:     isEnd,
			})
		}
		board[i] = nodes
	}

	// create connections
	for i, row := range board {
		for j := range row {
			board.connectNodes(Pos{
				X: i,
				Y: j,
			})
		}
	}

	return board, nil
}

func (b *Board) isValidPos(pos Pos) bool {

	if pos.X < 0 || pos.Y < 0 || pos.X > len(*b)-1 || pos.Y > len((*b)[0])-1 {
		return false
	}

	return true
}

func (b *Board) connectNodes(pos Pos) {
	node := b.getNode(pos)

	connectedNodes := []Pos{
		// up
		{
			pos.X,
			pos.Y + 1,
		},
		// right
		{
			pos.X + 1,
			pos.Y,
		},
		// down
		{
			pos.X,
			pos.Y - 1,
		},
		// left
		{
			pos.X - 1,
			pos.Y,
		},
	}

	for _, v := range connectedNodes {
		if b.isValidPos(v) {
			node.Neighbors = append(node.Neighbors, b.getNode(v))
		}
	}
}

func (b *Board) getNode(pos Pos) *Node {
	return &(*b)[pos.X][pos.Y]
}

func (b *Board) findStart() *Node {
	for i := range *b {
		for j := range (*b)[0] {
			node := b.getNode(Pos{
				X: i,
				Y: j,
			})
			if node.IsStart {
				return node
			}

		}
	}
	return nil
}

func (b *Board) findEnd() *Node {
	for i := range *b {
		for j := range (*b)[0] {
			node := b.getNode(Pos{
				X: i,
				Y: j,
			})
			if node.IsEnd {
				return node
			}

		}
	}
	return nil
}

func (b *Board) findShortestPath() uint {
	endNode := b.findEnd()
	// Update elevation to help algorithms
	endNode.Elevation = 'z'

	curNode := b.findStart()
	// Update elevation to help algorithms
	curNode.Elevation = 'a'

	for {

		if curNode == endNode {
			b.visualizeBoard(curNode)
			return curNode.Distance
		}

		// Step 1: update estimates
		curNode.Visited = true

		for _, neighbor := range curNode.Neighbors {
			if neighbor.Elevation-curNode.Elevation <= 1 {
				neighbor.Distance = curNode.Distance + 1
			}
		}

		// Step 2: choose next vertex
		shortest := &Node{
			Distance: ^uint(0),
		}

		for i := range *b {
			for j := range (*b)[0] {
				node := b.getNode(Pos{
					X: i,
					Y: j,
				})

				if !node.Visited && node.Distance < shortest.Distance {
					shortest = node
				}

			}
		}

		if shortest.Elevation == 0 {
			b.visualizeBoard(curNode)

			panic("could not find any neighbor and this should not happen")
		}

		curNode = shortest

		// b.visualizeBoard(curNode)

		// time.Sleep(time.Duration(1 * time.Second))
		// if runtime.GOOS == "windows" {
		// 	cmd := exec.Command("cmd", "/c", "cls") // Windows
		// 	cmd.Stdout = os.Stdout
		// 	cmd.Run()
		// } else {
		// 	cmd := exec.Command("clear") // Linux and macOS
		// 	cmd.Stdout = os.Stdout
		// 	cmd.Run()
		// }
	}

}

func (b *Board) visualizeBoard(curNode *Node) {
	// ANSI color codes
	red := "\033[31m"
	green := "\033[32m"
	white := "\033[37m"
	reset := "\033[0m"

	// Find the position of the current node
	var curPos Pos
	for i := range *b {
		for j := range (*b)[0] {
			if b.getNode(Pos{X: i, Y: j}) == curNode {
				curPos = Pos{X: i, Y: j}
				break
			}
		}
	}

	// Calculate boundaries for the 8x8 view
	startX := max(0, curPos.X-10)
	endX := min(len(*b), curPos.X+10)
	startY := max(0, curPos.Y-10)
	endY := min(len((*b)[0]), curPos.Y+10)

	// Adjust if the view is smaller than 8x8 due to board boundaries
	if endX-startX < 20 {
		if startX == 0 {
			endX = min(len(*b), startX+20)
		} else if endX == len(*b) {
			startX = max(0, endX-20)
		}
	}

	if endY-startY < 20 {
		if startY == 0 {
			endY = min(len((*b)[0]), startY+20)
		} else if endY == len((*b)[0]) {
			startY = max(0, endY-20)
		}
	}

	// Define the format for each cell, ensuring a fixed width
	cellFormat := "%-10s"

	// Render the view
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			node := b.getNode(Pos{X: i, Y: j})
			if node == nil {
				continue
			}
			color := red // Default color for unvisited nodes
			if node == curNode {
				color = white // White color for the current node
			} else if node.Visited {
				color = green // Green color for visited nodes
			}
			if curNode == node {
				fmt.Printf(color+cellFormat+reset, "X")
			} else if node.Distance == ^uint(0) {
				fmt.Printf(color+cellFormat+reset, fmt.Sprintf("%d(∞)", node.Elevation))
			} else {
				fmt.Printf(color+cellFormat+reset, fmt.Sprintf("%d(%d)", node.Elevation, node.Distance))
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func main() {
	board, err := createBoard("C:/Users/nikos/Projects/advent-of-code/2022/day_12/go/input/prod.txt")
	if err != nil {
		panic(err)
	}

	steps := board.findShortestPath()
	fmt.Printf("Reached destination node in %d steps\n", steps)
}
