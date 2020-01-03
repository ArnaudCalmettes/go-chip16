package core

// Rx = Rx | HHLL
func oriRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] | int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx | Ry
func orRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] | v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rz = Rx | Ry
func orRxRyRz(v *VirtualMachine, o Opcode) error {
	res := v.Regs[o.X()] | v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.Z()] = res
	return nil
}

func init() {
	setOp(0x70, "ORI Rx, HHLL", oriRxHHLL)
	setOp(0x71, "OR Rx, Ry", orRxRy)
	setOp(0x72, "OR Rx, Ry, Rz", orRxRyRz)
}
