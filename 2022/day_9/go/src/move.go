package main

import (
	"strconv"
	"strings"
)

type MoveOrder struct {
	Direction string
	Amount    int
}

func NewMoveOrder(direction string, amount int) *MoveOrder {
	return &MoveOrder{
		Direction: direction,
		Amount:    amount,
	}
}

func ParseMove(input string) *MoveOrder {
	direction, amountStr, found := strings.Cut(input, " ")
	if !found {
		panic("Invalid input")
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		panic("Invalid input: " + err.Error())
	}

	return NewMoveOrder(direction, amount)
}

func (m *MoveOrder) Execute(grid *Grid) {
	grid.Rope.Move(m, grid.Update)
}
