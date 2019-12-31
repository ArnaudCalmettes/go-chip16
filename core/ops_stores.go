package core

// Store Rx at [HHLL]
func stmRxHHLL(v *VirtualMachine, o Opcode) error {
	return v.PutInt16At(v.Regs[o.X()], Pointer(o.HHLL()))
}

// Store Rx at [Ry]
func stmRxRy(v *VirtualMachine, o Opcode) error {
	return v.PutInt16At(v.Regs[o.X()], Pointer(v.Regs[o.Y()]))
}

func init() {
	setOp(0x30, "STM Rx, HHLL", stmRxHHLL)
	setOp(0x31, "STM Rx, Ry", stmRxRy)
}
