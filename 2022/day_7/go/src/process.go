package main

import (
	"strconv"
	"strings"
)

// A struct named Command that has a Name field of type string
type Command struct {
	Line string
}

// A member function named Process that takes a reference to a file system
// and applies the command to the file system
func (c *Command) Process(fs *FileSystem) {

	if strings.HasPrefix(c.Line, "$") {

		data := strings.Split(c.Line, " ")
		command := data[1]

		var arg string

		if len(data) > 2 {
			arg = data[2]
		} else {
			arg = ""
		}

		switch command {
		case "cd":
			if arg == ".." {
				fs.Current = fs.Current.Parent
			} else if arg == "/" {
				fs.Current = fs.Root
			} else {
				fs.ChangeDirectory(arg)
			}
		case "ls":
			return
		default:
			panic("Unknown command")
		}
	} else {
		data := strings.Split(c.Line, " ")

		if data[0] == "dir" {
			fs.CreateDirectory(data[1])
		} else {
			size, err := strconv.Atoi(data[0])
			if err != nil {
				panic(err)
			}

			fs.CreateFile(data[1], size)
		}

	}
}
