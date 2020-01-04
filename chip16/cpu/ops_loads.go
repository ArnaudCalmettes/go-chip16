package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Set Rx to HHLL
func ldiRxHHLL(v *vm.State, o vm.Opcode) error {
	v.Regs[o.X()] = int16(o.HHLL())
	return nil
}

// Set SP to HHLL
func ldiSPHHLL(v *vm.State, o vm.Opcode) error {
	v.SP = vm.Pointer(o.HHLL())
	return nil
}

// Set Rx to [HHLL]
func ldmRxHHLL(v *vm.State, o vm.Opcode) error {
	var err error
	v.Regs[o.X()], err = v.Int16At(vm.Pointer(o.HHLL()))
	return err
}

// Set Rx to [Ry]
func ldmRxRy(v *vm.State, o vm.Opcode) error {
	var err error
	v.Regs[o.X()], err = v.Int16At(vm.Pointer(v.Regs[o.Y()]))
	return err
}

// Set Rx to Ry
func mov(v *vm.State, o vm.Opcode) error {
	v.Regs[o.X()] = v.Regs[o.Y()]
	return nil
}

func init() {
	setOp(0x20, "LDI Rx, HHLL", ldiRxHHLL)
	setOp(0x21, "LDI SP, HHLL", ldiSPHHLL)
	setOp(0x22, "LDM Rx, HHLL", ldmRxHHLL)
	setOp(0x23, "LDM Rx, Ry", ldmRxRy)
	setOp(0x24, "MOV Rx, Ry", mov)
}
