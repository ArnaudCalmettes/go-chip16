package core

// Utility function to compute flag-setting bitwise XOR on 16-bit integers
func xor16(x, y int16) (res int16, flags CPUFlags) {
	res = x ^ y
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

// Rx = Rx ^ HHLL
func xoriRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = xor16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Rx = Rx ^ Ry
func xorRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = xor16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Rz = Rx ^ Ry
func xorRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()], v.Flags = xor16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x80, "XORI Rx, HHLL", xoriRxHHLL)
	setOp(0x81, "XOR Rx, Ry", xorRxRy)
	setOp(0x82, "XOR Rx, Ry, Rz", xorRxRyRz)
}
