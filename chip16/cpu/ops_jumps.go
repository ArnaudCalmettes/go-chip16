package cpu

import "github.com/ArnaudCalmettes/go-chip16/chip16/vm"

func call(v *vm.State, p vm.Pointer) {
	v.PutPointerAt(v.PC, v.SP)
	v.SP += 2
	v.PC = p
}

// Unconditional jump to HHLL
func jmpHHLL(v *vm.State, o vm.Opcode) error {
	v.PC = vm.Pointer(o.HHLL())
	return nil
}

// Jump to HHLL if carry is set
func jmc(v *vm.State, o vm.Opcode) error {
	if v.Flags.Carry() {
		v.PC = vm.Pointer(o.HHLL())
	}
	return nil
}

// Jump to HHLL if a flag condition X is met
func jx(v *vm.State, o vm.Opcode) error {
	if cond, err := v.Flags.Condition(o.X()); err != nil {
		return err
	} else if cond {
		v.PC = vm.Pointer(o.HHLL())
	}
	return nil
}

// Jump to HHLL if Rx == RY
func jme(v *vm.State, o vm.Opcode) error {
	if v.Regs[o.X()] == v.Regs[o.Y()] {
		v.PC = vm.Pointer(o.HHLL())
	}
	return nil
}

// Perform an inconditional call to HHLL
func callHHLL(v *vm.State, o vm.Opcode) error {
	call(v, vm.Pointer(o.HHLL()))
	return nil
}

// Return from function call
func ret(v *vm.State, _ vm.Opcode) error {
	var err error
	v.SP -= 2
	v.PC, err = v.PointerAt(v.SP)
	return err
}

// Unconditional jump to Rx
func jmpRx(v *vm.State, o vm.Opcode) error {
	v.PC = vm.Pointer(v.Regs[o.X()])
	return nil
}

// Conditional call to HHLL
func cx(v *vm.State, o vm.Opcode) error {
	if cond, err := v.Flags.Condition(o.X()); err != nil {
		return err
	} else if cond {
		call(v, vm.Pointer(o.HHLL()))
	}
	return nil
}

// Perform an inconditional call to Rx
func callRx(v *vm.State, o vm.Opcode) error {
	call(v, vm.Pointer(v.Regs[o.X()]))
	return nil
}

func init() {
	setOp(0x10, "JMP HHLL", jmpHHLL)
	setOp(0x11, "JMC HHLL", jmc)
	setOp(0x12, "Jx HHLL", jx)
	setOp(0x13, "JME Rx, Ry, HHLL", jme)
	setOp(0x14, "CALL HHLL", callHHLL)
	setOp(0x15, "RET", ret)
	setOp(0x16, "JMP Rx", jmpRx)
	setOp(0x17, "Cx HHLL", cx)
	setOp(0x18, "CALL Rx", callRx)
}
