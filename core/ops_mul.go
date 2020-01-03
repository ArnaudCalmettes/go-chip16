package core

import "math"

// Utility flag-setting multiplication for signed 16-bit integers
func mul16(x, y int16, flags *CPUFlags) int16 {
	res32 := int32(x) * int32(y)
	res := int16(res32)

	flags.SetCarry(res32 > math.MaxInt16 || res32 < math.MinInt16)
	flags.SetZN(res)
	return res
}

// Rx = Rx * HHLL
func muliRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x] = mul16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return nil
}

// Rx = Rx * Ry
func mulRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	v.Regs[x] = mul16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return nil
}

// Rz = Rx * Ry
func mulRxRyRz(v *VirtualMachine, o Opcode) error {
	v.Regs[o.Z()] = mul16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return nil
}

func init() {
	setOp(0x90, "MULI Rx, HHLL", muliRxHHLL)
	setOp(0x91, "MUL Rx, Ry", mulRxRy)
	setOp(0x92, "MUL Rx, Ry, Rz", mulRxRyRz)
}
