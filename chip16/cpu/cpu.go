package cpu

import (
	"fmt"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

type opCallback func(*vm.State, vm.Opcode) error

type operation struct {
	Code        byte
	Description string
	Execute     opCallback
}

var cpuOps [256]*operation // All CPU operations

// SetOp registers a CPU operation to the VM
func setOp(code byte, desc string, exec opCallback) {
	c := int(code)
	if cpuOps[c] != nil {
		panic(fmt.Sprintf("Instruction %#02x already exists", code))
	}
	cpuOps[c] = &operation{code, desc, exec}
}

// Eval evaluates an Opcode
func Eval(v *vm.State, o vm.Opcode) error {
	op := o.Op()
	if inst := cpuOps[op]; inst != nil {
		if err := inst.Execute(v, o); err != nil {
			return fmt.Errorf("opcode (%#08x) %s", o, err)
		}
	} else {
		return fmt.Errorf("Unknown Opcode: %#08x", o)
	}

	return v.Check()
}
