package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Rx = Rx ^ HHLL
func xoriRxHHLL(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := v.Regs[x] ^ int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rx = Rx ^ Ry
func xorRxRy(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := v.Regs[x] ^ v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Rz = Rx ^ Ry
func xorRxRyRz(v *vm.State, o vm.Opcode) error {
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
