package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MinimaAssemblyInterpreter struct {
	registers map[string]int
	memory    map[int]int
}

const (
	MOV   = "MOV"
	ADD   = "ADD"
	SUB   = "SUB"
	LDR   = "LDR"
	STR   = "STR"
	PRINT = "PRINT"
	IF    = "IF"
)

func NewMinimaAssemblyInterpreter() *MinimaAssemblyInterpreter {
	return &MinimaAssemblyInterpreter{
		registers: make(map[string]int),
		memory:    make(map[int]int),
	}
}

func (interp *MinimaAssemblyInterpreter) run(code string) {
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		if line = strings.TrimSpace(line); line != "" {
			interp.execute(line)
		}
	}
}

func (interp *MinimaAssemblyInterpreter) execute(instruction string) {
	parts := strings.Fields(instruction)

	if len(parts) == 0 {
		return
	}

	n := len(parts)
	register := interp.registers

	switch parts[0] {
	case MOV:
		if n < 3 {
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			register[parts[1]] = register[parts[2]]
		} else {
			register[parts[1]] = value
		}

	case ADD:
		if n < 3 {
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			register[parts[1]] += register[parts[2]]
		} else {
			register[parts[1]] += value
		}

	case SUB:
		if n < 3 {
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			register[parts[1]] -= register[parts[2]]
		} else {
			register[parts[1]] -= value
		}

	case LDR:
		if n < 3 {
			return
		}
		address, err := strconv.Atoi(parts[2])
		if err != nil {
			address = register[parts[2]]
		}
		register[parts[1]] = interp.memory[address]

	case STR:
		if n < 3 {
			return
		}
		address, err := strconv.Atoi(parts[2])
		if err != nil {
			address = register[parts[2]]
		}
		interp.memory[address] = register[parts[1]]

	case PRINT:
		if n < 2 {
			return
		}
		fmt.Println(register[parts[1]])

	case IF:
		if n < 5 {
			return
		}
		n := parts[1]
		cmp := parts[2]
		value, err := strconv.Atoi(parts[3])
		if err != nil {
			value = register[parts[3]]
		}

		condition := false
		switch cmp {
		case "==":
			condition = register[n] == value
		case "!=":
			condition = register[n] != value
		case ">":
			condition = register[n] > value
		case ">=":
			condition = register[n] >= value
		case "<":
			condition = register[n] < value
		case "<=":
			condition = register[n] <= value
		default:
			panic(fmt.Sprintf("Invalid comparison operator: %s", cmp))
		}

		if condition {
			interp.execute(strings.Join(parts[4:], " "))
		}

	default:
		panic(fmt.Sprintf("Invalid instruction: %s", parts[0]))
	}
}

func main() {
	code := `
		MOV X 10
		STR X 1
		LDR Y 1
		PRINT Y
		IF X == Y PRINT X
	`

	interpreter := NewMinimaAssemblyInterpreter()
	interpreter.run(code)
}
