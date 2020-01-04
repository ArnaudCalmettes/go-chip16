package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Rx = Rx & HHLL
func andiRxHHLL(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := v.Regs[x] & int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx & Ry
func andRxRy(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := v.Regs[x] & v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rz = Rx & Ry
func andRxRyRz(v *vm.State, o vm.Opcode) error {
	res := v.Regs[o.X()] & v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.Z()] = res
	return nil
}

// Compute Rx & HHLL, discard result
func tstiRxHHLL(v *vm.State, o vm.Opcode) error {
	v.Flags.SetZN(v.Regs[o.X()] & int16(o.HHLL()))
	return nil
}

// Compute Rx & Ry, discard result
func tstRxRy(v *vm.State, o vm.Opcode) error {
	v.Flags.SetZN(v.Regs[o.X()] & v.Regs[o.Y()])
	return nil
}

func init() {
	setOp(0x60, "ANDI Rx, HHLL", andiRxHHLL)
	setOp(0x61, "AND Rx, Ry", andRxRy)
	setOp(0x62, "AND Rx, Ry, Rz", andRxRyRz)
	setOp(0x63, "TSTI Rx, HHLL", tstiRxHHLL)
	setOp(0x64, "TST Rx, Ry", tstRxRy)
}
