package core

// Utility function to compute flag-setting bitwise OR on 16-bit integers
func or16(x, y int16) (res int16, flags CPUFlags) {
	res = x | y
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

// Rx = Rx | HHLL
func oriRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = or16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Rx = Rx | Ry
func orRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = or16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Rz = Rx | Ry
func orRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()], v.Flags = or16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x70, "ORI Rx, HHLL", oriRxHHLL)
	setOp(0x71, "OR Rx, Ry", orRxRy)
	setOp(0x72, "OR Rx, Ry, Rz", orRxRyRz)
}
