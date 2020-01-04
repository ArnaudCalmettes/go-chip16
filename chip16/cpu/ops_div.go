package cpu

import (
	"fmt"

	"github.com/ArnaudCalmettes/go-chip16/chip16/vm"
)

var errDivideByZero = fmt.Errorf("Division by 0")

func div16(x, y int16, flags *vm.CPUFlags) (int16, error) {
	if y == 0 {
		return 0, errDivideByZero
	}
	res := x / y
	flags.SetCarry(x%y != 0)
	flags.SetZN(res)
	return res, nil
}

func rem16(x, y int16, flags *vm.CPUFlags) (int16, error) {
	if y == 0 {
		return 0, errDivideByZero
	}
	res := x % y
	flags.SetZN(res)
	return res, nil
}

func mod16(x, y int16, flags *vm.CPUFlags) (int16, error) {
	if y == 0 {
		return 0, errDivideByZero
	}
	res := x % y
	// The sign (top bit) of the result must agree to that of the divisor.
	// If they differ (res^y has top bit set), then adding negative divisor
	// to positive remainder, or positive divisor to negative remainder
	// yields correct modulus behavior.
	if res^y < 0 {
		res += y
	}
	flags.SetZN(res)
	return res, nil
}

// Rx = Rx / HHLL
func diviRxHHLL(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = div16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return
}

// Rx = Rx / Ry
func divRxRy(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = div16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return
}

// Rz = Rx / Ry
func divRxRyRz(v *vm.State, o vm.Opcode) (err error) {
	v.Regs[o.Z()], err = div16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return
}

// Rx = Rx MOD HHLL
func modiRxHHLL(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = mod16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return
}

// Rx = Rx MOD Ry
func modRxRy(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = mod16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return
}

// Rz = Rx MOD Ry
func modRxRyRz(v *vm.State, o vm.Opcode) (err error) {
	v.Regs[o.Z()], err = mod16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return
}

// Rx = Rx % HHLL
func remiRxHHLL(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = rem16(v.Regs[x], int16(o.HHLL()), &v.Flags)
	return
}

// Rx = Rx / Ry
func remRxRy(v *vm.State, o vm.Opcode) (err error) {
	x := o.X()
	v.Regs[x], err = rem16(v.Regs[x], v.Regs[o.Y()], &v.Flags)
	return
}

// Rz = Rx / Ry
func remRxRyRz(v *vm.State, o vm.Opcode) (err error) {
	v.Regs[o.Z()], err = rem16(v.Regs[o.X()], v.Regs[o.Y()], &v.Flags)
	return
}

func init() {
	setOp(0xA0, "DIVI Rx, HHLL", diviRxHHLL)
	setOp(0xA1, "DIV Rx, Ry", divRxRy)
	setOp(0xA2, "DIV Rx, Ry, Rz", divRxRyRz)
	setOp(0xA3, "MODI Rx, HHLL", modiRxHHLL)
	setOp(0xA4, "MOD Rx, Ry", modRxRy)
	setOp(0xA5, "MOD Rx, Ry, Rz", modRxRyRz)
	setOp(0xA6, "REMI Rx, HHLL", remiRxHHLL)
	setOp(0xA7, "REM Rx, Ry", remRxRy)
	setOp(0xA8, "REM Rx, Ry, Rz", remRxRyRz)
}
