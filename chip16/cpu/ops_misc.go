package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

func nop(*vm.State, vm.Opcode) error { return nil }

func init() {
	setOp(0x00, "NOP", nop)
}
