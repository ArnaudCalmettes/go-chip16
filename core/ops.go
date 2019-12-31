package core

import "fmt"

var cpuOps [256]*operation // All CPU operations

type opCallback func(*VirtualMachine, Opcode) error

type operation struct {
	Code        byte
	Description string
	Execute     opCallback
}

func setOp(code byte, desc string, exec opCallback) {
	c := int(code)
	if cpuOps[c] != nil {
		panic(fmt.Sprintf("Instruction %#02x already exists", code))
	}
	cpuOps[c] = &operation{code, desc, exec}
}
