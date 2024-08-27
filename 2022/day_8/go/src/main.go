package main

import (
	"flag"
	"fmt"
	"os"
)

func readInput(filename string) (string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {

	var mode bool

	// Read flag from command line to know if we are in test mode
	flag.BoolVar(&mode, "test", false, "Runs the program with test input")
	flag.Parse()

	var dataFilepath string
	if mode {
		// Test mode
		dataFilepath = "input/test.txt"
	} else {
		dataFilepath = "input/prod.txt"
	}

	data, err := readInput(dataFilepath)

	if err != nil {
		panic(err)
	}

	forest := NewForest(&data)
	visibleTrees := forest.CountVisibleTrees()
	fmt.Println("Visible trees:", visibleTrees)

	maxScenicScore, coords := forest.MaxScenicScore()
	fmt.Printf("Max scenic score is %d at tree (%d, %d)\n", maxScenicScore, coords[0], coords[1])
}
