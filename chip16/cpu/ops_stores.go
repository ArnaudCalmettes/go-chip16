package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Store Rx at [HHLL]
func stmRxHHLL(v *vm.State, o vm.Opcode) error {
	return v.PutInt16At(v.Regs[o.X()], vm.Pointer(o.HHLL()))
}

// Store Rx at [Ry]
func stmRxRy(v *vm.State, o vm.Opcode) error {
	return v.PutInt16At(v.Regs[o.X()], vm.Pointer(v.Regs[o.Y()]))
}

func init() {
	setOp(0x30, "STM Rx, HHLL", stmRxHHLL)
	setOp(0x31, "STM Rx, Ry", stmRxRy)
}
