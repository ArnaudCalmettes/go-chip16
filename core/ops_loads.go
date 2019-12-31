package core

// Set Rx to HHLL
func ldiRxHHLL(v *VirtualMachine, o Opcode) error {
	v.Regs[o.X()] = int16(o.HHLL())
	return nil
}

// Set SP to HHLL
func ldiSPHHLL(v *VirtualMachine, o Opcode) error {
	v.SP = Pointer(o.HHLL())
	return nil
}

// Set Rx to [HHLL]
func ldmRxHHLL(v *VirtualMachine, o Opcode) error {
	var err error
	v.Regs[o.X()], err = v.Int16At(Pointer(o.HHLL()))
	return err
}

// Set Rx to [Ry]
func ldmRxRy(v *VirtualMachine, o Opcode) error {
	var err error
	v.Regs[o.X()], err = v.Int16At(Pointer(v.Regs[o.Y()]))
	return err
}

// Set Rx to Ry
func mov(v *VirtualMachine, o Opcode) error {
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
