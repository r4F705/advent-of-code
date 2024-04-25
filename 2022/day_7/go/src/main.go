package main

import (
	"fmt"
	"os"
	"strings"
)

var MARKER map[string]int

func main() {
	dataFilepath := "input/prod.txt"
	data, err := os.ReadFile(dataFilepath)

	root := Directory{
		Name: "/",
	}

	fs := FileSystem{
		Root:    &root,
		Current: &root,
	}

	if err != nil {
		panic(err)
	}

	for i, v := range strings.Split(string(data), "\n") {

		if i == 0 {
			continue
		}

		command := Command{
			Line: v,
		}

		command.Process(&fs)
	}

	fs.PrintFileSystem(fs.Root, "")

	fmt.Printf("\n\n")

	solveOne(&fs)
	solveTwo(&fs)
}

func solveOne(fs *FileSystem) {
	MARKER = make(map[string]int)
	fs.Current = fs.Root
	predicate := func(size int) bool {
		return size <= 100000
	}
	fs.LocateDirForDelete(&MARKER, predicate, fs.Current.Name)

	sum := 0
	for _, v := range MARKER {
		sum += v
	}
	fmt.Println("Part one solution:", sum)
}

func solveTwo(fs *FileSystem) {
	MARKER = make(map[string]int)
	fs.Current = fs.Root
	usedSpace := fs.CalculateSize(fs.Current)
	fs.Current = fs.Root

	totalSpace := 70000000
	desiredSpace := 30000000
	freeSpace := totalSpace - usedSpace

	predicate := func(size int) bool {
		return freeSpace+size >= desiredSpace
	}

	fs.LocateDirForDelete(&MARKER, predicate, fs.Current.Name)

	min := totalSpace
	for _, v := range MARKER {
		if v < min {
			min = v
		}
	}
	fmt.Println("Part two solution:", min)
}
