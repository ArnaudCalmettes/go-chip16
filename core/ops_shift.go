package core

// Rx = Rx << N
func shlRxN(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] << o.N()
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx >> N, logical shift
func shrRxN(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := int16(uint16(v.Regs[x]) >> o.N())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx >> N, copying leading bit
func sarRxN(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] >> o.N()
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx << Ry
func shlRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] << v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx >> Ry, logical shift
func shrRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := int16(uint16(v.Regs[x]) >> v.Regs[o.Y()])
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx >> Ry, copying leading bit
func sarRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	res := v.Regs[x] >> v.Regs[o.Y()]
	v.Regs[x] = res
	v.Flags.SetZN(res)
	return nil
}

func init() {
	setOp(0xB0, "SHL Rx, N", shlRxN)
	setOp(0xB1, "SHR Rx, N", shrRxN)
	setOp(0xB2, "SAR Rx, N", sarRxN)
	setOp(0xB3, "SHL Rx, Ry", shlRxRy)
	setOp(0xB4, "SHR Rx, Ry", shrRxRy)
	setOp(0xB5, "SAR Rx, Ry", sarRxRy)
}