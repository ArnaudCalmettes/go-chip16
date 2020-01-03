package core

// Rx = Rx ^ HHLL
func xoriRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] ^ int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx ^ Ry
func xorRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] ^ v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rz = Rx ^ Ry
func xorRxRyRz(v *VirtualMachine, o Opcode) error {
	res := v.Regs[o.X()] ^ v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.Z()] = res
	return nil
}

func init() {
	setOp(0x80, "XORI Rx, HHLL", xoriRxHHLL)
	setOp(0x81, "XOR Rx, Ry", xorRxRy)
	setOp(0x82, "XOR Rx, Ry, Rz", xorRxRyRz)
}
