package main

import (
	"fmt"
	"strconv"
)

// Simple file struct that holds a reference to its parent directory
// and the name of the file and its size
type File struct {
	Name   string
	Parent *Directory
	Size   int
}

// Simple directory struct that hold a reference to its parent directory and
// a list of its children directories and anyfiles in the directory
type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []*File
}

// Simple struct that represents a filesystem that holds a reference to the root and the current directory
type FileSystem struct {
	Root    *Directory
	Current *Directory
}

func (fs *FileSystem) CreateDirectory(name string) {
	dir := Directory{
		Name:   name,
		Parent: fs.Current,
	}

	fs.Current.Children = append(fs.Current.Children, &dir)
}

func (fs *FileSystem) CreateFile(name string, size int) {
	file := File{
		Name:   name,
		Parent: fs.Current,
		Size:   size,
	}

	fs.Current.Files = append(fs.Current.Files, &file)
}

func (fs *FileSystem) ChangeDirectory(name string) {
	for _, v := range fs.Current.Children {
		if v.Name == name {
			fs.Current = v
			return
		}
	}
}

func (fs *FileSystem) PrintFileSystem(dir *Directory, prefix string) {
	fmt.Println(prefix+dir.Name, "(Size: "+strconv.Itoa(fs.CalculateSize(dir))+" bytes)")

	for _, file := range dir.Files {
		fmt.Println(prefix + "├── " + file.Name + " (Size: " + strconv.Itoa(file.Size) + " bytes)")
	}

	for _, childDir := range dir.Children {
		fs.PrintFileSystem(childDir, prefix+"│   ")
	}
}

func (fs *FileSystem) CalculateSize(dir *Directory) int {
	size := 0

	for _, file := range dir.Files {
		size += file.Size
	}

	for _, childDir := range dir.Children {
		size += fs.CalculateSize(childDir)
	}

	return size
}

// Recursive function that locates directories that can be deleted based on a predicate which takes the size of the directory as an argument
func (fs *FileSystem) LocateDirForDelete(marker *map[string]int, predicate func(int) bool, prefix string) {
	dir := fs.Current
	for _, childDir := range dir.Children {
		size := fs.CalculateSize(childDir)
		if predicate(size) {
			(*marker)[prefix+childDir.Name] = size
		}
		fs.Current = childDir
		fs.LocateDirForDelete(marker, predicate, prefix+childDir.Name)
	}
}
