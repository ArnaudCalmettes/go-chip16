package core

// Utility function to compute flag-setting bitwise AND on 16-bit integers
func and16(x, y int16) (res int16, flags CPUFlags) {
	res = x & y
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

// Rx = Rx & HHLL
func andiRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = and16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Rx = Rx & Ry
func andRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = and16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Rz = Rx & Ry
func andRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()], v.Flags = and16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

// Compute Rx & HHLL, discard result
func tstiRxHHLL(v *VirtualMachine, o Opcode) error {
	_, v.Flags = and16(v.Regs[o.X()], int16(o.HHLL()))
	return nil
}

// Compute Rx & Ry, discard result
func tstRxRy(v *VirtualMachine, o Opcode) error {
	_, v.Flags = and16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x60, "ANDI Rx, HHLL", andiRxHHLL)
	setOp(0x61, "AND Rx, Ry", andRxRy)
	setOp(0x62, "AND Rx, Ry, Rz", andRxRyRz)
	setOp(0x63, "TSTI Rx, HHLL", tstiRxHHLL)
	setOp(0x64, "TST Rx, Ry", tstRxRy)
}
