package core

// Rx = Rx & HHLL
func andiRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] & int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx & Ry
func andRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] & v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rz = Rx & Ry
func andRxRyRz(v *VirtualMachine, o Opcode) error {
	res := v.Regs[o.X()] & v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.Z()] = res
	return nil
}

// Compute Rx & HHLL, discard result
func tstiRxHHLL(v *VirtualMachine, o Opcode) error {
	v.Flags.SetZN(v.Regs[o.X()] & int16(o.HHLL()))
	return nil
}

// Compute Rx & Ry, discard result
func tstRxRy(v *VirtualMachine, o Opcode) error {
	v.Flags.SetZN(v.Regs[o.X()] & v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x60, "ANDI Rx, HHLL", andiRxHHLL)
	setOp(0x61, "AND Rx, Ry", andRxRy)
	setOp(0x62, "AND Rx, Ry, Rz", andRxRyRz)
	setOp(0x63, "TSTI Rx, HHLL", tstiRxHHLL)
	setOp(0x64, "TST Rx, Ry", tstRxRy)
}
