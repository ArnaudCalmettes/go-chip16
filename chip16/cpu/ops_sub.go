package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Utility flag-setting difference between signed 16-bit integers
func sub16(x, y int16) (diff int16, flags vm.CPUFlags) {
	diff = x - y

	// Carry flag is set if the difference underflows.
	// The difference underflows if the top bit of x is not set and the top
	// bit of y is set (^x & y) or if they are the same (^(x ^ y)) and a borrow
	// from the lower place happens. If that borrow happens, the result
	// will be 1 - 1 - 1 = 0 - 0 - 1 = 1 (& diff).
	flags.SetCarry(((^x & y) | (^(x ^ y) & diff)) < 0)

	// Overflow flag is set when:
	// diff > 0 && x < 0 && y > 0,
	// diff < 0 && x > 0 && y < 0.
	// i.e. top bit is the same in diff and y (^(diff^y)) and differs between x
	// and y (x^y)
	flags.SetOverflow((x^y)&^(diff^y) < 0)

	flags.SetZN(diff)
	return
}

// Set Rx to Rx - HHLL
func subiRxHHLL(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = sub16(v.Regs[x], int16(o.HHLL()))
	return nil
}

// Set Rx to Rx - Ry
func subRxRy(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x], v.Flags = sub16(v.Regs[x], v.Regs[o.Y()])
	return nil
}

// Set Rz to Rx - Ry
func subRxRyRz(v *vm.State, o vm.Opcode) error {
	v.Regs[o.Z()], v.Flags = sub16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

// Compute Rx - HHLL, discard the result
func cmpiRxHHLL(v *vm.State, o vm.Opcode) error {
	_, v.Flags = sub16(v.Regs[o.X()], int16(o.HHLL()))
	return nil
}

// Compute Rx - Ry, discard the result
func cmpRxRy(v *vm.State, o vm.Opcode) error {
	_, v.Flags = sub16(v.Regs[o.X()], v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x50, "SUBI Rx, HHLL", subiRxHHLL)
	setOp(0x51, "SUB Rx, Ry", subRxRy)
	setOp(0x52, "SUB Rx, Ry, Rz", subRxRyRz)
	setOp(0x53, "CMPI Rx, HHLL", cmpiRxHHLL)
	setOp(0x54, "CMP Rx, Ry", cmpRxRy)
}
