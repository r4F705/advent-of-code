package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"supply_stacks/src/cargo"
)

func load_data(filepath string) ([]string, []string) {
	fd, err := os.OpenFile(filepath, os.O_RDONLY, 0440)
	if err != nil {
		log.Fatal(err.Error())
	}

	fileInfo, _ := fd.Stat()
	buf := make([]byte, fileInfo.Size())
	fd.Read(buf)

	data := strings.Split(string(buf[:]), "\n")
	var cargoData []string
	var instructionData []string

	for i := 0; i < len(data); i += 1 {
		if len(data[i]) == 0 {
			cargoData = data[:i]
			instructionData = data[i+1:]
			break
		}
	}

	return cargoData, instructionData
}

func main() {
	input := "./data/input.txt"
	cargoData, instructionData := load_data(input)

	crg := new(cargo.Cargo)
	crane := new(cargo.Crane)
	crane.Model = 9001
	crane.Debug = true

	crg.Init(cargoData)
	crane.LoadInstructions(instructionData)

	crg.ShowCargo()
	crane.ShowInstructions()

	fmt.Printf("Execute instructions:\n")
	crane.ExecuteInstructions(crg)

	crg.ShowCargo()

	crg.ShowTopLayer()
}
