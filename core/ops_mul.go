package core

import "math"

// Utility flag-setting multiplication for signed 16-bit integers
func mul16(x, y int16) (res int16, flags CPUFlags) {
	res32 := int32(x) * int32(y)
	res = int16(res32)

	flags.SetCarry(res32 > math.MaxInt16 || res32 < math.MinInt16)
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

// Rx = Rx * HHLL
func muliRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = mul16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Rx = Rx * Ry
func mulRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = mul16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Rz = Rx * Ry
func mulRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()], v.Flags = mul16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x90, "MULI Rx, HHLL", muliRxHHLL)
	setOp(0x91, "MUL Rx, Ry", mulRxRy)
	setOp(0x92, "MUL Rx, Ry, Rz", mulRxRyRz)
}
