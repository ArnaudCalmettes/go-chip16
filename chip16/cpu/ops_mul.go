package cpu

import (
	"math"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

// Utility flag-setting multiplication for signed 16-bit integers
func mul16(x, y int16, flags *vm.CPUFlags) int16 {
	res32 := int32(x) * int32(y)
	res := int16(res32)

	flags.SetCarry(res32 > math.MaxInt16 || res32 < math.MinInt16)
	flags.SetZN(res)
	return res
}

// Rx = Rx * HHLL
func muliRxHHLL(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x] = mul16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return nil
}

// Rx = Rx * Ry
func mulRxRy(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x] = mul16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return nil
}

// Rz = Rx * Ry
func mulRxRyRz(v *vm.State, o vm.Opcode) error {
	v.Regs[o.Z()] = mul16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return nil
}

func init() {
	setOp(0x90, "MULI Rx, HHLL", muliRxHHLL)
	setOp(0x91, "MUL Rx, Ry", mulRxRy)
	setOp(0x92, "MUL Rx, Ry, Rz", mulRxRyRz)
}
