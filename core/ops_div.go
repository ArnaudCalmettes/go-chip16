package core

import "fmt"

var errDivideByZero = fmt.Errorf("Division by 0")

func div16(x, y int16) (res int16, flags CPUFlags) {
	res = x / y
	flags.SetCarry(x%y != 0)
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

func rem16(x, y int16) (res int16, flags CPUFlags) {
	res = x % y
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

func mod16(x, y int16) (res int16, flags CPUFlags) {
	res = x % y
	// The sign (top bit) of the result must agree to that of the divisor.
	// If they differ (res^y has top bit set), then adding negative divisor
	// to positive remainder, or positive divisor to negative remainder
	// yields correct modulus behavior.
	if res^y < 0 {
		res += y
	}
	flags.SetZero(res == 0)
	flags.SetNegative(res < 0)
	return
}

// Rx = Rx / HHLL
func diviRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := int16(o.HHLL())
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = div16(v.Regs[x], d)
	return nil
}

// Rx = Rx / Ry
func divRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = div16(v.Regs[x], d)
	return nil
}

// Rz = Rx / Ry
func divRxRyRz(v *VirtualMachine, o Opcode) error {
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[o.Z()], v.Flags = div16(v.Regs[o.X()], d)
	return nil
}

// Rx = Rx MOD HHLL
func modiRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := int16(o.HHLL())
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = mod16(v.Regs[x], d)
	return nil
}

// Rx = Rx MOD Ry
func modRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = mod16(v.Regs[x], d)
	return nil
}

// Rz = Rx MOD Ry
func modRxRyRz(v *VirtualMachine, o Opcode) error {
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[o.Z()], v.Flags = mod16(v.Regs[o.X()], d)
	return nil
}

// Rx = Rx % HHLL
func remiRxHHLL(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := int16(o.HHLL())
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = rem16(v.Regs[x], d)
	return nil
}

// Rx = Rx / Ry
func remRxRy(v *VirtualMachine, o Opcode) error {
	x := o.X()
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[x], v.Flags = rem16(v.Regs[x], d)
	return nil
}

// Rz = Rx / Ry
func remRxRyRz(v *VirtualMachine, o Opcode) error {
	d := v.Regs[o.Y()]
	if d == 0 {
		return errDivideByZero
	}
	v.Regs[o.Z()], v.Flags = rem16(v.Regs[o.X()], d)
	return nil
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
