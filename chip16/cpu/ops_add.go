package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Utility flag-setting add function for signed 16-bit integers
func add16(x, y int16, flags *vm.CPUFlags) int16 {
	sum := x + y

	// Carry flag is set if the sum overflows.
	// The sum overflows if both top bits are set (x & y) or if one of them
	// is (x | y), and a carry from the lower place happened. If such a carry
	// happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
	flags.SetCarry(((x & y) | ((x | y) &^ sum)) < 0)

	// Overflow flag is set if both operands have the same sign and the sign of
	// the sum disagrees with that of the operands. i.e top bit is the same in
	// x and y (^(x^y)), and differs between x and the sum (x^sum).
	flags.SetOverflow((x^sum)&^(x^y) < 0)

	flags.SetZN(sum)
	return sum
}

// Set Rx to Rx + HHLL
func addiRxHHLL(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x] = add16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return nil
}

// Set Rx to Rx + Ry
func addRxRy(v *vm.State, o vm.Opcode) error {
	x := o.X()
	v.Regs[x] = add16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return nil
}

// Set Rz to Rx + Ry
func addRxRyRz(v *vm.State, o vm.Opcode) error {
	v.Regs[o.Z()] = add16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return nil
}

func init() {
	setOp(0x40, "ADDI RX, HHLL", addiRxHHLL)
	setOp(0x41, "ADD RX, RY", addRxRy)
	setOp(0x42, "ADD RX, RY, RZ", addRxRyRz)
}
