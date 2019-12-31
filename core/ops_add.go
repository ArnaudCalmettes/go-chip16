package core

// Utility flag-setting add function for signed 16-bit integers
func add16(x, y int16) (sum int16, flags CPUFlags) {
	sum = x + y
	// The sum will overflow if both top bits are set (x & y) or if one of them
	// is (x | y), and a carry from the lower place happened. If such a carry
	// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
	flags.SetCarry(((x&y)|((x|y)&^sum))>>15 != 0)
	flags.SetOverflow((x > 0 && y > 0 && sum < 0) || (x < 0 && y < 0 && sum > 0))
	flags.SetNegative(sum < 0)
	flags.SetZero(sum == 0)
	return
}

// Set Rx to Rx + HHLL
func addiRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = add16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Set Rx to Rx + Ry
func addRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = add16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Set Rz to Rx + Ry
func addRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()], v.Flags = add16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x40, "ADDI RX, HHLL", addiRxHHLL)
	setOp(0x41, "ADD RX, RY", addRxRy)
	setOp(0x42, "ADD RX, RY, RZ", addRxRyRz)
}
