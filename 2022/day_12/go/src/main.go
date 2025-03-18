package main

import (
	"fmt"
	"os"
	"strings"
)

type Board [][]Node

type Pos struct {
	X int
	Y int
}

type Node struct {
	Elevation uint8
	Distance  int
	Visited   bool
	IsStart   bool
	IsEnd     bool
}

func createBoard(filename string) (Board, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
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

			nodes = append(nodes, Node{
				Elevation: uint8(v),
				Distance:  ^int(0),
				Visited:   false,
				IsStart:   isStart,
				IsEnd:     isEnd,
			})
		}
		board[i] = nodes
	}

	return board, nil
}

func (b *Board) getNode(pos Pos) *Node {
	return &(*b)[pos.X][pos.Y]
}

func (b *Board) findStart() Pos {
	for i, row := range *b {
		for j, node := range row {
			if node.IsStart {
				return Pos{
					X: i,
					Y: j,
				}
			}
		}
	}

	return Pos{
		X: -1,
		Y: -1,
	}
}

func (b *Board) findEnd() Pos {
	for i, row := range *b {
		for j, node := range row {
			if node.IsEnd {
				return Pos{
					X: i,
					Y: j,
				}
			}
		}
	}

	return Pos{
		X: -1,
		Y: -1,
	}
}

func (b *Board) findShortestPath() int {
	startPos := b.findStart()
	endPos := b.findEnd()

	curPos := startPos
	// Step 1: Move to unvisited node at shortest distance that we can actually reach
	// Step 2: Look at all the connected nodes and update distance at reachable nodes
}

func main() {
	board, err := createBoard("input/test.txt")
	if err != nil {
		panic(err)
	}

	for _, nodes := range board {
		for _, node := range nodes {
			fmt.Printf("%d    ", node.Elevation)
		}
		fmt.Printf("\n")
	}
}
