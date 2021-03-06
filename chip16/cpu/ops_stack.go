package cpu

import (
	"encoding/binary"
	"fmt"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

var errStackOverflow = fmt.Errorf("Stack overflow")
var errStackUnderflow = fmt.Errorf("Stack underflow")

// Push Rx onto the stack
func pushRx(v *vm.State, o vm.Opcode) error {
	if v.SP >= vm.IOStart {
		return errStackOverflow
	}
	v.PutInt16At(v.Regs[o.X()], v.SP)
	v.SP += 2
	return nil
}

// Pop Rx off the stack
func popRx(v *vm.State, o vm.Opcode) error {
	if v.SP <= vm.StackStart {
		return errStackUnderflow
	}
	v.SP -= 2
	v.Regs[o.X()], _ = v.Int16At(v.SP)
	return nil
}

// Push all registers to the stack
func pushAll(v *vm.State, o vm.Opcode) error {
	if v.SP > vm.IOStart-32 {
		return errStackOverflow
	}
	for _, rx := range v.Regs {
		binary.LittleEndian.PutUint16(v.RAM[v.SP:], uint16(rx))
		v.SP += 2
	}
	return nil
}

// Pop all registers off the stack
func popAll(v *vm.State, o vm.Opcode) error {
	if v.SP < vm.StackStart+32 {
		return errStackUnderflow
	}
	for i := len(v.Regs) - 1; i >= 0; i-- {
		v.SP -= 2
		v.Regs[i] = int16(binary.LittleEndian.Uint16(v.RAM[v.SP:]))
	}
	return nil
}

// Push flags to the stack
func pushF(v *vm.State, o vm.Opcode) error {
	if v.SP >= vm.IOStart {
		return errStackOverflow
	}
	v.RAM[v.SP] = uint8(v.Flags)
	v.SP += 2
	return nil
}

// Pop flags off the stack
func popF(v *vm.State, o vm.Opcode) error {
	if v.SP <= vm.StackStart {
		return errStackUnderflow
	}
	v.SP -= 2
	v.Flags = vm.CPUFlags(v.RAM[v.SP])
	return nil
}

func init() {
	setOp(0xC0, "PUSH Rx", pushRx)
	setOp(0xC1, "POP Rx", popRx)
	setOp(0xC2, "PUSHALL", pushAll)
	setOp(0xC3, "POPALL", popAll)
	setOp(0xC4, "PUSHF", pushF)
	setOp(0xC5, "POPF", popF)
}
