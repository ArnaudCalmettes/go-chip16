package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

// Set Rx to ^HHLL
func notiRxHHLL(v *vm.State, o vm.Opcode) error {
	res := ^int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[o.X()] = res
	return nil
}

// Set Rx to ^Rx
func notRx(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := ^v.Regs[x]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Set Rx to ^Ry
func notRxRy(v *vm.State, o vm.Opcode) error {
	res := ^v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.X()] = res
	return nil
}

// Set Rx to -HHLL
func negiRxHHLL(v *vm.State, o vm.Opcode) error {
	res := -int16(o.HHLL())
	v.Flags.SetZN(res)
	v.Regs[o.X()] = res
	return nil
}

// Set Rx to -Rx
func negRx(v *vm.State, o vm.Opcode) error {
	x := o.X()
	res := -v.Regs[x]
	v.Flags.SetZN(res)
	v.Regs[x] = res
	return nil
}

// Set Rx to -Ry
func negRxRy(v *vm.State, o vm.Opcode) error {
	res := -v.Regs[o.Y()]
	v.Flags.SetZN(res)
	v.Regs[o.X()] = res
	return nil
}

func init() {
	setOp(0xE0, "NOTI RX, HHLL", notiRxHHLL)
	setOp(0xE1, "NOT Rx", notRx)
	setOp(0xE2, "NOT Rx, Ry", notRxRy)
	setOp(0xE3, "NEGI Rx, HHLL", negiRxHHLL)
	setOp(0xE4, "NEG Rx", negRx)
	setOp(0xE5, "NEG Rx, Ry", negRxRy)
}
